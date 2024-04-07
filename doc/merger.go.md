File name: merger.go
Language: Go
Purpose: This file automates the summarization of pull request (PR) details and diff content using an AI model from the OpenAI API, specifically tailored for GitHub repositories.

Important parts: 

- The `PRDetail` struct defines the structure for pull request details (lines 17-26).

- The `Chunk` struct holds a part of the diff content (lines 28-30).

- Function `getFromGPT` sends a prompt to the OpenAI API and returns a summary provided by GPT (lines 32-53).

- Function `summarizeSummary` formats a given summary to markdown using the AI API (lines 55-68).

- Function `chunkToSummary` converts diff chunks into summarized markdown text with the help of the AI API (lines 70-88).

- Function `splitDiff` processes the raw diff content into separate chunks (lines 90-119).

- Function `getDiff` fetches the diff from a GitHub pull request as chunks (lines 121-147).

- Function `getPRDetail` retrieves the details of a GitHub pull request (lines 149-178).

- The `main` function is the entry point where command-line arguments are parsed, PR details and diffs are fetched, and summarization through the AI API is performed for each diff chunk, followed by creating a combined summary output (lines 180-221).