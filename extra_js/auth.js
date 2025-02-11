let toast = new Toasty();

/*=============================================
    =   Función Para Acceder (Login)          =
=============================================*/

async function validateAccess(e) {

    e.preventDefault();
    
    let data = new FormData(e.target);

    let additionalData = Object.fromEntries(data.entries());

    let local_url = global_url+'clients';
    
    let response = await sendData('access', local_url, additionalData);

    if (response) {
        window.location.href = global_url + 'account/dashboard';
    } 

    return false;
}

/*=============================================
    =   Función Para Validar Registro         =
============================================= */

async function validateRegister(e) {
    
    e.preventDefault();

    let data = new FormData(e.target);
    let additionalData = Object.fromEntries(data.entries());

    if (additionalData.password != additionalData.password_confirmation) {
        toast.error("Las contraseñas no coinciden"); 
        return false;
    }
    let local_url = global_url+'clients';

    let response = await sendData('register_client_account', local_url, additionalData);

    if (response) {
        window.location.href = global_url + 'account/dashboard';
    }

    return false; 
}

/*=============================================
    =   Funcion Para Actualizar perfil    =
=============================================*/

async function validateUpdate(e) {
    
    e.preventDefault();

    let data = new FormData(e.target);
    let additionalData = Object.fromEntries(data.entries());

    let local_url = global_url+'clients';

    let response = await sendData('update_profile', local_url, additionalData);

    if (response) { 
    }

    return false; 
}

/*===========================================================================
    =   Funcion para Agregar nueva contraseña dentro de editar cuenta   =
===========================================================================*/
async function validateChangePassword(e) {
    
    e.preventDefault();
    let toast = new Toasty();

    let data = new FormData(e.target);
    let additionalData = Object.fromEntries(data.entries());

    if (additionalData.password_current != additionalData.password_1) {
        if (additionalData.password_1 == additionalData.password_2) {
            
            let local_url = global_url+'clients';

            let response = await sendData('change_new_password', local_url, additionalData);

            if (response) { 
                e.target.reset();
            }

        } else { 
            toast.error("La nueva contraseña no coincide");
            return false;
        }
    } else { 
        toast.error("La nueva contraseña no debe ser igual a la actual");
        return false;
    }

    return false; 
}

/*=============================================
    =   Funcion Para Validar Contacto         =
=============================================*/

async function validateContact(e) {
    
    e.preventDefault();

    let data = new FormData(e.target);
    let additionalData = Object.fromEntries(data.entries());

    let local_url = global_url+'sendcontact';

    let response = await sendData('contact', local_url, additionalData);

    if (response) { 

        e.target.reset();
    }

    return false; 
}

/*=============================================
    =   Funcion Para Recuperar contraseña   =
=============================================*/

async function validateReset(e) {
    
    e.preventDefault();

    let data = new FormData(e.target);
    let additionalData = Object.fromEntries(data.entries());

    let local_url = global_url+'clients';
    
    let response = await sendData('reset_client_account', local_url, additionalData);

    if (response) {
        e.target.reset();
    } 

    return false;
}

/*====================================================
    =   Funcion para Agregar nueva contraseña    =
====================================================*/

async function validateChange(e) {
    
    e.preventDefault();

    let data = new FormData(e.target);
    let additionalData = Object.fromEntries(data.entries());
    
    if (additionalData.password == additionalData.repeat_password) {

        let local_url = global_url+'clients';
        
        let response = await sendData('change_password', local_url, additionalData);

        if (response) {
            
            e.target.reset();
            window.location.href = global_url+"login/";
        } 

    } else {
        toast.error("Las contraseñas no coinciden"); 
        return false;
    }

    return false;
}

function toggleBtn(e, color){
    let btn = e.target;

    if(btn.dataset.state == '0'){
        btn.classList.replace(`btn-outline-${color}`, `btn-${color}`);
        btn.dataset.state = '1'
    }else if(btn.dataset.state == '1'){
        btn.classList.replace(`btn-${color}`, `btn-outline-${color}`);
        btn.dataset.state = '0'
    }

}

function shareSns(type, href=null){
    let url = href ?? encodeURIComponent(window.location.href);

    switch(type){
        case 'facebook':
            window.open(`https://www.facebook.com/sharer/sharer.php?u=${url}`, '_blank');
        break;
        case 'whatsapp':
            window.open(`https://api.whatsapp.com/send?text=${url}`, '_blank');
        break;
    }
}

/*=============================================
    =   Función Para Datos de facturación    =
=============================================*/

async function validateBilling(e) {

    e.preventDefault();
    
    let data = new FormData(e.target);

    let additionalData = Object.fromEntries(data.entries());
    
    let local_url = global_url+'clients';
    
    let response = await sendData('billing_data', local_url, additionalData, true);

    if (response) {
        window.location.reload();
    } 

    return false;
}
