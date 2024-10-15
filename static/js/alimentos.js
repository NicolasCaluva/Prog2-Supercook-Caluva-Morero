document.addEventListener('DOMContentLoaded', function () {
    const listaAlimentos = document.getElementById('lista-alimentos');


    const tr = document.createElement('tr');
    let token = localStorage.getItem('authToken')

    makeRequest('http://localhost:8080/alimentos/', 'GET', null, JSON, CallType.PRIVATE, successFn, errorFn);
});


function successFn(response) {
    let lista_alimentos = document.getElementById('lista-alimentos');
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
        lista_alimentos.appendChild(tr);
    });
    console.log("Éxito", response);
}

function errorFn(response) {
    console.log("Ocurrió un error:", response);
}