package web_files_manipulation

import (
	"fmt"
	"os"
	"strings"

	arraylist "github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/array_list"
	"github.com/sunshineplan/node"
	"github.com/yosssi/gohtml"
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
		insertScripts(doc)

		if modifications[file.fileParentDir] != nil {
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
				case "HEAD":
					fallthrough
				case "BODY":
					{
						targetContainer = QuerySelector(doc, strings.ToLower(firstClass))
						InsertBeforeLastChild(f(modification.AppendHTML, IDS[".sizes_items_details select"]), &targetContainer)
					}
				case "tallas_presentations":
					// HARDCODED more like HARDCORE jajjajjajjajajajjajajjajajjajajajja
					{
						targetContainer = QuerySelector(doc, query)
						sizesParent := targetContainer.Parent()
						// THIS COULD INCREMENT IN A FUTURE
						notAirChildren := make([]node.Node, 2)
						idx := 0
						for _, child := range sizesParent.Children() {
							if !IsBlank(child.HTML()) {
								notAirChildren[idx] = child
								idx++
							}
						}
						itemsContainer := notAirChildren[1]
						itemSamples := QuerySelectorAll(itemsContainer, ".ed-element")
						itemFirstSample := itemSamples[0]
						HandleHTMLModifications(modification, itemFirstSample)
						DeleteOtherElementCopies(itemSamples)
						AttributesChanges(modification, itemFirstSample)
						HandleContainerHTMLChanges(modification, itemFirstSample)
					}
				default:
					{
						var targets []node.Node

						if query != "" {
							targets = QuerySelectorAll(doc, query)
						} else {
							targets = []node.Node{doc}
						}

						if len(targets) == 0 {
							panik(" not found query: `%s`", query)
						}
						targetContainer = targets[0]
						StoreID(targetContainer, firstClass)
						HandleHTMLModifications(modification, targetContainer)
						if modification.DeleteSiblings {
							RemoveAllChildrenExceptThis(targetContainer.Parent(), targetContainer)
						}
						DeleteOtherElementCopies(targets)
						AttributesChanges(modification, targetContainer)
						HandleContainerHTMLChanges(modification, targetContainer)
					}
				}
			}
		}
		buildPHPFile(file, doc)
	}

	for i := 0; i < int(directories.Length); i++ {
		directory := directories.ArrayList[i]
		innerDirectories, innerFiles, err := getFilesInDirectory(directory.(string))
		if err != nil {
			// TODO: manage a better error recovery here in c i should free all the memory or just keep going and just log it idrk
			panic(fmt.Sprintf("Failed at directory: %s", directory))
		}
		NewAPITrans(innerDirectories, innerFiles, modifications)
	}
}
