package web_files_manipulation

var rootDirModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: homeHeader,
	},
	Modification{
		Query: ".product_item_coffe",
		PrependHTML: `
		<?php if (isset($cafes) && count($cafes)): ?>
			<?php foreach (array_slice($cafes, 0, 4) as $cafe): ?> 
				<?php $product = $productsController->getQuickView($cafe) ?>
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
					Value: `<?= $cafe->cover->full_path ?>`,
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
				Query: ".product_name_home",
				Mode:  INNER_HTML,
				HTML:  `<?= $product->name ?>`,
			},
			HTMLChange{
				Query: ".product_price_home",
				Mode:  INNER_HTML,
				HTML:  `DESDE $<?= $product->price ?>`,
			},
		},
	},
	Modification{
		Query: ".product_item_favorite",
		PrependHTML: `
		<?php if (isset($merchs) && count($merchs)): ?>
			<?php foreach (array_slice($merchs, 0, 3) as $merch): ?> 
				<?php $product = $productsController->getQuickView($merch) ?>
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
				Query: ".product_name_home",
				Mode:  INNER_HTML,
				HTML:  `<?= $product->name ?>`,
			},
			HTMLChange{
				Query: ".product_price_home",
				Mode:  INNER_HTML,
				HTML:  `$<?= $product->price ?>`,
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
