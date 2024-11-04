package Middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"supercook/Errors"
)

func ErrorMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err
		var errorCodigo *Errors.ErrorCodigo
		if ok := errors.As(err, &errorCodigo); ok {
			c.JSON(httpStatusFromCode(errorCodigo.Codigo), gin.H{"error": modificarMensajeID500(errorCodigo)})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}
}
func modificarMensajeID500(codigo *Errors.ErrorCodigo) string {
	if codigo.Codigo == "ERR_500" {
		return "Error interno del servidor"
	}
	return codigo.Mensaje
}
func httpStatusFromCode(codigo string) int {
	switch codigo {
	case "ERR_500":
		return http.StatusInternalServerError
	case "ERR_401":
		return http.StatusUnauthorized
	case "ERR_404":
		return http.StatusNotFound
	case "ERR_400":
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
