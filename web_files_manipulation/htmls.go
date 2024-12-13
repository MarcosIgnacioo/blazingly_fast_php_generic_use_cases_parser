package web_files_manipulation

var instructions = map[string][]map[string]string{
	"app": {
		{
			"class":        "html",
			"htmlToInsert": homeHeader,
		},
		{
			"class":       "product_item_coffe",
			"href":        homeProductHref,
			"productName": homeProductName,
			"img":         homeProductIMG,
			"foreach":     homeForeachWrapper,
			"price":       homeProductPrice,
			"classPrice":  "product_price_home",
			"className":   "product_name_home",
		},
		{
			"class":       "product_item_favorite",
			"href":        homeProductFavHref,
			"productName": homeProductFavName,
			"img":         homeProductFavIMG,
			"foreach":     homeForeachFavWrapper,
			"price":       homeProductFavPrice,
			"classPrice":  "product_price_home",
			"className":   "product_name_home",
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
			"foreach":     storeForeachWrapper,
			"price":       storeProductPrice,
			"classPrice":  "product_price",
			"className":   "product_name",
		},
	},
	"cafe": {
		{
			"class":        "html",
			"htmlToInsert": coffeeHeader,
		},
		{
			"class":       "product_item_coffee_mezclas",
			"href":        coffeeProductHref,
			"productName": coffeeProductName,
			"img":         coffeeProductIMG,
			"foreach":     coffeeForeachWrapper,
			"price":       coffeeProductPrice,
			"classPrice":  "product_price",
			"className":   "product_name",
		},
		{
			"class":       "product_item_coffee_autor",
			"href":        authorProductHref,
			"productName": authorProductName,
			"img":         authorProductIMG,
			"foreach":     authorForeachWrapper,
			"price":       authorProductPrice,
			"classPrice":  "product_price",
			"className":   "product_name",
		},
		{
			"class":       "product_item_coffee_origen",
			"href":        originProductHref,
			"productName": originProductName,
			"img":         originProductIMG,
			"foreach":     originForeachWrapper,
			"price":       originProductPrice,
			"classPrice":  "product_price",
			"className":   "product_name",
		},
	},
}
