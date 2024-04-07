File name: merger.go
Language: Go
Purpose: The file is designed to facilitate automated summarization of pull request details and diff contents from GitHub repositories using the OpenAI API.

Important parts:

- The struct `PRDetail` outlines the expected structure to hold pull request information (lines 17-26).

- The struct `Chunk` is used to represent sections of the diff content (lines 28-30).

- The function `getFromGPT` communicates with the OpenAI API to obtain a summary based on a provided prompt (lines 32-53).

- The function `summarizeSummary` transforms a summary into markdown format leveraging the AI API (lines 55-68).

- The function `chunkToSummary` creates summarized markdown text from the diff chunks via the AI API (lines 70-88).

- The function `splitDiff` turns raw diff content into individual chunks for processing (lines 90-119).

- The function `getDiff` extracts the diff from a GitHub pull request into chunks (lines 121-147).

- The function `getPRDetail` fetches the specific details for a GitHub pull request (lines 149-178).

- The `main` function operates as the program’s entry point. It parses input arguments, acquires PR details and diffs, invokes AI API for summarization of each diff chunk, and compiles these summaries into a final output (lines 180-221).