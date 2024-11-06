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
	momento := c.Query("momento")
	nombre := c.Query("nombre")
	tipoAlimento := c.Query("tipoAlimento")

	filtro := [3]string{momento, nombre, tipoAlimento}
	log.Printf("Iniciando consulta de recetas con filtros: momento=%s, nombre=%s, tipoAlimento=%s, usuario=%s")
	recetas, err := handler.RecetaService.ObtenerRecetas(&filtro, &userInfo.Codigo)
	if err != nil {
		log.Printf("Error: %v\n", err)
		c.Error(err)
		return
	}
	log.Printf("Recetas obtenidas: %v\n", recetas)
	c.JSON(http.StatusOK, recetas)
}

func (handler *RecetaHandler) ObtenerRecetaPorID(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	id := c.Param("id")
	log.Printf("Iniciando consulta de receta con ID: %s, usuario: %s", id, userInfo)
	receta, err := handler.RecetaService.ObtenerRecetaPorID(&id, &userInfo.Codigo)
	if err != nil {
		log.Printf("Error: %v\n", err)
		c.Error(err)
		return
	}
	log.Printf("Receta obtenida: %v\n", receta)
	c.JSON(http.StatusOK, receta)
}

func (handler *RecetaHandler) CrearReceta(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	var recetaDto Dto.RecetaDto
	if err := c.BindJSON(&recetaDto); err != nil {
		c.Error(Errors.ErrorJsonInvalidoReceta)
		return
	}
	recetaDto.IDUsuario = userInfo.Codigo
	log.Printf("Iniciando creación de receta con ID: %s, usuario: %s", recetaDto.IDUsuario, userInfo)
	resultado := handler.RecetaService.CrearReceta(&recetaDto)
	if resultado != nil {
		log.Printf("Error: %v\n", resultado)
		c.Error(resultado)
		return
	}
	log.Printf("Receta creada correctamente\n")
	c.JSON(http.StatusOK, gin.H{"mensaje": "Receta creada correctamente"})
}

func (handler *RecetaHandler) EliminarReceta(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	id := c.Param("id")
	log.Printf("Iniciando eliminación de receta con ID: %s, usuario: %s", id, userInfo)
	resultado := handler.RecetaService.EliminarReceta(&id, &userInfo.Codigo)
	if resultado != nil {
		log.Printf("Error: %v\n", resultado)
		c.Error(resultado)
		return
	}
	log.Printf("Receta eliminada correctamente\n")
	c.JSON(http.StatusOK, gin.H{"mensaje": "Receta eliminada correctamente"})
}
