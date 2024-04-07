File Name: internal/directory.go
Language: Go
Purpose: This file provides utilities for directory operations, specifically checking for the existence of a directory and creating it if necessary, as well as writing files to a directory.
Important Parts: 
- The `FileTools` struct is defined to handle directory operations with the `Directory` attribute (line 7).
- The `CheckOrCreateDirectory` method checks for the existence of a directory and creates it if it doesn't exist, generating an error if the directory field is empty (lines 12-20).
- The `WriteFile` method is responsible for writing a file with given content to a directory, constructing a file path by concatenating `Directory` and `fileName` (lines 22-27).