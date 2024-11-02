package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"supercook/Dto"
	"supercook/Services"
	"supercook/Utils"
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
	userInfo := Utils.GetUserInfoFromContext(c)

	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	momento := c.Query("momento")
	nombre := c.Query("nombre")
	tipoAlimento := c.Query("tipoAlimento")

	filtro := [3]string{momento, nombre, tipoAlimento}
	var recetas = handler.RecetaService.ObtenerRecetas(&filtro, &userInfo.Codigo)
	c.JSON(http.StatusOK, recetas)
}

func (handler *RecetaHandler) ObtenerRecetaPorID(c *gin.Context) {
	userInfo := c.Request.Header.Get("Authorization")
	id := c.Param("id")
	receta := handler.RecetaService.ObtenerRecetaPorID(&id, &userInfo)
	c.JSON(http.StatusOK, receta)
}

func (handler *RecetaHandler) CrearReceta(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)

	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	var recetaDto Dto.RecetaDto

	if err := c.BindJSON(&recetaDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recetaDto.IDUsuario = userInfo.Codigo
	resultado := handler.RecetaService.CrearReceta(&recetaDto)
	if !resultado.BoolResultado {
		c.JSON(http.StatusBadRequest, resultado)
		return
	}
	c.JSON(http.StatusOK, resultado)
}

func (handler *RecetaHandler) ActualizarReceta(c *gin.Context) {
	//userInfo := c.Request.Header.Get("Authorization")
	var recetaDto Dto.RecetaDto
	err := c.BindJSON(&recetaDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resultado := handler.RecetaService.ActualizarReceta(&recetaDto)
	c.JSON(http.StatusOK, resultado)
}

func (handler *RecetaHandler) EliminarReceta(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	id := c.Param("id")
	resultado := handler.RecetaService.EliminarReceta(&id, &userInfo.Codigo)
	if !resultado.BoolResultado {
		c.JSON(http.StatusBadRequest, resultado)
		return
	}
	c.JSON(http.StatusOK, resultado)
}
