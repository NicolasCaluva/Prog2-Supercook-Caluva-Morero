package Handlers

import (
	"github.com/gin-gonic/gin"
	"supercook/Dto"
	"supercook/Services"
)

type RecetaHandler struct {
	RecetaService Services.RecetaInterface
}

func NuevoRecetaHandler(recetaService Services.RecetaInterface) *RecetaHandler {
	return &RecetaHandler{
		RecetaService: recetaService,
	}
}

func (handler *RecetaHandler) ObtenerRecetas(c *gin.Context) {
	filtro := [3]string{c.Query("tipo"), c.Query("nombre")}
	userInfo := c.Request.Header.Get("Authorization")
	var recetas = handler.RecetaService.ObtenerRecetas(&filtro, &userInfo)
	c.JSON(200, recetas)
}

func (handler *RecetaHandler) ObtenerRecetaPorID(c *gin.Context) {
	userInfo := c.Request.Header.Get("Authorization")
	id := c.Param("id")
	receta := handler.RecetaService.ObtenerRecetaPorID(&id, &userInfo)
	c.JSON(200, receta)
}

func (handler *RecetaHandler) CrearReceta(c *gin.Context) {
	//userInfo := c.Request.Header.Get("Authorization")
	var recetaDto Dto.RecetaDto
	err := c.BindJSON(&recetaDto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.RecetaService.CrearReceta(&recetaDto)
	c.JSON(200, resultado)
}

func (handler *RecetaHandler) ActualizarReceta(c *gin.Context) {
	//userInfo := c.Request.Header.Get("Authorization")
	var recetaDto Dto.RecetaDto
	err := c.BindJSON(&recetaDto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.RecetaService.ActualizarReceta(&recetaDto)
	c.JSON(200, resultado)
}

func (handler *RecetaHandler) EliminarReceta(c *gin.Context) {
	userInfo := c.Request.Header.Get("Authorization")
	id := c.Param("id")
	resultado := handler.RecetaService.EliminarReceta(&id, &userInfo)
	c.JSON(200, resultado)
}
