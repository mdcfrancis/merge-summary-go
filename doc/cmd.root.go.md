# cmd/root.go
## Language: Go
## Purpose:
This file defines the base command for a command-line application that provides tools for summarizing Pull Requests (PRs) on GitHub.

## Important parts:
- The package declaration and imports indicate the file is part of the `cmd` package and uses the `cobra` library and `os` package for command-line interactions and operating system functionality, respectively:

    ```go
    package cmd

    import (
    	"os"

    	"github.com/spf13/cobra"
    )
    ```

- The `rootCmd` is a global variable defined which represents the base command when no subcommands are called. It includes `Use`, `Short`, and `Long` descriptions:

    ```go
    var rootCmd = &cobra.Command{
        Use:   "merge-summary-go",
        Short: "A set of tools to summarise PRs on GitHub",
        Long:  `A set of tools to summarise PRs on GitHub.`,
    }
    ```

- The `Execute` function is responsible for executing the `rootCmd`. It handles errors by exiting the program with a non-zero status code to indicate failure:

    ```go
    func Execute() {
        err := rootCmd.Execute()
        if err != nil {
            os.Exit(1)
        }
    }
    ```

- There are three global variables declared (`repoOwner`, `repoName`, and `gptAuth`) which store information about the repository owner, repository name, and an API key for GPT (which can also be provided via an environment variable):

    ```go
    var (
        repoOwner string
        repoName  string
        gptAuth   string
    )
    ```

- The `init` function initializes the command flags for the `rootCmd` allowing users to specify the repository owner, name, and GPT API key using flags or an environment variable:

    ```go
    func init() {
        rootCmd.Flags().StringVarP(&repoOwner, "owner", "o", "mdcfrancis", "The owner of the repository")
        rootCmd.Flags().StringVarP(&repoName, "repo", "r", "merge-summary-go", "The name of the repository")
        rootCmd.Flags().StringVarP(&gptAuth, "gpt-auth", "a", os.Getenv("GPT_AUTH"), "The GPT API key (not recommended, use environment")
    }
    ```