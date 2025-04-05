document.addEventListener('DOMContentLoaded', function () {
    obtenerRecetasPorMomento();
    obtenerRecetasPorTipoDeAlimento();
    obtenerRecetasPorBeneficio();
    let botonEnviar = document.getElementById('botonEnviar');
    botonEnviar.addEventListener('click', obtenerRecetasPorBeneficio);
});
let myChart;

async function obtenerRecetasPorMomento() {
    const URL = 'http://localhost:8080/recetas/contarRecetasPorMomento/';
    await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, function (response) {
        if (response && typeof response === 'object') {
            const labels = Object.keys(response);
            const values = Object.values(response);

            const ctx = document.getElementById('graficoBarraTipoUso').getContext('2d');
            var graficoTipoUso = new Chart(ctx, {
                type: 'bar',
                data: {
                    labels: labels,
                    datasets: [{
                        label: 'Cantidad de Recetas por Momento',
                        data: values,
                        backgroundColor: [
                            'rgba(255, 0, 0, 0.2)',
                            'rgba(0, 0, 255, 0.2)',
                            'rgba(255, 255, 0, 0.2)',
                            'rgba(0, 255, 0, 0.2)',
                        ],
                    }]
                },
            });
        } else {
            console.error('Invalid response format:', response);
        }
    }, function (status, response) {
        console.error('Error consultado:', response);
    });
}

async function obtenerRecetasPorTipoDeAlimento() {
    const URL = 'http://localhost:8080/recetas/contarRecetasPorTipoAlimento/';
    await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, function (response) {
        if (response && typeof response === 'object') {
            const labels = Object.keys(response);
            const values = Object.values(response);

            const ctx = document.getElementById('graficoBarraTipoAlimentos').getContext('2d');
            var graficoTipoAlimento = new Chart(ctx, {
                type: 'bar',
                data: {
                    labels: labels,
                    datasets: [{
                        label: 'Cantidad de Recetas por Tipo de Alimento',
                        data: values,
                        backgroundColor: [
                            'rgba(255, 0, 0, 0.2)',
                            'rgba(0, 0, 255, 0.2)',
                            'rgba(255, 255, 0, 0.2)',
                            'rgba(0, 255, 0, 0.2)',
                            'rgba(128, 0, 128, 0.2)',
                            'rgba(255, 165, 0, 0.2)'
                        ],
                    }]
                },
            });
        } else {
            console.error('Invalid response format:', response);
        }
    }, function (status, response) {
        console.error('Error consultado:', response);
    });
}

async function obtenerRecetasPorBeneficio() {
    let fechaInicio = document.getElementById('FechaInicial').value;
    let fechaFinal = document.getElementById('FechaFinal').value;
    const URL = `http://localhost:8080/compras/montoTotalEntreFechas/?fechaInicio=${fechaInicio}&fechaFin=${fechaFinal}`;

    await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, function (response) {
        if (response && typeof response === 'object') {
            const labels = Object.keys(response.montoTotal);
            const values = Object.values(response.montoTotal);
            const ctx = document.getElementById('graficoBarraBeneficio').getContext('2d');
            if (myChart) {
                myChart.destroy();
            }
            myChart = new Chart(ctx, {
                type: 'bar',
                data: {
                    labels: labels,
                    datasets: [{
                        label: 'Montos Totales de Beneficio',
                        data: values,
                        backgroundColor: [
                            'rgba(255, 0, 0, 0.2)',
                            'rgba(0, 0, 255, 0.2)',
                            'rgba(255, 255, 0, 0.2)',
                            'rgba(0, 255, 0, 0.2)',
                            'rgba(128, 0, 128, 0.2)',
                            'rgba(255, 165, 0, 0.2)',
                            'rgba(0, 255, 255, 0.2)',
                            'rgba(255, 0, 255, 0.2)',
                            'rgba(128, 128, 128, 0.2)',
                            'rgba(255, 255, 255, 0.2)',
                            'rgba(0, 0, 0, 0.2)',
                            'rgba(128, 128, 0, 0.2)',
                        ],
                    }]
                },
            });
        } else {
            console.error('Invalid response format:', response);
        }
    }, function (status, response) {
        console.error('Error consultado:', response);
    });
}
