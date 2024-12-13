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
			Class:                 "html",
			PrependHTML:           storeHeader,
			ForEach:               "",
			TagsAttributes:        nil,
			InnerHtmlReplacements: nil,
		},
		Instruction{
			Class:   "product_item_coffe",
			ForEach: productItemAllForeachWrapper,
			TagsAttributes: []TagAttribute{
				productAnchorHrefTagAtrr,
				productAnchorTitleTagAtrr,
				productImgSrcSetTagAtrr,
				productImgSrcTagAtrr,
				productImgAltTagAtrr,
			},
			// HTMLManipulation: []Manipulation { {ClassName:string, inner:string, outer:string, append:string, prepend:string}}
			InnerHtmlReplacements: []HTMLReplacement{
				HTMLReplacement{
					ClassName: "product_price_home",
					HTML:      productPrice,
				},
				HTMLReplacement{
					ClassName: "product_name_home",
					HTML:      productName,
				},
			},
		},
	},
}

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
