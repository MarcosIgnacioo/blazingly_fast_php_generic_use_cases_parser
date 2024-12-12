package web_files_manipulation

import "fmt"

type readingFilesInDirectoryError struct {
	fileName string
}

type FileType string

var fileTypesMap = map[string]FileType{
	".html": "html",
	".js":   "js",
	".css":  "css",
	".php":  "php",
	".jpg":  "img",
	".png":  "img",
	".svg":  "img",
}

const (
	HTML  FileType = "html"
	JS             = "js"
	CSS            = "css"
	PHP            = "php"
	IMAGE          = "img"
	FONT           = "font"
)

type File struct {
	fileName      string
	filePath      string
	fileParentDir string
	fileType      FileType
	nestedLevel   int
}

func (f *File) String() string {
	return fmt.Sprintf("file_name:\t%s\nfile_path:\t%s\nfile_type:\t%v\nnested_level:\t%d\nparent_dir:\t%s\n", f.fileName, f.filePath, f.fileType, f.nestedLevel, f.fileParentDir)
}

func (self *readingFilesInDirectoryError) Error() string {
	return fmt.Sprintf("error reading directory %s", self.fileName)
}

const (
	default_array_size = 12
)
