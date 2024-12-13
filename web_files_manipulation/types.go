package web_files_manipulation

import (
	"fmt"

	"github.com/sunshineplan/node"
)

type readingFilesInDirectoryError struct {
	fileName string
}

// "accesorios": {
// 	{
// 		"class":        "html",
// 		"htmlToInsert": accesoriesHeader,
// 	},
// 	{
// 		"class":      "product_item_accessories",
// 		"foreach":    accesoriesForeachWrapper,
// 		"classPrice": "product_price",
// 		"className":  "product_name",
// 		"attributesManipulation" : {
// 			"tagsToChangeAttributes":{"aasdf,asdf,asdf"},
// 			"attributesInTagsToChange":{"href":"asdf"},
// 		},
// 		"innerHtmlReplacements" : {
// 			"product_price": "<?=$product->details?>"
// 			"product_name" : "<?=$foo?>"
// 		}
// 	},
// },

type Instruction struct {
	Class                 string            `json:"class"`
	InnerHTML             string            `json:"inner_html"`
	OuterHTML             string            `json:"outer_html"`
	AppendHTML            string            `json:"append_html"`
	PrependHTML           string            `json:"prepend_html"`
	ForEach               string            `json:"for_each"`
	TagsAttributes        []TagAttribute    `json:"tags_attributes"`
	InnerHtmlReplacements []HTMLReplacement `json:"inner_html_replacements"`
}

type Attribute struct {
	Name  string
	Value string
}

type TagAttribute struct {
	Tag  node.TagFilter
	Attr Attribute
}

type HTMLReplacement struct {
	ClassName string
	HTML      string
}

type InnerHTMLReplacement struct {
	ClassNameReplacementHTML map[string]string
}

type FileType string

var fileTypesMap = map[string]FileType{
	".html": "html",
	".js":   "js",
	".css":  "css",
	".php":  "php",
	".jpg":  "img",
	".png":  "img",
	".svg":  "img",
}

const (
	HTML  FileType = "html"
	JS             = "js"
	CSS            = "css"
	PHP            = "php"
	IMAGE          = "img"
	FONT           = "font"
)

type File struct {
	fileName      string
	filePath      string
	fileParentDir string
	fileType      FileType
	nestedLevel   int
}

func (f *File) String() string {
	return fmt.Sprintf("file_name:\t%s\nfile_path:\t%s\nfile_type:\t%v\nnested_level:\t%d\nparent_dir:\t%s\n", f.fileName, f.filePath, f.fileType, f.nestedLevel, f.fileParentDir)
}

func (self *readingFilesInDirectoryError) Error() string {
	return fmt.Sprintf("error reading directory %s", self.fileName)
}

const (
	default_array_size = 12
)

// file:///home/happy/Downloads/%3C?=%20$image-%3Efull_path%20?%3E
// "diablo": {
// {
// 	"class" :"product_description_details"
// 	"foreach":    accesoriesForeachWrapper,
// 	"attributesManipulation" : {
// 		"tagsToChangeAttributes":{node.A, node.Img}
// 		"attributesInTagsToChange":{"href":"asdf"}
// 	},
// 	"innerHtmlReplacements" : {
// 		"product_description_details": {"<?=$product->details?>"}
// 	}
// }
// }
