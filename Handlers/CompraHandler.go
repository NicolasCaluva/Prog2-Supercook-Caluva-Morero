package Handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"supercook/Dto"
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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	var compraDto Dto.CompraDto
	err := c.BindJSON(&compraDto)
	compraDto.IDUsuario = userInfo.Codigo
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error en el JSON"})
		return
	}
	resultado := handler.CompraService.AgregarCompra(&compraDto)
	if !resultado.BoolResultado {
		c.JSON(http.StatusBadRequest, resultado)
		return
	}
	c.JSON(http.StatusOK, resultado)
}
func (handler *CompraHandler) ObtenerListaAlimentosStockMenorStockMinimo(c *gin.Context) {
	log.Printf("ObtenerListaAlimentosStockMenorStockMinimo")
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	alimentos := handler.CompraService.ObtenerListaAlimentosStockMenorStockMinimo(&userInfo.Codigo)
	c.JSON(http.StatusOK, alimentos)
}
