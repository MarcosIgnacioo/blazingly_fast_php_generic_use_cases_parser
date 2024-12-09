package web_files_manipulation

import (
	"fmt"
	"os"

	arraylist "github.com/MarcosIgnacioo/blazingly_fast_php_generic_use_cases_parser/array_list"
)

func Init(sourceDirectory string) {
	directories, files, err := getFilesInDirectory(sourceDirectory)
	if err != nil {
		panic(fmt.Sprintf("Failed at initial directory %s", sourceDirectory))
	}
	contentTransformation(directories, files)
}

func contentTransformation(directories *arraylist.ArrayList, files *arraylist.ArrayList) {
	// html a php
	fmt.Println(directories)
	fmt.Println(files)
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
			files.Enqueue(e.Name())
		}
	}
	return
}

// func scannDirectory() error {
// 	// scannDirectory(directories[i])
// }

// func HTMLToPHP(FILE* file) void {
// }
