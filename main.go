package main

import (
	"archive/zip"
	"fmt"
	"github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/web_files_manipulation"
	"github.com/sunshineplan/node"
	"github.com/yosssi/gohtml"
	"golang.org/x/net/html"
	"io"
	"log"
	"os"
)

var testhtml = `
<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title></title>
    <link href="css/style.css" rel="stylesheet" />
  </head>
  <body>
    <p>
      <span><a href="">popipip</a></span>
    </p>
    <p>
      <span><a href="">iopopo</a></span>
    </p>
    <p>
      <span><a href="">iopopo</a></span>
    </p>
  </body>
</html>
	`

func p(a ...any) {
	fmt.Println(a)
}

// possible improvement make this function return true or false if it can change it or not
func setAttribute(htmlNode *node.Node, attribute string, value string) {
	attributes := (*htmlNode).Raw().Attr
	// c goat
	for i := 0; i < len(attributes); i++ {
		attr := &attributes[i]
		if attr.Key == attribute {
			attr.Val = value
			break
		}
	}
}

func main() {
	// f, _ := os.ReadFile("./asdf.html")
	// doc, _ := node.ParseHTML(string(f))
	// fmt.Println(string(f))
	// spans := doc.FindAll(node.Descendant, node.P)
	// for _, v := range spans {
	// 	fmt.Println("###########")
	// 	fmt.Println(v.HTML())
	// }

	// popo := web_files_manipulation.QuerySelectorAll(doc, "p span a")
	// foo := `asdf %d asdf`
	// fmt.Printf(foo, 123)
	run()
	// return
	// textNodeToInsert := &html.Node{
	// 	Parent:      nil,
	// 	PrevSibling: nil,
	// 	NextSibling: nil,
	// 	Data:        "DESDE $<?= $product->price ?>",
	// 	Type:        html.RawNode, // XD
	// 	Attr:        []html.Attribute{},
	// }
	//
	// span.Raw().AppendChild(textNodeToInsert)
	// println(a)
}

func printHtml(doc node.Node) {
	fmt.Println(gohtml.Format(doc.HTML()))
}

func run() {
	unzip("./1240.zip", "./app")
	web_files_manipulation.Init("./app")
}

func insertingIndexPHPHead(doc node.Node) string {
	ogHtml := gohtml.Format(doc.HTML())
	phpHead := `<?php 
	  include_once "app/config.php"; 
	  include_once "app/ProductsController.php"; 
	  $productsController = new ProductsController();   
	  $cafes = $productsController->getByCategory('cafe');
	  $merchs = $productsController->getByCategory('merch');
?>
`
	return fmt.Sprintf("%s%s", phpHead, ogHtml)
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

func insertRawTextInsideTag(text string, tag node.Node) {
	textNodeToInsert := &html.Node{
		Parent:      nil,
		PrevSibling: nil,
		NextSibling: nil,
		Data:        text,
		Type:        html.RawNode, // XD
		Attr:        []html.Attribute{},
	}
	tag.Raw().AppendChild(textNodeToInsert)
}

func scannDirectory(directory string) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range entries {
		fmt.Println(e.Name())
	}
	return nil
}

func WriteHtmlToPhp(fileSource string, fileDestiny string) {
	dat, _ := os.ReadFile(fileSource)
	formattedHtml := gohtml.FormatBytes((dat))
	// check if permissions should be like this, because maybe we fuck it up
	os.WriteFile(fileDestiny, formattedHtml, 0777)
}

func unzip(zipPath string, destinyPath string) {
	os.RemoveAll(destinyPath)
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		log.Fatalf("impossible to open zip reader: %s", err)
	}

	defer r.Close()

	for k, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatalf("impossible to open file n°%d in archive: %s", k, err)
		}
		defer rc.Close()
		newFilePath := fmt.Sprintf("%s/%s", destinyPath, f.Name)

		if f.FileInfo().IsDir() {
			err = os.MkdirAll(newFilePath, 0777)
			if err != nil {
				log.Fatalf("impossible to MkdirAll: %s", err)
			}
			continue
		}

		uncompressedFile, err := os.Create(newFilePath)
		if err != nil {
			log.Fatalf("impossible to create uncompressed: %s", err)
		}
		_, err = io.Copy(uncompressedFile, rc)
		if err != nil {
			log.Fatalf("impossible to copy file n°%d: %s", k, err)
		}
	}
}
