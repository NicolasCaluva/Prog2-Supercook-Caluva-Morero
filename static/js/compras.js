let listaCompras = document.getElementById('lista-compras');
let PaginaActual = 1;
const ItemsPorPagina = 10;
let comprasData = [];

document.addEventListener('DOMContentLoaded', async function () {
    await obtenerListaAlimentosPocoStock();

    const confirmarButton = document.getElementById('cargarNuevaCompra');
    confirmarButton.addEventListener('click', enviarCompraDto);
    document.getElementById('aplicarFiltros').addEventListener('click', aplicarFiltros);
    document.getElementById('pagAnterior').addEventListener('click', IrPaginaPrevia);
    document.getElementById('pagSiguiente').addEventListener('click', IrPaginaSiguiente);
});

function successCargarListaCompras(response) {
    comprasData = response;
    RenderizarPagina(PaginaActual);
}

function RenderizarPagina(page) {
    listaCompras.innerHTML = '';
    const start = (page - 1) * ItemsPorPagina;
    const end = start + ItemsPorPagina;
    const pageItems = comprasData.slice(start, end);

    pageItems.forEach(compra => {
        const tr = document.createElement('tr');
        tr.dataset.idAlimento = compra.IdAlimento;

        tr.innerHTML = `
            <td>${compra.Nombre}</td>
            <td>$${compra.PrecioUnitario}</td>
            <td>${compra.Stock}</td>
            <td>${compra.CantMinimaStock}</td>
            <td>${compra.TipoAlimento.charAt(0).toUpperCase() + compra.TipoAlimento.slice(1)}</td>
            <td>${compra.MomentoDelDia.map(momento => momento.charAt(0).toUpperCase() + momento.slice(1)).join(' - ')}</td>
        `;

        const tdAcciones = document.createElement('td');
        const inputCantidad = document.createElement('input');
        inputCantidad.type = 'number';
        inputCantidad.min = '1';
        inputCantidad.placeholder = 'Cantidad';
        inputCantidad.className = 'form-control';
        tdAcciones.appendChild(inputCantidad);
        tr.appendChild(tdAcciones);
        listaCompras.appendChild(tr);
    });

    document.getElementById('indicadorPagina').textContent = `Página ${PaginaActual}`;
    updatePaginationButtons();
}

function updatePaginationButtons() {
    document.getElementById('pagAnterior').disabled = PaginaActual === 1;
    document.getElementById('pagSiguiente').disabled = PaginaActual * ItemsPorPagina >= comprasData.length;
}

function IrPaginaPrevia() {
    if (PaginaActual > 1) {
        PaginaActual--;
        RenderizarPagina(PaginaActual);
    }
}

function IrPaginaSiguiente() {
    if (PaginaActual * ItemsPorPagina < comprasData.length) {
        PaginaActual++;
        RenderizarPagina(PaginaActual);
    }
}

function errorCargarListaAlimentosPocoStock(status, response) {
    console.log("Falla al cargar la lista:", response);
}

async function aplicarFiltros() {
    debugger;
    const tipoAlimento = document.getElementById('tipoAlimento').value;
    const nombre = document.getElementById('nombre').value;
    const momentoDelDiaSelect = document.getElementById('momentoDelDia');
    const momentosSeleccionados = Array.from(momentoDelDiaSelect.selectedOptions)
        .map(option => `momentoDelDia=${option.value}`)
        .join('&');

    const url = `http://localhost:8080/alimentos/?${momentosSeleccionados}&tipoAlimento=${tipoAlimento}&nombre=${nombre}&StockMenorCantidadMinima=True`;
    await makeRequest(url, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successCargarListaCompras, errorCargarListaAlimentosPocoStock);
}

async function obtenerListaAlimentosPocoStock() {
    const url = 'http://localhost:8080/alimentos/?StockMenorCantidadMinima=True';
    await makeRequest(url, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successCargarListaCompras, errorCargarListaAlimentosPocoStock);
}

async function enviarCompraDto() {
    const alimentosComprados = [];
    const filas = document.querySelectorAll('#lista-compras tr');
    filas.forEach(fila => {
        const idAlimento = fila.dataset.idAlimento;
        const inputCantidad = fila.querySelector('input[type="number"]');
        if (idAlimento && inputCantidad && inputCantidad.value && inputCantidad.value > 0) {
            alimentosComprados.push({
                IDAlimento: idAlimento,
                CantComprada: parseInt(inputCantidad.value)
            });
        }
    });

    const compraDto = {
        Alimentos: alimentosComprados
    };

    console.log('CompraDto a enviar:', compraDto);

    let url = 'http://localhost:8080/compras/';
    await makeRequest(url, Method.POST, compraDto, ContentType.JSON, CallType.PRIVATE, successEnviarCompra, errorEnviarCompra);
}

function successEnviarCompra() {
    alert('Compra realizada con éxito.');
    location.reload();
}

function errorEnviarCompra(status, response) {
    console.log(response.error);
    alert(response.error);
}
