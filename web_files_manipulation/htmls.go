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
		Value: "<?= $product->name ?>",
	},
}

var productHref = `<?= $product->url ?>`
var productPrice = `DESDE $<?= $product->price ?>`
var productThumbnail = `<?= $product->thumbnail_path ?>`
var productName = `<a href="<?= $product->url ?>"><?= $product->name ?></a>`

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
			productName,
			productPrice,
			productImgCoffeCoverFullPathSrcSetTagAtrr,
			productImgCoffeCoverFullPathSrcTagAtrr,
		),
		getBasicInstruction(
			"product_item_favorite",
			productItemAllForeachWrapper,
			"product_price_home",
			"product_name_home",
			productName,
			`<?= $product->price ?>`,
		),
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
			"product_name",
			productPrice,
			productName,
		),
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
			"product_name",
			productPrice,
			`<a href="<?= $product->url ?>" title=""><span style="font-family: 'daisywhl';"><?= $product->name ?></span></a>`,
		),
		getBasicInstruction(
			"product_item_coffee_autor",
			authorForeachWrapper,
			"product_price",
			"product_name",
			productPrice,
			productName,
		),
		getBasicInstruction(
			"product_item_coffee_origen",
			originForeachWrapper,
			"product_price",
			"product_name",
			productPrice,
			productName,
		),
		getBasicInstruction(
			"product_item_favorite",
			productItemFavoriteForeachWrapper,
			"product_price",
			"product_name",
			productPrice,
			productName,
		),
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
			"product_name",
			productPrice,
			`<a href="<?= $product->url ?>" title=""> <span style="font-family: 'daisywhl';"><?= $product->name ?></span></a>`,
		),
	},
	"accesorios": {
		Instruction{
			Class:       "html",
			PrependHTML: accesoriesHeader,
		},
		getBasicInstruction(
			"product_item_accessories",
			accesoriesForeachWrapper,
			"product_price",
			"product_name",
			productPrice,
			productName,
		),
	},
	"diablo": {
		Instruction{
			Class:       "html",
			PrependHTML: devilHeader,
		},
		Instruction{
			Class: ".sizes_items_details select",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: node.Select,
					Attr: Attribute{
						Name:  "onchange",
						Value: "updateSelection(this)",
					},
				},
			},
			InnerHTML: `
				<?php if (isset($grand_product->presentations) && count($grand_product->presentations)): ?> 
					<?php foreach ($grand_product->presentations as $presentation): ?> 
							<option data-id="<?=$presentation->id ?>" data-price="<?=$presentation
										->current_price->amount ?>" value="<?=$presentation->id ?>" data-product='<?=json_encode($productsController->getQuickView($grand_product, $presentation->id)) ?>' >
									<?=$presentation->description
				?>
							</option> 
					<?php
						endforeach ?>
					<?php
				endif ?>
			`,
		},
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
