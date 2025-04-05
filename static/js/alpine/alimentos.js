document.addEventListener('alpine:init', () => {
    Alpine.data('alimentos', () => ({
        alimentos: [],

        // init
        async init() {
           await this.obtenerAlimentos()
        },

        // methods
        async obtenerAlimentos() {
            const URL = url.ALIMENTOS;
            await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PRIVATE, (response) => {
                this.alimentos = response;
                },
                () => {
                    console.error("Error al obtener la lista de alimentos");
                });
        },

    }));
});