package web_files_manipulation

import (
	"fmt"
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
	// esto es query pero me da hueva hacer refactor xd
	Class                              string            `json:"class"`
	ShouldRemoveAllChildren            bool              `json:"should_removeall_children"`
	ShouldRemoveAllChildrenExceptFirst bool              `json:"should_remove_all_childrene_xcept_first"`
	ShouldReplaceOld                   bool              `json:"should_replace_old"`
	ShouldRemoveTagsWithSameClassName  bool              `json:"should_removetags_with_same_classname"`
	ShouldAppendAttributes             bool              `json:"should_append_attributes"`
	TargetParent                       bool              `json:"target_parent"`
	IsParent                           bool              `json:"is_parent"`
	TargetItemID                       string            `json:"target_item_id"`
	InnerHTML                          string            `json:"inner_html"`
	OuterHTML                          string            `json:"outer_html"`
	AppendHTML                         string            `json:"append_html"`
	PrependHTML                        string            `json:"prepend_html"`
	ForEach                            string            `json:"for_each"`
	TagsAttributes                     []TagAttribute    `json:"tags_attributes"`
	InnerHtmlReplacements              []HTMLReplacement `json:"inner_html_replacements"`
}

type Attribute struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type TagAttribute struct {
	Tag  string
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

// new API

// if query is empty or not initialized then the other instructions (AttributesChanges and HTMLChanges) use the whole document as the base search
type Modification struct {
	Query             string            `json:"query"`
	ChangeAll         bool              `json:"change_all"`
	DeleteSiblings    bool              `json:"delete_siblings"`
	InnerHTML         string            `json:"inner_html"`
	OuterHTML         string            `json:"outer_html"`
	AppendHTML        string            `json:"append_html"`
	PrependHTML       string            `json:"prepend_html"`
	AttributesChanges []AttributeChange `json:"attributes_changes"`
	HTMLChanges       []HTMLChange      `json:"html_changes"`
}

type AttributeMode string
type HTMLMode string

// const (
// 	INNER_HTML        = 0
// 	OUTER_HTML        = 1
// 	APPEND_HTML       = 3
// 	PREPEND_HTML      = 4
// 	REPLACE_ATTRIBUTE = 5
// 	APPEND_ATTRIBUTE  = 6
// )

const (
	INNER_HTML        = "inner"
	OUTER_HTML        = "outer"
	APPEND_HTML       = "append"
	PREPEND_HTML      = "prepend"
	REPLACE_ATTRIBUTE = "replace_attribute"
	APPEND_ATTRIBUTE  = "append_attribute"
)

// if query is empty or not initialized the upper query from the modfication is used as the targeting container
type AttributeChange struct {
	Query     string        `json:"query"`
	Mode      AttributeMode `json:"mode"`
	Attribute Attribute     `json:"attribute"`
}

type HTMLChange struct {
	Query string   `json:"query"`
	Mode  HTMLMode `json:"mode"`
	HTML  string   `json:"html"`
}

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
