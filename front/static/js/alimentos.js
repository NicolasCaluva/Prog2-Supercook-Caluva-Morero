// noinspection LanguageDetectionInspection

document.addEventListener('DOMContentLoaded', async function () {
    // Variables globales y carga de alimentos con el DOM

    let listaAlimentos;

    listaAlimentos = document.getElementById('lista-alimentos');
    const modal = document.getElementById('cargarAlimento');

    await obtenerListaAlimentos();

    const agregarAlimentoBtn = document.getElementById('agregarNuevoAlimento');
    agregarAlimentoBtn.addEventListener('click', function () {
        document.getElementById('form-alimento').reset();
        const confirmarAlimentoBtn = document.getElementById('confirmarAlimento');
        confirmarAlimentoBtn.addEventListener('click', () => confirmarNuevoAlimento())
    });

    modal.addEventListener('hidden.bs.modal', function () {
        document.getElementById('form-alimento').reset();
        const confirmarAlimentoBtn = document.getElementById('confirmarAlimento');
        confirmarAlimentoBtn.replaceWith(confirmarAlimentoBtn.cloneNode(true));
    });


    // Funciones

    async function obtenerListaAlimentos(page) {
        page = page || 1;
        const URL = 'http://localhost:8080/alimentos/?page=' + page;
        await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successCargarListaAlimentos, response => console.error('Error: ', response.error));
    }

    function successCargarListaAlimentos(response) {
        listaAlimentos.innerHTML = '';
        let alimentos = response.AlimentosDto;
        let paginaTotal = response.PaginasTotales;
        let pagActual = response.NroPagina;

        let pagAnteriorButton = document.getElementById('pagAnterior');
        pagAnteriorButton.disabled = pagActual === 1;
        if (pagActual !== 1) {
            pagAnteriorButton.addEventListener('click', () => obtenerListaAlimentos(pagActual - 1));
        }

        let pagSiguienteButton = document.getElementById('pagSiguiente');
        pagSiguienteButton.disabled = pagActual === paginaTotal;
        if (pagActual !== paginaTotal) {
            pagSiguienteButton.addEventListener('click', () => obtenerListaAlimentos(pagActual + 1));
        }

        let contadorPaginas = document.getElementById('numeroPagina');
        contadorPaginas.textContent = `${pagActual}/${paginaTotal}`;
        contadorPaginas.classList.add('fs-8');


        alimentos.forEach(alimento => {
            const tr = document.createElement('tr');
            tr.innerHTML = `
                            <td class="text-secondary text-sm text-center">${alimento.Nombre}</td>
                            <td class="text-secondary text-sm text-center">$${alimento.PrecioUnitario}</td>
                            <td class="text-secondary text-sm text-center">${alimento.Stock}</td>
                            <td class="text-secondary text-sm text-center">${alimento.CantMinimaStock}</td>
                            <td class="text-secondary text-sm text-center">${alimento.TipoAlimento.charAt(0).toUpperCase() + alimento.TipoAlimento.slice(1)}</td>
                            <td class="text-secondary text-sm text-center">${alimento.MomentoDelDia.map(momento => momento.charAt(0).toUpperCase() + momento.slice(1)).join(' - ')}</td>
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
            const divBotones = document.createElement('div');
            divBotones.setAttribute('class', 'd-flex gap-2 pt-3');
            tdBotones.appendChild(divBotones);
            divBotones.appendChild(botonEditar);
            divBotones.appendChild(botonEliminar);
            tr.appendChild(tdBotones);
            listaAlimentos.appendChild(tr);
        });
    }

    function successObtenerAlimento(alimento) {
        let nombreAlimento = document.getElementById('nombreAlimento');
        let precioUnitario = document.getElementById('precioUnitario');
        let stock = document.getElementById('stock');
        let cantMinimaStock = document.getElementById('cantMinimaStock');
        let tipoAlimento = document.getElementById('tipoAlimento');
        let momentoDelDia = document.getElementById('momentoDelDia');

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

        confirmarAlimentoBtn.addEventListener('click', () => confirmarEdicionAlimento(alimento));
    }

    function errorObtenerAlimento(status, response) {
        alert(response.error);
    }

    function successEditarAlimento() {
        alert("Alimento editado correctamente");
        obtenerListaAlimentos();
    }

    function errorEditarAlimento(status, response) {
        alert(response.error);
    }

    function successEliminarAlimento() {
        alert("Alimento eliminado correctamente");
        obtenerListaAlimentos();
    }

    function errorEliminarAlimento(status, response) {
        alert(response.error);
    }

    async function confirmarNuevoAlimento() {
        let nombreAlimento = document.getElementById('nombreAlimento');
        let precioUnitario = document.getElementById('precioUnitario');
        let stock = document.getElementById('stock');
        let cantMinimaStock = document.getElementById('cantMinimaStock');
        let tipoAlimento = document.getElementById('tipoAlimento');
        let momentoDelDia = document.getElementById('momentoDelDia');


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
        await makeRequest(URL, 'POST', nuevoAlimento, ContentType.JSON, CallType.PRIVATE, () => {
            alert("Alimento cargado correctamente");
        }, () => {
            alert("Error al cargar el alimento");
        });

        location.reload();
    }

    async function confirmarEdicionAlimento(alimento) {
        let nombreAlimento = document.getElementById('nombreAlimento');
        let precioUnitario = document.getElementById('precioUnitario');
        let stock = document.getElementById('stock');
        let cantMinimaStock = document.getElementById('cantMinimaStock');
        let tipoAlimento = document.getElementById('tipoAlimento');

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
        const modal = document.getElementById('cargarAlimento');
        modal.addEventListener('hidden.bs.modal', function () {
            document.getElementById('form-alimento').reset();
            const confirmarAlimentoBtn = document.getElementById('confirmarAlimento');
            confirmarAlimentoBtn.removeEventListener('click', null);
        });
        document.getElementById('form-alimento').reset();
        location.reload();
    }
});

