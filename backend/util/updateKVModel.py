import json
import sys
from transformers import T5Tokenizer, T5ForConditionalGeneration, pipeline

checkpoint = "MBZUAI/LaMini-Flan-T5-248M"
tokenizer = T5Tokenizer.from_pretrained(checkpoint)
base_model = T5ForConditionalGeneration.from_pretrained(
    checkpoint, device_map='cpu')


def updateKVModel(info: str):
    pipe_sum = pipeline(
        'summarization',
        model=base_model,
        tokenizer=tokenizer,
        max_length=500,
        min_length=50)
    result = pipe_sum(info)
    result = result[0]['summary_text']

    return result


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python script.py <string>")
        sys.exit(1)

    kvstring = sys.argv[1]
    summary = updateKVModel(kvstring)

    print(json.dumps(summary, indent=4))
