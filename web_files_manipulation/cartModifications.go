package web_files_manipulation

var cartModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: cartHeader,
	},
	Modification{
		Query: "",
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: ".cart_subtotal",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format( $total ,2) ?>`,
			},
			HTMLChange{
				Query: ".cart_total",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format( $total ,2) ?>`,
			},
		},
	},
	// falta un separator pero  no lo dijeron en el readme y me da un poco d hueva meterlo la vddd
	Modification{
		Query: ".product_item_cart",
		PrependHTML: `
		<?php $total = 0; ?>
			<?php if (isset($_SESSION['cart']) && count($_SESSION['cart'])): ?>
				<?php foreach ($_SESSION['cart'] as $product): ?> 
		`,
		AppendHTML: `
				<?php $total += ($product->cantidad * $product->price) ?>
			<?php endforeach ?>
		<?php endif ?> 
		`,
		AttributesChanges: []AttributeChange{
			NewAttributeChange(
				``,
				APPEND_ATTRIBUTE,
				`class="item_<?= $product->id ?>_<?= str_replace(' ', '', $product->feature) ?>"`,
			),
			NewAttributeChange(
				".cart_product_quantity",
				REPLACE_ATTRIBUTE,
				`value="<?= $product->cantidad ?>"`,
			),
			NewAttributeChange(
				".cart_product_delete",
				REPLACE_ATTRIBUTE,
				`value="<?= $product->cantidad ?>"`,
			),
			NewAttributeChange(
				".cart_product_delete",
				REPLACE_ATTRIBUTE,
				`onclick="removeItemCart(this)"`,
			),
			NewAttributeChange(
				".cart_product_delete",
				REPLACE_ATTRIBUTE,
				`data-id="<?= $product->id ?>"`,
			),

			NewAttributeChange(
				".cart_product_delete",
				REPLACE_ATTRIBUTE,
				`data-parent="item_<?= $product->id ?>_<?= str_replace(' ', '', $product->feature) ?>`,
			),
			NewAttributeChange(
				".cart_product_delete",
				REPLACE_ATTRIBUTE,
				`role="button"`,
			),
		},
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: ".cart_product_img",
				Mode:  OUTER_HTML,
				HTML: `
				<img
					src="<?= $product->cover ?>"
					alt=""
					class="ed-lazyload cart_product_img"
					style="object-fit: cover;"
				/>`,
			},
			HTMLChange{
				Query: ".cart_product_name",
				Mode:  INNER_HTML,
				HTML:  `<?= $product->name ?>`,
			},
			HTMLChange{
				Query: ".cart_product_price",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(floatval($product->price) ?? 0, 2)?>`,
			},
			HTMLChange{
				Query: ".cart_product_subtotal",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(floatval($product->cantidad * $product->price) ?? 0, 2) ?>`,
			},
			HTMLChange{
				Query: ".cart_product_description",
				Mode:  INNER_HTML,
				HTML:  `<?= $product->feature ?>`,
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
