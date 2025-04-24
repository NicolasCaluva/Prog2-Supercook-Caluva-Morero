document.addEventListener('DOMContentLoaded', async function () {
    const LOGIN_URL = '/Prog2-Supercook-Caluva-Morero/templates/login.html'
    const nav_alimentos = document.getElementById('nav-alimentos');
    const nav_recetas = document.getElementById('nav-recetas');
    const nav_compras = document.getElementById('nav-compras');
    const userinfo = document.getElementById('user-info');
    const logout = document.getElementById('logout');
    userinfo.innerHTML = localStorage.getItem('username');
    logout.addEventListener('click', () => {
        localStorage.clear();
        location.href = LOGIN_URL;
    });

    const lastSegment = location.href.split('/').pop();
    if (lastSegment.includes('alimentos.html')) {
        nav_alimentos.classList.add('nav-link', 'active', 'bg-gradient-dark', 'text-white');
    }
    if (lastSegment.includes('recetas.html')) {
        nav_recetas.classList.add('nav-link', 'active', 'bg-gradient-dark', 'text-white');
    }
    if (lastSegment.includes('compras.html')) {
        nav_compras.classList.add('nav-link', 'active', 'bg-gradient-dark', 'text-white');
    }
});