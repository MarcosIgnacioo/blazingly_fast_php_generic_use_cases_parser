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
)

func Init(sourceDirectory string) {
	directories, files, err := getFilesInDirectory(sourceDirectory)
	if err != nil {
		panic(fmt.Sprintf("Failed at initial directory %s", sourceDirectory))
	}
	contentTransformation3D(directories, files)
}

func InsertHeader() {
}

// probablemente esta funcion pase a ser unicamente la que mete los productos en las tiendas, las otras cosas que se ocupan hacer son mas especificas y complicadas que encontrar una manera generica de hacerlo es algo que requiere bastante mas tiempo
func contentTransformation3D(directories *arraylist.ArrayList, files *arraylist.ArrayList) {
	for i := 0; i < int(files.Length); i++ {
		file := files.ArrayList[i].(*File)
		fileReaded, err := os.ReadFile(file.filePath)
		fileContent := gohtml.Format(string(fileReaded))
		doc, err := node.ParseHTML(fileContent)
		insertCartMobile(doc, file.nestedLevel)
		if instructionsTYPED[file.fileParentDir] != nil {
			if err != nil {
				panic(fmt.Sprintf("tried to open this file but for some reason crashed %s %s", file.filePath))
			}
			if err != nil {
				panic(fmt.Sprintf("error parsing this file", file.filePath))
			}
			for _, instruction := range instructionsTYPED[file.fileParentDir] {
				classToSearch := instruction.Class

				var targetDiv node.Node
				var parent node.Node

				switch classToSearch {

				case "html":
					targetDiv = doc.Find(node.Descendant, node.Html)
					// we PREPEND the html in this case instead of replacing the innerhtml
					if instruction.PrependHTML == "" {
						panic(fmt.Sprintf("preppending html for header doesnt exist in file %", file.filePath))
					}
					preppendHTMLToNode(instruction.PrependHTML, targetDiv)

				default:
					queries := parseQuery(classToSearch)
					var targetDiv node.Node
					if classToSearch == "" {
						targetDiv = doc
					} else {
						targetDiv = QuerySelector(doc, classToSearch)
					}
					if instruction.TargetParent {
						targetDiv = targetDiv.Parent()
					}

					firstQuery := getClassNameOrTagName(queries[0])
					IDS[firstQuery], _ = targetDiv.Attrs().Get("id")
					// change this at the end
					// XDDDD
					if firstQuery == "body" || firstQuery == "head" {
						instruction.TargetItemID = IDS["sizes_items_details"]
						instruction.AppendHTML = fmt.Sprintf(instruction.AppendHTML, instruction.TargetItemID)
					}

					// targetDiv = doc.Find(node.Descendant, nil, node.Class(classToSearch))
					if targetDiv == nil {
						panic(fmt.Sprintf(" targetDiv is nil probably wrong class for the wrong directory\nthis is the class we are trying to look for: `%s`\n inside file: `%s`", classToSearch, file.filePath))
					}

					HTMLManipulation(instruction, targetDiv)

					if instruction.TagsAttributes != nil {
						attributesManipulation(instruction, classToSearch, targetDiv)
					}

					if instruction.InnerHtmlReplacements != nil {
						innerHTMLManipulation(instruction, classToSearch, targetDiv)
					}
					constructHTML(&parent, targetDiv, instruction, firstQuery)
				}

				buildPHPFile(file, doc)
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
		contentTransformation3D(innerDirectories, innerFiles)
	}
}

func prepareHTMLForFile(doc node.Node) []byte {
	html := (doc.HTML())
	// var greaterThan = regexp.MustCompile(`(\?="*.)(&gt;)(?=.*")`)
	// var lessThan = regexp.MustCompile(`(\?="*.)(&lt;)(?=.*")`)
	// fixedGt := greaterThan.ReplaceAll([]byte(html), []byte(">"))
	// finalHtml := lessThan.ReplaceAll(fixedGt, []byte("<"))
	coolerHtml := strings.Replace(strings.Replace(strings.Replace(html, "&lt;", "<", -1), "&gt;", ">", -1), "&#39;", "'", -1)
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
