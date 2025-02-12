
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

function emptyCart() {
    carrito = [];
    localStorage.setItem('carrito', JSON.stringify(carrito));
    saveCarrito(getCart(), "", false, false)
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

    if (document.querySelector('.feature_product') !== null) {
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
        .then(function(response) {

            //actualizar los carritos pequeños
            updateMinCart(carrito);

            //actualizar los carritos del detalle
            //updateDetailCart(carrito); 

            if (alert) {

                var toast = new Toasty();
                toast.success(string);
            }
            if (reload) window.location.reload()

        })
        .catch(function(error) {
            console.log(error);
        });
}


/*=============================================
    =   Funcion Para Guardar el carrito     =
=============================================*/

function updateMinCart(carrito) {
    const parser = new DOMParser();

    fetch(global_url + 'mobile_cart/').then(function(response) {

        return response.text();

    }).then(function(new_code) {

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

    fetch(global_url + 'detail_cart/').then(function(response) {

        return response.text();

    }).then(function(new_code) {

        if (document.getElementById('cart_container_details') !== null) {

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

function updateSelection(target, type) {

    //cambir el producto

    if (type == 'cafe') {

        var select = document.getElementById(id_selection);
        var option = select.options[select.selectedIndex];
        product = option.dataset.product;

        //document.getElementById('buttonAdd').dataset.product = product
        document.querySelector('.boton_add_to_cart').dataset.product = product;


    }

    if (type == 'merch') {
        let item = target.children[target.selectedIndex].dataset.product;
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

    indice = carrito.findIndex(el => (el.id == remove_item && el.feature == feature));

    carrito.splice(indice, 1)

    localStorage.setItem('carrito', JSON.stringify(carrito));

    document.querySelectorAll(`.${destroy_item}`).forEach(el => {
        el.remove()
    })

    saveCarrito(carrito, "Producto removido", false, true)

}

/*=============================================
    =   Funcion Para Actualizar el carrito      =
=============================================*/

function updateCart() {
    let carrito = getCart();

    document.querySelectorAll('.cart_product_quantity > input').forEach(inputEl => {
        let indice = null;
        indice = carrito.findIndex(el => (el.id == inputEl.dataset.id));

        if (indice >= 0) carrito[indice].cantidad = inputEl.value
    })

    localStorage.setItem('carrito', JSON.stringify(carrito));

    saveCarrito(carrito, "", true, false)
}
/*=============================================
    =   Funcion Para Crear un pedido      =
=============================================*/

let can_send_order = true

function validateOrder(e) {
    let toast = new Toasty();
    e.preventDefault()

    if (!getCart().length) {
        Swal.fire(
            "Carrito vacío",
            "Agregue productos para poder procesar su orden.",
            "info"
        );
        return false;
    }

    let bodyFormData = new FormData(e.target);
    bodyFormData.append('token', global_token);
    bodyFormData.append('action', 'quote_shipment');

    Swal.fire({
        title: 'Cargando',
        text: 'Espere un momento',
        icon: "warning",
        showConfirmButton: false,
    })
    axios({
        method: "post",
        url: global_url + 'shopcart',
        data: bodyFormData,
        headers: { "Content-Type": "multipart/form-data" },
    }).then(function(response) {
        Swal.close();

        if (response.data?.length > 0) {
            const selectElement = document.createElement('select');
            selectElement.style = `
                display: block;
                width: 100%;
                padding: .375rem 2.25rem .375rem .75rem;
                -moz-padding-start: calc(0.75rem - 3px);
                font-size: 1rem;
                font-weight: 400;
                line-height: 1.5;
                color: #212529;
                background-color: #fff;
                background-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16'%3e%3cpath fill='none' stroke='%23343a40' stroke-linecap='round' stroke-linejoin='round' stroke-width='2' d='M2 5l6 6 6-6'/%3e%3c/svg%3e");
                background-repeat: no-repeat;
                background-position: right .75rem center;
                background-size: 16px 12px;
                border: 1px solid #ced4da;
                border-radius: .25rem;
                transition: border-color .15s ease-in-out,box-shadow .15s ease-in-out;
                -webkit-appearance: none;
                -moz-appearance: none;
                appearance: none 
            `;

            response.data.forEach((el, index) => {
                selectElement.innerHTML += `<option ${index == 0 ? 'selected' : ''} value="${el.amount}">
                $${el.amount} - ${el.provider} ${el.servicelevel} (${el.duration_terms})
                </option>`
            })
            Swal.fire({
                title: 'Seleccione una tarifa de envío',
                html: selectElement,
                icon: "warning",
                showCancelButton: true,
                confirmButtonColor: "#3085d6",
                cancelButtonText: "Editar dirección",
                confirmButtonText: "Continuar al pago"
            }).then((result) => {
                if (result.isConfirmed) {
                    if (le && can_send_order) {
                        can_send_order = false

                        Swal.fire({
                            title: 'Cargando',
                            text: 'Espere un momento',
                            icon: "warning",
                            showConfirmButton: false,
                        })

                        bodyFormData.set('action', 'make_order');
                        bodyFormData.append('delivery_type_id', 2);
                        bodyFormData.append('shipping_cost', selectElement.value);
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
                            .then(function(response) {

                                if (response.data[0].code > 0) {


                                    window.location.href = global_url + "tienda/payment/" + response.data[0].folio + "/";

                                } else {

                                    toast.warning(response.data[0].message);

                                    can_send_order = true

                                    return false;
                                }

                            })
                            .catch(function(error) {
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
        else {
            Swal.fire(
                "Sin envío",
                "No hay tarifas de envío disponibles para este código postal.",
                "info"
            );
        }
    }).catch(function(error) {
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

    let local_url = global_url + 'shopcart';

    let response = await sendData('make_payment', local_url, additionalData, true);

    if (response) {
        document.open();
        document.write(response.data);
        document.close();
    }

    return false;
}
