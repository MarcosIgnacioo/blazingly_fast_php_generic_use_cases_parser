package web_files_manipulation

var individualOrderModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: individualOrderHeader,
	},
	Modification{
		Query: "",
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: ".order_folio",
				Mode:  INNER_HTML,
				HTML:  `ORDEN #<?= $_GET['folio'] ?>`,
			},
			HTMLChange{
				Query: ".order_subtotal",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(($order->total),2) ?>`,
			},
			HTMLChange{
				Query: ".order_shipping_cost",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(($order->shipping_cost ?? 0),2) ?>`,
			},
			HTMLChange{
				Query: ".order_total",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(($order->total),2) ?>`,
			},
			HTMLChange{
				Query: ".button_payment",
				Mode:  OUTER_HTML,
				HTML: `
				<?php if($order->order_status->id == 2): ?>
					<a href="<?= BASE_PATH ?>tienda/payment/<?= $order->folio ?>/" 
					class="button center button-variant-secondary button-payment">
						CONTINUAR CON EL PAGO
					</a>
				<?php endif ?>
`,
			},
		},
	},
	Modification{
		Query: ".cart_item_row",
		PrependHTML: `
		<?php if (isset($order->presentations) && count($order->presentations)): ?>
			<?php foreach ($order->presentations as $presentation): ?>                 
		`,
		AppendHTML: `
				<div class="ed-element ed-separator"><hr class="bg-primary" /></div>
			<?php endforeach ?>
		<?php endif ?> 
		`,
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: ".cart_product_img",
				Mode:  OUTER_HTML,
				HTML: `
				<img
					src="<?= $presentation->cover?->full_path ?>"
					alt=""
					class="ed-lazyload cart_product_img"
					style="object-fit: cover;"
				/>
`,
			},

			HTMLChange{
				Query: ".cart_product_name",
				Mode:  INNER_HTML,
				HTML: `
				<?php if(isset($presentation->product)): ?>
					<a href="<?= BASE_PATH ?>shop/details/<?= $presentation->product->slug; ?>/">
						<?= $presentation->product->name ?>  
					</a>
				<?php endif ?>  
`,
			},
			HTMLChange{
				Query: ".cart_product_description",
				Mode:  INNER_HTML,
				HTML:  `<?= $presentation->pivot->comment ?>`,
			},
			HTMLChange{
				Query: ".cart_product_price",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(($presentation->pivot->price->amount ?? 0) * $presentation->pivot->quantity,2) ?>`,
			},
			HTMLChange{
				Query: ".cart_product_quantity",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(($presentation->pivot->price->amount ?? 0),2) ?> x <?= $presentation->pivot->quantity ?>`,
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
