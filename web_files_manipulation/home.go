package web_files_manipulation

var homeHeader = `
	<?php
	  include_once "app/config.php"; 
	  include_once "app/ProductsController.php"; 

	  $productsController = new ProductsController();   

	  $cafes = $productsController->getByCategory('cafe');
	  $merchs = $productsController->getByCategory('merch');
	?>
`

var homeProductHref = `<?= $product->url ?>`
var homeProductName = `<?= $product->name ?>`
var homeProductIMG = `<?= $cafe->cover->full_path ?>`
var homeProductPrice = `<?= $product->price ?>`
var homeForeachWrapper = `
<?php if (isset($cafes) && count($cafes)): ?>
	<?php foreach (array_slice($cafes, 0, 4) as $cafe): ?> 
		<?php $product = $productsController->getQuickView($cafe) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`

var homeProductFavHref = `<?= $product->url ?>`
var homeProductFavName = `<?= $product->name ?>`
var homeProductFavIMG = `<?= $product->thumbnail_path ?>`
var homeProductFavPrice = `<?= $product->price ?>`
var homeForeachFavWrapper = `
<?php if (isset($merchs) && count($merchs)): ?>
	<?php foreach (array_slice($merchs, 0, 3) as $merch): ?> 
		<?php $product = $productsController->getQuickView($merch) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`
