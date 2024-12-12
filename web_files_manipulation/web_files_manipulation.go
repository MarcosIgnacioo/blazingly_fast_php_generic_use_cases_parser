package web_files_manipulation

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	arraylist "github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/array_list"
	"github.com/sunshineplan/node"
	"github.com/yosssi/gohtml"
	"golang.org/x/net/html"
)

func Init(sourceDirectory string) {
	directories, files, err := getFilesInDirectory(sourceDirectory)
	if err != nil {
		panic(fmt.Sprintf("Failed at initial directory %s", sourceDirectory))
	}
	contentTransformation(directories, files)
}

func InsertHeader() {
}

func contentTransformation(directories *arraylist.ArrayList, files *arraylist.ArrayList) {
	// html a php

	for i := 0; i < int(files.Length); i++ {
		file := files.ArrayList[i].(*File)
		if instructions[file.fileParentDir] != nil {
			fmt.Println(file.filePath)
			fileReaded, err := os.ReadFile(file.filePath)
			fileContent := gohtml.Format(string(fileReaded))
			if err != nil {
				panic(fmt.Sprintf("tried to open this file but for some reason crashed %s %s", file.filePath))
			}

			for _, instruction := range instructions[file.fileParentDir] {
				classToSearch := instruction["class"]
				htmlToInsert := instruction["htmlToInsert"]
				// switch classToSearch {
				// case "head":
				// }
				fmt.Println(fileContent)
				fmt.Println("class", classToSearch)
				fmt.Println("html", htmlToInsert)
			}
		}
	}

	for i := 0; i < int(directories.Length); i++ {
		directory := directories.ArrayList[i]
		innerDirectories, innerFiles, err := getFilesInDirectory(directory.(string))
		if err != nil {
			// TODO: manage a better error recovery here in c i should free all the memory or just keep going and just log it idrk
			panic(fmt.Sprintf("Failed at directory: %s", directory))
		}
		contentTransformation(innerDirectories, innerFiles)
	}
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
		nestedLevel := (len(strings.Split(fullFilePath, "/"))/2 - 1)
		parentDirectory := filepath.Base(directory)
		newFile := &File{fileName: fileName, filePath: fullFilePath, fileType: fileTypesMap[fileExtension], nestedLevel: nestedLevel, fileParentDir: parentDirectory}
		files.Enqueue(newFile)
	}
}

func insertRawTextBeforeNode(text string, beforeNode node.Node) {
	textNodeToInsert := &html.Node{
		Parent:      nil,
		PrevSibling: nil,
		NextSibling: nil,
		Data:        text,
		Type:        html.RawNode, // XD
		Attr:        []html.Attribute{},
	}
	beforeNode.Raw().InsertBefore(textNodeToInsert, beforeNode.Raw())
}

// func scannDirectory() error {
// 	// scannDirectory(directories[i])
// }

// func HTMLToPHP(FILE* file) void {
// }
