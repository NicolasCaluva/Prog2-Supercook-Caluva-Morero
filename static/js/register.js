// Ejemplo de uso
// TODO: Cambiar URL
const url = "http://w230847.ferozo.com/tp_prog2/api/account/register";

document.addEventListener("DOMContentLoaded", function (eventDOM) {
  document
    .getElementById("btnRegistrarse")
    .addEventListener("click", async function (eventClick) {
      eventClick.preventDefault();

      const emailInput = document.getElementById('email');

      // Validación de email con una expresión regular
      const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
      if (!emailRegex.test(emailInput.value)) {
        emailInput.classList.add('is-invalid');
        return false;
      } else {
        emailInput.classList.remove('is-invalid');
      }

      const inputPassword = document.getElementById("password").value;
      const inputConfirmPassword =
        document.getElementById("ConfirmPassword").value;

      if (inputPassword === "" || inputConfirmPassword === "" || inputPassword !== inputConfirmPassword) {
        alert("Password and Confirm Password must be the same and can not be empty");
        return false;
      }

      const data = {
        grant_type: "password",
        email: document.getElementById("email").value,
        password: inputPassword,
        ConfirmPassword: inputConfirmPassword,
        Role: "ADMIN" // TODO: Borrar a futuro, por ahora está OK
      };

      try {
        await makeRequest(
          url,
          Method.POST,
          data,
          ContentType.URL_ENCODED,
          CallType.PUBLIC,
          successFn,
          errorFn
        );
      } catch (error) {
        console.error("Error registering: ", error?.message || error)
      }

      return false;
    });
});

function successFn(response) {
  console.log("Éxito:", response);
  window.location = "index.html";
}

function errorFn(status, response) {
  console.log("Falla:", response);
  alert(response?.ModelState?.["model.Password"]?.[0] ||
    response?.ModelState?.["model.Role"]?.[0] ||
    response?.ModelState?.[""]?.[0] ||
    "Intente nuevamente")
}
