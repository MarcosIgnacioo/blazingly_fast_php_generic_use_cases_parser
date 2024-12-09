package main

import (
	"archive/zip"
	"context"
	"fmt"
	"github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/copy"
	"github.com/yosssi/gohtml"
	"io"
	"log"
	"os"
)

func main() {
	unzip("./bundles.zip")
	CreatePhpDirectories()
	WriteHtmlToPhp("./bundles/index.html", "./exported/views/index.php")
	WriteHtmlToPhp("./bundles/tienda/index.html", "./exported/views/shop/index.php")
	WriteHtmlToPhp("./bundles/menu-malecon/index.html", "./exported/views/menus/menu_malecon.php")
	popo.Copy(context.TODO(), "./bundles/css/", "./exported/public/")
	popo.Copy(context.TODO(), "./bundles/js/", "./exported/public/")
	popo.Copy(context.TODO(), "./bundles/facturar/index.html", "./exported/public/files/")
	popo.Copy(context.TODO(), "./bundles/images/", "./exported/public/")
}

func WriteHtmlToPhp(fileSource string, fileDestiny string) {
	dat, _ := os.ReadFile(fileSource)
	CreatePhpDirectories()
	formattedHtml := gohtml.FormatBytes((dat))
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
			log.Fatalf("impossible to open file n°%d in archine: %s", k, err)
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
