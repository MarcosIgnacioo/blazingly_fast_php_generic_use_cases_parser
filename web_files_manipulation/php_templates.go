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
<?php if (isset($products) && count($products)): ?>
	<?php foreach ($products as $product): ?> 
		<?php $product = $productsController->getQuickView($product) ?>
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

// # Coffee

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

// # Merch

var merchHeader = `
<?php

	  include_once "../../app/config.php"; 
	  include_once "../../app/ProductsController.php";
	  include_once "../../app/CategoriesController.php";

	  $productsController = new ProductsController(); 
	  
	  $products = $productsController -> getByCategory('merch');

	?>
`
var merchForeachWrapper = `
<?php if (isset($products) && count($products)): ?>
	<?php foreach ($products as $product): ?> 
		<?php $product = $productsController->getQuickView($product) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`

// # Accesories
var accessoriesHeader = `
<?php

	  include_once "../../app/config.php"; 
	  include_once "../../app/ProductsController.php";
	  include_once "../../app/CategoriesController.php";

	  $productsController = new ProductsController(); 
	  
	  $products = $productsController -> getByCategory('accesorios');

	?>
`
var accesoriesForeachWrapper = `
<?php if (isset($products) && count($products)): ?>
	<?php foreach ($products as $product): ?> 
		<?php $product = $productsController->getQuickView($product) ?>
			%s
	<?php endforeach ?>
<?php endif ?>
`

// # devil

var devilHeader = `
<?php

  include_once "../../../app/config.php"; 
  include_once "../../../app/ProductsController.php";
  include_once "../../../app/CategoriesController.php";

  $productsController = new ProductsController(); 
  $categoriesController = new CategoriesController(); 

  if(isset($_GET['slug'])) {

    $grand_product = $productsController -> getBySlug($_GET['slug']); 

  }else{
    header("Location:".$_SERVER['HTTP_REFERER']);
  }

  if (is_null($grand_product)) {
    header("Location:".$_SERVER['HTTP_REFERER']);
  } 

  $features = $productsController ->getFeatures($grand_product,"molienda");
  $backgroud = $productsController ->getFile($grand_product,"background_".$_GET['slug']);
  $stock = $productsController ->getStock($grand_product->presentations);;


  $tueste = $productsController ->getFile($grand_product,"tueste_".$_GET['slug']);

  $perfil_de_sabor = $productsController ->getFile($grand_product,"perfil_de_sabor_".$_GET['slug']); 

?>
`

// $ sudadera

var sweaterHeader = `
 	<?php
 	    
 	  include_once "../../../app/config.php"; 
 	  include_once "../../../app/ProductsController.php";
 	  include_once "../../../app/CategoriesController.php";

 	  $productsController = new ProductsController(); 
 	  $categoriesController = new CategoriesController(); 

 	  if(isset($_GET['slug'])) {

 	    $grand_product = $productsController -> getBySlug($_GET['slug']); 

 	  }else{
 	    header("Location:".$_SERVER['HTTP_REFERER']);
 	  }

 	  if (is_null($grand_product)) {
 	    header("Location:".$_SERVER['HTTP_REFERER']);
 	  } 

 	  $stock = $productsController ->getStock($grand_product->presentations);;

 	?>
`

// # login

var loginHeader = `
	<?php
		include_once "../app/config.php";    

		if (isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'account/dashboard');
		}
	?>
	`

// # dashboard
var dashBoardHeader = `
	<?php
		include_once "../../app/config.php";    

		if (!isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'login/');
		}
	?>
`

// # details
var detailsHeader = `
	<?php
		include_once "../../app/config.php";    

		if (!isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'login');
		}
	?>
`

// # addresseds

var addressesHeader = `
	<?php
		include_once "../../app/config.php";    
		include_once "../../app/AuthController.php";

		$authController = new AuthController();

		$authController = new AuthController();

		if (!isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'login/');
		}

		$addresses = $authController->getAddress();

	?>
`

var ordersHeader = `
	<?php
		include_once "../../app/config.php";    
		include_once "../../app/AuthController.php";

		$authController = new AuthController();
		if (!isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'login/');
		}

		$orders = array_reverse($authController->getOrders($_SESSION['client_id']));
	?>
	`

var cartHeader = `
	<?php
		include_once "../app/config.php"; 
		include_once "../app/ProductsController.php";

		$productsController = new ProductsController(); 
	?>
	`

var formCarritoHeader = `
	<?php
		include_once "../app/config.php"; 
		include_once "../app/ProductsController.php";
		$productsController = new ProductsController(); 
	?>
	`

var paymentHeader = `
	<?php 
		include_once "../app/config.php";
		include_once "../app/ProductsController.php";
		include_once "../app/ShopController.php";

		if (!isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'login/');
		}

		$shopController = new ShopController();
		$productsController = new ProductsController();

		$order = $shopController->getDataOrder($_GET['folio']); 
	?>
`

var paymentForEachWrapper = `
<?php if (isset($order->presentations) && count($order->presentations)): ?>
		<?php foreach ($order->presentations as $presentation): ?>                 
			%s
		<div class="ed-element ed-separator"><hr class="bg-primary" /></div>
	<?php endforeach ?>
<?php endif ?> 
	`
