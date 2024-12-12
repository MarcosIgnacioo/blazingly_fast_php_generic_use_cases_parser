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

var testhtml = `<div class="ed-element ed-container columns-box wv-overflow_visible product_item_all" id="ed-456730256"><div class="inner"><div class="ed-element ed-image" id="ed-456730259"><a href="/tienda/cafe/diablo"><img src="/images/1024/9706732/Coyote-Over.jpg" alt="" srcset="/images/576/9706732/Coyote-Over.jpg 576w, /images/975/9706732/Coyote-Over.jpg 975w"></a></div><div class="ed-element ed-text custom-theme" id="ed-456730262"><p style="line-height: 1.15; text-align: center;"><span style="font-size: 18px; font-family: &quot;daisywhl&quot;;"><a href="/tienda/cafe/diablo" title="">MEZCLA COYOTE</a></span></p><p style="line-height: 1; text-align: center;"><span style="font-size: 14px; color: rgb(94, 94, 94); font-family: &quot;daisywhl&quot;;">DESDE $240</span></p></div></div></div>`

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
	run()
	// doc, _ := node.ParseHTML(testhtml)
	// anchors := doc.FindAll(node.Descendant, node.A)
	// // imgs := doc.FindAll(node.Descendant, node.Img)
	// for _, a := range anchors {
	// 	setAttribute(&a, "href", "caca")
	// }
	//
	// for _, a := range anchors {
	// 	fmt.Println(a.Raw().Attr)
	// }
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
