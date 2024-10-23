function successFn(response) {
    console.log(response)
    const listaAlimentos = document.getElementById('lista-alimentos');
    response.forEach(alimento => {
        const tr = document.createElement('tr');
        tr.innerHTML = `
                            <td>${alimento.Nombre}</td>
                            <td>${alimento.PrecioUnitario}</td>
                            <td>${alimento.Stock}</td>
                            <td>${alimento.CantMinimaStock}</td>
                            <td>${alimento.TipoAlimento}</td>
                            <td>${alimento.MomentoDelDia}</td>
                        `;
        const botonEditar = document.createElement('button');
        const iconoEditar = document.createElement('i');
        iconoEditar.setAttribute('class', 'fa-solid fa-pencil');

        Object.assign(botonEditar, {
            className: 'btn btn-primary',
            type: 'button',
            'data-bs-toggle': 'modal',
            'data-bs-target': '#cargarAlimento'
        });
        botonEditar.appendChild(iconoEditar);
        const botonEliminar = document.createElement('button');
        const iconoEliminar = document.createElement('i');

        Object.assign(botonEliminar, {
            className: 'btn btn-danger',
            type: 'button',
        });
        iconoEliminar.setAttribute('class', 'fa-solid fa-trash');
        botonEliminar.appendChild(iconoEliminar);

        const tdBotones = document.createElement('td');

        tdBotones.appendChild(botonEditar);
        tdBotones.appendChild(botonEliminar);
        tdBotones.setAttribute('class', 'd-flex gap-2')
        tr.appendChild(tdBotones);
        listaAlimentos.appendChild(tr);
    })
}

function errorFn(status, response) {
    console.log("Falla:", response);
}

document.addEventListener('DOMContentLoaded', async function () {
    let url = 'http://localhost:8080/alimentos/';
    await makeRequest(url, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successFn, errorFn);

    const agregarAlimentoBtn = document.querySelector('button[data-bs-target="#cargarAlimento"]');

    agregarAlimentoBtn.addEventListener('click', function () {
        document.getElementById('form-alimento').reset();
    });

    const confirmarAlimentoBtn = document.getElementById('confirmarAlimento');

    confirmarAlimentoBtn.addEventListener('click', async function () {
        const nombreAlimento = document.getElementById('nombreAlimento').value;
        const precioUnitario = document.getElementById('precioUnitario').value;
        const stock = document.getElementById('stock').value;
        const cantMinimaStock = document.getElementById('cantMinimaStock').value;
        const tipoAlimento = document.getElementById('tipoAlimento').value;

        const momentosSeleccionados = Array.from(document.getElementById('momentoDelDia').selectedOptions).map(option => option.value);

        if (nombreAlimento && precioUnitario && stock && cantMinimaStock && tipoAlimento && momentosSeleccionados.length > 0) {
            const nuevoAlimento = {
                Nombre: nombreAlimento,
                PrecioUnitario: parseFloat(precioUnitario),
                Stock: parseInt(stock),
                CantMinimaStock: parseInt(cantMinimaStock),
                TipoAlimento: tipoAlimento,
                MomentoDelDia: momentosSeleccionados
            };

            try {
                const urlPost = 'http://localhost:8080/alimentos/';
                await makeRequest(urlPost, 'POST', nuevoAlimento, ContentType.JSON, CallType.PRIVATE, successFnCargar, errorFn);
                const modal = bootstrap.Modal.getInstance(document.getElementById('cargarAlimento'));
                modal.hide();
                document.getElementById('form-alimento').reset();
            } catch (error) {
                console.error('Error al agregar el alimento:', error);
                alert('Hubo un error al agregar el alimento.');
            }
        } else {
            alert('Todos los campos son obligatorios');
        }
    });
});
function successFnCargar(response) {
    console.log(response);

    if (response.BoolResultado) {
        alert("Operación exitosa: \n" + response.ListaMensaje.join("\n"));
    } else {
        alert("Error: \n" + response.ListaMensaje.join("\n"));
    }
}