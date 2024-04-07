# cmd/root.go
## Language: Go
## Purpose: 
	This file serves as the main command setup for a CLI application, providing a framework to add subcommands related to summarizing PRs on GitHub.
## Important parts: 
- Declaration of the base command:
	The variable `rootCmd` is initiated as a `cobra.Command` struct which forms the base for the application's command-line interface. (Lines 10-13)

	```go
	var rootCmd = &cobra.Command{
		Use:   "merge-summary-go",
		Short: "A set of tools to summarise PRs on GitHub",
		Long:  `A set of tools to summarise PRs on GitHub.`,
	}
	```

- Execution function:
	The `Execute` function, which is called by the package's `main` function, is responsible for invoking the command-line interface and handling any errors that occur during its execution. (Lines 17-23)

	```go
	func Execute() {
		err := rootCmd.Execute()
		if err != nil {
			os.Exit(1)
		}
	}
	```

- Command flag initialization:
	In the `init` function, the code defines flags for the `rootCmd` to accept command-line arguments for the repository owner, repository name, and GPT API key. These inputs default to specified values or the `GPT_AUTH` environment variable. (Lines 25-39)

	```go
	func init() {
		rootCmd.Flags().StringVarP(&repoOwner, "owner", "o", "mdcfrancis", "The owner of the repository")
		rootCmd.Flags().StringVarP(&repoName, "repo", "r", "merge-summary-go", "The name of the repository")
		rootCmd.Flags().StringVarP(&gptAuth, "gpt-auth", "a", os.Getenv("GPT_AUTH"), "The GPT API key (not recommended, use environment")
	}
	```