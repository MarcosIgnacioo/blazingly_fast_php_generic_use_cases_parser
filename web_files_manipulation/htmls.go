package web_files_manipulation

var instructions = map[string][]map[string]string{
	"app": {
		{
			"class":        "html",
			"htmlToInsert": homeHeader,
		},
		{
			"class":        "product_item_coffe",
			"htmlToInsert": homeProducts,
		},
	},
	"tienda": {
		{
			"class":        "html",
			"htmlToInsert": storeHeader,
		},
		{
			"class":       "product_item_all",
			"href":        storeProductHref,
			"productName": storeProductName,
			"img":         storeProductIMG,
		},
	},
}
