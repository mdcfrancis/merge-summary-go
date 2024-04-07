# cmd/root.go
## Language: Go
## Purpose: 
	This file establishes the command-line interface structure for a tool designed to summarize pull requests on GitHub.
## Important parts: 
	- The `rootCmd` variable (Line 11): defines the base command with its usage, description, and long description.
	```go
	var rootCmd = &cobra.Command{
		Use:   "merge-summary-go",
		Short: "A set of tools to summarise PRs on GitHub",
		Long:  `A set of tools to summarise PRs on GitHub.`,
	}
	```
	- The `Execute` function (Line 18): is the entry point to execute the base command which includes error handling.
	```go
	func Execute() {
		err := rootCmd.Execute()
		if err != nil {
			os.Exit(1)
		}
	}
	```
	- Global variables (Lines 22-24): store command-line flag values such as the repository owner, repository name, and GPT API key.
	- The `init` function (Line 26): initializes the command-line flags with default values and specifies how to parse them.
	```go
	func init() {
		rootCmd.Flags().StringVarP(&repoOwner, "owner", "o", "mdcfrancis", "The owner of the repository")
		rootCmd.Flags().StringVarP(&repoName, "repo", "r", "merge-summary-go", "The name of the repository")
		rootCmd.Flags().StringVarP(&gptAuth, "gpt-auth", "a", os.Getenv("GPT_AUTH"), "The GPT API key (not recommended, use environment")
	}
	```