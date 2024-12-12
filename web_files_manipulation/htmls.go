package web_files_manipulation

var instructions = map[string][]map[string]string{
	"app": {
		{"class": "head", "htmlToInsert": homeHeader},
		{"class": "product_item_coffe", "htmlToInsert": homeProducts},
	},
	"tienda": {
		{"class": "head", "htmlToInsert": homeHeader},
		{"class": "product_item_coffe", "htmlToInsert": homeProducts},
	},
}
