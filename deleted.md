	// file, _ := os.ReadFile("./out/index.html")
	// htmlContent := string(file)
	// doc, err := node.ParseHTML(htmlContent)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// doc.Find(node.Descendant, node.Div)
	// insertingProductStoreIndex(doc)
	// os.WriteFile("popo.html", []byte(insertingIndexPHPHead(doc)), 0777)
	// poop := doc.Find(node.Descendant, node.Div, node.Class("separator_initial_cart"))
	// insertRawTextBeforeNode("<?php include \"layouts/cart_mobile.template.php\"; ?>", poop)
	// fmt.Println(gohtml.Format(doc.HTML()))
// creo que voy a tener que parsearlo de nuevo cada vez que haga un cambio asi para que se guarden XD
// no XD primero se hace todo lo que sea manipulacion del dom y al final ya se hacen las que sean concatenar textos asi directo a lo puerkito
// func insertingProductStoreIndex(doc node.Node) string {
// 	ogHtml := gohtml.Format(doc.HTML())
// 	prod := `
// <div class="ed-element ed-container columns-box wv-overflow_visible product_item_coffe" id="ed-456729794">
//   <div class="inner">
//     <div class="ed-element ed-image" id="ed-456729797">
//       <a href="<?= $product->url ?>">
// 				<img src="<?= $cafe->cover->full_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
// 			</a>
//     </div>
//     <div class="ed-element ed-text custom-theme" id="ed-456729800">
//       <p style="line-height: 1.15; text-align: center;">
//         <span style="font-size: 18px; font-family: &#34;daisywhl&#34;;">
//           <a href="/tienda/cafe/diablo" title="">
//             MEZCLA COYOTE
//           </a>
//         </span>
//       </p>
//       <p style="line-height: 1; text-align: center;">
//         <span style="font-size: 14px; color: rgb(94, 94, 94); font-family: &#34;daisywhl&#34;;">
//           DESDE $240
//         </span>
//       </p>
//     </div>
//   </div>
// </div>
// 	`
// 	productHtml := `
// 	<div class="ed-element ed-container columns-box wv-overflow_visible product_item_coffe" id="ed-456729794">
// 									<div class="inner">
// 	                        <div class="ed-element ed-image product_coffe_home_item" id="ed-456729788">
//
// 	                        </div>
// 	                        <div class="ed-element ed-text custom-theme product_coffe_home_item" id="ed-456729791">
// 	                          <p style="line-height: 1.15; text-align: center;">
// 	                            <span style="font-size: 18px; font-family: &#34;daisywhl&#34;;">
// 	                              <a href="<?= $product->url ?>"><?= $product->name ?></a>
// 	                            </span>
// 	                          </p>
// 	                          <p style="line-height: 1; text-align: center;">
// 	                            <span style="font-size: 14px; color: rgb(94, 94, 94); font-family: &#34;daisywhl&#34;;" class="product_price_home">
// 	                              DESDE $<?= $product->price ?>
// 	                            </span>
// 	                          </p>
// 	                        </div>
// 	                      </div>
// 	</div>
// 	`
// 	productComponent := `
// 	<?php if (isset($cafes) && count($cafes)): ?>
// 	        <?php foreach (array_slice($cafes, 0, 4) as $cafe): ?>
// 	            <?php $product = $productsController->getQuickView($cafe) ?>
// 	        <?php endforeach ?>
// 	    <?php endif ?>
// 	`
// 	products := doc.FindAll(node.Descendant, node.Div, node.Class("product_item_coffe"))
//
// 	for _, product := range products {
// 		fmt.Println(gohtml.Format(product.HTML()))
// 		fmt.Println("")
// 	}
//
// 	return "foo"
// }
