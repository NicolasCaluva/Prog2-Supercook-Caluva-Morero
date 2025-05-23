document.addEventListener('DOMContentLoaded', async function () {
    // Variables y primer carga de datos

    let nombre = document.getElementById('nombre');
    let listaAlimentos = document.getElementById('lista-alimentos');
    let momento = document.getElementById('momento');
    let listaAlimentosSeleccionados = document.getElementById('lista-alimentos-seleccionados');
    let confirmarRecetaBtn = document.getElementById('confirmarReceta');
    const agregarRecetaBtn = document.getElementById('agregarNuevaReceta');

    try {
        await obtenerListaRecetas();
    } catch (error) {
        console.log('Error:', error);
    }

    agregarRecetaBtn.addEventListener('click', function () {
        momento.addEventListener('change', momentoOnChange);
        confirmarRecetaBtn.addEventListener('click', () => confirmarFormularioReceta('POST'));
    });

    document.getElementById('filtrarRecetas').addEventListener('click', aplicarFiltros);

    const modal = document.getElementById('cargarReceta');
    modal.addEventListener('hidden.bs.modal', function () {
        nombre.value = '';
        momento.value = '';
        listaAlimentosSeleccionados.innerHTML = '';
        listaAlimentos.innerHTML = '';
        confirmarRecetaBtn.value = '';
        confirmarRecetaBtn.removeEventListener('click', confirmarFormularioReceta);
    });

    document.getElementById('pagAnterior').addEventListener('click', () => cambiarPagina(-1));
    document.getElementById('pagSiguiente').addEventListener('click', () => cambiarPagina(1));

    // Funciones

    async function obtenerListaRecetas(page) {
        page = page || 1;
        let URL = 'http://localhost:8080/recetas/?page=' + page;
        await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successObtenerListaRecetas, errorObtenerListaRecetas);
    }

    function successObtenerListaRecetas(response) {
        let listaRecetas = document.getElementById('lista-recetas');
        listaRecetas.innerHTML = '';
        let recetas = response.RecetaDto;
        let paginaTotal = response.PaginasTotales;
        let pagActual = response.NroPagina;

        let pagAnteriorButton = document.getElementById('pagAnterior');
        pagAnteriorButton.disabled = pagActual === 1;
        if (pagActual !== 1) {
            pagAnteriorButton.addEventListener('click', () => obtenerListaRecetas(pagActual - 1));
        }

        let pagSiguienteButton = document.getElementById('pagSiguiente');
        pagSiguienteButton.disabled = pagActual === paginaTotal;
        if (pagActual !== paginaTotal) {
            pagSiguienteButton.addEventListener('click', () => obtenerListaRecetas(pagActual + 1));
        }

        let contadorPaginas = document.getElementById('indicadorPagina');
        contadorPaginas.textContent = `${pagActual}/${paginaTotal}`;
        contadorPaginas.classList.add('fs-8');

        recetas.forEach(receta => {
            const tr = document.createElement('tr');
            tr.innerHTML = `
            <td class="text-secondary text-sm text-center">${receta.Nombre}</td>
            <td class="text-secondary text-sm text-center">${receta.Alimentos.map(alimento => alimento.Nombre).join(', ')}</td>
            <td class="text-secondary text-sm text-center">${receta.Momento.charAt(0).toUpperCase() + receta.Momento.slice(1)}</td>
        `;

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
            const divBotones = document.createElement('div');

            tdBotones.appendChild(divBotones);
            divBotones.appendChild(botonEliminar);
            divBotones.setAttribute('class', 'd-flex align-items-center');
            tdBotones.setAttribute('class', 'd-flex justify-content-center align-items-center pt-2 pb-0 mb-0');
            tr.appendChild(tdBotones);
            listaRecetas.appendChild(tr);
        });
    }

    function errorObtenerListaRecetas(status, response) {
        console.error("Falla al cargar la lista:", response);
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
        // Recibe todos los alimentos que se pueden agregar a la receta
        // Carga las opciones de alimentos en el selector de alimentos según el momento del día recibido
        let listaAlimentosSeleccionados = document.getElementById('lista-alimentos-seleccionados');
        let listaAlimentos = document.getElementById('lista-alimentos');

        response.AlimentosDto.forEach(alimento => {
            if (Array.from(listaAlimentosSeleccionados.querySelectorAll('input')).find(input => input.id === alimento.IdAlimento)) {
                return;
            }
            const option = document.createElement('option');
            option.value = alimento.IdAlimento;
            option.innerText = alimento.Nombre;
            option.addEventListener('click', () => agregarAlimentoAReceta(option, alimento));
            listaAlimentos.appendChild(option);
        });
    }

    function agregarAlimentoAReceta(option, alimento) {
        // Agrega los alimentos a la lista de alimentos seleccionados y permite eliminarlos de la lista de alimentos
        // y agregarlos nuevamente al selector de alimentos
        let listaAlimentosSeleccionados = document.getElementById('lista-alimentos-seleccionados');
        let listaAlimentos = document.getElementById('lista-alimentos');

        if (option) {
            listaAlimentos.removeChild(option);
        }

        const div = document.createElement('div');
        div.setAttribute('class', 'row mt-3');
        const div2 = document.createElement('div');
        div2.setAttribute('class', 'col-2 d-flex align-items-center justify-content-center');

        const label = document.createElement('label');
        label.innerText = alimento.Nombre;
        label.setAttribute('for', alimento.IdAlimento);
        label.setAttribute('class', 'col-9 d-flex flex-column');

        const input = document.createElement('input');
        input.setAttribute('type', 'number');
        input.setAttribute('id', alimento.IdAlimento);
        input.setAttribute('min', 1);
        input.setAttribute('value', alimento.Cantidad ? alimento.Cantidad : 1);
        input.setAttribute('class', 'form-control form-control-sm custom-border');
        input.setAttribute('required', true);

        const botonEliminarSeleccionado = document.createElement('button');
        const iconoEliminar = document.createElement('i');
        botonEliminarSeleccionado.setAttribute('class', 'btn btn-danger btn-sm mb-0');
        iconoEliminar.setAttribute('class', 'fa-solid fa-trash fs-8');
        botonEliminarSeleccionado.appendChild(iconoEliminar);
        botonEliminarSeleccionado.setAttribute('value', alimento.IdAlimento);
        botonEliminarSeleccionado.addEventListener('click', function () {
            if (!option) {
                option = document.createElement('option');
                option.value = this.value;
                option.innerText = alimento.Nombre;
                option.addEventListener('click', () => agregarAlimentoAReceta(option, alimento));
            }
            listaAlimentos.appendChild(option);
            div.remove();
        });
        label.appendChild(input);
        div.appendChild(label);
        div2.appendChild(botonEliminarSeleccionado);
        div.appendChild(div2);
        listaAlimentosSeleccionados.appendChild(div);
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


    async function confirmarFormularioReceta() {
        let nombre = document.getElementById('nombre');
        let listaAlimentosSeleccionados = document.getElementById('lista-alimentos-seleccionados');
        let momento = document.getElementById('momento');

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
    }

    async function momentoOnChange() {
        let listaAlimentosSeleccionados = document.getElementById('lista-alimentos-seleccionados');
        let momento = document.getElementById('momento');
        let listaAlimentos = document.getElementById('lista-alimentos');

        if (listaAlimentosSeleccionados.hasChildNodes()) {
            listaAlimentosSeleccionados.innerHTML = '';
        }

        if (listaAlimentos.hasChildNodes()) {
            listaAlimentos.innerHTML = '';
        }
        await cargarAlimentos(momento.value)
    }

    async function aplicarFiltros() {
        const nombre = document.getElementById('filtro-nombre').value;
        const momento = document.getElementById('filtro-momento').value;
        const tipo_alimento = document.getElementById('filtro-tipo-alimento').value;
        const URL = `http://localhost:8080/recetas/?nombre=${nombre}&momento=${momento}&tipoAlimento=${tipo_alimento}`;
        await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, successObtenerListaRecetas, errorObtenerListaRecetas);
    }
});