package web_files_manipulation

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
			ShouldRemoveAllChildren:           true,
			ShouldRemoveTagsWithSameClassName: true,
			Class:                             ".sizes_items_details select",
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
			Class:                   ".cover_product_details",
			InnerHTML: `
	<img
        src="<?= $productsController->getCover($grand_product) ?>"
        alt="<?= $grand_product->name ?>" 
    />
`,
		},
		Instruction{
			ShouldRemoveAllChildren:           true,
			ShouldRemoveTagsWithSameClassName: true,
			Class:                             `.product_description_details`,
			InnerHTML: `
				<?= $grand_product->description ?>
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
			ShouldRemoveAllChildren:           false,
			ShouldRemoveTagsWithSameClassName: true,
			Class:                             ".price_big_product",
			InnerHTML:                         "$<?= number_format( $productsController->getPrice($grand_product->presentations) ,2) ?>",
		},
		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".description_big_product",
			InnerHTML:               "<?= $grand_product->description ?>",
		},
		Instruction{
			ShouldRemoveAllChildren: true,
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
			ShouldRemoveAllChildren:           false,
			ShouldRemoveTagsWithSameClassName: true,
			Class:                             ".tallas_presentations",
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
			ForEach: `
			<?php if (isset($grand_product->related_products) && count($grand_product->related_products)): ?>
				<?php foreach ($grand_product->related_products as $product): ?>
					<?php 
							$product->categories = $grand_product->categories;
					?>
						<?php $product = $productsController->getQuickView($product) ?>
							%s
				<?php endforeach ?>
			<?php endif ?> 
			`,
			InnerHTML: `
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
			Class:            ".login_email",
			ShouldReplaceOld: true,
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
			Class:            ".login_password",
			ShouldReplaceOld: true,
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
			Class:            ".register_email",
			ShouldReplaceOld: true,
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
			Class:            ".register_password",
			ShouldReplaceOld: true,
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
			Class:            ".register_password_confirmation",
			ShouldReplaceOld: true,
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
			Class:            ".register_name",
			ShouldReplaceOld: true,
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
			Class:            ".register_phone",
			ShouldReplaceOld: true,
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

	"dashboard": {
		Instruction{
			Class:       "html",
			PrependHTML: dashBoardHeader,
		},
		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".header_welcome",
			InnerHTML:               `Hola <?= $_SESSION['name'] ?> (no eres <?= $_SESSION['name'] ?>? <a href="/clients?action=logout">Cerrar sesión</a> )`,
		},
	},
	"details": generateInputTagAttributes(detailsInputs, "input", detailsHeader),

	"addresses": {
		Instruction{
			Class:       "html",
			PrependHTML: addressesHeader,
		},
		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".address-data",
			InnerHTML: `
			<?php if (isset($addresses) && count($addresses)) : ?>
				<?php foreach ($addresses as $address) : ?>
						<div class="ed-element ed-container image-boxes-box wv-overflow_visible address-data" data-min-count="1" data-max-count="127" data-mappable="1" id="ed-456888664">
							<div class="inner">       
								<div class="ed-element ed-text custom-theme" data-min-count="0" data-max-count="1" data-mappable="1" id="ed-456888667">
									<p id="isPasted"><strong>Nombre:&nbsp;</strong></p>
									<p><span style="text-align: right; font-size: 1rem; letter-spacing: 0px;">
										<?= $address->first_name ?> <?= $address->last_name ?>
									</span></p>
									<p><strong>Calle:</strong><span style="white-space: pre;">&nbsp; &nbsp;&nbsp;</span></p>
									<p><?= $address->street_and_use_number ?></p>
									<p>
										<strong>Apartamento:</strong><span style="white-space: pre;"><strong>&nbsp; &nbsp;&nbsp;</strong></span>
									</p>
									<p><span style="white-space: pre;"><?= $address->apartment ?></span></p>
									<p>
										<strong>Provincia:</strong><span style="white-space: pre;"><strong>&nbsp;</strong> &nbsp;&nbsp;</span>
									</p>
									<p><?= $address->city ?> <?= $address->state ?></p>
									<p><strong>Código postal:</strong><span style="white-space: pre;">&nbsp; &nbsp;&nbsp;</span></p>
									<p><?= $address->postal_code ?> <?= $address->country->code ?></p>
									<p>
										<strong>Teléfono:</strong><span style="white-space: pre;"><strong>&nbsp;</strong> &nbsp;&nbsp;</span>
									</p>
													<p><span style="white-space: pre;"><?= $address->phone_number ?></span></p>
									<p><strong>Paí</strong><strong>s:</strong></p>
									<p><?= $address->references ?></p>
								</div>
							</div>
							<style>
								#ed-456888664 > .inner {
									padding: 10px;
								}
							</style>
						</div>
					<?php endforeach ?>
				<?php endif ?>
			`,
		},
	},
	"orders": {
		Instruction{
			Class:       "html",
			PrependHTML: ordersHeader,
		},
		Instruction{
			ShouldRemoveAllChildren: true,
			Class:                   ".table-text table tbody",
			InnerHTML: `
				<?php if (isset($orders) && count($orders)) : ?>
				<?php foreach ($orders as $order) : ?>
				<tr>
					<td style="width: 17.42%;"><div style="text-align: center;">#<?= $order->folio ?></div></td>
					<td style="width: 19.2363%;">
						<div style="text-align: center;"><?= (new DateTime($order->order_date))->format('F d, Y') ?> <span style="white-space: pre;" id="isPasted">&nbsp; &nbsp;&nbsp;</span>&nbsp;</div>
					</td>
					<td style="width: 19.1538%;"><div style="text-align: center;"><?= $order->order_status->name ?></div></td>
					<td style="width: 20%; text-align: center;">
						<div style="text-align: center;">
							$ <?= number_format($order->total,2) ?> <br>
							<small>
								por <?= count($order->presentations) ?> artículos
							</small>
						</div>
					</td>
					<td style="width: 24.1899%;">
						<div style="text-align: center;">
							<a href="<?= BASE_PATH ?>account/order/<?= $order->folio ?>/">
								Detalles
							</a>
						</div>
					</td>
				</tr>
				<?php endforeach ?>
				<?php endif ?>`,
		},
	},
	"carrito": {
		Instruction{
			Class:       "html",
			PrependHTML: cartHeader,
		},
		Instruction{
			Class:                              ".product_item_cart",
			ShouldAppendAttributes:             true,
			ShouldRemoveAllChildrenExceptFirst: true,
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag: "",
					Attr: Attribute{
						Name:  "class",
						Value: `item_<?= $product->id ?>_<?= str_replace(' ', '', $product->feature) ?>`,
					},
				},
			},
			InnerHtmlReplacements: newHTMLReplacements(cartStuff),
		},
		Instruction{
			Class:                   ".product_item_cart .cart_product_delete",
			ShouldAppendAttributes:  true,
			ShouldRemoveAllChildren: true,
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag:  "",
					Attr: CreateAttribute("onclick", "removeItemCart(this)"),
				},
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag:  "",
					Attr: CreateAttribute("data-id", "<?= $product->id ?>"),
				},
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag:  "",
					Attr: CreateAttribute("data-feature", "<?=$product->feature?>"),
				},
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag:  "",
					Attr: CreateAttribute("data-parent", `item_<?= $product->id ?>_<?= str_replace(" ", "", $product->feature) ?>`),
				},
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag:  "",
					Attr: CreateAttribute("data-parent", `item_<?= $product->id ?>_<?= str_replace(" ", "", $product->feature) ?>`),
				},
			},
		},
		Instruction{
			Class:                   ".cart_product_quantity input",
			ShouldAppendAttributes:  true,
			ShouldRemoveAllChildren: true,
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag:  "",
					Attr: CreateAttribute("value", `<?= $product->cantidad ?>`),
				},
			},
		},
		Instruction{
			Class:                   ".cart_product_description",
			ShouldRemoveAllChildren: true,
			InnerHTML: `
			<?= $product->feature ?>
			`,
		},
		Instruction{
			Class:                   ".cart_subtotal",
			ShouldRemoveAllChildren: true,
			InnerHTML: `
			$<?= number_format( $total ,2) ?>
			`,
		},
		Instruction{
			Class:                   ".cart_total",
			ShouldRemoveAllChildren: true,
			InnerHTML: `
				$<?= number_format( $total ,2) ?>
			`,
		},
	},
	"form-carrito": {
		Instruction{
			Class:       "html",
			PrependHTML: formCarritoHeader,
		},
		Instruction{
			Class:                              ".cart_item_row",
			ShouldRemoveAllChildrenExceptFirst: true,
			ForEach: `
				<?php $total = 0; ?>
					<?php if (isset($_SESSION['cart']) && count($_SESSION['cart'])): ?>
						<?php foreach ($_SESSION['cart'] as $product): ?> 
							<div class="ed-element ed-separator"><hr class="bg-primary" />
								<hr class="bg-primary" />
							</div>
							%s
						<?php $total += ($product->cantidad * $product->price) ?>
					<?php endforeach ?>
				<?php endif ?> 
			`,
			InnerHtmlReplacements: []HTMLReplacement{
				newHTMLReplacement("cart_product_name", `<?= $product->name ?>`),
				newHTMLReplacement("cart_product_price", `$<?= number_format($product->price * $product->cantidad, 2)?>`),
				newHTMLReplacement("cart_product_quantity", `$<?= number_format($product->price,2) ?> x <?= $product->cantidad ?> `),
				// no existen aun creo
				// newHTMLReplacement("checkout_subtotal", `$<?= number_format( $total ,2) ?> `),
				// newHTMLReplacement("checkout_total", `$<?= number_format( $total ,2) ?>`),
			},
		},
		Instruction{
			Class:                   `.submit_checkout`,
			ShouldRemoveAllChildren: true,
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag:  "",
					Attr: CreateAttribute("onclick", "document.querySelector('.checkout_form')?.dispatchEvent(new Event('submit', { cancelable: true }))"),
				},
			},
		},
		Instruction{
			Class: `.checkout_name`,
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag:  "input",
					Attr: CreateAttribute("name", "name"),
				},
			},
		},
		Instruction{
			Class: `.checkout_email`,
			TagsAttributes: []TagAttribute{
				TagAttribute{
					// with empty tag we use the Target div from above
					Tag:  "input",
					Attr: CreateAttribute("name", "email"),
				},
			},
		},
		CreateInstructionReplacementForAttributeOfTag(
			`checkout_phone`,
			`name="phone"`,
			`input`,
		),
		CreateInstructionReplacementForAttributeOfTag(
			`checkout_street_and_use_number`,
			`name="street_and_use_number"`,
			`input`,
		),
		CreateInstructionReplacementForAttributeOfTag(
			`checkout_street_and_use_number`,
			`name="street_and_use_number"`,
			`input`,
		),
		CreateInstructionReplacementForAttributeOfTag(
			`checkout_estate`,
			`name="estate"`,
			`select`,
		),
		CreateInstructionReplacementForAttributeOfTag(
			`checkout_city`,
			`name="city"`,
			`select`,
		),
		CreateInstructionReplacementForAttributeOfTag(
			`checkout_postal_code`,
			`name="postal_code"`,
			`input`,
		),
		CreateInstructionReplacementForAttributeOfTag(
			`checkout_references`,
			`name="references"`,
			`textarea`,
		),
	},
	"payment": {
		Instruction{
			Class:       "html",
			PrependHTML: paymentHeader,
		},
		// doesnt exist yet i think
		// Instruction{
		// 	ShouldRemoveAllChildren: true,
		// 	Class:                   ".order_folio span",
		// 	InnerHTML:               `ORDEN #<?= $_GET['folio'] ?>`,
		// },
		Instruction{
			Class:                             ".cart_item_row",
			ShouldRemoveTagsWithSameClassName: true,
			ForEach:                           paymentForEachWrapper,
			// TagsAttributes: []TagAttribute{
			// 	productAnchorHrefTagAtrr,
			// 	productAnchorTitleTagAtrr,
			// 	productImgAltTagAtrr,
			// },
			InnerHtmlReplacements: []HTMLReplacement{
				HTMLReplacement{
					ClassName: "cart_product_name",
					HTML: `
					<?php if(isset($presentation->product)): ?>
						<a href="<?= BASE_PATH ?>shop/details/<?= $presentation->product->slug; ?>/">
							<?= $presentation->product->name ?>  
						</a>
					<?php endif ?>  
`,
				},
				HTMLReplacement{
					ClassName: "cart_product_quantity",
					HTML: `
$<?= number_format(($presentation->pivot->price->amount ?? 0),2) ?> x <?= $presentation->pivot->quantity ?>
`,
				},
				HTMLReplacement{
					ClassName: "cart_product_price",
					HTML: `
$<?= number_format(($presentation->pivot->price->amount ?? 0) * $presentation->pivot->quantity,2) ?>
`,
				},
				// HTMLReplacement{
				// 	ClassName: `cart_product_description`,
				// 	HTML: `
				// 		<?= $presentation->pivot->comment ?>
				// 	`,
				// },
			},
		},
		Instruction{
			Class:                             ".total_item_cart",
			ShouldRemoveTagsWithSameClassName: true,
			InnerHtmlReplacements: []HTMLReplacement{
				HTMLReplacement{
					ClassName: "checkout_subtotal",
					HTML:      ` $<?= number_format(($order->total),2) ?> `,
				},
				HTMLReplacement{
					ClassName: "checkout_shipping_cost",
					HTML:      `$<?= number_format(($order->shipping_cost ?? 0),2) ?>`,
				},
				HTMLReplacement{
					ClassName: "checkout_total",
					HTML:      `$<?= number_format(($order->total),2) ?>`,
				},
			},
		},
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`card_number`,
		// 	`name="numero_tarjeta"`,
		// 	`input`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`card_exp_date`,
		// 	`name="fecha_exp"`,
		// 	`input`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`card_company`,
		// 	`name="marca_tarjeta"`,
		// 	`select`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`card_type`,
		// 	`name="tipo_tarjeta"`,
		// 	`select`,
		// ),
		// Instruction{
		// 	Class:                   ".card_company select",
		// 	ShouldRemoveAllChildren: true,
		// 	InnerHTML: `
		// 	<option value="" selected disabled>Selecciona una opción</option>
		// 	<option value="VISA">Visa</option>
		// 	<option value="MC">Mastercard</option>
		// 	`,
		// },
		// Instruction{
		// 	Class:                   ".card_type select",
		// 	ShouldRemoveAllChildren: true,
		// 	InnerHTML: `
		// 	<option value="" selected disabled>Selecciona una opción</option>
		// 	<option value="CR">Crédito</option>
		// 	<option value="DB">Débito</option>
		// 	`,
		// },
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_city`,
		// 	`name="ciudad"`,
		// 	`select`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_city`,
		// 	`value="<?= $order->address->city ?>"`,
		// 	`select`,
		// ),
		//
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_postal_code`,
		// 	`name="codigo_postal"`,
		// 	`select`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_postal_code`,
		// 	`value="<?= $order->address->postal_code ?>"`,
		// 	`select`,
		// ),
		//
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_first_name`,
		// 	`name="nombre"`,
		// 	`select`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_first_name`,
		// 	`value="<?= $order->address->first_name ?>"`,
		// 	`select`,
		// ),
		//
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_phone_number`,
		// 	`name="numero_celular"`,
		// 	`select`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_phone_number`,
		// 	`value="<?= $order->address->phone_number ?>"`,
		// 	`select`,
		// ),
		//
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_state`,
		// 	`name="estado"`,
		// 	`select`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_state`,
		// 	`value="<?= $order->address->state ?>"`,
		// 	`select`,
		// ),
		//
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_street_and_use_number`,
		// 	`name="calle"`,
		// 	`select`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_street_and_use_number`,
		// 	`value="<?= $order->address->street_and_use_number ?>`,
		// 	`select`,
		// ),
		//
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_last_name`,
		// 	`name="apellido"`,
		// 	`select`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_last_name`,
		// 	`value="<?= $order->address->last_name ?>"`,
		// 	`select`,
		// ),
		//
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_email`,
		// 	`name="correo"`,
		// 	`select`,
		// ),
		// CreateInstructionReplacementForAttributeOfTag(
		// 	`address_email`,
		// 	`value="<?= $order->client->email ?>"`,
		// 	`select`,
		// ),
	},
}

// CreateInstructionReplacementForAttributeOfTag(
// 	``,
// 	``,
// 	`input`,
// ),

var detailsInputs map[string][]string = map[string][]string{
	".details_email":            {"name", "email"},
	".details_name":             {"name", "name"},
	".details_lastname":         {"name", "lastname"},
	".details_phone":            {"name", "phone"},
	".details_password_current": {"name", "password_current"},
	".details_password_1":       {"name", "password_1"},
	".details_password_2":       {"name", "password_2"},
}

var cartStuff map[string]string = map[string]string{
	"cart_product_name":     "<?= $product->name ?>",
	"cart_product_price":    "$<?= number_format(floatval($product->price) ?? 0, 2)?>",
	"cart_product_subtotal": "$<?= number_format(floatval($product->cantidad * $product->price) ?? 0, 2) ?>",
}

// TagsAttributes: []TagAttribute{
// 	TagAttribute{
// 		Tag: "input",
// 		Attr: Attribute{
// 			Name:  "name",
// 			Value: "phone",
// 		},
// 	},
// },
