## License Agreement Metadata Extraction and Summarization

This repository contains a Python model that extracts key metadata and generates a summary from legal license agreements in PDF format.
A backend is build in golang to serve the model as a REST API.
Frontend is build in ReactJS.

## Key Features

Text Extraction: Extracts text from PDF files using PyPDFLoader.
Question Answering: Answers specific questions about the agreement using a pre-trained BERT model for question answering.
Summarization: Generates a concise summary of the agreement's key points using a T5 text-to-text model.

## Installation
```bash

    # Run the backend
    cd backend
    go run main.go
    ./word-extractor-apis

    # Run the frontend
    cd frontend
    npm install
    npm start
```

## Usage
A user can upload a PDF file and the model will extract the text, answer questions, and generate a summary of the agreement.
A admin can edit the answered questions and summary will be generated again.

## About the Model

This model streamlines the process of understanding legal license agreements by automatically extracting key metadata and generating a summary.

Here's a quick overview of its core operations:

1. Text Extraction from PDFs:

Utilizes the PyPDFLoader library to efficiently extract text content from PDF files.
Breaks down the extracted text into smaller segments for better processing using a RecursiveCharacterTextSplitter.

2. Answering Your Questions:

Harnesses the power of a pre-trained BERT model, BertForQuestionAnswering, to answer specific questions about the agreement.
Takes your questions as input and searches for answers within the extracted text.
Provides tailored answers for each individual question.

3. Generating a Concise Summary:

Leverages the text-to-text capabilities of a T5 model, T5ForConditionalGeneration, to create a summary of the agreement's key points.
Inputs the generated question-answer pairs and relevant text snippets to produce a succinct overview.