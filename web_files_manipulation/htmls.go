package web_files_manipulation

var productHref = `<?= $product->url ?>`
var productName = `<?= $product->name ?>`
var productThumbnail = `<?= $product->thumbnail_path ?>`
var productPrice = `<?= $product->price ?>`

var instructions = map[string][]map[string]string{
	"app": {
		{
			"class":        "html",
			"htmlToInsert": homeHeader,
		},
		{
			"class":      "product_item_coffe",
			"foreach":    productItemCoffeeForeachWrapper,
			"img":        productItemCoffeCoverFullPath,
			"classPrice": "product_price_home",
			"className":  "product_name_home",
		},
		{
			"class":      "product_item_favorite",
			"foreach":    productItemFavoriteForeachWrapper,
			"classPrice": "product_price_home",
			"className":  "product_name_home",
		},
	},
	"tienda": {
		{
			"class":        "html",
			"htmlToInsert": storeHeader,
		},
		{
			"class":      "product_item_all",
			"foreach":    productItemAllForeachWrapper,
			"classPrice": "product_price",
			"className":  "product_name",
		},
	},
	"cafe": {
		{
			"class":        "html",
			"htmlToInsert": coffeeHeader,
		},
		{
			"class":      "product_item_coffee_mezclas",
			"foreach":    coffeeMixesForeachWrapper,
			"classPrice": "product_price",
			"className":  "product_name",
		},
		{
			"class":      "product_item_coffee_autor",
			"foreach":    authorForeachWrapper,
			"classPrice": "product_price",
			"className":  "product_name",
		},
		{
			"class":      "product_item_coffee_origen",
			"foreach":    originForeachWrapper,
			"classPrice": "product_price",
			"className":  "product_name",
		},
		{
			"class":      "product_item_favorite",
			"foreach":    productItemFavoriteForeachWrapper,
			"classPrice": "product_price_asdf",
			"className":  "product_name",
		},
	},
}
