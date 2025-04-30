document.addEventListener('DOMContentLoaded', async function () {
    const eye_password = document.getElementById('eye-password');
    const eye_icon = document.getElementById('eye-icon');
    const password = document.getElementById('password');
    eye_password.addEventListener('click', () => {
        if (password.type === 'password') {
            password.type = 'text';
            eye_icon.classList.remove('fa-eye-slash');
            eye_icon.classList.add('fa-eye');
        } else {
            password.type = 'password';
            eye_icon.classList.remove('fa-eye');
            eye_icon.classList.add('fa-eye-slash');
        }
    });

    password.addEventListener('focus', () => {
        eye_icon.classList.remove('fc-white');
        eye_icon.classList.add('fc-black');
    });

    password.addEventListener('blur', () => {
        eye_icon.classList.remove('fc-black');
        eye_icon.classList.add('fc-white');
    });

});