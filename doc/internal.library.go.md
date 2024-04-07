# internal/library.go
## Language: Go
## Purpose: 
This file provides multiple functions to interact with GitHub APIs, parse pull request (PR) details and diffs, and to communicate with OpenAI's API to generate summaries for PRs and code changes.

## Important parts: 
- The `PRDetail` struct defines fields relevant to a GitHub pull request (lines 13-23).
- `Chunk` struct holds a string that presumably contains a part of the diff (lines 25-27).
- `Config` struct contains configuration settings and has methods like `getFromGPT` for using the OpenAI API, and `SummarizeSummary`, `ChunkToSummary`, and `FileToSummary` for generating different types of summaries (lines 29-108).
- `IndexPage` compiles summaries into a single index page format (lines 110-139).
- `splitDiff` function takes an `io.ReadCloser` and splits the data into diff chunks (lines 141-173).
- `GetDiff` fetches the diff for a PR from GitHub and uses `splitDiff` to parse it (lines 175-194).
- `GetPRDetail` fetches the details of a PR from GitHub and unmarshals them into the `PRDetail` struct (lines 196-219). 

Code snippets include:
- Struct for PR details:
	```go
	type PRDetail struct {
		Number            int    `json:"number"`
		// other fields
	}
	```
	
- Config method that interacts with OpenAI's API:
	```go
	func (cfg *Config) getFromGPT(prompt string) (string, error) {
		// Interaction with OpenAI API
	}
	```
	
- Function to parse diffs:
	```go
	func splitDiff(closer io.ReadCloser) ([]Chunk, error) {
		// Parsing logic
	}
	```