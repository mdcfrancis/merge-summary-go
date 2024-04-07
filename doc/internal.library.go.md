# internal/library.go
## Language: Go
## Purpose: 
The file defines a Go package within the 'internal' directory, intended to provide functionality for interacting with both GitHub's API (fetching pull request details and diffs) and with OpenAI's GPT models (summarizing code, diffs, and summaries). It includes data structures to represent pull request details and textual chunks, as well as functions for fetching GitHub data, processing diffs, and generating summaries with AI.

## Important parts:
- `PRDetail` struct (Lines 13-23): Data structure to represent details of a GitHub pull request.
- `Chunk` struct (Lines 25-27): Represents a block of text, such as a diff.
- `Config` struct (Lines 29-32): Holds configuration, such as whether to provide qualitative analysis or authentication for OpenAI's GPT models.
- `getFromGPT` method (Lines 34-57): Uses the OpenAI client to generate completion from a given prompt using GPT models.
- `SummarizeSummary`, `ChunkToSummary`, `FileToSummary`, `IndexPage` methods (Lines 59-128): Functions to summarize different types of inputs with prompts structured for the GPT model.
- `splitDiff` function (Lines 130-156): Splits the diff text into separate chunks.
- `GetDiff` and `GetPRDetail` methods (Lines 158-200): Fetch details and diffs from GitHub for a particular pull request.

Some snippets from the code:
```go
type PRDetail struct { ... } // Struct definition for PR Detail (Line 13)

client := openai.NewClient(cfg.GptAuth) // Creating a new OpenAI client with authentication (Line 36)

chunks, err := splitDiff(resp.Body) // Splitting the body of the response into chunks (Line 194)
```