package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"supercook/Dto"
	"supercook/Errors"
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
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}

	momento := c.Query("momento")
	nombre := c.Query("nombre")
	tipoAlimento := c.Query("tipoAlimento")

	filtro := [3]string{momento, nombre, tipoAlimento}
	recetas, err := handler.RecetaService.ObtenerRecetas(&filtro, &userInfo.Codigo)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, recetas)
}

func (handler *RecetaHandler) ObtenerRecetaPorID(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}
	id := c.Param("id")
	receta, err := handler.RecetaService.ObtenerRecetaPorID(&id, &userInfo.Codigo)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, receta)
}

func (handler *RecetaHandler) CrearReceta(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)

	if userInfo == nil {
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}

	var recetaDto Dto.RecetaDto

	if err := c.BindJSON(&recetaDto); err != nil {
		c.Error(Errors.ErrorJsonInvalidoReceta)
		return
	}

	recetaDto.IDUsuario = userInfo.Codigo
	resultado := handler.RecetaService.CrearReceta(&recetaDto)
	if resultado != nil {
		c.Error(resultado)
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Receta creada correctamente"})
}

func (handler *RecetaHandler) EliminarReceta(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}
	id := c.Param("id")
	resultado := handler.RecetaService.EliminarReceta(&id, &userInfo.Codigo)
	if resultado != nil {
		c.Error(resultado)
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Receta eliminada correctamente"})
}
