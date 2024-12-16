package web_files_manipulation

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
			ShouldRemoveAllChildren: true,
			Class:                   ".sizes_items_details select",
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
			ShouldRemoveAllChildren: true,
			Class:                   ".quantity_items_product select",
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
			Class:                   ".product_item_info",
			ShouldRemoveAllChildren: false,
			InnerHtmlReplacements: []HTMLReplacement{
				HTMLReplacement{
					ClassName: "product_title_details",
					HTML:      "<?= $grand_product->name ?>",
				},
				HTMLReplacement{
					ClassName: "product_price_details",
					HTML:      "$<?= number_format( $productsController->getPrice($grand_product->presentations) ,2) ?>",
				},
			},
		},
		Instruction{
			Class: ".background-image-holder",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "",
					Attr: Attribute{
						Name:  "style",
						Value: "background-image: 'url('<?= ($backgroud != '') ? $backgroud : '/images/0/9983852/Graficos_BolsasCafe_DIABLO.svg' ?>');'",
					},
				},
			},
		},
		Instruction{
			Class:                   ".molienda_container select",
			ShouldRemoveAllChildren: true,
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
			ShouldRemoveAllChildren: true,
			Class:                   ".boton_add_to_cart a",
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
			ShouldRemoveAllChildren: true,
			Class:                   ".tueste_item_details",
			InnerHTML: `
			<?php if ($tueste!=""): ?>
				<img src="<?= $tueste ?>" alt="" />
				<?php else: ?>
				<img src="/images/250/10441024/escala_tueste.png" alt="" />
			<?php endif ?>
			`,
		},
		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".sabor_item_details",
			InnerHTML: `
			<?php if ($perfil_de_sabor!=""): ?>
					<img src="<?= $perfil_de_sabor ?>" alt="" />
				<?php else: ?>
					<img src="/images/250/10441024/escala_tueste.png" alt="" />
			<?php endif ?>
			`,
		},
		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".ed-gallery-thumb",
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
			ShouldRemoveAllChildren: true,
			Class:                   ".product_item_recomendation",
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
		// TODO
		// Instruction{
		// 	Class:                   "body",
		// 	ShouldRemoveAllChildren: false,
		// 	AppendHTML:              `<script type="text/javascript"> document.getElementById('%s').onchange() </script>`,
		// },
		// Instruction{
		// 	Class:                   "head",
		// 	ShouldRemoveAllChildren: false,
		// 	AppendHTML:              `<script type="text/javascript"> var id_selection = '%s'; </script> `,
		// },
	},

	"sudadera": {
		Instruction{
			Class:       "html",
			PrependHTML: sweaterHeader,
		},
		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".name_big_product",
			InnerHTML:               "<?= $grand_product->name ?>",
		},
		Instruction{
			ShouldRemoveAllChildren: false,
			Class:                   ".price_big_product",
			InnerHTML:               "$<?= number_format( $productsController->getPrice($grand_product->presentations) ,2) ?>",
		},
		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".description_big_product",
			InnerHTML:               "<?= $grand_product->description ?>",
		},
		Instruction{
			ShouldRemoveAllChildren: false,
			Class:                   ".cover_product_details",
			InnerHTML: `<?php  
						$imagenes = '';

						if (isset($grand_product) && isset($grand_product->images)):
								foreach ($grand_product->images as $image): 
										 $imagenes .= ',{"image":"'.$image->full_path.'","title":""}';
								endforeach;
						endif;

						$imagenes = substr($imagenes, 1);
				?> 

				<div
						class="slider-container has-dots slick-dotted slick-initialized slick-slider"
						data-parameters='{"items":[<?= $imagenes ?>],"adaptiveHeight":true,"slidesToShow":1,"slidesToScroll":1,"rows":1,"slidesPerRow":1,"height":null,"animation":"slide","animationSpeed":"500ms","direction":"horizontal","autoplay":true,"autoplaySpeed":"5s","pauseOnHover":true,"loop":true,"nav":false,"dots":true,"enlarge":true,"retinaImages":true,"lazyLoad":"progressive","variableWidth":false,"centerMode":false,"centerPadding":"0px","asNavFor":"","responsive":[{"breakpoint":976,"settings":{"slidesToShow":1,"slidesToScroll":1,"centerPadding":"0px"}},{"breakpoint":576,"settings":{"slidesToShow":1,"slidesToScroll":1,"centerPadding":"0px"}}],"insideContainer":false}'
						style="max-width: 100%;"
				></div>`,
		},
		Instruction{
			ShouldRemoveAllChildren: false,
			Class:                   ".tallas_presentations",
			InnerHTML: `
	<?php if (isset($grand_product->presentations) && count($grand_product->presentations)): ?>
    <?php foreach ($grand_product->presentations as $presentation): ?>
    
    <div class="ed-element ed-container columns-box wv-overflow_visible" id="ed-456731279">
        <div class="inner">
            <div class="ed-element ed-button custom-theme" id="ed-456731273">
                <a
                    href="#"
                    onclick="updateSelection(this,'merch')"
                    data-id="<?= $presentation->id ?>" 
                    data-price="<?= $presentation->current_price->amount ?>"
                    data-product='<?= json_encode($productsController->getQuickView($grand_product,$presentation->id)) ?>'
                    class="button center button-variant-primary bg-background color-default border-color-primary bg-active-primary color-active-background border-color-active-primary presentations"
                >
                    <?= $presentation->description ?>
                </a>
            </div>
        </div>
    </div>

    <?php endforeach ?>
    <?php endif ?> 
`,
		},
		Instruction{
			// esta fakin shit se quita yu no se por q se duplica lam amada d abajo si le quito el ShouldRemoveAllChildren
			// <label for="form_456731219_ed-f-456731222">
			//                                 Cantidad
			//                               </label>
			ShouldRemoveAllChildren: true,
			Class:                   ".quantity_items_details select",
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
		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".button_add_big_product",
			TargetParent:            true,
			IsParent:                true,
			InnerHTML: `
								<?php if ($stock>0): ?>
														
												
												<a
														href="#"
														onclick="QuickAdd(this)" 
														class="button center button-variant-primary button-xlarge color-background bg-active-background color-active-default border-color-active-primary button_add_big_product"
														style="background-color: rgb(72, 79, 71);"
												>
														Agregar al carrito
												</a>
												<?php else: ?>

														<h3>
																Este producto se encuentra agotado
														</h3>

												<?php endif ?>
			`,
		},

		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".ed-gallery-items",
			// TargetParent:            true,
			// IsParent:                true,
			InnerHTML: `
					<?php if (isset($grand_product) && isset($grand_product->images)): ?>
							<?php foreach ($grand_product->images as $image): ?>  
							<?php if (!$image->is_cover): ?> 
									<li id="ed-gallery-item_9970970" class="ed-gallery-thumb" style="width: 800px; padding: 0px;">
											<a href="<?= $image->full_path ?>" target="_blank" title="Caption">
													<img
															src="<?= $image->full_path ?>"
															alt="Caption"
															srcset="
																	data:image/svg+xml,
																	%3Csvg%20width='800'%20viewBox='0%200%20800%20800'%20xmlns='http://www.w3.org/2000/svg'%3E%3Crect%20width='800'%20height='800'%20style='fill:%20%23F7F7F7'%20/%3E%3C/svg%3E
															"
															data-src="<?= $image->full_path ?>"
															class="ed-lazyload"
															data-srcset="<?= $image->full_path ?>"
													/>
											</a>
									</li>
								<?php endif ?>
							<?php endforeach ?> 
						<?php endif ?>  
			`,
		},

		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".product_item_recomendation",
			// TargetParent:            true,
			// IsParent:                true,
			InnerHTML: `
			<?php if (isset($grand_product->related_products) && count($grand_product->related_products)): ?>
				<?php foreach ($grand_product->related_products as $product): ?>
					<?php 
							$product->categories = $grand_product->categories;
					?>
					<?php $product = $productsController->getQuickView($product) ?>
						<div class="ed-element ed-container columns-box wv-overflow_visible product_item_recomendation" id="ed-456731237">
								<div class="inner">
										<div class="ed-element ed-image" id="ed-456731240">
												<a href="<?= $product->url ?>">
														<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>"  />
												</a>
										</div>
										<div class="ed-element ed-text custom-theme" id="ed-456731243">
												<p style="line-height: 1.15; text-align: center;">
														<span style="font-size: 18px; font-family: 'daisywhl';"><a href="<?= $product->url ?>" title=""><?= $product->name ?></a></span>
												</p>
												<p style="line-height: 1; text-align: center;"><span style="font-size: 14px; color: rgb(94, 94, 94); font-family: 'daisywhl';">DESDE $<?= $product->price ?></span></p>
										</div>
								</div>
						</div>
				<?php endforeach ?>
				<?php endif ?> 
			`,
		},

		// Instruction{
		// 	Class: "body",
		// 	AppendHTML: `
		// 	<input type="hidden" id="g-recaptcha-response" name="g-recaptcha-response">
		// 	<script src="https://unpkg.com/axios@1.1.2/dist/axios.min.js"></script>
		// 	<script src="<?= BASE_PATH ?>js/toasty/dist/toasty.min.js" type="text/javascript"></script>
		// 	<script src="<?= BASE_PATH ?>js/functions.js"></script>
		// 		<script src="<?= BASE_PATH ?>js/shop.js"></script>
		// 		<script src="<?= BASE_PATH ?>js/auth.js"></script>
		// 	<script src='https://www.google.com/recaptcha/api.js?render=<?= PUBLIC_KEY_GG ?>'></script>
		// 	<script>
		// 		grecaptcha.ready(function() {
		// 		grecaptcha.execute('<?= PUBLIC_KEY_GG ?>', {action: 'homepage'})
		// 			.then(function(token) {
		// 			document.getElementById('g-recaptcha-response').value=token;
		// 		});
		// 		});
		// 	// %s xd
		// 	</script>
		// 	`,
		// },

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

	},
	"login": {
		Instruction{
			Class:       "html",
			PrependHTML: loginHeader,
		},
		Instruction{
			Class: ".login_email",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "input",
					Attr: Attribute{
						Name:  "name",
						Value: "email",
					},
				},
			},
		},
		Instruction{
			Class: ".login_password",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "input",
					Attr: Attribute{
						Name:  "name",
						Value: "password",
					},
				},
			},
		},
		Instruction{
			Class: ".register_email",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "input",
					Attr: Attribute{
						Name:  "name",
						Value: "email",
					},
				},
			},
		},
		Instruction{
			Class: ".register_password",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "input",
					Attr: Attribute{
						Name:  "name",
						Value: "password",
					},
				},
				TagAttribute{
					Tag: "input",
					Attr: Attribute{
						Name:  "type",
						Value: "password",
					},
				},
			},
		},
		Instruction{
			Class: ".register_password_confirmation",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "input",
					Attr: Attribute{
						Name:  "name",
						Value: "password_confirmation",
					},
				},
				TagAttribute{
					Tag: "input",
					Attr: Attribute{
						Name:  "type",
						Value: "password",
					},
				},
			},
		},
		Instruction{
			Class: ".register_name",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "input",
					Attr: Attribute{
						Name:  "name",
						Value: "name",
					},
				},
			},
		},

		Instruction{
			Class: ".register_phone",
			TagsAttributes: []TagAttribute{
				TagAttribute{
					Tag: "input",
					Attr: Attribute{
						Name:  "name",
						Value: "phone",
					},
				},
			},
		},
	},
}
