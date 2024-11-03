package Handlers

import (
	"github.com/gin-gonic/gin"
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
	if userInfo == nil {
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}
	var compraDto Dto.CompraDto
	err := c.BindJSON(&compraDto)
	compraDto.IDUsuario = userInfo.Codigo
	if err != nil {
		c.Error(Errors.ErrorJsonInvalidoCompras)
		return
	}
	resultado := handler.CompraService.AgregarCompra(&compraDto)
	if resultado != nil {
		c.Error(*resultado)
		return
	}
	c.JSON(http.StatusOK, resultado)
}
