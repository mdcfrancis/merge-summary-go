File name: cmd/root.go
Language: Go
Purpose: This file defines the base command for a command-line application and sets up command-line flags for configuration.
Important parts: The `rootCmd` variable on line 9 declares the base command and its descriptions. The `Execute` function on line 19 is the entry point that executes the base command with error handling. Command-line flags are initialized in the `init` function from line 31 to 34, allowing users to specify repository owner, repository name, and GPT API key.