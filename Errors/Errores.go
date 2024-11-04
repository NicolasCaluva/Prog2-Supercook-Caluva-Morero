package Errors

import "fmt"

// ErrorCodigo define una estructura para errores personalizados
type ErrorCodigo struct {
	Codigo  string
	Mensaje string
}

// Error implementa la interfaz de error para ErrorCodigo
func (error *ErrorCodigo) Error() string {
	return fmt.Sprintf("%s: %s", error.Codigo, error.Mensaje)
}

// New crea una nueva instancia de ErrorCodigo
func NuevoErrorCodigo(codigo, mensaje string) *ErrorCodigo {
	return &ErrorCodigo{
		Codigo:  codigo,
		Mensaje: mensaje,
	}
}

// Definimos algunos errores personalizados
var (
	//Errores Repositories
	ErrorConectarBD          = NuevoErrorCodigo("ERR_500", "Error al conectar a la base de datos")
	ErrorDecodificarAlimento = NuevoErrorCodigo("ERR_500", "Error al decodificar el alimento desde la Base de Datos")
	ErrorDecodificarReceta   = NuevoErrorCodigo("ERR_500", "Error al decodificar la receta desde la Base de Datos")
	// Errores Handlers
	ErrorUsuarioNoAutenticado = NuevoErrorCodigo("ERR_401", "Usuario no autenticado")
	// Errores Alimentos
	ErrorListaVaciaDeAlimentos               = NuevoErrorCodigo("ERR_404", "Lista vacia de alimentos")
	ErrorAlimentoNoEncontrado                = NuevoErrorCodigo("ERR_404", "Alimento no encontrado")
	ErrorAlimentoNoEncontradoActualizar      = NuevoErrorCodigo("ERR_404", "Alimento no encontrado para actualizar")
	ErrorAlimentoNoEncontradoEliminar        = NuevoErrorCodigo("ERR_404", "Alimento no encontrado para eliminar")
	ErrorAlimentoNombreMalIngresado          = NuevoErrorCodigo("ERR_400", "El nombre del alimento no puede estar vacío")
	ErrorAlimentoPrecioUnitarioMalIngresado  = NuevoErrorCodigo("ERR_400", "El precio unitario del alimento no puede ser menor a 0")
	ErrorAlimentoStockMalIngresado           = NuevoErrorCodigo("ERR_400", "El stock del alimento no puede ser menor a 0")
	ErrorAlimentoCantMinimaStockMalIngresado = NuevoErrorCodigo("ERR_400", "La cantidad mínima de stock del alimento no puede ser menor a 0")
	ErrorAlimentoTipoAlimentoMalIngresado    = NuevoErrorCodigo("ERR_400", "El tipo de alimento no puede estar vacío")
	ErrorAlimentoMomentoDelDiaMalIngresado   = NuevoErrorCodigo("ERR_400", "El momento del día no puede estar vacío")
	ErrorJsonInvalidoAlimento                = NuevoErrorCodigo("ERR_400", "Error en el JSON de alimento")

	// Errores Compras
	ErrorListaVaciaDeCompras = NuevoErrorCodigo("ERR_404", "Lista vacia de alimentos a comprar")
	ErrorJsonInvalidoCompras = NuevoErrorCodigo("ERR_400", "Error en el JSON de compra")
	// Errores Recetas
	ErrorListaVaciaDeRecetas             = NuevoErrorCodigo("ERR_404", "Lista vacia de recetas")
	ErrorRecetaNoEncontrada              = NuevoErrorCodigo("ERR_404", "Receta no encontrada")
	ErrorRecetaNoEncontradoActualizar    = NuevoErrorCodigo("ERR_404", "Receta no encontrada para actualizar")
	ErrorRecetaNoEncontradoEliminar      = NuevoErrorCodigo("ERR_404", "Receta no encontrado para eliminar")
	ErrorRecetaNombreMalIngresado        = NuevoErrorCodigo("ERR_400", "El nombre de la receta no puede estar vacío")
	ErrorRecetaAlimentosMalIngresados    = NuevoErrorCodigo("ERR_400", "Debe haber al menos un alimento en la receta")
	ErrorRecetaMomentoDelDiaMalIngresado = NuevoErrorCodigo("ERR_400", "Debe haber un momento del día en la receta")
	ErrorJsonInvalidoReceta              = NuevoErrorCodigo("ERR_400", "Error en el JSON de receta")
)
