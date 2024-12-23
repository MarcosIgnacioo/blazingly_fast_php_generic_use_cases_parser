
/*=============================================
    =   Funcion Para Obtener el carrito     =
=============================================*/

function getCart() {
    let carrito = localStorage.getItem('carrito');
    if (carrito) { carrito = JSON.parse(carrito); }
    if (localStorage.getItem('carrito') == null) {
        carrito = [];
        localStorage.setItem('carrito', JSON.stringify(carrito));
    }
    return carrito;
}

/*=============================================
    =   Funcion Para Calcular el total del carrito     =
=============================================*/

function getTotal(carrito) {
    let total = 0;
    for (var i = carrito.length - 1; i >= 0; i--) {
        total += carrito[i].price * carrito[i].cantidad;
    }

    return total;
} 

/*=============================================
    =   Funcion Para Añadir al carrito     =
=============================================*/

function QuickAdd(target) {

    let data = JSON.parse(target.dataset.product);
    let carrito = getCart();

    const cantidadProducto = document.querySelector('.cantidad_producto').value;  

    let feature_product = ""; 

    if ( document.querySelector('.feature_product') !== null ) {
        feature_product = document.querySelector('.feature_product').value; 
    }
    

    var producto = {
        "id": data.id,
        "cantidad": cantidadProducto,
        "name": data.name,
        "brand": data.name,
        "categories": data.categories,
        "category_parent": "",
        "category_child": "",
        "price": data.original_price,
        "weight": 0,
        "feature": feature_product,
        "inventory": data.inventory,
        "size": data.size,
        "cover": data.thumbnail_path,
        "slug": data.slug,
        "max": data.stock
    };

    var bool = true;

    if (carrito.length > 0) {
        carrito.forEach((value) => {
            if (value.id == producto.id && value.inventory == producto.inventory && value.feature == producto.feature) {
                value.cantidad = parseInt(value.cantidad) + parseInt(producto.cantidad);
                bool = false;
            }
        })
    }
    if (bool) {
        carrito.push(producto);
    }
    localStorage.setItem('carrito', JSON.stringify(carrito));

    saveCarrito(carrito, "Producto añadido al carrito", false)

}

/*=============================================
    =   Funcion Para Guardar el carrito     =
=============================================*/

function saveCarrito(carrito, string, reload, alert = true) {

    //enviar el carrito a guardar en la sesión
    var bodyFormData = new FormData();
    bodyFormData.append('action', 'update_cart');
    bodyFormData.append('token', global_token);
    bodyFormData.append('data', JSON.stringify(carrito));

    axios({
        method: "post",
        url: global_url + 'shopcart',
        data: bodyFormData,
        headers: { "Content-Type": "multipart/form-data" },
    })
        .then(function (response) {

            //actualizar los carritos pequeños
            updateMinCart(carrito);

            //actualizar los carritos del detalle
            //updateDetailCart(carrito); 
            
            if (alert) { 

                var toast = new Toasty();  
                toast.success(string);
            }

        })
        .catch(function (error) {
            console.log(error);
        });
}


/*=============================================
    =   Funcion Para Guardar el carrito     =
=============================================*/

function updateMinCart(carrito) {
    const parser = new DOMParser(); 
 
    fetch(global_url + 'mobile_cart/').then(function (response) {

        return response.text();

    }).then(function (new_code) {

        document.querySelectorAll('.remove_item_on_update').forEach((item) => {
            item.remove();
        }); 
        
        var nuevoElementoHTML = new_code;

        document.getElementById("cart_container_mobile").insertAdjacentHTML('beforeend', nuevoElementoHTML); 

    })
}

/*=============================================
    =   Funcion Para Guardar el carrito     =
=============================================*/

function updateDetailCart(carrito) {
    const parser = new DOMParser(); 

    fetch(global_url + 'detail_cart/').then(function (response) {

        return response.text();

    }).then(function (new_code) {

        if ( document.getElementById('cart_container_details') !== null ) { 

            document.getElementById("cart_container_details").innerHTML = ''; 
            var nuevoElementoHTML = new_code + document.getElementById('cart_container_details').innerHTML; 
            document.getElementById("cart_container_details").innerHTML = nuevoElementoHTML;

            let total = getTotal(carrito);

            document.getElementById("final_amount_1").innerHTML = '$' + new Intl.NumberFormat("en-IN").format(total) + ' MXN';
            document.getElementById("final_amount_2").innerHTML = '$' + new Intl.NumberFormat("en-IN").format(total) + ' MXN';

        }

    })

}

/*=============================================
    =   Funcion Para Guardar el carrito     =
=============================================*/

function updateSelection(target,type) {

    //cambir el producto

    if (type == 'cafe') { 

        var select = document.getElementById(id_selection);
        var option = select.options[select.selectedIndex];
        product = option.dataset.product;

        //document.getElementById('buttonAdd').dataset.product = product
        document.querySelector('.boton_add_to_cart').dataset.product = product;

        
    }

    if (type == 'merch') {

        let item = target.dataset.product; 

        document.querySelector('.button_add_big_product').dataset.product = item; 
        document.querySelector('.button_add_big_product').dataset.feature_product = ""; 



    }
}


/*=============================================
    =   Funcion Para Eliminar producto del carrito      =
=============================================*/

function removeItemCart(target) {
    let carrito = getCart();

    let remove_item = target.dataset.id;
    let feature = target.dataset.feature;
    let destroy_item = target.dataset.parent;

    let indice = 0;

    indice = carrito.findIndex(el=> (el.id == remove_item && el.feature == feature));

    carrito.splice(indice, 1)

    localStorage.setItem('carrito', JSON.stringify(carrito));

    document.querySelectorAll(`.${destroy_item}`).forEach(el=>{
        el.remove()
    })

    saveCarrito(carrito, "Producto removido", false, true)

}

/*=============================================
    =   Funcion Para Crear un pedido      =
=============================================*/

let can_send_order = true

function validateOrder(e) { 
    let toast = new Toasty();  
    e.preventDefault()
    
    let bodyFormData = new FormData(e.target);
    bodyFormData.append('token', global_token);
    bodyFormData.append('action', 'quote_shipment');

    axios({
        method: "post",
        url: global_url + 'shopcart',
        data: bodyFormData,
        headers: { "Content-Type": "multipart/form-data" },
    }).then(function (response) {
        if(response.data >= 0){
            Swal.fire({
                title: 'Envío',
                text: `El envío calculado para esta dirección es de $${response.data}`,
                icon: "warning",
                showCancelButton: true,
                confirmButtonColor: "#3085d6",
                cancelButtonText: "Editar dirección",
                confirmButtonText: "Continuar al pago"
            }).then((result) => {
                if (result.isConfirmed) {
                    if (le && can_send_order) { 
                        can_send_order = false
                        toast.warning("Espere mientras procesamos tu pedido");
                
                        bodyFormData.set('action', 'make_order');
                        bodyFormData.append('delivery_type_id', 2);
                        bodyFormData.append('shipping_cost', response.data);
                        bodyFormData.append('password', "123");
                        bodyFormData.append('ship_lastname', "");
                
                        let additionalData = Object.fromEntries(bodyFormData.entries());
                
                        bodyFormData.append('ship_name', additionalData.name);
                
                        axios({
                            method: "post",
                            url: global_url + 'shopcart',
                            data: bodyFormData,
                            headers: { "Content-Type": "multipart/form-data" },
                        })
                            .then(function (response) {
                
                                if (response.data[0].code > 0) {
                
                
                                    window.location.href = global_url + "tienda/payment/" + response.data[0].folio + "/";
                
                                } else {
                
                                    toast.warning(response.data[0].message);
                
                                    can_send_order = true
                
                                    return false;
                                }
                
                            })
                            .catch(function (error) {
                                console.log(error);
                                toast.error("Verifique su información");
                
                                can_send_order = true
                                return false;
                            });
                
                    } else {
                        toast.error("Debe verificar la casilla");
                    }
                }
            });
        }
        else{
            toast.error("Verifique su información");
        }
    }).catch(function (error) {
        console.log(error);
        toast.error("Verifique su información");
    });

    return false;
}

/*=============================================
    =   Función Para Pagar (Banorte)          =
=============================================*/

async function validatePayment(e) {

    e.preventDefault();
    
    let data = new FormData(e.target);

    let additionalData = Object.fromEntries(data.entries());

    additionalData.folio = order_folio;
    
    let local_url = global_url+'shopcart';
    
    let response = await sendData('make_payment', local_url, additionalData, true);

    if (response) {
        document.open();
        document.write(response.data);
        document.close();
    } 

    return false;
}
