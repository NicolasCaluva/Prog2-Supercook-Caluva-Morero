let comprasData = [];

document.addEventListener('DOMContentLoaded', async function () {
    await obtenerListaAlimentosPocoStock();

    const confirmarButton = document.getElementById('cargarNuevaCompra');
    confirmarButton.addEventListener('click', enviarCompraDto);
    document.getElementById('aplicarFiltros').addEventListener('click', aplicarFiltros);
});

function successCargarListaCompras(response) {
    comprasData = response;
    RenderizarLista();
}

function RenderizarLista() {
    let listaCompras = document.getElementById('lista-compras');
    listaCompras.innerHTML = '';

    comprasData.AlimentosDto.forEach(compra => {
        const tr = document.createElement('tr');
        tr.dataset.idAlimento = compra.IdAlimento;

        tr.innerHTML = `
            <td class="text-secondary text-sm text-center">${compra.Nombre}</td>
            <td class="text-secondary text-sm text-center">$${compra.PrecioUnitario}</td>
            <td class="text-secondary text-sm text-center">${compra.Stock}</td>
            <td class="text-secondary text-sm text-center">${compra.CantMinimaStock}</td>
            <td class="text-secondary text-sm text-center">${compra.TipoAlimento.charAt(0).toUpperCase() + compra.TipoAlimento.slice(1)}</td>
            <td class="text-secondary text-sm text-center">${compra.MomentoDelDia.map(momento => momento.charAt(0).toUpperCase() + momento.slice(1)).join(' - ')}</td>
        `;

        const tdAcciones = document.createElement('td');
        const inputCantidad = document.createElement('input');
        inputCantidad.type = 'number';
        inputCantidad.min = '1';
        inputCantidad.placeholder = 'Cantidad';
        inputCantidad.className = 'form-control form-control-sm custom-border';
        tdAcciones.appendChild(inputCantidad);
        tr.appendChild(tdAcciones);
        listaCompras.appendChild(tr);
    });
}

function errorCargarListaAlimentosPocoStock(status, response) {
    console.log("Falla al cargar la lista:", response);
}

async function aplicarFiltros() {
    const tipoAlimento = document.getElementById('tipoAlimento').value;
    const nombre = document.getElementById('nombre').value;
    const momentoDelDiaSelect = document.getElementById('momentoDelDia');
    const momentosSeleccionados = Array.from(momentoDelDiaSelect.selectedOptions)
        .map(option => `momentoDelDia=${option.value}`)
        .join('&');

    const url = `http://localhost:8080/alimentos/?${momentosSeleccionados}&tipoAlimento=${tipoAlimento}&nombre=${nombre}&StockMenorCantidadMinima=true`;
    await makeRequest(url, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successCargarListaCompras, errorCargarListaAlimentosPocoStock);
}

async function obtenerListaAlimentosPocoStock() {
    const url = 'http://localhost:8080/alimentos/?StockMenorCantidadMinima=true';
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