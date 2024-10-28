let listaRecetas = document.getElementById('lista-recetas');
let nombre = document.getElementById('nombre');
let listaAlimentos = document.getElementById('lista-alimentos');
let momento = document.getElementById('momento');

document.addEventListener('DOMContentLoaded', async function () {
    const agregarRecetaBtn = document.getElementById('agregarNuevaReceta');

    await obtenerListaRecetas();

    agregarRecetaBtn.addEventListener('click',async function () {
        debugger
        document.getElementById('form-receta').reset();
        await cargarAlimentos();
        const confirmarRecetaBtn = document.getElementById('confirmarReceta');

        confirmarRecetaBtn.addEventListener('click', async function () {
            const listaAlimentosSeleccionados = Array.from(listaAlimentos.selectedOptions).map(option => option.value);
            const nuevoReceta = {
                Nombre: nombre.value,
                Alimentos: listaAlimentosSeleccionados,
                Momento: momento.value
            };
            const URL = 'http://localhost:8080/recetas/';
            await makeRequest(URL, Method.POST, nuevoReceta, ContentType.JSON, CallType.PRIVATE, successCargarNuevaReceta, errorCargarNuevaReceta);
            const modal = bootstrap.Modal.getInstance(document.getElementById('cargarReceta'));
            modal.hide();
            document.getElementById('form-receta').reset();
        });
    });
});

async function obtenerListaRecetas() {
    let URL = 'http://localhost:8080/recetas/';
    // Acá está el error, al ejecutar el MakeRequest, no se ejecuta ni el success ni el error
    await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successObtenerListaRecetas, errorObtenerListaRecetas);
}

function successObtenerListaRecetas(response) {
    console.log('Recetas:', response);
    response.forEach(receta => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
                            <td>${receta.nombre}</td>
                            <td>$${receta.alimentos}</td>
                            <td>${receta.momento}</td>
                        `;
        const botonEditar = document.createElement('button');
        const iconoEditar = document.createElement('i');
        iconoEditar.setAttribute('class', 'fa-solid fa-pencil');

        Object.assign(botonEditar, {
            className: 'btn btn-warning',
            innerText: 'Editar',
            value: receta.id
        });
        botonEditar.setAttribute('data-bs-toggle', 'modal');
        botonEditar.setAttribute('data-bs-target', '#editarReceta');
        botonEditar.appendChild(iconoEditar);
        botonEditar.addEventListener('click', async function () {
            let idReceta = this.value;
            const URL = 'http://localhost:8080/recetas/' + idReceta + '/';
            await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successObtenerReceta, errorObtenerReceta);
        });

        const botonEliminar = document.createElement('button');
        const iconoEliminar = document.createElement('i');

        Object.assign(botonEliminar, {
            className: 'btn btn-danger',
            innerText: 'Eliminar',
            value: receta.id
        });
        iconoEliminar.setAttribute('class', 'fa-solid fa-trash');
        botonEliminar.appendChild(iconoEliminar);
        botonEliminar.addEventListener('click', async function () {
            let idReceta = receta.id;
            const URL = 'http://localhost:8080/recetas/' + idReceta + '/';
            await makeRequest(URL, Method.DELETE, null, ContentType.JSON, CallType.PRIVATE, successEliminarReceta, errorEliminarReceta);
        });

        const tdBotones = document.createElement('td');
        tdBotones.appendChild(botonEliminar);
        tdBotones.appendChild(botonEditar);
        tdBotones.setAttribute('class', 'd-flex gap-2');
        tr.appendChild(tdBotones);
        listaRecetas.appendChild(tr);
    });
}

function errorObtenerListaRecetas(status, response) {
    console.log("Falla:", response);
}

function successObtenerReceta(response) {
    alert('Receta obtenida');
    console.log(response);
}

function errorObtenerReceta(status, response) {
    console.log("Falla:", response);
}

function successEliminarReceta(response) {
    alert('Receta eliminada');
    console.log(response);
}

function errorEliminarReceta(status, response) {
    console.log("Falla:", response);
}

async function cargarAlimentos() {
    const URL = 'http://localhost:8080/alimentos/';
    await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successCargarAlimentos, errorCargarAlimentos);
}

function successCargarAlimentos(response) {
    console.log('Alimentos:', response)
    debugger
    response.forEach(alimento => {
        const option = document.createElement('option');
        option.value = alimento.id;
        option.innerText = alimento.nombre;
        listaAlimentos.appendChild(option);
    });
}

function errorCargarAlimentos(status, response) {
    console.log("Falla:", response);
}

function successCargarNuevaReceta(response) {
    alert('Receta cargada');
    console.log(response);
}

function errorCargarNuevaReceta(status, response) {
    console.log("Falla:", response);
}