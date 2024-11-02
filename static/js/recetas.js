let listaRecetas;
let nombre = document.getElementById('nombre');
let listaAlimentos = document.getElementById('lista-alimentos');
let momento = document.getElementById('momento');
let listaAlimentosSeleccionados = document.getElementById('lista-alimentos-seleccionados');

document.addEventListener('DOMContentLoaded', async function () {
    const agregarRecetaBtn = document.getElementById('agregarNuevaReceta');
    listaRecetas = document.getElementById('lista-recetas');
    try {
        await obtenerListaRecetas();
    } catch (error) {
        console.log('Error:', error);
    }

    agregarRecetaBtn.addEventListener('click', function () {
        momento.addEventListener('change', async function () {
            if (listaAlimentosSeleccionados.hasChildNodes()) {
                listaAlimentosSeleccionados.innerHTML = '';
            }

            if (listaAlimentos.hasChildNodes()) {
                listaAlimentos.innerHTML = '';
            }
            await cargarAlimentos(momento.value)
        });

        const confirmarRecetaBtn = document.getElementById('confirmarReceta');

        confirmarRecetaBtn.addEventListener('click', async function () {
            const lista_inputs = listaAlimentosSeleccionados.querySelectorAll('input');
            const lista_alimentos_seleccionados = [];
            lista_inputs.forEach(input => {
                const idAlimento = input.id;
                const cantidad = input.value;
                lista_alimentos_seleccionados.push({
                    IdAlimento: idAlimento,
                    Cantidad: parseInt(cantidad)
                });
            });

            const nuevoReceta = {
                Nombre: nombre.value,
                Alimentos: lista_alimentos_seleccionados,
                Momento: momento.value
            };
            const URL = 'http://localhost:8080/recetas/';
            await makeRequest(URL, Method.POST, nuevoReceta, ContentType.JSON, CallType.PRIVATE, successCargarNuevaReceta, errorCargarNuevaReceta);
            const modal = bootstrap.Modal.getInstance(document.getElementById('cargarReceta'));
            modal.hide();
            location.reload();
        });
    });
});

async function obtenerListaRecetas() {
    let URL = 'http://localhost:8080/recetas/';
    await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successObtenerListaRecetas, errorObtenerListaRecetas);
}

function successObtenerListaRecetas(response) {
    console.log('Recetas:', response);
    response.forEach(receta => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
                            <td>${receta.Nombre}</td>
                            <td>${receta.Alimentos.map(alimento => alimento.Nombre).join(', ')}</td>
                            <td>${receta.Momento.charAt(0).toUpperCase() + receta.Momento.slice(1)}</td>
                        `;
        const botonEditar = document.createElement('button');
        const iconoEditar = document.createElement('i');
        iconoEditar.setAttribute('class', 'fa-solid fa-pencil');

        Object.assign(botonEditar, {
            className: 'btn btn-primary',
            value: receta.ID
        });
        botonEditar.setAttribute('data-bs-toggle', 'modal');
        botonEditar.setAttribute('data-bs-target', '#cargarReceta');
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
            type: 'button',
        });
        iconoEliminar.setAttribute('class', 'fa-solid fa-trash');
        botonEliminar.appendChild(iconoEliminar);
        botonEliminar.addEventListener('click', async function () {
            let idReceta = receta.ID;
            const URL = 'http://localhost:8080/recetas/' + idReceta + '/';
            await makeRequest(URL, Method.DELETE, null, ContentType.JSON, CallType.PRIVATE, successEliminarReceta, errorEliminarReceta);
        });

        const tdBotones = document.createElement('td');
        tdBotones.appendChild(botonEditar);
        tdBotones.appendChild(botonEliminar);
        tdBotones.setAttribute('class', 'd-flex gap-2');
        tr.appendChild(tdBotones);
        listaRecetas.appendChild(tr);
    });
}

function errorObtenerListaRecetas(status, response) {
    console.log("Falla:", response);
}

function successObtenerReceta(response) {
    console.log('Receta:', response);
    nombre.value = response.Nombre;
    momento.value = response.Momento;
    
    response.Alimentos.forEach(alimento => {
        const div = document.createElement('div');
        div.setAttribute('class', 'row mt-3');
        const div2 = document.createElement('div');
        div2.setAttribute('class', 'col-2 d-flex align-items-center');

        const label = document.createElement('label');
        label.innerText = alimento.Nombre;
        label.setAttribute('for', alimento.id);
        label.setAttribute('class', 'col-10');

        const input = document.createElement('input');
        input.setAttribute('type', 'number');
        input.setAttribute('id', alimento.IdAlimento);
        input.setAttribute('min', 1);
        input.setAttribute('value', alimento.Cantidad);
        input.setAttribute('class', 'form-control form-control-sm');
        input.setAttribute('required', true);

        const botonEliminar = document.createElement('button');
        const iconoEliminar = document.createElement('i');
        botonEliminar.setAttribute('class', 'btn btn-danger btn-sm');
        iconoEliminar.setAttribute('class', 'fa-solid fa-trash');
        botonEliminar.appendChild(iconoEliminar);
        botonEliminar.addEventListener('click', function () {
            const option = document.createElement('option');
            option.value = alimento.IdAlimento;
            option.innerText = alimento.Nombre;
            listaAlimentos.appendChild(option);
            div.remove();
        });
        label.appendChild(input);
        div.appendChild(label);
        div2.appendChild(botonEliminar);
        div.appendChild(div2);
        listaAlimentosSeleccionados.appendChild(div);
    });
}

function errorObtenerReceta(status, response) {
    console.log("Falla:", response);
}

function successEliminarReceta(response) {
    alert('Receta eliminada');
    location.reload();
}

function errorEliminarReceta(status, response) {
    console.log("Falla:", response);
}

function cargarAlimentos(momento) {
    const URL = 'http://localhost:8080/alimentos/?momentoDelDia=' + momento;
    makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successCargarAlimentos, errorCargarAlimentos);
}

function successCargarAlimentos(response) {
    console.log('Alimentos:', response)
    response.forEach(alimento => {
        const option = document.createElement('option');
        option.value = alimento.id;
        option.innerText = alimento.Nombre;
        option.addEventListener('click', function () {
            listaAlimentos.removeChild(option);
            const div = document.createElement('div');
            div.setAttribute('class', 'row mt-3');
            const div2 = document.createElement('div');
            div2.setAttribute('class', 'col-2 d-flex align-items-center');

            const label = document.createElement('label');
            label.innerText = alimento.Nombre;
            label.setAttribute('for', alimento.id);
            label.setAttribute('class', 'col-10');

            const input = document.createElement('input');
            input.setAttribute('type', 'number');
            input.setAttribute('id', alimento.IdAlimento);
            input.setAttribute('min', 1);
            input.setAttribute('value', 1);
            input.setAttribute('class', 'form-control form-control-sm');
            input.setAttribute('required', true);

            const botonEliminar = document.createElement('button');
            const iconoEliminar = document.createElement('i');
            botonEliminar.setAttribute('class', 'btn btn-danger btn-sm');
            iconoEliminar.setAttribute('class', 'fa-solid fa-trash');
            botonEliminar.appendChild(iconoEliminar);
            botonEliminar.addEventListener('click', function () {
                listaAlimentos.appendChild(option);
                div.remove();
            });
            label.appendChild(input);
            div.appendChild(label);
            div2.appendChild(botonEliminar);
            div.appendChild(div2);
            listaAlimentosSeleccionados.appendChild(div);
        });
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