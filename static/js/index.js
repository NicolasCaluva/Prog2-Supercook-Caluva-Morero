document.addEventListener('DOMContentLoaded', function () {
    const ctx = document.getElementById('graficoBarraTipoUso').getContext('2d');
    var myChart = new Chart(ctx, {
        type: 'bar',
        data: {
            labels: ['Red', 'Blue', 'Yellow', 'Green', 'Purple', 'Orange'],
            datasets: [{
                label: '# of Votes',
                data: [12, 19, 3, 5, 2, 3],
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
});
