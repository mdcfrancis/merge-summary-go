# Repository [merge-summary-go](https://github.com/mdcfrancis/merge-summary-go)
- This repository contains an application developed in Go, designed to provide command-line interface tools for documenting, reviewing, and summarizing pull requests on GitHub. It leverages tools such as Cobra for command-line structure and OpenAI's API for summarization tasks.

## cmd/doc.go
- [View file in repository](../cmd/doc.go)
- [View Doc File](../doc/cmd.doc.go.md)
- This file implements the 'doc' command to generate summary documents for GitHub repos, with functionality for cloning repositories, filtering files, processing content, and writing summaries.

## cmd/message.go
- [View file in repository](../cmd/message.go)
- [View Doc File](../doc/cmd.message.go.md)
- This file defines a command-line command to retrieve and provide summaries of GitHub Pull Requests, using Cobra commands and various parsing and summarization functions.

## cmd/root.go
- [View file in repository](../cmd/root.go)
- [View Doc File](../doc/cmd.root.go.md)
- This file sets up the command-line interface for the tool and is central in defining the base 'merge-summary-go' command and parsing command-line flags.

## internal/directory.go
- [View file in repository](../internal/directory.go)
- [View Doc File](../doc/internal.directory.go.md)
- The file focuses on directory operations such as checking directory existence, creating directories, and writing files to the specified directory.

## internal/library.go
- [View file in repository](../internal/library.go)
- [View Doc File](../doc/internal.library.go.md)
- This file provides a collection of functions for interacting with GitHub APIs, parsing PR diffs, and communicating with the OpenAI API to produce summaries for PRs and code changes.

## main.go
- [View file in repository](../main.go)
- [View Doc File](../doc/main.go.md)
- The file is the entry point for the Go application, initiating the application's command-line interface execution process.