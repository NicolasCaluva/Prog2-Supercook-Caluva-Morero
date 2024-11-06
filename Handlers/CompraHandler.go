package Handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"supercook/Dto"
	"supercook/Errors"
	"supercook/Services"
	"supercook/Utils"
)

type CompraHandler struct {
	CompraService Services.CompraInterfaz
}

func NuevoCompraHandler(compraService Services.CompraInterfaz) *CompraHandler {
	return &CompraHandler{
		CompraService: compraService,
	}
}
func (handler *CompraHandler) CrearCompra(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	var compraDto Dto.CompraDto
	err := c.BindJSON(&compraDto)
	compraDto.IDUsuario = userInfo.Codigo
	if err != nil {
		c.Error(Errors.ErrorJsonInvalidoCompras)
		return
	}
	log.Printf("Iniciando creación de compra con ID: %s, usuario: %s", compraDto.IDUsuario, userInfo)
	resultado := handler.CompraService.AgregarCompra(&compraDto)
	if resultado != nil {
		log.Printf("Error: %v\n", resultado)
		c.Error(resultado)
		return
	}
	log.Printf("Compra realizada con éxito\n")
	c.JSON(http.StatusOK, gin.H{"mensaje": "Compra realizada con éxito"})
}
