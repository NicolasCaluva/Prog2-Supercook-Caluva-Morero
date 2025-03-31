document.addEventListener('alpine:init', () => {
    Alpine.data('alimentos', () => ({
        alimentos: [],

        // init
        async init() {
            const URL = '/alimentos/';
            await makeRequest(URL, Method.GET, null, ContentType.JSON, CallType.PUBLIC, this.successObtenerAlimentos, this.errorObtenerAlimentos);
        },

        // obtener alimentos

    }));
});