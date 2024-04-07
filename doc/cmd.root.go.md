# cmd/root.go
## Language: Go
## Purpose: 
	This file implements the command-line interface for a tool that summarizes pull requests (PRs) on GitHub.
## Important parts: 
- The `rootCmd` variable is a Cobra command which acts as the base for the CLI (line 11):
	```go
	var rootCmd = &cobra.Command{...}
	```
- The `Execute` function prepares and executes the root command, handling any errors (line 22):
	```go
	func Execute() { ... }
	```
- Global variables `repoOwner`, `repoName`, and `gptAuth` are declared to store command line flag values (line 30).
- The `init` function configures `rootCmd` flags for repository owner, name, and GPT API authentication (line 35):
	```go
	func init() { ... }
	```