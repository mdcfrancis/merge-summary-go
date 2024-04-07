# internal/library.go
## Language: Go
## Purpose: 
	The file defines a Go package responsible for interfacing with GitHub to fetch pull request (PR) details and diffs, and processing the data using the GPT model to generate summaries and other documentation aids.

## Important parts: 
1. **Struct Definitions (Lines 12-40):**
   Structures `PRDetail`, `Chunk`, and `Config` are defined to hold information about pull request details, content chunks, and configuration settings, respectively.

2. **GPT Interaction Methods (Lines 42-67, 69-91, 93-117, 119-146, 148-176):**
   These methods (`getFromGPT`, `SummarizeSummary`, `ChunkToSummary`, `FileToSummary`, `IndexPage`) use a GPT client to generate different text outputs based on prompts configured within each method. They serve different summarization and content processing purposes, such as summarizing diffs, files and providing a repository index page.

   Example from `getFromGPT` method:
   ```go
   client := openai.NewClient(cfg.GptAuth)
   ```
   
3. **Diff Processing (Lines 178-216):**
   The `splitDiff` function takes an `io.ReadCloser` as input, which likely represents the contents of a diff file, and processes it line by line to chunk the diffs accordingly.

4. **GitHub PR Details and Diff Retrieval (Lines 218-261):**
   Methods `GetDiff` and `GetPRDetail` are responsible for making HTTP requests to the GitHub API to fetch the diff and detailed information of a PR. The methods leverage Go's HTTP package to send the requests and process the responses.

   Example from `GetPRDetail` method:
   ```go
   req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/pulls/"+pr, nil)
   ```