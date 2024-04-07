File name: cmd/doc.go
Language: Go
Purpose: This file defines a command-line command to generate a summary document for a specified GitHub repository, including all of its non-excluded files.
Important parts: 
- `docCmd` is a Cobra command with `Run` function that clones a repository, processes its files for generating summaries, and creates an index page as README.md (lines 27-89).
```go
var docCmd = &cobra.Command{
	// ...
	Run: func(cmd *cobra.Command, args []string) {
		// ...
	},
}
```
- Repository cloning uses the `git.Clone` function from the `go-git` package, specifying an in-memory storage (lines 32-37).
- Summary creation excludes specific files and files in dot directories (lines 50-81), and for each file included, it generates a summary, formats, and writes it to the specified output directory.
- The `init` function adds the `docCmd` to the root command and sets up the `outputDirectory` flag, making it a required flag for the command (lines 91-95).
```go
func init() {
	rootCmd.AddCommand(docCmd)
	// ...
}
```