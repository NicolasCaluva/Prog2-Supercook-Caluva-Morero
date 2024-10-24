document.addEventListener('DOMContentLoaded', async function () {
    const listaCompras = document.getElementById('lista-compras');
    await obtenerListaAlimentosPocoStock();

    const confirmarButton = document.querySelector('.modal-footer .btn-primary');
    confirmarButton.addEventListener('click', enviarCompraDto);
});

function successCargarListaCompras(response) {
    const listaCompras = document.getElementById('lista-compras');
    response.forEach(compra => {
        console.log(compra);
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
}

function errorCargarListaAlimentosPocoStock(status, response) {
    console.log("Falla al cargar la lista:", response);
}

async function obtenerListaAlimentosPocoStock() {
    let url = 'http://localhost:8080/compras/';
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

    if (alimentosComprados.length === 0) {
        alert('No has seleccionado alimentos o cantidades válidas.');
        return;
    }

    const compraDto = {
        Alimentos: alimentosComprados
    };

    console.log('CompraDto a enviar:', compraDto);

    let url = 'http://localhost:8080/compras/';
    await makeRequest(url, Method.POST, compraDto, ContentType.JSON, CallType.PRIVATE, successEnviarCompra, errorEnviarCompra);
}

function successEnviarCompra(response) {
    alert('Compra realizada con éxito.');
    location.reload();
}

function errorEnviarCompra(status, response) {
    console.log('Error al enviar la compra:', response);
    alert('Hubo un error al realizar la compra.');
}