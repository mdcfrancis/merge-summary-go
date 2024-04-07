# internal/directory.go
## Language: Go
## Purpose:
The purpose of this Go file is to provide functionalities for working with directories, such as ensuring a directory exists or creating it if necessary, and writing files to a specified directory.

## Important parts:
- Struct Definition: Defines a `FileTools` struct with a `Directory` field to store the path of a directory.
  ```go
  type FileTools struct {
      Directory string
  }
  ```

- CheckOrCreateDirectory Method (Line 13): This method checks if the `Directory` exists, and if not, it creates the directory with proper file permissions. This function returns an error if the `Directory` field is empty.
  ```go
  func (f FileTools) CheckOrCreateDirectory() error {
      // ... implementation ...
  }
  ```

- WriteFile method (Line 22): This method takes a file name and content (`summary`) as parameters and writes the file to the `Directory` specified in the `FileTools` struct.
  ```go
  func (f FileTools) WriteFile(fileName, summary string) error {
      // ... implementation ...
  }
  ```