// COMPLETED
package web_files_manipulation

var devilModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: devilHeader,
	},
	Modification{
		Query:     `.ed-gallery-thumb`,
		SelectAll: true,
		PrependHTML: `
	 <?php if (isset($grand_product) && isset($grand_product->images)): ?>
			<?php foreach ($grand_product->images as $image): ?>  
        <?php if (!$image->is_cover): ?> 
		`,
		AppendHTML: `
					<?php endif ?>
				<?php endforeach ?> 
			<?php endif ?>
		`,
		InnerHTML: `
			<a href="<?= $image->full_path ?>" target="_blank" title="Caption">
						<img
								src="<?= $image->full_path ?>"
								alt="Caption"
								srcset="
										data:image/svg+xml,
										%3Csvg%20width='800'%20viewBox='0%200%20800%20800'%20xmlns='http://www.w3.org/2000/svg'%3E%3Crect%20width='800'%20height='800'%20style='fill:%20%23F7F7F7'%20/%3E%3C/svg%3E
								"
								data-src="<?= $image->full_path ?>"
								class="ed-lazyload lazyload"
								data-srcset="<?= $image->full_path ?>"
						/>
				</a>
		`,
	},

	Modification{
		Query:     `.product_item_recomendation`,
		SelectAll: true,
		PrependHTML: `
		<?php if (isset($grand_product->related_products) && count($grand_product->related_products)): ?>
			<?php foreach ($grand_product->related_products as $product): ?>
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
	},
	Modification{
		Query: "",
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: ".cover_product_details",
				Mode:  INNER_HTML,
				HTML:  `<img src="<?= $productsController->getCover($grand_product) ?>" alt="<?= $grand_product->name ?>"/>`,
			},
			HTMLChange{
				Query: ".product_title_details",
				Mode:  INNER_HTML,
				HTML:  `<?= $grand_product->name ?>`,
			},
			HTMLChange{
				Query: ".product_price_details",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format( $productsController->getPrice($grand_product->presentations) ,2) ?>`,
			},
			HTMLChange{
				Query: ".product_description_details",
				Mode:  INNER_HTML,
				HTML:  `<?= $grand_product->description ?>`,
			},

			HTMLChange{
				Query: ".sizes_items_details select",
				Mode:  INNER_HTML,
				HTML: `
				<?php if (isset($grand_product->presentations) && count($grand_product->presentations)): ?>
					<?php foreach ($grand_product->presentations as $presentation): ?> 
						<option data-id="<?= $presentation->id ?>" data-price="<?= $presentation->current_price->amount ?>" value="<?= $presentation->id ?>" data-product='<?= json_encode($productsController->getQuickView($grand_product,$presentation->id)) ?>' >
								<?= $presentation->description ?>
						</option> 
					<?php endforeach ?>
        <?php endif ?>`,
			},
			HTMLChange{
				Query: ".quantity_items_product select",
				Mode:  INNER_HTML,
				HTML: `
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
			HTMLChange{
				Query: ".molienda_container select",
				Mode:  INNER_HTML,
				HTML: `
					<?php foreach ($features as $feature): ?>
							<option value="<?= $feature  ?>">
									<?= $feature  ?>
							</option>
					<?php endforeach ?>
        `,
			},
			HTMLChange{
				Query: ".tueste_item_details",
				Mode:  INNER_HTML,
				HTML: `
				<?php if ($tueste!=""): ?>
					<img src="<?= $tueste ?>" alt="" />
				<?php else: ?>
					<img src="/images/250/10441024/escala_tueste.png" alt="" />
        <?php endif ?> 
        `,
			},
			HTMLChange{
				Query: ".sabor_item_details",
				Mode:  INNER_HTML,
				HTML: `
				<?php if ($perfil_de_sabor!=""): ?>
					<img src="<?= $perfil_de_sabor ?>" alt="" />
        <?php else: ?>
					<img src="/images/250/10441024/escala_tueste.png" alt="" />
        <?php endif ?> 
        `,
			},
			HTMLChange{
				Query: ".item_recomendation_name",
				Mode:  OUTER_HTML,
				HTML: `
				<a href="<?= $product->url ?>">
							<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>"  />
				</a>
				<a href="<?= $product->url ?>" title=""class="item_recomendation_name" title=""><?= $product->name ?></a>
        `,
			},
			HTMLChange{
				Query: ".item_recomendation_price",
				Mode:  INNER_HTML,
				HTML: `
					DESDE $<?= $product->price ?>
        `,
			},
		},

		AttributesChanges: []AttributeChange{
			AttributeChange{
				Query: ".background-image-holder",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "style",
					Value: `background-image: 'url('<?= ($backgroud != '') ? $backgroud : '/images/0/9983852/Graficos_BolsasCafe_DIABLO.svg' ?>');'`,
				},
			},
			AttributeChange{
				Query: ".sizes_items_details select",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "onchange",
					Value: `updateSelection(this, 'cafe')`,
				},
			},
			AttributeChange{
				Query: ".quantity_items_product select",
				Mode:  APPEND_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "class",
					Value: `cantidad_producto`,
				},
			},
			AttributeChange{
				Query: ".molienda_container select",
				Mode:  APPEND_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "class",
					Value: `feature_product`,
				},
			},
			AttributeChange{
				Query: ".boton_add_to_cart a",
				Mode:  REPLACE_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "onclick",
					Value: `QuickAdd(this)`,
				},
			},
		},
	},

	Modification{
		Query: `BODY`,
		AppendHTML: `
		<script type="text/javascript">
				document.getElementById('%s').onchange()
		</script>
		`,
	},
	Modification{
		Query: `HEAD`,
		AppendHTML: `
		<script type="text/javascript">
			var id_selection = '%s';
		</script>
		`,
	},
}
