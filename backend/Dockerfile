# Imagen base oficial de Go
FROM golang:1.23-alpine

# Crea directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos del backend
COPY . ./

# Descarga las dependencias
RUN go mod download

# Compila el binario
RUN go build -o app

# Expone el puerto que usará el backend
EXPOSE 8080

# Comando por defecto al iniciar el contenedor
CMD ["./app"]