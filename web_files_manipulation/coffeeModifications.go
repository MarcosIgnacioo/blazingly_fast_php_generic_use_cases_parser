package web_files_manipulation

//TOIMPROVE all of this use the same template literally everything so maybe a template for it could be good

var coffeeModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: coffeeHeader,
	},
	Modification{
		Query: ".product_item_coffee_mezclas",
		PrependHTML: `
		<?php if (isset($mezclas) && count($mezclas)): ?>
				<?php foreach ($mezclas as $product): ?> 
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
		Query: ".product_item_coffee_autor",
		PrependHTML: `
	<?php if (isset($autores) && count($autores)): ?>
		<?php foreach ($autores as $product): ?> 
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
		Query: ".product_item_coffee_origen",
		PrependHTML: `
		<?php if (isset($origenes) && count($origenes)): ?>
			<?php foreach ($origenes as $product): ?> 
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
		Query: ".product_item_favorite",
		PrependHTML: `
		<?php if (isset($merchs) && count($merchs)): ?>
			<?php foreach ($merchs as $product): ?> 
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
