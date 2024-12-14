package web_files_manipulation

import "github.com/sunshineplan/node"

// anchor tag attributes
var productAnchorHrefTagAtrr = TagAttribute{
	Tag: node.A,
	Attr: Attribute{
		Name:  "href",
		Value: productHref,
	},
}

var productAnchorTitleTagAtrr = TagAttribute{
	Tag: node.A,
	Attr: Attribute{
		Name:  "title",
		Value: productName,
	},
}

var productImgSrcTagAtrr = TagAttribute{
	Tag: node.Img,
	Attr: Attribute{
		Name:  "src",
		Value: productThumbnail,
	},
}

// img tags attributes
var productImgSrcSetTagAtrr = TagAttribute{
	Tag: node.Img,
	Attr: Attribute{
		Name:  "srcset",
		Value: productThumbnail,
	},
}

var productImgCoffeCoverFullPathSrcSetTagAtrr = TagAttribute{
	Tag: node.Img,
	Attr: Attribute{
		Name:  "srcset",
		Value: productItemCoffeCoverFullPath,
	},
}

var productImgCoffeCoverFullPathSrcTagAtrr = TagAttribute{
	Tag: node.Img,
	Attr: Attribute{
		Name:  "src",
		Value: productItemCoffeCoverFullPath,
	},
}

var productImgAltTagAtrr = TagAttribute{
	Tag: node.Img,
	Attr: Attribute{
		Name:  "alt",
		Value: productName,
	},
}

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
			"classPrice": "product_price",
			"className":  "product_name",
		},
	},
	"merch": {
		{
			"class":        "html",
			"htmlToInsert": merchHeader,
		},
		{
			"class":      "product_item_merch",
			"foreach":    merchForeachWrapper,
			"classPrice": "product_price",
			"className":  "product_name",
		},
	},
	"accesorios": {
		{
			"class":        "html",
			"htmlToInsert": accesoriesHeader,
		},
		{
			"class":      "product_item_accessories",
			"foreach":    accesoriesForeachWrapper,
			"classPrice": "product_price",
			"className":  "product_name",
		},
	},
	// "diablo": {
	// 	{
	// 		"class":        "html",
	// 		"htmlToInsert": devilHeader,
	// 	},
	// 	{
	// 		"class":      "PRODUCT_CLASS",
	// 		"foreach":    ForeachWrapper,
	// 		"img":        "<?= $productsController->getCover($grand_product) ?>",
	// 		"classPrice": "product_price",
	// 		"className":  "product_name",
	// 	},
	// },
}

var htmlManipulationInstructionSet = map[string][]map[string]string{
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
			"classPrice": "product_price",
			"className":  "product_name",
		},
	},
	"merch": {
		{
			"class":        "html",
			"htmlToInsert": merchHeader,
		},
		{
			"class":      "product_item_merch",
			"foreach":    merchForeachWrapper,
			"classPrice": "product_price",
			"className":  "product_name",
		},
	},
	"accesorios": {
		{
			"class":        "html",
			"htmlToInsert": accesoriesHeader,
		},
		{
			"class":      "product_item_accessories",
			"foreach":    accesoriesForeachWrapper,
			"classPrice": "product_price",
			"className":  "product_name",
		},
	},
	// "diablo": {
	// 	{
	// 		"class":        "html",
	// 		"htmlToInsert": devilHeader,
	// 	},
	// 	{
	// 		"class":      "PRODUCT_CLASS",
	// 		"foreach":    ForeachWrapper,
	// 		"img":        "<?= $productsController->getCover($grand_product) ?>",
	// 		"classPrice": "product_price",
	// 		"className":  "product_name",
	// 	},
	// },
}

var instructionsTYPED map[string][]Instruction = map[string][]Instruction{
	"app": {
		Instruction{
			Class:       "html",
			PrependHTML: homeHeader,
		},
		getBasicInstruction(
			"product_item_coffe",
			productItemCoffeeForeachWrapper,
			"product_price_home",
			"product_name_home",
			productImgCoffeCoverFullPathSrcSetTagAtrr,
			productImgCoffeCoverFullPathSrcTagAtrr,
		),
		getBasicInstruction(
			"product_item_favorite",
			productItemAllForeachWrapper,
			"product_price_home",
			"product_name_home"),
	},

	"tienda": {
		Instruction{
			Class:       "html",
			PrependHTML: storeHeader,
		},
		getBasicInstruction(
			"product_item_all",
			productItemAllForeachWrapper,
			"product_price",
			"product_name"),
	},

	"cafe": {
		Instruction{
			Class:       "html",
			PrependHTML: coffeeHeader,
		},
		getBasicInstruction(
			"product_item_coffee_mezclas",
			coffeeMixesForeachWrapper,
			"product_price",
			"product_name"),
		getBasicInstruction(
			"product_item_coffee_autor",
			authorForeachWrapper,
			"product_price",
			"product_name"),
		getBasicInstruction(
			"product_item_coffee_origen",
			originForeachWrapper,
			"product_price",
			"product_name"),
		getBasicInstruction(
			"product_item_favorite",
			productItemFavoriteForeachWrapper,
			"product_price",
			"product_name"),
	},
	"merch": {
		Instruction{
			Class:       "html",
			PrependHTML: merchHeader,
		},
		getBasicInstruction(
			"product_item_merch",
			merchForeachWrapper,
			"product_price",
			"product_name"),
	},

	// DIRECTORY: {
	// 	Instruction{
	// 		Class:       "html",
	// 		PrependHTML: HEADER,
	// 	},
	// 	getBasicInstruction(
	// 		productClass,
	// 		forEachWrapper,
	// 		"product_price",
	// 		"product_name"),
	// },

}

// HTMLManipulation: []Manipulation { {ClassName:string, inner:string, outer:string, append:string, prepend:string}}

// {
// 	"class":        "html",
// 	"htmlToInsert": storeHeader,
// },
// {
// 	"class":      "product_item_all",
// 	"foreach":    productItemAllForeachWrapper,
// 	"classPrice": "product_price",
// 	"className":  "product_name",
// },
