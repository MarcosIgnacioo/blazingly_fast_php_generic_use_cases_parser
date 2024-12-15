package web_files_manipulation

import "fmt"

var IDS map[string]string = map[string]string{
	"": "",
}

// anchor tag attributes
var productAnchorHrefTagAtrr = TagAttribute{
	Tag: "a",
	Attr: Attribute{
		Name:  "href",
		Value: productHref,
	},
}

var productAnchorTitleTagAtrr = TagAttribute{
	Tag: "a",
	Attr: Attribute{
		Name:  "title",
		Value: productName,
	},
}

var productImgSrcTagAtrr = TagAttribute{
	Tag: "img",
	Attr: Attribute{
		Name:  "src",
		Value: productThumbnail,
	},
}

// img tags attributes
var productImgSrcSetTagAtrr = TagAttribute{
	Tag: "img",
	Attr: Attribute{
		Name:  "srcset",
		Value: productThumbnail,
	},
}

var productImgCoffeCoverFullPathSrcSetTagAtrr = TagAttribute{
	Tag: "img",
	Attr: Attribute{
		Name:  "srcset",
		Value: productItemCoffeCoverFullPath,
	},
}

var productImgCoffeCoverFullPathSrcTagAtrr = TagAttribute{
	Tag: "img",
	Attr: Attribute{
		Name:  "src",
		Value: productItemCoffeCoverFullPath,
	},
}

var productImgAltTagAtrr = TagAttribute{
	Tag: "img",
	Attr: Attribute{
		Name:  "alt",
		Value: "<?= $product->name ?>",
	},
}

var productHref = `<?= $product->url ?>`
var productPrice = `DESDE $<?= $product->price ?>`
var productThumbnail = `<?= $product->thumbnail_path ?>`
var productName = `<a href="<?= $product->url ?>"><?= $product->name ?></a>`

var instructionsTYPED map[string][]Instruction = map[string][]Instruction{
	"app": {
		Instruction{
			Class:       "html",
			PrependHTML: homeHeader,
		},
		getBasicInstruction(
			".product_item_coffe",
			productItemCoffeeForeachWrapper,
			"product_price_home",
			"product_name_home",
			productName,
			productPrice,
			productImgCoffeCoverFullPathSrcSetTagAtrr,
			productImgCoffeCoverFullPathSrcTagAtrr,
		),
		getBasicInstruction(
			".product_item_favorite",
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
			".product_item_all",
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
			".product_item_coffee_mezclas",
			coffeeMixesForeachWrapper,
			"product_price",
			"product_name",
			productPrice,
			`<a href="<?= $product->url ?>" title=""><span style="font-family: 'daisywhl';"><?= $product->name ?></span></a>`,
		),
		getBasicInstruction(
			".product_item_coffee_autor",
			authorForeachWrapper,
			"product_price",
			"product_name",
			productPrice,
			productName,
		),
		getBasicInstruction(
			".product_item_coffee_origen",
			originForeachWrapper,
			"product_price",
			"product_name",
			productPrice,
			productName,
		),
		getBasicInstruction(
			".product_item_favorite",
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
			".product_item_merch",
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
			".product_item_accessories",
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
		// como lo reworkearia sweria que en la clase solamente estuviera el contenedor padre y dentro de los html manipulation inner outer append prepend meter la clase o tag de donde se va a meter ese html
		Instruction{
			Class: ".sizes_items_details select",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag: "",
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
		Instruction{
			Class: ".quantity_items_product select",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag: "",
					Attr: Attribute{
						Name:  "class",
						Value: "cantidad_producto",
					},
				},
			},
			InnerHTML: `
			<?php if (isset($stock) && $stock > 0): ?>

				<?php for ($i = 1; $i < $stock; $i++): ?>
						<option value="<?= $i ?>">
								<?= $i ?>
						</option>
						<?php if ($i > 9) { break; } ?>
				<?php endfor; ?>

			<?php endif; ?> 
			`,
		},
		// FALTA el producttile detail product price details product description details PONERLO
		Instruction{
			Class: ".molienda_container select",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag: "",
					Attr: Attribute{
						Name:  "class",
						Value: "feature_product",
					},
				},
			},
			InnerHTML: `
			<?php foreach ($features as $feature): ?>
				<option value="<?= $feature  ?>">
						<?= $feature  ?>
				</option>
			<?php endforeach ?> 
			`,
		},

		Instruction{
			Class: ".boton_add_to_cart a",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag: "",
					Attr: Attribute{
						Name:  "onclick",
						Value: "QuickAdd(this)",
					},
				},
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag: "",
					Attr: Attribute{
						Name:  "data-product",
						Value: "<?= json_encode($productsController->getQuickView($grand_product)) ?>",
					},
				},
			},
		},
		Instruction{
			Class: ".tueste_item_details",
			InnerHTML: `
			<?php if ($tueste!=""): ?>
				<img src="<?= $tueste ?>" alt="" />
				<?php else: ?>
				<img src="/images/250/10441024/escala_tueste.png" alt="" />
			<?php endif ?> 
			`,
		},
		Instruction{
			Class: ".sabor_item_details",
			InnerHTML: `
			<?php if ($perfil_de_sabor!=""): ?>
					<img src="<?= $perfil_de_sabor ?>" alt="" />
				<?php else: ?>
					<img src="/images/250/10441024/escala_tueste.png" alt="" />
			<?php endif ?> 
			`,
		},
		Instruction{
			Class: ".ed-gallery-thumb",
			ForEach: `
				<?php if (isset($grand_product) && isset($grand_product->images)): ?>
					<?php foreach ($grand_product->images as $image): ?>  
						<?php if (!$image->is_cover): ?> 
							%s
						<?php endif ?>
					<?php endforeach ?> 
				<?php endif ?>
			`,
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "a",
					Attr: Attribute{
						Name:  "href",
						Value: "<?= $image->full_path ?>",
					},
				},
				TagAttribute{
					Tag: "img",
					Attr: Attribute{
						Name:  "src",
						Value: "<?= $image->full_path ?>",
					},
				},
				TagAttribute{
					Tag: "img",
					Attr: Attribute{
						Name:  "srcset",
						Value: "<?= $image->full_path ?>",
					},
				},
			},
		},
		Instruction{
			Class: ".product_item_recomendation",
			ForEach: `
		<?php if (isset($grand_product->related_products) && count($grand_product->related_products)): ?> 
				<?php foreach ($grand_product->related_products as $product): ?>
					<?php $product = $productsController->getQuickView($product) ?>
				%s
				<?php endforeach ?> 
			<?php endif ?> 
			`,
			InnerHtmlReplacements: []HTMLReplacement{
				HTMLReplacement{
					ClassName: "item_recomendation_name",
					HTML:      "<?= $product->name ?>",
				},
				HTMLReplacement{
					ClassName: "item_recomendation_price",
					HTML:      "DESDE $<?= $product->price ?>",
				},
			},
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "a",
					Attr: Attribute{
						Name:  "href",
						Value: "<?= $product->url ?>",
					},
				},
				TagAttribute{
					Tag: "img",
					Attr: Attribute{
						Name:  "src",
						Value: "<?= $image->full_path ?>",
					},
				},
				TagAttribute{
					Tag: "img",
					Attr: Attribute{
						Name:  "srcset",
						Value: "<?= $image->full_path ?>",
					},
				},
			},
		},
		Instruction{
			Class:      "body",
			AppendHTML: fmt.Sprintf("<script type=\"text/javascript\"> document.getElementById('%s').onchange() </script>", IDS["sizes_items_details"]),
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
