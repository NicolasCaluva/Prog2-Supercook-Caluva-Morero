// Variables globales
let listaAlimentos;
let nombreAlimento = document.getElementById('nombreAlimento');
let precioUnitario = document.getElementById('precioUnitario');
let stock = document.getElementById('stock');
let cantMinimaStock = document.getElementById('cantMinimaStock');
let tipoAlimento = document.getElementById('tipoAlimento');
let momentoDelDia = document.getElementById('momentoDelDia');

// DOM
document.addEventListener('DOMContentLoaded', async function () {
    listaAlimentos = document.getElementById('lista-alimentos');

    await obtenerListaAlimentos();

    const agregarAlimentoBtn = document.querySelector('button[id="agregarNuevoAlimento"]');

    agregarAlimentoBtn.addEventListener('click', function () {
        document.getElementById('form-alimento').reset();

        const confirmarAlimentoBtn = document.getElementById('confirmarAlimento');

        confirmarAlimentoBtn.addEventListener('click', async function () {
            const momentosSeleccionados = Array.from(momentoDelDia.selectedOptions).map(option => option.value);
            const nuevoAlimento = {
                Nombre: nombreAlimento.value,
                PrecioUnitario: parseFloat(precioUnitario.value),
                Stock: parseInt(stock.value),
                CantMinimaStock: parseInt(cantMinimaStock.value),
                TipoAlimento: tipoAlimento.value,
                MomentoDelDia: momentosSeleccionados
            };
            const URL = 'http://localhost:8080/alimentos/';
            await makeRequest(URL, 'POST', nuevoAlimento, ContentType.JSON, CallType.PRIVATE, successCargarNuevoAlimento, errorCargarNuevoAlimento);
            const modal = bootstrap.Modal.getInstance(document.getElementById('cargarAlimento'));
            modal.hide();
            document.getElementById('form-alimento').reset();
            listaAlimentos.appendChild(tr)
        });
    });
});

// Funciones de éxito y error de obtención del listado de alimentos

function successCargarListaAlimentos(response) {
    response.forEach(alimento => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
                            <td>${alimento.Nombre}</td>
                            <td>$${alimento.PrecioUnitario}</td>
                            <td>${alimento.Stock}</td>
                            <td>${alimento.CantMinimaStock}</td>
                            <td>${alimento.TipoAlimento.charAt(0).toUpperCase() + alimento.TipoAlimento.slice(1)}</td>
                            <td>${alimento.MomentoDelDia.map(momento => momento.charAt(0).toUpperCase() + momento.slice(1)).join(' - ')}</td>
                        `;
        const botonEditar = document.createElement('button');
        const iconoEditar = document.createElement('i');
        iconoEditar.setAttribute('class', 'fa-solid fa-pencil');

        Object.assign(botonEditar, {
            className: 'btn btn-primary editar-alimento',
            type: 'button',
            value: alimento.IdAlimento,
        });
        botonEditar.setAttribute('data-bs-toggle', 'modal');
        botonEditar.setAttribute('data-bs-target', '#cargarAlimento');
        botonEditar.appendChild(iconoEditar);
        botonEditar.addEventListener('click', async function () {
            let idAlimento = this.value;
            const URL = 'http://localhost:8080/alimentos/' + idAlimento + '/';
            await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successObtenerAlimento, errorObtenerAlimento);
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
            let idAlimento = alimento.IdAlimento;
            const URL = 'http://localhost:8080/alimentos/' + idAlimento + '/';
            await makeRequest(URL, Method.DELETE, null, ContentType.JSON, CallType.PRIVATE, successEliminarAlimento, errorEliminarAlimento);
        });

        const tdBotones = document.createElement('td');

        tdBotones.appendChild(botonEditar);
        tdBotones.appendChild(botonEliminar);
        tdBotones.setAttribute('class', 'd-flex gap-2')
        tr.appendChild(tdBotones);
        listaAlimentos.appendChild(tr);
    })
}

function errorCargarListaAlimentos(status, response) {
    console.log("Falla:", response);
}

async function obtenerListaAlimentos() {
    let url = 'http://localhost:8080/alimentos/';
    await makeRequest(url, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successCargarListaAlimentos, errorCargarListaAlimentos);
}


// Funciones para la carga de un nuevo alimento en el sistema

function successCargarNuevoAlimento(response) {
    alert("Operación exitosa: \n" + response.ListaMensaje.join("\n"));
}

function errorCargarNuevoAlimento(status, response) {
    alert("Hubo un error al cargar el alimento");
}

// Funciones para la obtención de un alimento del sistema

function successObtenerAlimento(alimento) {
    nombreAlimento.value = alimento.Nombre;
    precioUnitario.value = alimento.PrecioUnitario;
    stock.value = alimento.Stock;
    cantMinimaStock.value = alimento.CantMinimaStock;
    tipoAlimento.value = alimento.TipoAlimento;

    Array.from(momentoDelDia.options).forEach(option => {
        option.selected = false;
    });

    Array.from(momentoDelDia.options).forEach(option => {
        if (alimento.MomentoDelDia.includes(option.value)) {
            option.selected = true;
        }
    });

    const confirmarAlimentoBtn = document.getElementById('confirmarAlimento');

    confirmarAlimentoBtn.addEventListener('click', async function () {
        const momentosSeleccionados = Array.from(document.getElementById('momentoDelDia').selectedOptions).map(option => option.value);
        const nuevoAlimento = {
            IdAlimento: alimento.IdAlimento,
            Nombre: nombreAlimento.value,
            PrecioUnitario: parseFloat(precioUnitario.value),
            Stock: parseInt(stock.value),
            CantMinimaStock: parseInt(cantMinimaStock.value),
            TipoAlimento: tipoAlimento.value,
            MomentoDelDia: momentosSeleccionados
        };

        const URL = 'http://localhost:8080/alimentos/';
        await makeRequest(URL, Method.PUT, nuevoAlimento, ContentType.JSON, CallType.PRIVATE, successEditarAlimento, errorEditarAlimento);
        const modal = bootstrap.Modal.getInstance(document.getElementById('cargarAlimento'));
        modal.hide();
        document.getElementById('form-alimento').reset();
    });
}

function errorObtenerAlimento(response) {
    alert("Hubo un error al obtener el alimento" + response.ListaMensaje.join("\n"));
}


// Funciones para la edición de un alimento en el sistema
function successEditarAlimento(response) {
    alert("Operación exitosa: \n" + response.ListaMensaje.join("\n"));
    location.reload();
}

function errorEditarAlimento(response) {
    alert("Hubo un error al editar el alimento" + response.ListaMensaje.join("\n"));
}

// Funciones para la eliminación de un alimento en el sistema
function successEliminarAlimento(response) {
    alert("Operación exitosa: \n" + response.ListaMensaje.join("\n"));
    location.reload();
}

function errorEliminarAlimento(response) {
    alert("Hubo un error al eliminar el alimento" + response.ListaMensaje.join("\n"));
}
