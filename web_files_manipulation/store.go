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

var storeForeachWrapper = `
<?php if (isset($cafes) && count($cafes)): ?>
  <?php foreach (array_slice($cafes, 0, 4) as $cafe): ?>
    <?php $product = $productsController->getQuickView($cafe) ?>
				%s
  <?php endforeach ?>
<?php endif ?>
`
var storeProductHref = `<?= $product->url ?>`
var storeProductName = `<?= $product->name ?>`
var storeProductIMG = `<?= $product->thumbnail_path ?>`
var storeProductPrice = `<?= $product->price ?>`
