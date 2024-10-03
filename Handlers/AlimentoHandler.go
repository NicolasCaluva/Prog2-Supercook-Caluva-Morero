package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"supercook/Dto"
	"supercook/Services"
)

type AlimentoHandler struct {
	AlimentoService Services.AlimentoInteface
}

func NuevoAlimentoHandler(alimentoService Services.AlimentoInteface) *AlimentoHandler {
	return &AlimentoHandler{
		AlimentoService: alimentoService,
	}
}
func (handler *AlimentoHandler) ObtenerAlimentos(c *gin.Context) {
	alimentos := handler.AlimentoService.ObtenerAlimentos()
	c.JSON(http.StatusOK, alimentos)
}

func (handler *AlimentoHandler) ObtenerAlimentoPorID(c *gin.Context) {
	id := c.Param("id")
	alimento := handler.AlimentoService.ObtenerAlimentoPorID(id)
	c.JSON(http.StatusOK, alimento)
}

func (handler *AlimentoHandler) CrearAlimento(c *gin.Context) {
	var alimentoDto Dto.AlimentoDto

	if err := c.BindJSON(&alimentoDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.AlimentoService.CrearAlimento(&alimentoDto)
	c.JSON(http.StatusOK, resultado)
}

func (handler *AlimentoHandler) ActualizarAlimento(c *gin.Context) {
	id := c.Param("id")
	var alimentoDto Dto.AlimentoDto
	c.BindJSON(&alimentoDto)
	resultado := handler.AlimentoService.ActualizarAlimento(id, &alimentoDto)
	c.JSON(http.StatusOK, resultado)
}

func (handler *AlimentoHandler) EliminarAlimento(c *gin.Context) {
	id := c.Param("id")
	resultado := handler.AlimentoService.EliminarAlimento(id)
	c.JSON(http.StatusOK, resultado)
}
