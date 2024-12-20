# doce40 shop store

app[x]
tienda[x]
    cafe[x]
        diablo[x]
    merch[x]
        sudadera[]
    accesorios[x]
login[]
dashboard[]
addresses[]
orders[]
carrito[]
form-carrito[]
payment[]


FALTA eliminar la funcion del archivo de js 

### Paso 1: aplicar a la lista de carpetas: 

- Desminificar todos los archivos
- Cambiar la extención de .html a .php

### Paso 2: 

- Crear una carpeta en la raíz con el nombre  ``` app  ```
- Descomprimir el zip app dentro de la carpeta

### Paso 3: 

- En la carpeta js de la raiz descromprimir el zip js dentro de el 

### Paso 4: 

- En todos los archivos ubicar el item con la clase ``` separator_initial_cart  ``` 

    y sobre ese item colocar el código siguiente: 

    ``` 
     
                                    
	<?php include "layouts/cart_mobile.template.php"; ?>

	 
	``` 

	para los archivos de un nivel más profundo colocar un extra ../

### Paso 4:

- En el index.php de la raiz colocar el siguiente código en la línea principal: 
	```
	<?php

	  include_once "app/config.php"; 
	  include_once "app/ProductsController.php"; 

	  $productsController = new ProductsController();   

	  $cafes = $productsController -> getByCategory('cafe');
	  $merchs = $productsController -> getByCategory('merch');

	?>
	```
- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` product_item_coffe  ```
- Se deben borrar o comentar los últimos 3 items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
	```

	<?php if (isset($cafes) && count($cafes)): ?>
    <?php foreach (array_slice($cafes, 0, 4) as $cafe): ?> 
    <?php $product = $productsController->getQuickView($cafe) ?>

    Aquí va el código del item

    <?php endforeach ?>
    <?php endif ?>
	```
- Dentro del item 
	```
	<a href="<?= $product->url ?>">
	```
- Dentro del item hay que ubicar la imagen y reemplazar el código por el siguiente: 
	```
	<img src="<?= $cafe->cover->full_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
	```
- Dentro del item 
	```
	<a href="<?= $product->url ?>"><?= $product->name ?></a>
	```
- Dentro del item
	```
	DESDE $<?= $product->price ?>
	```

- Lo mismo pasa con la clase ``` product_item_favorite  ```
- Se deben borrar o comentar los últimos 2 items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
	```
	<?php if (isset($merchs) && count($merchs)): ?>
    <?php foreach (array_slice($merchs, 0, 3) as $merch): ?> 
    <?php $product = $productsController->getQuickView($merch) ?>

    Aquí va el código del item

	<?php endforeach ?>
    <?php endif ?>

    ```
- Dentro del item
	```
    <a href="<?= $product->url ?>">
 	<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>"  />
 	```

- Dentro del item
	```
 	<a href="<?= $product->url ?>"><?= $product->name ?></a>
 	```
- Dentro del item
	```
 	<?= $product->price ?>
 	```

### Paso 5:
- En la carpeta tienda/index.php colocar el siguiente código en la cabecera 
	```
	<?php
		include_once "../app/config.php"; 
		include_once "../app/ProductsController.php";
		include_once "../app/CategoriesController.php";

		$productsController = new ProductsController(); 
		$categoriesController = new CategoriesController(); 

		$products = $productsController -> getMostSold(12);
	?>
	```

- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` product_item_all  ```
- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
	```
	<?php if (isset($products) && count($products)): ?>
    <?php foreach ($products as $product): ?> 
    <?php $product = $productsController->getQuickView($product) ?>

    Aquí va el código del item

    <?php endforeach ?>
    <?php endif ?>
	```

- Dentro del item 
	```
	<a href="<?= $product->url ?>" >
	```
- Dentro del item hay que ubicar la imagen y reemplazar el código por el siguiente: 
	```
	<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
	```
- Dentro del item 
	```
	<a href="<?= $product->url ?>" title=""><?= $product->name ?></a>
	```
- Dentro del item
	```
	DESDE $<?= $product->price ?>
	```

### Paso 6: 

- En la carpeta tienda/cafe/index.php colocar el siguiente código en la cabecera 
	```
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
	```

- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` product_item_coffee_mezclas  ```

- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
	```

	<?php if (isset($mezclas) && count($mezclas)): ?>
    <?php foreach ($mezclas as $product): ?> 
    <?php $product = $productsController->getQuickView($product) ?>

    Aquí va el código del item

    <?php endforeach ?>
    <?php endif ?>
	```


- Dentro del item 
	```
	<a href="<?= $product->url ?>" >
	```
- Dentro del item hay que ubicar la imagen y reemplazar el código por el siguiente: 
	```
	<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
	```
- Dentro del item 
	```
	<a href="<?= $product->url ?>" title=""><span style="font-family: 'daisywhl';"><?= $product->name ?></span></a>
	```
- Dentro del item
	```
	DESDE $<?= $product->price ?>
	```



- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` product_item_coffee_autor  ```

- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
```

<?php if (isset($autores) && count($autores)): ?>
<?php foreach ($autores as $product): ?> 
<?php $product = $productsController->getQuickView($product) ?>

Aquí va el código del item

<?php endforeach ?>
<?php endif ?>
```


- Dentro del item 
	```
	<a href="<?= $product->url ?>" >
	```
- Dentro del item hay que ubicar la imagen y reemplazar el código por el siguiente: 
	```
	<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
	```
- Dentro del item 
	```
	<a href="<?= $product->url ?>" title=""><?= $product->name ?></a>
	```
- Dentro del item
	```
	DESDE $<?= $product->price ?>
	```



- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` product_item_coffee_origen  ```

- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
```

<?php if (isset($origenes) && count($origenes)): ?>
<?php foreach ($origenes as $product): ?> 
<?php $product = $productsController->getQuickView($product) ?>

Aquí va el código del item

<?php endforeach ?>
<?php endif ?>
```


- Dentro del item 
	```
	<a href="<?= $product->url ?>" >
	```
- Dentro del item hay que ubicar la imagen y reemplazar el código por el siguiente: 
	```
	<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
	```
- Dentro del item 
	```
	<a href="<?= $product->url ?>" title=""><?= $product->name ?></a>
	```
- Dentro del item
	```
	DESDE $<?= $product->price ?>
	```

- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` product_item_favorite  ```

- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
```

<?php if (isset($merchs) && count($merchs)): ?>
<?php foreach ($merchs as $product): ?> 
<?php $product = $productsController->getQuickView($product) ?>

Aquí va el código del item

<?php endforeach ?>
<?php endif ?>
```


- Dentro del item 
	```
	<a href="<?= $product->url ?>" >
	```
- Dentro del item hay que ubicar la imagen y reemplazar el código por el siguiente: 
	```
	<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
	```
- Dentro del item 
	```
	<a href="<?= $product->url ?>" title=""><?= $product->name ?></a>
	```
- Dentro del item
	```
	$<?= $product->price ?>
	```


### Paso 7: 

- En la carpeta tienda/merch/index.php colocar el siguiente código en la cabecera 
	```
	<?php

		  include_once "../../app/config.php"; 
		  include_once "../../app/ProductsController.php";
		  include_once "../../app/CategoriesController.php";

		  $productsController = new ProductsController(); 
		  
		  $products = $productsController -> getByCategory('merch');

		?>
	```

- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` product_item_merch  ```

- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
```

<?php if (isset($products) && count($products)): ?>
<?php foreach ($products as $product): ?> 
<?php $product = $productsController->getQuickView($product) ?>

Aquí va el código del item

<?php endforeach ?>
<?php endif ?>
```


- Dentro del item 
	```
	<a href="<?= $product->url ?>" >
	```
- Dentro del item hay que ubicar la imagen y reemplazar el código por el siguiente: 
	```
	<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
	```
- Dentro del item 
	```
	<a href="<?= $product->url ?>" title=""> <span style="font-family: 'daisywhl';"><?= $product->name ?></span></a>
	```
- Dentro del item
	```
	DESDE $<?= $product->price ?>
	```

### Paso 8: 

- En la carpeta tienda/accesorios/index.php colocar el siguiente código en la cabecera 
	```
	<?php

		  include_once "../../app/config.php"; 
		  include_once "../../app/ProductsController.php";
		  include_once "../../app/CategoriesController.php";

		  $productsController = new ProductsController(); 
		  
		  $products = $productsController -> getByCategory('accesorios');

		?>
	```

- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` product_item_accessories  ```

- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
```

<?php if (isset($products) && count($products)): ?>
<?php foreach ($products as $product): ?> 
<?php $product = $productsController->getQuickView($product) ?>

Aquí va el código del item

<?php endforeach ?>
<?php endif ?>
```


- Dentro del item 
	```
	<a href="<?= $product->url ?>" >
	```
- Dentro del item hay que ubicar la imagen y reemplazar el código por el siguiente: 
	```
	<img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>" title="<?= $product->name ?>" />
	```
- Dentro del item 
	```
	<a href="<?= $product->url ?>" title=""><?= $product->name ?></a>
	```
- Dentro del item
	```
	DESDE $<?= $product->price ?>
	```

### Paso 9:

- En la carpeta tienda/cafe/diablo/index.php colocar el siguiente código en la cabecera 
	```
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
	```
- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` cover_product_details  ``` este contiene la portada del item a reemplazar, reemplazar el código contenido por el siguiente: 

	```
	<img
        src="<?= $productsController->getCover($grand_product) ?>"
        alt="<?= $grand_product->name ?>" 
    />
	```

- Dentro del mismo archivo ubicar el item con la clase ``` background-image-holder```  

- Ubicar ``` product_title_details```  
- Ubicar ``` product_price_details```  
- Ubicar ``` product_description_details```  

- Dentro de ``` sizes_items_details``` ubicar el select
 EN EL PHP original viene con un segundo argumento esta funcion
	colocarle onchange="updateSelection(this)"
	eliminar los options
	y colocar el siguiete código

	<?php if (isset($grand_product->presentations) && count($grand_product->presentations)): ?>
        <?php foreach ($grand_product->presentations as $presentation): ?> 
            <option data-id="<?= $presentation->id ?>" data-price="<?= $presentation->current_price->amount ?>" value="<?= $presentation->id ?>" data-product='<?= json_encode($productsController->getQuickView($grand_product,$presentation->id)) ?>' >
                <?= $presentation->description ?>
            </option> 
        <?php endforeach ?>
        <?php endif ?> 

    importate: este item tiene un id similar al siguiente: form_456731144_ed-f-456731147 reserva ese id porque se necesitará 

- Dentro de quantity_items_product ubicar el select 
	colocar la clase class="cantidad_producto" al select

	eliminar los options 
	y colocar el siguiente código

	<?php if (isset($stock) && $stock > 0): ?>
        <?php for ($i = 1; $i < $stock; $i++): ?>
            <option value="<?= $i ?>">
                <?= $i ?>
            </option>
            <?php if ($i > 9) { break; } ?>
        <?php endfor; ?>
    <?php endif; ?>

- Dentro del item molienda_container ubicar el select 
	colocarle la clase al select class="feature_product" 
	eliminar los options
	y colocar el siguiente código 

	<?php foreach ($features as $feature): ?>
                                                            
        <option value="<?= $feature  ?>">
            <?= $feature  ?>
        </option>
    
    <?php endforeach ?>

- Dentro del item buscar el objeto con la clase boton_add_to_cart
	añadir el siguiente código como atributo
	onclick="QuickAdd(this)"
	data-product="<?= json_encode($productsController->getQuickView($grand_product)) ?>"


- Ubicar el item con la clase tueste_item_details
	eliminar el código interno y colocar el siguiente 

	<?php if ($tueste!=""): ?>

        <img src="<?= $tueste ?>" alt="" />

        <?php else: ?>

        <img src="/images/250/10441024/escala_tueste.png" alt="" />

        <?php endif ?> 

- Ubicar el item con la clase sabor_item_details 
	eliminar el código interno y colocar el siguiente 

	<?php if ($perfil_de_sabor!=""): ?>

        <img src="<?= $perfil_de_sabor ?>" alt="" />

        <?php else: ?>

        <img src="/images/250/10441024/escala_tueste.png" alt="" />

        <?php endif ?> 


- Buscar la galeria, son 4 items con la clase ed-gallery-thumb 
	borrar todos menos 1 y colocar el código sobre y debajo siguiente 

	 <?php if (isset($grand_product) && isset($grand_product->images)): ?>
        <?php foreach ($grand_product->images as $image): ?>  
        <?php if (!$image->is_cover): ?> 

         
        
        <?php endif ?>
        <?php endforeach ?> 
        <?php endif ?>


     dentro del li reemplazar el código interno por lo siguiente: 

     <a href="<?= $image->full_path ?>" target="_blank" title="Caption">
        <img
            src="<?= $image->full_path ?>"
            alt="Caption"
            srcset="
                data:image/svg+xml,
                %3Csvg%20width='800'%20viewBox='0%200%20800%20800'%20xmlns='http://www.w3.org/2000/svg'%3E%3Crect%20width='800'%20height='800'%20style='fill:%20%23F7F7F7'%20/%3E%3C/svg%3E
            "
            data-src="<?= $image->full_path ?>"
            class="ed-lazyload lazyload"
            data-srcset="<?= $image->full_path ?>"
        />
    </a> 

- ubicar también: product_item_recomendation 
	-  eliminar todos dejando el primero 

	<?php if (isset($grand_product->related_products) && count($grand_product->related_products)): ?>
    <?php foreach ($grand_product->related_products as $product): ?>
    <?php $product = $productsController->getQuickView($product) ?>
    <?php endforeach ?>
    <?php endif ?> 

	dentro del item contiene una clase item_recomendation_name colocar el siguiente código 

	<a href="<?= $product->url ?>">
        <img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>"  />
    </a>

    <a href="<?= $product->url ?>" title=""class="item_recomendation_name" title=""><?= $product->name ?></a>

	dentro del item contiene una clase item_recomendation_price colocar el siguietne código 

	DESDE $<?= $product->price ?>

- antes del cierre del body añadir el siguiente código:

	importante el item id es la variable reserva en pases anteriores: ejemplo form_456731144_ed-f-456731147

	<script type="text/javascript">

      document.getElementById('form_456731144_ed-f-456731147').onchange()
      
    </script>

- antes del cierre del head 

	importante el id es el reservado en pasos anteriores: ejemplo: form_456731144_ed-f-456731147
	<script type="text/javascript">
            
            var id_selection = 'form_456731144_ed-f-456731147';
            
        </script>

### Paso 10:

- En la raiz del proyecto crear una carpeta LAYOUTS

	- dentro de la carpeta layouts crear un archivo cart_mobile.template.php 
	y colocar el código

	```
	<?php 
	  if (!isset($_SESSION)) {  
	    include_once "../app/config.php";
	  }
	?>


    <?php $total = 0; ?>
    <?php if (isset($_SESSION['cart']) && count($_SESSION['cart'])): ?> 
    <?php foreach ($_SESSION['cart'] as $product): ?>
    <span id="cart_container_mobile"></span>

    aquí va el código

    <?php $total += ($product->cantidad*$product->price) ?> 
    <?php endforeach ?>
    <?php endif ?>

    ``` 

    - ubicar el carrito ```product_item_cart``` en el archivo index.php principal
    - reservar un item más el separador que está sobre el para complementar el código anterior 

    - la imagen del item debería verse así: 

    <img
        src="<?= $product->cover ?>"
        alt="<?= $product->name ?>"
        data-src="<?= $product->cover ?>"
        class="ed-lazyload"
        style="object-fit: cover;"
        srcset="<?= $product->cover ?>"
        data-srcset="<?= $product->cover ?>"
    />

    - ubicar el item con la clase ```resumen_carrito_nav``` se debe tomar ese item y se debe colocar dentro del layout debajo del ciclo, quedando de la siguiente manera: 

    ```
    <div class="ed-element ed-separator remove_item_on_update" id="ed-456730127"><hr class="bg-primary" /></div>
    <div class="ed-element ed-container wv-boxed wv-spacer preset-columns-two-v2-subtotal-carrito resumen_carrito_nav remove_item_on_update" id="ed-456730142">
        <div class="inner">
            <div class="ed-element ed-container columns-box wv-overflow_visible" id="ed-456730145">
                <div class="inner">
                    <div class="ed-element ed-text custom-theme" id="ed-456730130">
                        <p style="text-align: left;"><span style="font-family: 'Roboto';">Subtotal</span></p>
                    </div>
                </div>
            </div>
            <div class="ed-element ed-container columns-box wv-overflow_visible total_item_cart" id="ed-456730148">
                <div class="inner">
                    <div class="ed-element ed-text custom-theme" id="ed-456730151">
                        <p style="text-align: right;"><span style="font-family: 'Roboto';" class="total_item_carrito_nav">$1,740</span></p>
                    </div>
                </div>
            </div>
        </div>
    </div>
    ```

    - por último se deben eliminar todos los archivos index los items con la clase: ```remove_item_on_update```

### Paso 9:

- En la carpeta tienda/merch/sudadera/index.php colocar el siguiente código en la cabecera 

	```
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
```

- Dentro del mismo archivo index ubicar el item que contenga la clase ``` name_big_product  ``` 
	reemplazar el texto con el código siguiente
	``` 
	<?= $grand_product->name ?> 
	``` 

- Dentro del mismo archivo index ubicar el item que contenga la clase ``` price_big_product  ``` 
	reemplazar el texto con el código siguiente
	```
	$<?= number_format( $productsController->getPrice($grand_product->presentations) ,2) ?>
	```

- Dentro del mismo archivo index ubicar el item que contenga la clase ``` description_big_product  ``` 
	reemplazar el texto con el código siguiente
	```
	<?= $grand_product->description ?>
	``` 

- Dentro del item ``` cover_product_details ``` se debe ubicar el div con la configuración del sliders 
	
	```

	<?php  
	    $imagenes = '';

	    if (isset($grand_product) && isset($grand_product->images)):
	        foreach ($grand_product->images as $image): 
	             $imagenes .= ',{"image":"'.$image->full_path.'","title":""}';
	        endforeach;
	    endif;

	    $imagenes = substr($imagenes, 1);
	?> 

	<div
	    class="slider-container has-dots slick-dotted slick-initialized slick-slider"
	    data-parameters='{"items":[<?= $imagenes ?>],"adaptiveHeight":true,"slidesToShow":1,"slidesToScroll":1,"rows":1,"slidesPerRow":1,"height":null,"animation":"slide","animationSpeed":"500ms","direction":"horizontal","autoplay":true,"autoplaySpeed":"5s","pauseOnHover":true,"loop":true,"nav":false,"dots":true,"enlarge":true,"retinaImages":true,"lazyLoad":"progressive","variableWidth":false,"centerMode":false,"centerPadding":"0px","asNavFor":"","responsive":[{"breakpoint":976,"settings":{"slidesToShow":1,"slidesToScroll":1,"centerPadding":"0px"}},{"breakpoint":576,"settings":{"slidesToShow":1,"slidesToScroll":1,"centerPadding":"0px"}}],"insideContainer":false}'
	    style="max-width: 100%;"
	></div>

	```
- Dentro del item ``` talla_presnentations ``` se debe tomar el primer item y recorrer las presentaciones 
	
	```
	<?php if (isset($grand_product->presentations) && count($grand_product->presentations)): ?>
    <?php foreach ($grand_product->presentations as $presentation): ?>
    
    <div class="ed-element ed-container columns-box wv-overflow_visible" id="ed-456731279">
        <div class="inner">
            <div class="ed-element ed-button custom-theme" id="ed-456731273">

                

                <a
                    href="#"
                    onclick="updateSelection(this,'merch')"
                    data-id="<?= $presentation->id ?>" 
                    data-price="<?= $presentation->current_price->amount ?>"
                    data-product='<?= json_encode($productsController->getQuickView($grand_product,$presentation->id)) ?>'
                    class="button center button-variant-primary bg-background color-default border-color-primary bg-active-primary color-active-background border-color-active-primary presentations"
                >
                    <?= $presentation->description ?>
                </a>

                

            </div>
        </div>
    </div>

    <?php endforeach ?>
    <?php endif ?> 
	```

	como se puede observar al item anterior se le debe añadir la clase ```presentations```

- Dentro del item ```quantity_items_details``` buscar el select y añadir la clase ```cantidad_producto```
	y añadir este código dentro del select 

	```
	<?php if (isset($stock) && $stock > 0): ?>
        <?php for ($i = 1; $i < $stock; $i++): ?>
            <option value="<?= $i ?>">
                <?= $i ?>
            </option>
            <?php if ($i > 9) { break; } ?>
        <?php endfor; ?>
    <?php endif; ?> 
	```

- Dentro del item ```button_add_big_product```

	añadir el atributo ```onclick="QuickAdd(this)"```

	<a
        href="#"
        onclick="QuickAdd(this)" 
        class="button center button-variant-primary button-xlarge color-background bg-active-background color-active-default border-color-active-primary button_add_big_product"
        style="background-color: rgb(72, 79, 71);"
    >
        Agregar al carrito
    </a>
    <?php else: ?>

        <h3>
            Este producto se encuentra agotado
        </h3>

    <?php endif ?>

- Dentro del item ```ed-gallery-items```
	
	```
	<li id="ed-gallery-item_9970970" class="ed-gallery-thumb" style="width: 800px; padding: 0px;">
        <a href="<?= $image->full_path ?>" target="_blank" title="Caption">
            <img
                src="<?= $image->full_path ?>"
                alt="Caption"
                srcset="
                    data:image/svg+xml,
                    %3Csvg%20width='800'%20viewBox='0%200%20800%20800'%20xmlns='http://www.w3.org/2000/svg'%3E%3Crect%20width='800'%20height='800'%20style='fill:%20%23F7F7F7'%20/%3E%3C/svg%3E
                "
                data-src="<?= $image->full_path ?>"
                class="ed-lazyload"
                data-srcset="<?= $image->full_path ?>"
            />
        </a>
    </li>
    ```

- Buscar los item con la clase ```product_item_recomendation``` 
	
	```
	<?php if (isset($grand_product->related_products) && count($grand_product->related_products)): ?>
    <?php foreach ($grand_product->related_products as $product): ?>
    <?php 
        $product->categories = $grand_product->categories;
    ?>
    <?php $product = $productsController->getQuickView($product) ?>

    <div class="ed-element ed-container columns-box wv-overflow_visible product_item_recomendation" id="ed-456731237">
        <div class="inner">
            <div class="ed-element ed-image" id="ed-456731240">
                <a href="<?= $product->url ?>">
                    <img src="<?= $product->thumbnail_path ?>" alt="<?= $product->name ?>"  />
                </a>
            </div>
            <div class="ed-element ed-text custom-theme" id="ed-456731243">
                <p style="line-height: 1.15; text-align: center;">
                    <span style="font-size: 18px; font-family: 'daisywhl';"><a href="<?= $product->url ?>" title=""><?= $product->name ?></a></span>
                </p>
                <p style="line-height: 1; text-align: center;"><span style="font-size: 14px; color: rgb(94, 94, 94); font-family: 'daisywhl';">DESDE $<?= $product->price ?></span></p>
            </div>
        </div>
    </div>

    <?php endforeach ?>
    <?php endif ?> 
    ```

### Paso 2: 
	
	En todos los archivos colocar el siguiente código justo antes del cierre del </body>
	
	Scripts
	```
	<input type="hidden" id="g-recaptcha-response" name="g-recaptcha-response">

	<script src="https://unpkg.com/axios@1.1.2/dist/axios.min.js"></script>
	<script src="<?= BASE_PATH ?>js/toasty/dist/toasty.min.js" type="text/javascript"></script>
	<script src="<?= BASE_PATH ?>js/functions.js"></script>
    <script src="<?= BASE_PATH ?>js/shop.js"></script>
    <script src="<?= BASE_PATH ?>js/auth.js"></script>
	
	<script src='https://www.google.com/recaptcha/api.js?render=<?= PUBLIC_KEY_GG ?>'></script>
	<script>
		grecaptcha.ready(function() {
		grecaptcha.execute('<?= PUBLIC_KEY_GG ?>', {action: 'homepage'})
			.then(function(token) {
			document.getElementById('g-recaptcha-response').value=token;
		});
		});
	</script>
	```

### Paso 2: 
	
	En todos los archivos colocar dentro de la etiqueta <head></head> el siguiente código:

	```
	<link href="<?= BASE_PATH ?>js/toasty/dist/toasty.min.css" rel="stylesheet">

	<script type="text/javascript">
        var global_token = '<?= $_SESSION['token'] ?>';
        var global_url = '<?= BASE_PATH ?>';
        var le = <?= (DEVELOP_MODE) ?>;
        var lg = <?= (isset($_SESSION['client_id']) && $_SESSION['client_id'] != "") ? 'true' : 'false' ?>; 
    </script> 
	```

### Paso 2: 

	```
	Código del carrito
	```


### Paso 9:

- En la carpeta login/index.php colocar el siguiente código en la cabecera 
	```
	<?php
		include_once "../app/config.php";    

		if (isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'account/dashboard');
		}
	?>
	```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` login_email  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="email"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` login_password  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="password"  ``` y el type a ``` type="password" ```


- Dentro del mismo archivo index ubicar el item que contiene la clase ``` register_email  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="email"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` register_password  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="password"  ``` y el type a ``` type="password" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` register_password_confirmation  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="password_confirmation"  ``` y el type a ``` type="password" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` register_name  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="name"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` register_phone  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="phone"  ```

-Al final del mismo archivo index agregar lo siguiente, antes del ``` </body> ```
	```
        <script type="text/javascript">
            document.querySelector('.login_form').addEventListener('submit', validateAccess);
            document.querySelector('.register_form').addEventListener('submit', validateRegister);
        </script>

	```

### Paso 10: 

- Dentro del archivo index de las carpetas [account/dashboard,  account/details, account/addresses, account/orders] ubicar el item que tiene un href hacia el login ``` href="/login"  ```

- Remplazar el href con 
	```
		href="/clients?action=logout"
	```
### Paso 11: 

- En la carpeta account/dashboard/index.php colocar el siguiente código en la cabecera 
	```
	<?php
		include_once "../../app/config.php";    

		if (!isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'login/');
		}
	?>
	```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` header_welcome  ```

- Remplazar el interior con 
	```
	Hola <?= $_SESSION['name'] ?> (no eres <?= $_SESSION['name'] ?>? <a href="/clients?action=logout">Cerrar sesión</a> )
	```

### Paso 12:

- En la carpeta account/details/index.php colocar el siguiente código en la cabecera 
	```
	<?php
		include_once "../../app/config.php";    

		if (!isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'login');
		}
	?>
	```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` details_email  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="email"  ``` y el value a ``` value="<?= $_SESSION['email'] ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` details_name  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="name"  ``` y el value a ``` value="<?= $_SESSION['name'] ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` details_lastname  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="lastname"  ``` y el value a ``` value="<?= $_SESSION['lastname'] ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` details_phone  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="phone"  ``` y el value a ``` value="<?= $_SESSION['phone'] ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` details_password_current  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="password_current"  ``` y el type a ``` type="password" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` details_password_1  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="password_1"  ``` y el type a ``` type="password" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` details_password_2  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="password_2"  ``` y el type a ``` type="password" ```

-Al final del mismo archivo index agregar lo siguiente, antes del ``` </body> ```
	```
        <script type="text/javascript">
            document.querySelector('.details_edit_form').addEventListener('submit', validateUpdate);
            document.querySelector('.password_form').addEventListener('submit', validateChangePassword);
        </script>

	```

### Paso 13: 

- En la carpeta account/addresses/index.php colocar el siguiente código en la cabecera 
	```
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
	```

- Dentro del mismo archivo index ubicar los items que contienen la clase ``` address-data  ```

- Eliminar los elementos y en su lugar agregar lo siguiente 
	```
	<?php if (isset($addresses) && count($addresses)) : ?>
	<?php foreach ($addresses as $address) : ?>
	<div class="ed-element ed-container image-boxes-box wv-overflow_visible address-data" data-min-count="1" data-max-count="127" data-mappable="1" id="ed-456888664">
		<div class="inner">       
			<div class="ed-element ed-text custom-theme" data-min-count="0" data-max-count="1" data-mappable="1" id="ed-456888667">
				<p id="isPasted"><strong>Nombre:&nbsp;</strong></p>
				<p><span style="text-align: right; font-size: 1rem; letter-spacing: 0px;">
					<?= $address->first_name ?> <?= $address->last_name ?>
				</span></p>
				<p><strong>Calle:</strong><span style="white-space: pre;">&nbsp; &nbsp;&nbsp;</span></p>
				<p><?= $address->street_and_use_number ?></p>
				<p>
					<strong>Apartamento:</strong><span style="white-space: pre;"><strong>&nbsp; &nbsp;&nbsp;</strong></span>
				</p>
				<p><span style="white-space: pre;"><?= $address->apartment ?></span></p>
				<p>
					<strong>Provincia:</strong><span style="white-space: pre;"><strong>&nbsp;</strong> &nbsp;&nbsp;</span>
				</p>
				<p><?= $address->city ?> <?= $address->state ?></p>
				<p><strong>Código postal:</strong><span style="white-space: pre;">&nbsp; &nbsp;&nbsp;</span></p>
				<p><?= $address->postal_code ?> <?= $address->country->code ?></p>
				<p>
					<strong>Teléfono:</strong><span style="white-space: pre;"><strong>&nbsp;</strong> &nbsp;&nbsp;</span>
				</p>
                <p><span style="white-space: pre;"><?= $address->phone_number ?></span></p>
				<p><strong>Paí</strong><strong>s:</strong></p>
				<p><?= $address->references ?></p>
			</div>
		</div>
		<style>
			#ed-456888664 > .inner {
				padding: 10px;
			}
		</style>
	</div>
	<?php endforeach ?>
	<?php endif ?>
	```


### Paso 14: 

- En la carpeta account/orders/index.php colocar el siguiente código en la cabecera 
	```
	<?php
		include_once "../../app/config.php";    
		include_once "../../app/AuthController.php";

		$authController = new AuthController();
		if (!isset($_SESSION['client_id'])) {
			header("Location:" . BASE_PATH . 'login/');
		}

		$orders = array_reverse($authController->getOrders($_SESSION['client_id']));
	?>
	```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` table-text  ```

- Dentro del item anterior, hay un elemento ``` <table> ``` remplazar el interior del ``` <tbody> ``` con 
	```
	<?php if (isset($orders) && count($orders)) : ?>
	<?php foreach ($orders as $order) : ?>
	<tr>
		<td style="width: 17.42%;"><div style="text-align: center;">#<?= $order->folio ?></div></td>
		<td style="width: 19.2363%;">
			<div style="text-align: center;"><?= (new DateTime($order->order_date))->format('F d, Y') ?> <span style="white-space: pre;" id="isPasted">&nbsp; &nbsp;&nbsp;</span>&nbsp;</div>
		</td>
		<td style="width: 19.1538%;"><div style="text-align: center;"><?= $order->order_status->name ?></div></td>
		<td style="width: 20%; text-align: center;">
			<div style="text-align: center;">
				$ <?= number_format($order->total,2) ?> <br>
				<small>
					por <?= count($order->presentations) ?> artículos
				</small>
			</div>
		</td>
		<td style="width: 24.1899%;">
			<div style="text-align: center;">
				<a href="<?= BASE_PATH ?>account/order/<?= $order->folio ?>/">
					Detalles
				</a>
			</div>
		</td>
	</tr>
	<?php endforeach ?>
	<?php endif ?>
	```

### Paso 15:
- En la carpeta carrito/index.php colocar el siguiente código en la cabecera 

```
	<?php
		include_once "../app/config.php"; 
		include_once "../app/ProductsController.php";

		$productsController = new ProductsController(); 
	?>
```
- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` product_item_cart  ``` agregarle esto en las clases ``` item_<?= $product->id ?>_<?= str_replace(" ", "", $product->feature) ?> ```

- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
```

<?php $total = 0; ?>
<?php if (isset($_SESSION['cart']) && count($_SESSION['cart'])): ?>
<?php foreach ($_SESSION['cart'] as $product): ?> 

Aquí va el código del item

<div class="ed-element ed-separator item_<?= $product->id ?>_<?= str_replace(" ", "", $product->feature) ?>"><hr class="bg-primary" /></div>

<?php $total += ($product->cantidad * $product->price) ?>
<?php endforeach ?>
<?php endif ?> 
```

- Dentro del item buscar la imagen con clase ``` cart_product_img  ``` y reemplazarla con 

```
<img
	src="<?= $product->cover ?>"
	alt=""
	class="ed-lazyload cart_product_img"
	style="object-fit: cover;"
/>
```

- Dentro del item buscar el span con clase ``` cart_product_name  ``` y reemplazar el contenido con 

```
	<?= $product->name ?>
```

- Dentro del item buscar el span con clase ``` cart_product_price  ``` y reemplazar el contenido con 

```
$<?= number_format(floatval($product->price) ?? 0, 2)?> 
```

- Dentro del item buscar el span con clase ``` cart_product_subtotal  ``` y reemplazar el contenido con 

```
$<?= number_format(floatval($product->cantidad * $product->price) ?? 0, 2) ?>
```

- Dentro del item buscar el div con clase ``` cart_product_quantity  ``` y dentro hay un input agregarle el atributo value 

```
value="<?= $product->cantidad ?>"
```

- Dentro del item buscar el span con clase ``` cart_product_delete  ``` y agregarle los siguientes atributos 

```
onclick="removeItemCart(this)" data-id="<?= $product->id ?>" data-feature="<?=$product->feature?>" data-parent="item_<?= $product->id ?>_<?= str_replace(" ", "", $product->feature) ?>" role="button"
```

- Dentro del item buscar el span con clase ``` cart_product_description  ``` y reemplazar el contenido con 
```
<?= $product->feature ?>
```

-  Buscar el span con clase ``` cart_subtotal  ``` y reemplazar el contenido con 
```
$<?= number_format( $total ,2) ?>
```

-  Buscar el span con clase ``` cart_total  ``` y reemplazar el contenido con 
```
$<?= number_format( $total ,2) ?>
```

### Paso 16:
- En la carpeta form-carrito/index.php colocar el siguiente código en la cabecera 

```
	<?php
		include_once "../app/config.php"; 
		include_once "../app/ProductsController.php";

		$productsController = new ProductsController(); 
	?>
```
- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` cart_item_row  ```

- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
```

<?php $total = 0; ?>
<?php if (isset($_SESSION['cart']) && count($_SESSION['cart'])): ?>
<?php foreach ($_SESSION['cart'] as $product): ?> 

Aquí va el código del item

<div class="ed-element ed-separator"><hr class="bg-primary" /></div>

<?php $total += ($product->cantidad * $product->price) ?>
<?php endforeach ?>
<?php endif ?> 
```

- Dentro del item buscar el span con clase ``` cart_product_name  ``` y reemplazar el contenido con 

```
<?= $product->name ?>
```

- Dentro del item buscar el span con clase ``` cart_product_description  ``` y reemplazar el contenido con 
```
<?= $product->feature ?>
```

- Dentro del item buscar el span con clase ``` cart_product_price  ``` y reemplazar el contenido con 

```
$<?= number_format($product->price * $product->cantidad, 2)?> 
```

- Dentro del item buscar el span con clase ``` cart_product_quantity  ``` y reemplazar el contenido con 

```
$<?= number_format($product->price,2) ?> x <?= $product->cantidad ?>
```

- Buscar el span con clase ``` checkout_subtotal  ``` y reemplazar el contenido con 
```
$<?= number_format( $total ,2) ?>
```

- Buscar el span con clase ``` checkout_total  ``` y reemplazar el contenido con 
```
$<?= number_format( $total ,2) ?>
```

- Buscar los botones con clase ``` submit_checkout  ``` y agregarles esto
	```
	onclick="document.querySelector('.checkout_form')?.dispatchEvent(new Event('submit', { cancelable: true }))"
	```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` checkout_name  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="name"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` checkout_email  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="email"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` checkout_phone  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="phone"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` checkout_street_and_use_number  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="street_and_use_number"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` checkout_estate  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="estate"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` checkout_city  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="city"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` checkout_postal_code  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="postal_code"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` checkout_references  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="references"  ```

-Al final del mismo archivo index agregar lo siguiente, antes del ``` </body> ```
	```
	    <script src="https://cdn.jsdelivr.net/npm/sweetalert2@11"></script>
        <script type="text/javascript">
            document.querySelector('.checkout_form').addEventListener('submit', validateOrder);
        </script>

	```


### Paso 16:
- En la carpeta payment/index.php colocar el siguiente código en la cabecera 

```
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

```

- Buscar el span con clase ``` order_folio  ``` y reemplazar el contenido con 
```
ORDEN #<?= $_GET['folio'] ?>
```

- Dentro del mismo archivo index ubicar el primer item que contenga la clase ``` cart_item_row  ```

- Se deben borrar o comentar todos los items y dejar solo 1, el cual se repetirá en bucle a través del ciclo foreach con el siguiente código: 
```
<?php if (isset($order->presentations) && count($order->presentations)): ?>
<?php foreach ($order->presentations as $presentation): ?>                 

Aquí va el código del item

<div class="ed-element ed-separator"><hr class="bg-primary" /></div>
<?php endforeach ?>
<?php endif ?> 
```

- Dentro del item buscar el span con clase ``` cart_product_name  ``` y reemplazar el contenido con 

```
<?php if(isset($presentation->product)): ?>
	<a href="<?= BASE_PATH ?>shop/details/<?= $presentation->product->slug; ?>/">
		<?= $presentation->product->name ?>  
	</a>
<?php endif ?>  
```

- Dentro del item buscar el span con clase ``` cart_product_description  ``` y reemplazar el contenido con 
```
<?= $presentation->pivot->comment ?>
```

- Dentro del item buscar el span con clase ``` cart_product_price  ``` y reemplazar el contenido con 

```
$<?= number_format(($presentation->pivot->price->amount ?? 0) * $presentation->pivot->quantity,2) ?>
```

- Dentro del item buscar el span con clase ``` cart_product_quantity  ``` y reemplazar el contenido con 

```
$<?= number_format(($presentation->pivot->price->amount ?? 0),2) ?> x <?= $presentation->pivot->quantity ?>
```

- Buscar el span con clase ``` checkout_subtotal  ``` y reemplazar el contenido con 
```
$<?= number_format(($order->total),2) ?>
```

- Buscar el span con clase ``` checkout_shipping_cost  ``` y reemplazar el contenido con 
```
$<?= number_format(($order->shipping_cost ?? 0),2) ?>
```

- Buscar el span con clase ``` checkout_total  ``` y reemplazar el contenido con 
```
$<?= number_format(($order->total),2) ?>
```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` card_number  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="numero_tarjeta"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` card_exp_date  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="fecha_exp"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` card_cvv  ```, dentro de ese elemento ubicar el input y asignar el atributo name a  ``` name="codigo_seguridad"  ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` card_company  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="marca_tarjeta"  ``` y reemplazar las opciones con

```
	<option value="" selected disabled>Selecciona una opción</option>
	<option value="VISA">Visa</option>
	<option value="MC">Mastercard</option>
```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` card_type  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="tipo_tarjeta"  ``` y reemplazar las opciones con

```
	<option value="" selected disabled>Selecciona una opción</option>
	<option value="CR">Crédito</option>
	<option value="DB">Débito</option>
```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` address_city  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="ciudad"  ``` y asignar el atributo value a ``` value="<?= $order->address->city ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` address_postal_code  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="codigo_postal"  ``` y asignar el atributo value a ``` value="<?= $order->address->postal_code ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` address_first_name  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="nombre"  ``` y asignar el atributo value a ``` value="<?= $order->address->first_name ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` address_phone_number  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="numero_celular"  ``` y asignar el atributo value a ``` value="<?= $order->address->phone_number ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` address_state  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="estado"  ``` y asignar el atributo value a ``` value="<?= $order->address->state ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` address_street_and_use_number  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="calle"  ``` y asignar el atributo value a ``` value="<?= $order->address->street_and_use_number ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` address_last_name  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="apellido"  ``` y asignar el atributo value a ``` value="<?= $order->address->last_name ?>" ```

- Dentro del mismo archivo index ubicar el item que contiene la clase ``` address_email  ```, dentro de ese elemento ubicar el select y asignar el atributo name a  ``` name="correo"  ``` y asignar el atributo value a ``` value="<?= $order->client->email ?>" ```

-Al final del mismo archivo index agregar lo siguiente, antes del ``` </body> ```
	```
        <script type="text/javascript">
			const order_folio = '<?= $_GET['folio'] ?>'
            document.querySelector('.payment_form').addEventListener('submit', validatePayment);
        </script>

	```

### Home


- Seleccionar "product_item_coffe" y tomar al padre contenedor, borrar el contenido y recorrer el arreglo de productos

- Seleccionar "product_item_favorite" y tomar al padre contenedor, borrar el contenido y recorrer el arreglo de productos 

### Tienda

- Seleccionar "product_item_all" y tomar al padre contenedor, borrar el contenido y recorrer el arreglo de productos

### Tienda > Café

- Seleccionar "product_item_coffee" y tomar al padre contenedor, borrar el contenido y recorrer el arreglo de productos


### Tienda > category > detalle de producto

	```
	onclick="QuickAdd(this)"
    data-product=""
    id="buttonAdd"
	```

## app.bundle
En la carpeta webcard/static/app.bundle.(este número cambia con cada compilacion).js buscar la siguiente función 

```
submit:function(){var t=this;if("undefined"!=typeof cms&&cms)return!1;if(!this.ajax)return this.$form.submit(),this;if(!this.isSubmitting){this.isSubmitting=!0;var e=$('input[type="file"]:not([disabled])',this.$form);e.each((function(t,e){e.files.length>0||$(e).prop("disabled",!0)}));var n=new FormData(this.$form[0]);"undefined"!=typeof webcard&&"culture"in webcard&&webcard.culture&&n.append("culture",webcard.culture),e.prop("disabled",!1),n.append("id",this.getId()),n.append("tac",Math.floor(((new Date).getTime()-this.time.getTime())/1e3));var r=this.$form.attr("action")||("undefined"!=typeof webcard&&webcard.apiHost?"https://"+webcard.apiHost:"")+"/form_container/submit";return $.ajax({url:r,crossDomain:!0,xhrFields:{withCredentials:!0},data:n,processData:!1,contentType:!1,type:"POST",success:function(e){t.isSubmitting=!1,t.$element.find(".ed-element").each((function(){this.element.destroy()})),t.$element.html(e),t.getViewport().scrollTo(t.$element,"center",300),$(e).is(".wv-success")?t.getViewport().notify("form.success",[t,t.$form]):t.getViewport().notify("form.error",[t,t.$form]),t.wakeup();var n=t.$element.find(".ed-element"),r=[];n.each((function(t,e){var n=i.Z.getTypeClassName($(e).attr("class"));n in window&&r.push(new window[n](e))})),animations.initSubsequent(r)}}).fail((function(e){t.isSubmitting=!1,t.$form.find(".wv-message").remove();var n=$('<div class="wv-message wv-failure">'.concat((null==e?void 0:e.responseText)||"Error while sending the form","</div>"));t.$form.prepend(n).append(n.clone())})),this}}
```

Vaciarla y dejarla así 
```
submit:function(){}
```

## Validaciones:

## Otros

- Crear .htaccess con este contenido: 
	```
	 
	```
## Uúltimo
- El resto de los archivos index.php colocar: 
	```
	<?php
	    include_once "../app/config.php";   
	?>
	```
## Recursos

- Notificaciones: https://jakim.me/Toasty.js/ 

- Google recapcha https://www.easyappcode.com/como-utilizar-el-nuevo-recaptcha-v3-de-google-en-nuestro-proyecto-web#:~:text=Veamos%20cómo%20utilizarlo%3A&text=Deberemos%20primero%20añadir%20la%20etiqueta,de%20reCAPTCHA%20y%20presionamos%20Enviar.
