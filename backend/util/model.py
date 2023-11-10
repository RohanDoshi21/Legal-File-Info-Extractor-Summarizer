import os
import sys
from langchain.text_splitter import RecursiveCharacterTextSplitter
from langchain.document_loaders import PyPDFLoader
import spacy
import json
from transformers import pipeline, BertForQuestionAnswering, BertTokenizer
from transformers import T5Tokenizer, T5ForConditionalGeneration

nlp = spacy.load('en_core_web_sm')

model = BertForQuestionAnswering.from_pretrained(
    'bert-large-uncased-whole-word-masking-finetuned-squad')

tokenizer = BertTokenizer.from_pretrained(
    'bert-large-uncased-whole-word-masking-finetuned-squad')

nlp = pipeline('question-answering', model=model,
               tokenizer=tokenizer, device='cuda')

checkpoint = "MBZUAI/LaMini-Flan-T5-248M"
tokenizer = T5Tokenizer.from_pretrained(checkpoint)
base_model = T5ForConditionalGeneration.from_pretrained(checkpoint, device_map='cpu')

def file_preprocessing(file_path):
    loader = PyPDFLoader(file_path)
    pages = loader.load_and_split()
    text_splitter = RecursiveCharacterTextSplitter(
        chunk_size=200, chunk_overlap=50)
    texts = text_splitter.split_documents(pages)
    final_texts = ""
    for text in texts:
        final_texts = final_texts + text.page_content

    return final_texts


def get_file_metadata(file_path):
    generated_text = file_preprocessing(file_path)
    file_size = os.path.getsize(file_path)

    qna = {
        "file_name": os.path.basename(file_path),
        "file_size": file_size
    }

    qna_text = ""

    questions = [
        "What is the legal name of the studio?",
        "Who is the distributor?",
        "What are the key business terms of the license?",
        "What is the consumer facing name of the distributor's service?",
        "Which content titles are licensed?",
        "What are the geographic markets permitted by the license?",
        "What are the start and end dates for the agreement?",
        "What is the business model of the distributor service?",
        "What are the payments due to the studio?",
        "What are the technical specifications for content and metadata delivery?",
        "What is the content protection system expected to be used by the studio?",
        "How and where will the media be delivered?",
        "When are certain titles available for release on the platform?",
        "Is the license exclusive to a particular distributor, window, or geography?",
        "Are there any other business terms?"
    ]

    info = {}
    question_prefixes = ["What is", "Who is", "Are there", "What are", "Which"]
    for question in questions:
        for prefix in question_prefixes:
            if question.startswith(prefix):
                cleaned_question = question[len(prefix):].strip()
                result = nlp(question=question, context=generated_text)
                info[cleaned_question[:-1]] = result['answer']

    for _, (question, answer) in enumerate(info.items(), start=1):
        qna[question] = answer
        qna_text += f"Q: {question}, A: {answer}\n"

    pipe_sum = pipeline(
        'summarization',
        model = base_model,
        tokenizer = tokenizer,
        max_length = 500,
        min_length = 50)
    result = pipe_sum(qna_text)
    result = result[0]['summary_text']

    qna["summary"] = result

    return qna


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <file_path>")
        sys.exit(1)

    file_path = sys.argv[1]
    metadata = get_file_metadata(file_path)

    print(json.dumps(metadata, indent=4))
