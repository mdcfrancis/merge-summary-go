# internal/directory.go
## Language: Go
## Purpose: 
	The file defines a struct and methods for handling directory-related operations such as checking for the existence of a directory and creating it if necessary, as well as writing files to a directory.

## Important parts:
- The `FileTools` struct (line 8) holds information regarding the directory that will be operated on.
	```go
	type FileTools struct {
		Directory string
	}
	```
- The `CheckOrCreateDirectory` method (line 14) checks whether the Directory field of FileTools is empty, returns an error if it is, or attempts to create the directory including all necessary parent directories with the appropriate permissions.
	```go
	func (f FileTools) CheckOrCreateDirectory() error {
		...
	}
	```
- The `WriteFile` method (line 21) is designed to write text content to a file with the specified `fileName` within the `Directory` specified by the FileTools struct. It handles the creation of the file and ensures the content is written with the appropriate file permissions.
	```go
	func (f FileTools) WriteFile(fileName, summary string) error {
		...
	}
	```