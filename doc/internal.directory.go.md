# internal/directory.go
## Language: Go
## Purpose:
This file defines a Go structure `FileTools` with methods to manage file system directories, specifically checking for the existence of directories and creating them if necessary, as well as writing files.

## Important parts:
- The `FileTools` struct holds a directory path as its field (line 8).

```go
type FileTools struct {
	Directory string
}
```

- The method `CheckOrCreateDirectory` checks if the directory set in `FileTools` exists and creates it if it does not, making directories recursively as required (lines 13-18).

```go
func (f FileTools) CheckOrCreateDirectory() error {
	// ...
}
```

- The method `WriteFile` is responsible for writing content to a file within the specified directory (lines 20-25).

```go
func (f FileTools) WriteFile(fileName, summary string) error {
	// ...
}
```