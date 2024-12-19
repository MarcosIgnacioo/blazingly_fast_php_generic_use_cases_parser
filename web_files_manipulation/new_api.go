package web_files_manipulation

import (
	"fmt"
	arraylist "github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/array_list"
	"github.com/sunshineplan/node"
	"github.com/yosssi/gohtml"
	"os"
)

func InitNewAPI(sourceDirectory string) {
	directories, files, err := getFilesInDirectory(sourceDirectory)
	if err != nil {
		panic(fmt.Sprintf("Failed at initial directory %s", sourceDirectory))
	}
	NewAPITrans(directories, files, modifications)
}

// probablemente esta funcion pase a ser unicamente la que mete los productos en las tiendas, las otras cosas que se ocupan hacer son mas especificas y complicadas que encontrar una manera generica de hacerlo es algo que requiere bastante mas tiempo
func NewAPITrans(directories *arraylist.ArrayList, files *arraylist.ArrayList, modifications map[string][]Modification) {
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
			for _, modification := range modifications[file.fileParentDir] {
				query := modification.Query
				firstClass := parseClass(query)

				var targetContainer node.Node
				// var parent node.Node

				switch firstClass {
				case "html":
					targetContainer = doc.Find(node.Descendant, node.Html)
					// we PREPEND the html in this case instead of replacing the innerhtml
					if modification.PrependHTML == "" {
						panic(fmt.Sprintf("preppending html for header doesnt exist in file %", file.filePath))
					}
					preppendHTMLToNode(modification.PrependHTML, targetContainer)
				default:
					{
						// queries := parseQuery(query)
					}
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
