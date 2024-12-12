package web_files_manipulation

var homeHeader = `<?php
	  include_once "app/config.php"; 
	  include_once "app/ProductsController.php"; 

	  $productsController = new ProductsController();   

	  $cafes = $productsController -> getByCategory('cafe');
	  $merchs = $productsController -> getByCategory('merch');
	?>`

var homeProducts = `<?php if (isset($cafes) && count($cafes)): ?>
        <?php foreach (array_slice($cafes, 0, 4) as $cafe): ?> 
            <?php $product = $productsController->getQuickView($cafe) ?>
                    <div class="inner">
                        <div class="ed-element ed-image product_coffe_home_item" id="ed-456729788">
                        <a href="<?= $product->url ?>">
                            <img src="<?= $cafe->cover->full_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
                          </a>
                        </div>
                        <div class="ed-element ed-text custom-theme product_coffe_home_item" id="ed-456729791">
                          <p style="line-height: 1.15; text-align: center;">
                            <span style="font-size: 18px; font-family: 'daisywhl';">
                              <a href="<?= $product->url ?>"><?= $product->name ?></a>
                            </span>
                          </p>
                          <p style="line-height: 1; text-align: center;">
                            <span style="font-size: 14px; color: rgb(94, 94, 94); font-family: 'daisywhl';" class="product_price_home">
                              DESDE $<?= $product->price ?>
                            </span>
                          </p>
                        </div>
                      </div>
        <?php endforeach ?>
    <?php endif ?>
`
