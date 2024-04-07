# internal/library.go
## Language: Go
## Purpose: 
	This file is designed as a utility library to interact with GitHub data, particularly pull requests (PRs), and to generate summaries using the OpenAI service.

## Important parts: 
- Struct definitions for `PRDetail` (Line 14) and `Chunk` (Line 28) are used for storing PR details and code blocks respectively:
	```go
	type PRDetail struct {
	  // struct fields
	}

	type Chunk struct {
	  Content string
	}
	```

- Struct `Config` with methods to interact with the OpenAI service and GitHub:
	```go
	type Config struct {
	  // config fields
	}

	func (cfg *Config) getFromGPT(prompt string) (string, error) {
	  // interacts with OpenAI
	}

	func (cfg *Config) SummarizeSummary(summary string) (string, error) {
	  // uses OpenAI to summarize a summary
	}

	func (cfg *Config) ChunkToSummary(chunk Chunk) (string, error) {
	  // uses OpenAI to summarize chunks
	}

	func (cfg *Config) FileToSummary(chunk Chunk) (string, error) {
	  // uses OpenAI to summarize a file
	}

	func (cfg *Config) IndexPage(title string, summaries []string) (string, error) {
	  // uses OpenAI to generate an index page summary
	}
	```

- `splitDiff` function that parses diff output and splits it into chunks (Line 144):
	```go
	func splitDiff(closer io.ReadCloser) ([]Chunk, error) {
	  // code for splitting diffs
	}
	```

- HTTP requests for getting the difference of a PR and the details of a PR (`GetDiff` on Line 168 and `GetPRDetail` on Line 194):
	```go
	func (cfg *Config) GetDiff(owner string, repo string, pr string) ([]Chunk, error) {
	  // code to get PR diff
	}

	func (cfg *Config) GetPRDetail(owner string, repo string, pr string) (PRDetail, error) {
	  // code to get PR details
	}
	```