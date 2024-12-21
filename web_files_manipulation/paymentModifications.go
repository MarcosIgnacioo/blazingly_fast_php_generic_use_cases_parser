package web_files_manipulation

var paymentModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: paymentHeader,
	},
	Modification{
		Query: "append_scripts_to_end_of_body",
		AppendHTML: `
	    <script type="text/javascript">
			const order_folio = '<?= $_GET['folio'] ?>'
            document.querySelector('.payment_form').addEventListener('submit', validatePayment);
        </script>		`,
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
			// HTMLChange{
			// 	Query: ".cart_product_description",
			// 	Mode:  INNER_HTML,
			// 	HTML:  `<?= $presentation->pivot->comment ?>`,
			// },
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
		Query: "",
		AttributesChanges: []AttributeChange{
			NewAttributeChange(
				`.card_number input`,
				REPLACE_ATTRIBUTE,
				`name="numero_tarjeta"`,
			),
			NewAttributeChange(
				`.card_exp_date input`,
				REPLACE_ATTRIBUTE,
				`name="fecha_exp"`,
			),
			NewAttributeChange(
				`.card_company select`,
				REPLACE_ATTRIBUTE,
				`name="marca_tarjeta"`,
			),
			NewAttributeChange(
				`.card_type select`,
				REPLACE_ATTRIBUTE,
				`name="tipo_tarjeta"`,
			),
			// NewAttributeChange(
			// 	`.address_city select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`name="ciudad"`,
			// ),
			// NewAttributeChange(
			// 	`.address_city select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`value="<?= $order->address->city ?>"`,
			// ),
			// NewAttributeChange(
			// 	`.address_postal_code select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`name="codigo_postal"`,
			// ),
			// NewAttributeChange(
			// 	`.address_postal_code select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`value="<?= $order->address->postal_code ?>"`,
			// ),
			// NewAttributeChange(
			// 	`.address_first_name select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`name="nombre"`,
			// ),
			// NewAttributeChange(
			// 	`.address_first_name select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`value="<?= $order->address->first_name ?>"`,
			// ),
			// NewAttributeChange(
			// 	`.address_phone_number select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`name="numero_celular"`,
			// ),
			// NewAttributeChange(
			// 	`.address_phone_number select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`value="<?= $order->address->phone_number ?>"`,
			// ),
			// NewAttributeChange(
			// 	`.address_state select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`name="estado"`,
			// ),
			// NewAttributeChange(
			// 	`.address_state select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`value="<?= $order->address->state ?>"`,
			// ),
			// NewAttributeChange(
			// 	`.address_street_and_use_number select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`name="calle"`,
			// ),
			// NewAttributeChange(
			// 	`.address_street_and_use_number select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`value="<?= $order->address->street_and_use_number ?>`,
			// ),
			// NewAttributeChange(
			// 	`.address_last_name select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`name="apellido"`,
			// ),
			// NewAttributeChange(
			// 	`.address_last_name select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`value="<?= $order->address->last_name ?>"`,
			// ),
			// NewAttributeChange(
			// 	`.address_email select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`name="correo"`,
			// ),
			// NewAttributeChange(
			// 	`.address_email select`,
			// 	REPLACE_ATTRIBUTE,
			// 	`value="<?= $order->client->email ?>"`,
			// ),
		},
		HTMLChanges: []HTMLChange{
			// HTMLChange{
			// 	Query: ".order_folio",
			// 	Mode:  INNER_HTML,
			// 	HTML:  `ORDEN #<?= $_GET['folio'] ?>`,
			// },
			HTMLChange{
				Query: ".card_company select",
				Mode:  INNER_HTML,
				HTML: `
					<option value="" selected disabled>Selecciona una opción</option>
					<option value="VISA">Visa</option>
					<option value="MC">Mastercard</option>
			`,
			},
			HTMLChange{
				Query: ".card_type select",
				Mode:  INNER_HTML,
				HTML: `
					<option value="" selected disabled>Selecciona una opción</option>
					<option value="CR">Crédito</option>
					<option value="DB">Débito</option>
			`,
			},
			HTMLChange{
				Query: ".checkout_subtotal",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(($order->total),2) ?>`,
			},

			HTMLChange{
				Query: ".checkout_shipping_cost",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(($order->shipping_cost ?? 0),2) ?>`,
			},

			HTMLChange{
				Query: ".checkout_total",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format(($order->total),2) ?>`,
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
