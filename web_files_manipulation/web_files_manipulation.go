package web_files_manipulation

import (
	"fmt"
	arraylist "github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/array_list"
	"github.com/sunshineplan/node"
	"github.com/yosssi/gohtml"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func p(a ...any) {
	fmt.Println(a)
}

func Init(sourceDirectory string) {
	directories, files, err := getFilesInDirectory(sourceDirectory)
	if err != nil {
		panic(fmt.Sprintf("Failed at initial directory %s", sourceDirectory))
	}
	contentTransformation(directories, files)
}

func InsertHeader() {
}

// probablemente esta funcion pase a ser unicamente la que mete los productos en las tiendas, las otras cosas que se ocupan hacer son mas especificas y complicadas que encontrar una manera generica de hacerlo es algo que requiere bastante mas tiempo
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
			doc, err := node.ParseHTML(fileContent)
			if err != nil {
				panic(fmt.Sprintf("error parsing this file", file.filePath))
			}
			for _, instruction := range instructions[file.fileParentDir] {
				classToSearch := instruction["class"]
				htmlToInsert := instruction["htmlToInsert"]

				var targetDiv node.Node
				var deletingDivs []node.Node
				var parent node.Node
				// var htmlContent string

				switch classToSearch {
				case "html":
					targetDiv = doc.Find(node.Descendant, node.Html)
					insertRawTextBeforeNode(htmlToInsert, targetDiv)
				// poner que en vez de que el realizar el reemplazo de productos sea el default que sea la cosa que reemplaza la default nomas que no se como hacerle bien porque
				default:
					productHref := getData(instruction, "href")
					productName := getData(instruction, "productName")
					productImg := getData(instruction, "img")
					productPrice := getData(instruction, "price")
					productForeach := getData(instruction, "foreach") // ocd tockin rn fr
					productClassForName := getData(instruction, "className")
					productClassForPrice := getData(instruction, "classPrice")

					deletingDivs = doc.FindAll(node.Descendant, node.Div, node.Class(classToSearch))
					if deletingDivs == nil || len(deletingDivs) == 0 {
						panic(fmt.Sprintf("deletingDivs are nil or zero len in this instruction %s", classToSearch))
					}
					targetDiv = deletingDivs[0]
					// targetDiv.Find(node.Descendant, node.A, node.Class("product_name"))
					anchors := targetDiv.FindAll(node.Descendant, node.A)
					if anchors == nil || len(anchors) == 0 {
						panic(fmt.Sprintf("anchors are nil or zero len in this instruction %s", classToSearch))
					}
					imgs := targetDiv.FindAll(node.Descendant, node.Img)
					if imgs == nil || len(imgs) == 0 {
						panic(fmt.Sprintf("imgs are nil or zero len in this instruction %s", classToSearch))
					}
					setUpAnchors(&anchors, productHref, productName)
					setUpImages(&imgs, productImg, productName)

					// bruhhhh
					priceSpan := targetDiv.Find(node.Descendant, node.Span, node.Class(productClassForPrice))
					if priceSpan == nil {
						panic(fmt.Sprintf("priceSpan is nil, check the value of the `classPrice` field in this instruction %s", classToSearch))
					}
					removeAllChildren(&priceSpan)
					priceSpan.Raw().AppendChild(newTextHtmlNode(fmt.Sprintf("DESDE $%s", productPrice)))

					nameAnchor := targetDiv.Find(node.Descendant, nil, node.Class(productClassForName))
					if nameAnchor == nil {
						panic(fmt.Sprintf("nameAnchor is nil, check the value of the `className` field in this instruction %s", classToSearch))
					}

					removeAllChildren(&nameAnchor)
					nameAnchor.Raw().AppendChild(newTextHtmlNode(productName))

					parent = targetDiv.Parent()
					htmlToInsert = fmt.Sprintf(productForeach, targetDiv.HTML())
					removeAllChildren(&parent)
					parent.Raw().AppendChild(newTextHtmlNode(htmlToInsert))
				}

				// parent.Raw().AppendChild(newTextHtmlNode(htmlToInsert))

				os.WriteFile(file.filePath, prepareHTMLForFile(doc), 0777)
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

func prepareHTMLForFile(doc node.Node) []byte {
	html := (doc.HTML())
	coolerHtml := strings.Replace(strings.Replace(html, "&lt;", "<", -1), "&gt;", ">", -1)
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
		nestedLevel := (len(strings.Split(fullFilePath, "/"))/2 - 1)
		parentDirectory := filepath.Base(directory)
		newFile := &File{fileName: fileName, filePath: fullFilePath, fileType: fileTypesMap[fileExtension], nestedLevel: nestedLevel, fileParentDir: parentDirectory}
		files.Enqueue(newFile)
	}
}
