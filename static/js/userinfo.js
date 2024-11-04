const LOGIN_URL = '/Prog2-Supercook-Caluva-Morero/templates/login.html'

document.addEventListener('DOMContentLoaded', async function () {
    const nav_alimentos = document.getElementById('nav-alimentos');
    const nav_recetas = document.getElementById('nav-recetas');
    const nav_compras = document.getElementById('nav-compras');
    const userinfo = document.getElementById('user-info');
    const rol = document.getElementById('rol');
    const logout = document.getElementById('logout');
    userinfo.innerHTML = localStorage.getItem('username');
    rol.innerHTML = localStorage.getItem('rol');
    logout.addEventListener('click', () => {
        localStorage.clear();
        location.href = LOGIN_URL;
    });

    const lastSegment = location.href.split('/').pop();
    console.log(lastSegment)
    if (lastSegment === 'alimentos.html') {
        nav_alimentos.classList.add('active');
    }
    if (lastSegment === 'recetas.html') {
        nav_recetas.classList.add('active');
    }
    if (lastSegment === 'compras.html') {
        nav_compras.classList.add('active');
    }
});