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

func main() {
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
