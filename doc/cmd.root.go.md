# cmd/root.go
## Language: Go
## Purpose: 
	The file defines the base command for a command-line application that offers tools for summarizing pull requests on GitHub.
## Important parts: 
	- The base command (rootCmd) is defined using the cobra library (lines 11-15).
	  ```go
	  var rootCmd = &cobra.Command{...}
	  ```
	- The Execute function, which handles the execution of the root command and any child commands, is essential for the application's operation (lines 17-23).
	  ```go
	  func Execute() {...}
	  ```
	- Global variables (repoOwner, repoName, gptAuth) are declared to store flags for command-line options (lines 25-27).
	- Initialization function (init) that binds command-line flags to the aforementioned global variables (lines 29-35).
	  ```go
	  func init() {...}
	  ```