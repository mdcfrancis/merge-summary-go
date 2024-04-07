# cmd/message.go
## Language: Go
## Purpose:
This file defines a command-line command named "message" that retrieves changes from a pull request (PR) and produces a summary of the diffs.

## Important parts:
- Importing dependencies (lines 6-12):
  ```go
  import (
      "com.github/mdcfrancis/merge-summary-go/internal"
      "fmt"
      "github.com/spf13/cobra"
      "log"
      "strings"
  )
  ```
- Declaration of `messageCmd` using the `cobra` library (lines 14-48) sets up the "message" command with usage, description, arguments requirement, and the function to run when the command is executed. It fetches PR details and diffs, and consolidates them into a summary.
  ```go
  var messageCmd = &cobra.Command{...}
  ```
- Flag declaration for command `messageCmd` to include the option for qualitative summarization (lines 50-52).
  ```go
  messageCmd.Flags().BoolVarP(&qualitative, "qualitative", "q", false, "Use qualitative summarization")
  ```
- Registration of the new command to the root command within the `init` function, making `messageCmd` available for the command-line interface (lines 54-56).
  ```go
  rootCmd.AddCommand(messageCmd)
  ```