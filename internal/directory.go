package internal

import (
	"fmt"
	"os"
)

type FileTools struct {
	Directory string
}

// CheckOrCreateDirectory if a directory exists, if not create it (recursive if required)
// directory: the directory to check or create
func (f FileTools) CheckOrCreateDirectory() error {
	if f.Directory == "" {
		return fmt.Errorf("directory is empty, you must provide a directory")
	}
	return os.MkdirAll(f.Directory, os.ModePerm)
}

// WriteFile writes a file to a directory
// fileName: the name of the file to write
// summary: the content of the file
func (f FileTools) WriteFile(fileName, summary string) error {
	return os.WriteFile(f.Directory+"/"+fileName, []byte(summary), 0644)
}
