package main

import (
	"archive/zip"
	"fmt"

	// arraylist "github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/array_list"
	"io"
	"log"
	"os"

	// "github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/web_files_manipulation"
	"github.com/sunshineplan/node"
	"github.com/yosssi/gohtml"
	"golang.org/x/net/html"
	// "github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/web_files_manipulation"
)

func main() {
	// web_files_manipulation.Init("./out")
	file, _ := os.ReadFile("./out/index.html")
	htmlContent := string(file)
	doc, err := node.ParseHTML(htmlContent)
	if err != nil {
		fmt.Println(err)
		return
	}
	doc.Find(node.Descendant, node.Div)
	os.WriteFile("popo.html", []byte(insertingIndexPHPHead(doc)), 0777)
	// poop := doc.Find(node.Descendant, node.Div, node.Class("separator_initial_cart"))
	// insertRawTextBeforeNode("<?php include \"layouts/cart_mobile.template.php\"; ?>", poop)
	// fmt.Println(gohtml.Format(doc.HTML()))
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
	CreatePhpDirectories()
	formattedHtml := gohtml.FormatBytes((dat))
	// check if permissions should be like this, because maybe we fuck it up
	os.WriteFile(fileDestiny, formattedHtml, 0777)
}

func CreatePhpDirectories() {
	// views directories
	os.MkdirAll("./exported", 0777)
	os.MkdirAll("./exported/views/account", 0777)
	os.MkdirAll("./exported/views/terms", 0777)
	os.MkdirAll("./exported/views/menus", 0777)
	os.MkdirAll("./exported/views/shop", 0777)
	os.MkdirAll("./exported/views/shop", 0777)
	// public directories
	os.MkdirAll("./exported/public/css", 0777)
	os.MkdirAll("./exported/public/files", 0777)
	os.MkdirAll("./exported/public/js", 0777)
}

func unzip(path string) {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(path)
	if err != nil {
		log.Fatalf("impossible to open zip reader: %s", err)
	}

	defer r.Close()

	// Iterate through the files in the archive,
	for k, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatalf("impossible to open file n°%d in archive: %s", k, err)
		}
		defer rc.Close()
		// define the new file path
		newFilePath := fmt.Sprintf("bundles/%s", f.Name)

		// CASE 1 : we have a directory
		if f.FileInfo().IsDir() {
			// if we have a directory we have to create it
			err = os.MkdirAll(newFilePath, 0777)
			if err != nil {
				log.Fatalf("impossible to MkdirAll: %s", err)
			}
			// we can go to next iteration
			continue
		}

		// CASE 2 : we have a file
		// create new uncompressed file
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
