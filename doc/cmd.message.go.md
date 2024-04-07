# cmd/message.go
## Language: Go
## Purpose: 
	This file defines a command-line functionality to fetch details from a pull request (PR) and produce a summarized message.
## Important parts: 
	- The variable `messageCmd` (line 13) of type `*cobra.Command` represents the message command and is configured with usage, short and long descriptions.
		```go
		var messageCmd = &cobra.Command{
			// Command configuration...
		}
		```
	- The `Run` function (line 20) of `messageCmd` performs the core logic to process the PR details and generate summaries.
		```go
		Run: func(cmd *cobra.Command, args []string) {
			// Core logic for processing and summarizing PR details...
		},
		```
	- The `init` function (line 55) initializes the command, adding it to the root command and setting up a flag for qualitative summarization.
		```go
		func init() {
			rootCmd.AddCommand(messageCmd)
			// ...
		}
		```