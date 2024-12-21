package web_files_manipulation

var cartFormModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: formCarritoHeader,
	},
	Modification{
		Query: "append_scripts_to_end_of_body",
		AppendHTML: `
	    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
        <script type="text/javascript">
            document.querySelector('.checkout_form').addEventListener('submit', validateOrder);
        </script>
		`,
	},
	Modification{
		Query: ".cart_item_row",
		PrependHTML: `
		<?php $total = 0; ?>
			<?php if (isset($_SESSION['cart']) && count($_SESSION['cart'])): ?>
				<?php foreach ($_SESSION['cart'] as $product): ?> 
		`,
		AppendHTML: `
					<div class="ed-element ed-separator"><hr class="bg-primary" /></div>
				<?php $total += ($product->cantidad * $product->price) ?>
			<?php endforeach ?>
		<?php endif ?> 
		`,
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: ".cart_product_name",
				Mode:  INNER_HTML,
				HTML:  `<?= $product->name ?>`,
			},
			// HTMLChange{
			// 	Query: ".cart_product_description",
			// 	Mode:  INNER_HTML,
			// 	HTML:  `<?= $product->feature ?>`,
			// },
			HTMLChange{
				Query: ".cart_product_price",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format($product->price * $product->cantidad, 2)?>`,
			},
			HTMLChange{
				Query: ".cart_product_quantity",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format($product->price,2) ?> x <?= $product->cantidad ?>`,
			},
		},
	},

	Modification{
		Query:     "",
		ChangeAll: true,
		AttributesChanges: []AttributeChange{
			NewAttributeChange(
				`.submit_checkout`,
				REPLACE_ATTRIBUTE,
				`onclick="document.querySelector('.checkout_form')?.dispatchEvent(new Event('submit', { cancelable: true }))"`,
			),
			NewAttributeChange(
				`.checkout_name input`,
				REPLACE_ATTRIBUTE,
				`name="name"`,
			),
			NewAttributeChange(
				`.checkout_email input`,
				REPLACE_ATTRIBUTE,
				`name="email"`,
			),
			NewAttributeChange(
				`.checkout_phone input`,
				REPLACE_ATTRIBUTE,
				`name="phone"`,
			),
			NewAttributeChange(
				`.checkout_street_and_use_number input`,
				REPLACE_ATTRIBUTE,
				`name="street_and_use_number"`,
			),
			NewAttributeChange(
				`.checkout_street_and_use_number input`,
				REPLACE_ATTRIBUTE,
				`name="street_and_use_number"`,
			),
			NewAttributeChange(
				`.checkout_estate select`,
				REPLACE_ATTRIBUTE,
				`name="estate"`,
			),
			NewAttributeChange(
				`.checkout_city select`,
				REPLACE_ATTRIBUTE,
				`name="city"`,
			),
			NewAttributeChange(
				`.checkout_postal_code input`,
				REPLACE_ATTRIBUTE,
				`name="postal_code"`,
			),
			NewAttributeChange(
				`.checkout_references textarea`,
				REPLACE_ATTRIBUTE,
				`name="references"`,
			),
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
