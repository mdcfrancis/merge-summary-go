# Repository https://github.com/mdcfrancis/merge-summary-go
- The repository contains a collection of Go files used to create a command-line application that interacts with GitHub repositories for the purpose of cloning, summarizing file contents, and handling pull request data, utilizing libraries such as `cobra` for CLI structure and potentially OpenAI's GPT models for generating summaries. The main components are commands that can clone a repository, produce file summaries, and generate a `README.md` with indexed summaries. It also includes logic for processing GitHub pull requests and summarizing diffs.

## example.go
- [View file in repository](/example.go)
- [View Doc File](/doc/example.go.md)
- This file contains the main logic for the application.

## cmd/doc.go
- [View file in repository](/cmd/doc.go)
- [View Doc File](/doc/cmd.doc.go.md)
- This file defines a `doc` command using `cobra` for generating summaries of files in a GitHub repository.

## cmd/message.go
- [View file in repository](/cmd/message.go)
- [View Doc File](/doc/cmd.message.go.md)
- This file defines a `message` command for retrieving and summarizing diffs from GitHub pull requests.

## cmd/root.go
- [View file in repository](/cmd/root.go)
- [View Doc File](/doc/cmd.root.go.md)
- The file sets up the base (root) command for the CLI application using the `cobra` library.

## internal/directory.go
- [View file in repository](/internal/directory.go)
- [View Doc File](/doc/internal.directory.go.md)
- This file includes the `FileTools` structure with methods for directory management and file writing.

## internal/library.go
- [View file in repository](/internal/library.go)
- [View Doc File](/doc/internal.library.go.md)
- The file contains the internal library for interfacing with GitHub's API and OpenAI's GPT models, providing functions for pulling request data, processing diffs, and summarizing content.

## main.go
- [View file in repository](/main.go)
- [View Doc File](/doc/main.go.md)
- This file is the entry point of the application, where the main function calls `cmd.Execute()` to initiate the CLI application.