document.addEventListener('DOMContentLoaded', async function () {
    const userinfo = document.getElementById('user-info');
    const rol = document.getElementById('rol');
    userinfo.innerHTML = localStorage.getItem('username');
    rol.innerHTML = localStorage.getItem('rol');
});