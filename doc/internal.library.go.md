File name: internal/library.go
Language: Go
Purpose: This file provides functionality for interacting with GitHub APIs to fetch pull request details, divide diffs into chunks, and generate summaries using GPT (AI) models.
Important parts:
- Structure definitions for `PRDetail` (lines 16-25) and `Chunk` (lines 27-29), which hold pull request details and content chunks respectively.
- `Config` struct and methods to generate summaries using GPT (lines 31-118), like `getFromGPT` for fetching responses from GPT (lines 33-50), and `FileToSummary` method that formats a file's content into a structured summary (lines 90-109).
- `IndexPage` method to generate a summary page for the repository (lines 110-127).
- Utilizing `splitDiff` function to break up diff outputs into chunks (lines 129-164).
- Network requests are made to the GitHub API to fetch diffs (`GetDiff` at lines 166-190) and pull request details (`GetPRDetail` at lines 192-219).