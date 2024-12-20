package web_files_manipulation

var storeModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: storeHeader,
	},
	Modification{
		Query: ".product_item_all",
		PrependHTML: `
		<?php if (isset($products) && count($products)): ?>
			<?php foreach ($products as $product): ?> 
				<?php $product = $productsController->getQuickView($product) ?>
		`,
		AppendHTML: `
			<?php endforeach ?>
    <?php endif ?>
		`,
		AttributesChanges: []AttributeChange{
			AttributeChange{
				Query: "a",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "href",
					Value: `<?= $product->url ?>`,
				},
			},
			AttributeChange{
				Query: "img",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "src",
					Value: `<?= $product->thumbnail_path ?>`,
				},
			},
			AttributeChange{
				Query: "img",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "srcset",
					Value: ``,
				},
			},
			AttributeChange{
				Query: "img",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "alt",
					Value: `<?= $product->name ?>`,
				},
			},
			AttributeChange{
				Query: "img",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "title",
					Value: `<?= $product->name ?>`,
				},
			},
		},
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: ".product_name",
				Mode:  INNER_HTML,
				HTML:  `<?= $product->name ?>`,
			},
			HTMLChange{
				Query: ".product_price",
				Mode:  INNER_HTML,
				HTML:  `DESDE $<?= $product->price ?>`,
			},
		},
	},
	Modification{
		Query: ".remove_item_on_update",
		AttributesChanges: []AttributeChange{
			AttributeChange{
				Query: "",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "style",
					Value: `display: none;`,
				},
			},
		},
	},
}
