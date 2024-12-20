package web_files_manipulation

var sweaterModifications = []Modification{
	Modification{
		Query:       "html",
		PrependHTML: sweaterHeader,
	},
	Modification{
		Query: "",
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: ".name_big_product",
				Mode:  INNER_HTML,
				HTML:  `<?= $grand_product->name ?>`,
			},
			HTMLChange{
				Query: ".price_big_product",
				Mode:  INNER_HTML,
				HTML:  `$<?= number_format( $productsController->getPrice($grand_product->presentations) ,2) ?>`,
			},
			HTMLChange{
				Query: ".description_big_product",
				Mode:  INNER_HTML,
				HTML:  `<?= $grand_product->description ?>`,
			},
			HTMLChange{
				Query: ".cover_product_details",
				Mode:  INNER_HTML,
				HTML: `
					<?php  
							$imagenes = '';

							if (isset($grand_product) && isset($grand_product->images)):
									foreach ($grand_product->images as $image): 
											 $imagenes .= ',{"image":"'.$image->full_path.'","title":""}';
									endforeach;
							endif;

							$imagenes = substr($imagenes, 1);
					?> 
					<div class="slider-container has-dots slick-dotted slick-initialized slick-slider" data-parameters='{"items":[<?= $imagenes ?>],"adaptiveHeight":true,"slidesToShow":1,"slidesToScroll":1,"rows":1,"slidesPerRow":1,"height":null,"animation":"slide","animationSpeed":"500ms","direction":"horizontal","autoplay":true,"autoplaySpeed":"5s","pauseOnHover":true,"loop":true,"nav":false,"dots":true,"enlarge":true,"retinaImages":true,"lazyLoad":"progressive","variableWidth":false,"centerMode":false,"centerPadding":"0px","asNavFor":"","responsive":[{"breakpoint":976,"settings":{"slidesToShow":1,"slidesToScroll":1,"centerPadding":"0px"}},{"breakpoint":576,"settings":{"slidesToShow":1,"slidesToScroll":1,"centerPadding":"0px"}}],"insideContainer":false}' style="max-width: 100%;"
					></div>
				`,
			},
		},
	},
	Modification{
		Query: ".tallas_presentations",
		// DeleteSiblings: true,

		PrependHTML: `
			<?php if (isset($grand_product->presentations) && count($grand_product->presentations)): ?>
				<?php foreach ($grand_product->presentations as $presentation): ?>
		`,
		AppendHTML: `
			<?php endforeach ?>
    <?php endif ?>
		`,
		AttributesChanges: []AttributeChange{
			NewAttributeChange(
				`a`,
				APPEND_ATTRIBUTE,
				`class="presentations"`,
			),
			NewAttributeChange(
				`a`,
				REPLACE_ATTRIBUTE,
				`onclick="updateSelection(this,'merch')"`,
			),
			NewAttributeChange(
				`a`,
				REPLACE_ATTRIBUTE,
				`data-id="<?= $presentation->id ?>"`,
			),
			NewAttributeChange(
				`a`,
				REPLACE_ATTRIBUTE,
				`data-price="<?= $presentation->current_price->amount ?>"`,
			),
			NewAttributeChange(
				`a`,
				REPLACE_ATTRIBUTE,
				`data-product="<?= json_encode($productsController->getQuickView($grand_product,$presentation->id)) ?>"`,
			),
		},
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: "a",
				Mode:  INNER_HTML,
				HTML:  `<?= $presentation->description ?>`,
			},
		},
	},
	Modification{
		Query: ".quantity_items_details",
		AttributesChanges: []AttributeChange{
			AttributeChange{
				Query: "select",
				Mode:  APPEND_ATTRIBUTE,
				Attribute: Attribute{
					Name:  "class",
					Value: `cantidad_producto`,
				},
			},
		},
		HTMLChanges: []HTMLChange{
			HTMLChange{
				Query: "select",
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
		},
	},
	Modification{
		Query: ".button_add_big_product",
		PrependHTML: `
		<?php if ($stock>0): ?>
		`,
		AppendHTML: `
    <?php else: ?>
        <h3>
            Este producto se encuentra agotado
        </h3>
    <?php endif ?>
		`,
		AttributesChanges: []AttributeChange{
			NewAttributeChange(
				``,
				REPLACE_ATTRIBUTE,
				`onclick="QuickAdd(this)"`,
			),
		},
	},
	Modification{
		Query: ".ed-gallery-items li",
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
		AttributesChanges: []AttributeChange{
			NewAttributeChange(
				`a`,
				REPLACE_ATTRIBUTE,
				`href="<?= $image->full_path ?>"`,
			),
			NewAttributeChange(
				`img`,
				REPLACE_ATTRIBUTE,
				`src="<?= $image->full_path ?>"`,
			),

			NewAttributeChange(
				`img`,
				REPLACE_ATTRIBUTE,
				`data-src="<?= $image->full_path ?>"`,
			),

			NewAttributeChange(
				`img`,
				REPLACE_ATTRIBUTE,
				`data-srcset="<?= $image->full_path ?>"`,
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
