package web_files_manipulation

// store header
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

var productHref = `<?= $product->url ?>`
var productName = `<?= $product->name ?>`
var productThumbnail = `<?= $product->thumbnail_path ?>`
var productPrice = `<?= $product->price ?>`
var cafeCoverFullPath = `<?= $cafe->cover->full_path ?>`

// tienda
var storeProductHref = `<?= $product->url ?>`
var storeProductName = `<?= $product->name ?>`
var storeProductIMG = `<?= $product->thumbnail_path ?>`
var storeProductPrice = `<?= $product->price ?>`
var storeForeachWrapper = `
<?php if (isset($cafes) && count($cafes)): ?>
  <?php foreach (array_slice($cafes, 0, 4) as $cafe): ?>
    <?php $product = $productsController->getQuickView($cafe) ?>
				%s
  <?php endforeach ?>
<?php endif ?>
`

// origin
var originProductHref = `<?= $product->url ?>`
var originProductName = `<?= $product->name ?>"`
var originProductIMG = `<?= $product->thumbnail_path ?>`
var originProductPrice = `$<?= $product->price ?>`
var originForeachWrapper = `
<?php if (isset($origenes) && count($origenes)): ?>
	<?php foreach ($origenes as $product): ?> 
		<?php $product = $productsController->getQuickView($product) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`

// autores
var authorProductHref = `<?= $product->url ?>`
var authorProductName = `<?= $product->name ?>"`
var authorProductIMG = `<?= $product->thumbnail_path ?>`
var authorProductPrice = `$<?= $product->price ?>`
var authorForeachWrapper = `
<?php if (isset($autores) && count($autores)): ?>
	<?php foreach ($autores as $product): ?> 
		<?php $product = $productsController->getQuickView($product) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`
