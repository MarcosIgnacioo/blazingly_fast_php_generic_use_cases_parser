package web_files_manipulation

var storeHeader = `
	<?php
	  include_once "../app/config.php"; 
	  include_once "../app/ProductsController.php";
	  include_once "../app/CategoriesController.php";

	  $productsController = new ProductsController(); 
	  $categoriesController = new CategoriesController(); 

	  $products = $productsController->getMostSold(12);
?>
`

var storeProducts = `
<?php if (isset($cafes) && count($cafes)): ?>
  <?php foreach (array_slice($cafes, 0, 4) as $cafe): ?>
    <?php $product = $productsController->getQuickView($cafe) ?>
    <div class="ed-element ed-container columns-box wv-overflow_visible product_item_all" id="ed-456730247">
      <div class="inner">
        <div class="ed-element ed-image" id="ed-456730250">
          <a href="<?= $product->url ?>">
						<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
          </a>
        </div>
        <div class="ed-element ed-text custom-theme" id="ed-456730253">
          <p style="line-height: 1.15; text-align: center;"><span style="font-size: 18px; font-family: 'daisywhl';" class="product_name">
              <a href="<?= $product->url ?>" title=""><?= $product->name ?></a></span></p>
          <p style="line-height: 1; text-align: center;"><span style="font-size: 14px; color: rgb(94, 94, 94); font-family: 'daisywhl';" class="product_price">DESDE $<?= $product->price ?></span></p>
        </div>
      </div>
    </div>
  <?php endforeach ?>
<?php endif ?>
`

var storeProductHref = `<?= $product->url ?>`
var storeProductName = `<?= $product->name ?>`
var storeProductIMG = `<?= $product->thumbnail_path ?>`
