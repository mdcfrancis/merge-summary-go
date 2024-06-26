# merge-summary-go
Create a summary for a pull request based on the diffs in the pull request.

## Usage
```bash
merge-summary-go --help

A set of tools to summarise PRs on GitHub.

Usage:
  merge-summary-go [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  doc         Given a owner and repo name generate a summary document of the repo, use --output to specify the output directory
  help        Help about any command
  message     Gets the diffs from a PR and summarises them into a message

Flags:
  -a, --gpt-auth string   The GPT API key (not recommended, use environment (default from environment GPT_AUTH) 
  -h, --help              help for merge-summary-go
  -o, --owner string      The owner of the repository (default "mdcfrancis")
  -r, --repo string       The name of the repository (default "merge-summary-go")

Use "merge-summary-go [command] --help" for more information about a command.

```
# Copyright for examples remains with the original authors
## Example MD for mdcfrancis/merge-summary-go pull 1
# File: `merger.go` - Modifications

Several updates were made to `merger.go` largely aimed at enhancing summarization of changes.

## Function: 'chunkToSummary'

The 'chunkToSummary' function underwent a few alterations. Programmer comments were removed and the presentation of the 'prompt' variable has been changed.

## New Functions: 'getFromGPT' and 'summarizeSummary'

Two new functions were incorporated, 'getFromGPT' and 'summarizeSummary'. 'getFromGPT' takes over fetching data from the GPT API, which was previously a part of the 'chunkToSummary' function. 'SummarizeSummary', on the other hand, builds a 'prompt' variable by summarizing an existing summary, with the aim to offer a GPT API response later on.

## Modification: Prompt setup change in 'chunkToSummary'

The prompt setup in the 'chunkToSummary' function was updated to optionally allow a quality analysis of the amendments.

## Script-Level Changes

At the script's top level, a new 'qualitative' flag has been added that dictates whether a qualitative summary is provided or not. The 'body' variable's display in the main function is now commented, and the 'chunkToSummary' function has been superseded by 'summarizeSummary'.

# Quality Analysis

The updates made to the `merger.go` file have been intended to improve the processes of fetching data and summarizing changes. With the cleaned coder comments and new functions, the file is expected to perform more effectively. The quality analysis inclusion in the summary provides deeper insight into the changes. However, the functionality and impact of commenting out the 'body' variable display are yet to be evaluated.
