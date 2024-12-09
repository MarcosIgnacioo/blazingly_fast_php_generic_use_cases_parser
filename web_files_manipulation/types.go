package web_files_manipulation

import "fmt"

type readingFilesInDirectoryError struct {
	fileName string
}

func (self *readingFilesInDirectoryError) Error() string {
	return fmt.Sprintf("error reading directory %s", self.fileName)
}

const (
	default_array_size = 12
)
