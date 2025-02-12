
/*=============================================
    =   Funcion generica tools    =
=============================================*/

function sendData(action, url, additionalData = {}, responseReturn = false) {
    var toast = new Toasty();

    return new Promise((resolve, reject) => {
        if (true) {
            toast.warning("Estamos enviando sus datos");

            var bodyFormData = new FormData();
            bodyFormData.append('action', action);
            bodyFormData.append('token', global_token);
            bodyFormData.append('token_google', document.getElementById('g-recaptcha-response').value);

            // Añadir datos adicionales
            for (let key in additionalData) {
                if (additionalData.hasOwnProperty(key)) {
                    bodyFormData.append(key, additionalData[key]);
                }
            }

            axios({
                method: "post",
                url: url,
                data: bodyFormData,
                headers: { "Content-Type": "multipart/form-data" },
            })
            .then(function (response) {
                if(responseReturn){
                    resolve(response)
                }
                if (response.data[0].code > 0) {
                    
                    toast.success(response.data[0].message);
                    resolve(response.data[0].code);
                
                } else {
                    toast.warning("Verifique la información");
                    resolve(false);
                }
            })
            .catch(function (error) {
                console.log(error);
                toast.error("Verifique su información");
                resolve(false);
            });
        } else {
            toast.error("Debe verificar la casilla");
            resolve(false);
        }
    });
}

/*=============================================
    =   Funcion Para Validar Reseña         =
=============================================*/

function validateReview(target) { 

    if ( true ) { 

        //swal("Espere", "Estamos enviando su mensaje", "info");

        let strats = getHighestSelectedValue();

        var toast = new Toasty();  
        toast.warning("Estamos enviando su reseña");

        var bodyFormData = new FormData();
        bodyFormData.append('action', 'add_reivew');
        bodyFormData.append('token', global_token); 

        bodyFormData.append('name', document.getElementById('author').value);
        bodyFormData.append('email', document.getElementById('email').value);
        bodyFormData.append('comment', document.getElementById('comment').value); 
        bodyFormData.append('rating_stars', strats); 
        bodyFormData.append('product_id', document.getElementById('product_id').value);  

        axios({
            method: "post",
            url: global_url + 'clients',
            data: bodyFormData,
            headers: { "Content-Type": "multipart/form-data" },
        })
            .then(function (response) {

                console.log(response.data)

                if (response.data[0].code > 0) {

                	toast.success("Su reseña ha sido enviada");

                    //swal("Correcto", "Su reseña ha sido enviada", "success");
                    document.getElementById("comment-form").reset();

                } else {

                	toast.warning("Verifique la información");
                    //swal("", "Verifique la información", "warning")
                }

            })
            .catch(function (error) {
                console.log(error);

                toast.error("La reseña no pudo ser enviado, por favor revisar la información del formulario");
                //swal("Error", "La reseña no pudo ser enviado, por favor revisar la información del formulario", "error");

            });

    } else {

        swal("Error", "Debe verificar la casilla", "error");
        return false;
    }

    return false;
}

/*=============================================
    =   Funcion Para Otener estrellas        =
=============================================*/

function getHighestSelectedValue() {
    // Selecciona todos los radio buttons con el nombre 'rating_stars'
    const radios = document.querySelectorAll('input[name="rating_stars"]');
    
    // Inicializa una variable para almacenar el valor más alto seleccionado
    let highestValue = 0;

    // Itera sobre cada radio button
    radios.forEach((radio) => {
        // Si el radio button está seleccionado y su valor es mayor que el valor almacenado
        if (radio.checked && parseInt(radio.value) > highestValue) {
            highestValue = parseInt(radio.value);
        }
    });

    return highestValue;
}

/*=============================================
    =   Funcion Para Otener estrellas        =
=============================================*/

function setState(target)
{
    var bodyFormData = new FormData();
    bodyFormData.append('action', 'update_state');
    bodyFormData.append('token', global_token); 
    
    bodyFormData.append('state', target); 

    axios({
        method: "post",
        url: global_url + 'clients',
        data: bodyFormData,
        headers: { "Content-Type": "multipart/form-data" },
    })
    .then(function (response) {

        console.log(response.data) 

    })
    .catch(function (error) {
        console.log(error);  
    });
}

/*=============================================
    =   Funcion solo letras   =
=============================================*/

function soloLetras(e) {
    var tecla = e.key.toLowerCase();
    return /^[a-zñáéíóúü\s]$/i.test(tecla);
}

/*=============================================
    =   Funcion solo números    =
=============================================*/

function soloNumeros(e) {
    var tecla = e.key.toLowerCase();
    return /^[0-9.]$/.test(tecla);
}

/*=============================================
    =   Funcion letras y números   =
=============================================*/

function soloLetrasYnumeros(e) {
    var tecla = e.key.toLowerCase();
    return /^[a-zñáéíóúü\s0-9.]+$/.test(tecla);
}

/*=============================================
    =   Funcion para replicar nombre    =
=============================================*/

function repplyName(t) {
    let name = t.value
    console.log(name.length)
    if (name.length) {

        let name2 = document.getElementById('ship_name').value
        if (name2.length <= 0) {
            document.getElementById('ship_name').value = name
        }
    }
}