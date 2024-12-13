package web_files_manipulation

// # Home
var homeHeader = `
	<?php
	  include_once "app/config.php"; 
	  include_once "app/ProductsController.php"; 

	  $productsController = new ProductsController();   

	  $cafes = $productsController->getByCategory('cafe');
	  $merchs = $productsController->getByCategory('merch');
	?>
`

var productItemCoffeeForeachWrapper = `
<?php if (isset($cafes) && count($cafes)): ?>
	<?php foreach (array_slice($cafes, 0, 4) as $cafe): ?> 
		<?php $product = $productsController->getQuickView($cafe) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`

// favorite

var productItemFavoriteForeachWrapper = `
<?php if (isset($merchs) && count($merchs)): ?>
	<?php foreach (array_slice($merchs, 0, 3) as $merch): ?> 
		<?php $product = $productsController->getQuickView($merch) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`

var productItemCoffeCoverFullPath = `<?= $cafe->cover->full_path ?>`

// # Store
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
var productItemAllForeachWrapper = `
<?php if (isset($cafes) && count($cafes)): ?>
  <?php foreach (array_slice($cafes, 0, 4) as $cafe): ?>
    <?php $product = $productsController->getQuickView($cafe) ?>
				%s
  <?php endforeach ?>
<?php endif ?>
`

// origin
var originForeachWrapper = `
<?php if (isset($origenes) && count($origenes)): ?>
	<?php foreach ($origenes as $product): ?> 
		<?php $product = $productsController->getQuickView($product) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`

// author

var authorForeachWrapper = `
<?php if (isset($autores) && count($autores)): ?>
	<?php foreach ($autores as $product): ?> 
		<?php $product = $productsController->getQuickView($product) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`

// merch
var merchHeader = `
<?php

	  include_once "../../app/config.php"; 
	  include_once "../../app/ProductsController.php";
	  include_once "../../app/CategoriesController.php";

	  $productsController = new ProductsController(); 
	  
	  $products = $productsController -> getByCategory('merch');

	?>
`

// coffee

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

var coffeeMixesForeachWrapper = `
<?php if (isset($mezclas) && count($mezclas)): ?>
	<?php foreach ($mezclas as $product): ?> 
		<?php $product = $productsController->getQuickView($product) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`
