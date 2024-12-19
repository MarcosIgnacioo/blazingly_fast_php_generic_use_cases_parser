package web_files_manipulation

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	arraylist "github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/array_list"
	"github.com/sunshineplan/node"
	// "github.com/yosssi/gohtml"
)

func prepareHTMLForFile(doc node.Node) []byte {
	html := (doc.HTML())
	// dangerous
	// html = gohtml.Format(html)
	// var greaterThan = regexp.MustCompile(`(\?="*.)(&gt;)(?=.*")`)
	// var lessThan = regexp.MustCompile(`(\?="*.)(&lt;)(?=.*")`)
	// fixedGt := greaterThan.ReplaceAll([]byte(html), []byte(">"))
	// finalHtml := lessThan.ReplaceAll(fixedGt, []byte("<"))
	coolerHtml := strings.ReplaceAll(strings.Replace(strings.Replace(strings.Replace(html, "&lt;", "<", -1), "&gt;", ">", -1), "&#39;", "'", -1), "&#34;", "'")
	return []byte(coolerHtml)
}

func getFilesInDirectory(directory string) (directories *arraylist.ArrayList, files *arraylist.ArrayList, err error) {
	directories = arraylist.NewArrayList(default_array_size)
	files = arraylist.NewArrayList(default_array_size)

	entries, err := os.ReadDir(directory)

	if err != nil {
		return
	}

	for _, e := range entries {
		if e.IsDir() {
			directories.Enqueue(fmt.Sprintf("%s/%s", directory, e.Name()))
		} else {
			storeFilesInArray(e, directory, files)
		}
	}
	return
}

func storeFilesInArray(entry fs.DirEntry, directory string, files *arraylist.ArrayList) {
	fileExtension := filepath.Ext(entry.Name())
	switch fileExtension {
	case ".html":
		fileName := entry.Name()
		fullFilePath := fmt.Sprintf("%s/%s", directory, fileName)
		fileExtension = filepath.Ext(fileName)
		nestedLevel := (len(strings.Split(fullFilePath, "/")) / 2)
		parentDirectory := filepath.Base(directory)
		newFile := &File{fileName: fileName, filePath: fullFilePath, fileType: fileTypesMap[fileExtension], nestedLevel: nestedLevel, fileParentDir: parentDirectory}
		files.Enqueue(newFile)
	}
}
