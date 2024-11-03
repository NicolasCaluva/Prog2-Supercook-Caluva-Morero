package Middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"supercook/Errors"
)

// ErrorMiddleware es un middleware que maneja los errores
func ErrorMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		err := c.Errors.Last().Err
		var errorCodigo *Errors.ErrorCodigo
		if ok := errors.As(err, &errorCodigo); ok {
			c.JSON(httpStatusFromCode(errorCodigo.Codigo), gin.H{"error": errorCodigo.Mensaje})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
	}
}

// httpStatusFromCode convierte un código de error en un código de estado HTTP
func httpStatusFromCode(codigo string) int {
	switch codigo {
	case "ERR_1":
		return http.StatusInternalServerError
	case "ERR_2":
		return http.StatusBadRequest
	case "ERR_10":
		return http.StatusUnauthorized
	case "ERR_50":
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
