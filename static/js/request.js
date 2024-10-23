// Version: 20231017_2140
const Method = {
    GET: "GET",
    POST: "POST",
    PUT: "PUT",
    DELETE: "DELETE",
};

const ContentType = {
    JSON: "application/json",
    URL_ENCODED: "application/x-www-form-urlencoded",
};

const CallType = {
    PUBLIC: false,
    PRIVATE: true,
};

function makeBody(contentType, data) {
    if (!data) return null;

    if (contentType == ContentType.JSON)
        return data ? JSON.stringify(data) : null;

    if (contentType == ContentType.URL_ENCODED) {
        const keys = Object.keys(data);
        const values = Object.values(data);
        const pairs = [];
        for (let i = 0; i < keys.length; i++) {
            pairs.push(
                encodeURIComponent(keys[i]) + "=" + encodeURIComponent(values[i])
            );
        }
        return pairs.join("&");
    }
}

// Definir la función que realiza la solicitud HTTP
async function makeRequest(
    url,
    method,
    data,
    contentType,
    isPrivateCall,
    successCallback,
    errorCallback
) {
    // Validar que el método sea válido
    if (!["GET", "POST", "PUT", "DELETE"].includes(method)) {
        throw new Error("Método HTTP no válido");
        return;
    }
    debugger
    // Validar que los datos no estén vacíos en caso de POST o PUT
    if ((method === "POST" || method === "PUT") && !data) {
        throw new Error(
            "Los datos del cuerpo no pueden estar vacíos en POST o PUT"
        );
        return;
    }

    //Validar que si la pegada es privada exista bearer token, si existe lo agrega como header authorization, sino va a login.html
    const token = localStorage.getItem("authToken");
    if (isPrivateCall && !token) {
        window.location = "login.html?reason=private_call_without_token";
        return;
    }

    // Realizar la solicitud HTTP
    try {
        const response = await fetch(url, {
            method: method,
            body: makeBody(contentType, data),
            headers: {
                "Content-Type": contentType,
                Authorization: (token && isPrivateCall) ? `Bearer ${token}` : null,
            },
        });

        let responseBody = {};
        try {
            responseBody = await response.json();
        } catch {
        }

        if ("access_token" in responseBody) {
            localStorage.setItem("authToken", responseBody.access_token);
        }

        if (response.ok) {
            console.info(
                `Llamada OK: ${url}: status: ${response.status} - ${JSON.stringify(
                    responseBody
                )}`
            );
            successCallback(responseBody);
        } else {
            console.error(
                `Error en llamada: ${url}: status: ${
                    response.status
                } - ${JSON.stringify(responseBody)}`
            );

            if (response.status === 401) {
                window.location = "login.html?reason=token_invalid";
            }

            errorCallback(response.status, responseBody);
        }
    } catch (error) {
        console.error("Error: ", error?.message || error)
        throw new Error(`Request ERROR: ${error?.message || error}`);
    }
}

function isUserLogged() {
    return localStorage.getItem("authToken") != null;
}