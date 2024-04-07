# internal/library.go
## Language: Go
## Purpose: 
	The file serves as a library for processing GitHub Pull Requests (PRs) and summarizing them using an AI model.
## Important parts: 
	- Definition of types `PRDetail` (Line 14), `Chunk` (Line 28), and `Config` (Line 31).
	- Function `getFromGPT` (Line 38): Sends a prompt to an AI model and gets a response.
		```go
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				// ...
			},
		)
		```
	- Function `SummarizeSummary` (Line 55): Formats a summary using the AI model based on a given prompt.
	- Function `ChunkToSummary` (Line 69): Turns a given chunk of text into a summary.
	- Function `FileToSummary` (Line 86): Outputs formatted markdown for a file summary from the AI model.
	- Function `IndexPage` (Line 109): Creates an index page summary with a list of file summaries.
	- Function `splitDiff` (Line 135): Splits the diff from a pull request body into chunks.
		```go
		if strings.HasPrefix(line, "diff --git") {
			// ...
		}
		```
	- Function `GetDiff` (Line 165): Retrieves the diff from a GitHub PR.
		```go
		req, err := http.NewRequest("GET", "https://api.github.com/repos/"+owner+"/"+repo+"/pulls/"+pr, nil)
		```
	- Function `GetPRDetail` (Line 185): Fetches details of a GitHub PR.
		```go
		err = json.Unmarshal(body, &prDetail)
		```