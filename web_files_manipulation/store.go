package web_files_manipulation

var storeHeader = `<?php
	  include_once "app/config.php"; 
	  include_once "app/ProductsController.php"; 

	  $productsController = new ProductsController();   

	  $cafes = $productsController -> getByCategory('cafe');
	  $merchs = $productsController -> getByCategory('merch');
	?>
`

var storeProducts = `
<?php if (isset($products) && count($products)): ?>
<?php foreach ($products as $product): ?> 
<?php $product = $productsController->getQuickView($product) ?>

Aquí va el código del item

<?php endforeach ?>
<?php endif ?>
`
