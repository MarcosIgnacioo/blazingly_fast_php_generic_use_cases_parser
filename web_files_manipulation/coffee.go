package web_files_manipulation

// cafe

var coffeeHeader = `
<?php
	  include_once "../../app/config.php"; 
	  include_once "../../app/ProductsController.php";
	  include_once "../../app/CategoriesController.php";

	  $productsController = new ProductsController(); 
	  $categoriesController = new CategoriesController(); 

	  $mezclas = $productsController -> getByCategory('mezclas');
	  $autores = $productsController -> getByCategory('autor');
	  $origenes = $productsController -> getByCategory('origen');

	  $merchs = $productsController -> getByCategory('merch');
	?>
`
var coffeeProductHref = `<?= $product->url ?>`
var coffeeProductName = `<?= $product->name ?>`
var coffeeProductIMG = `<?= $cafe->cover->full_path ?>`
var coffeeProductPrice = `<?= $product->price ?>`
var coffeeForeachWrapper = `
<?php if (isset($mezclas) && count($mezclas)): ?>
	<?php foreach ($mezclas as $product): ?> 
		<?php $product = $productsController->getQuickView($product) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`
