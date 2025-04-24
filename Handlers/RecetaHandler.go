package Handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	var alimentoDto Dto.AlimentoDto
	var filtro Dto.FiltroAlimentoDto
	var momentoDelDia = make([]string, 1)
	momentoDelDia[0] = c.Query("momento")
	errores := alimentoDto.ValidarFiltroMomentoDelDia(&momentoDelDia)
	if errores != nil {
		log.Printf("Valor de momento no válido en el filtro: %s", Errors.ErrorFiltroMomentoInvalido)
		c.Error(errores)
	}
	filtro.MomentoDelDiaDto = convertirAListaDeMomentos(momentoDelDia)

	tipoAlimento := c.Query("tipoAlimento")
	errores = alimentoDto.ValidarFiltroTipoAlimento(&tipoAlimento)
	if errores != nil {
		log.Printf("Valor de tipo de alimento no válido en el filtro: %s", Errors.ErrorFiltroAlimentoTipoAlimentoMalIngresado)
		c.Error(Errors.ErrorFiltroAlimentoTipoAlimentoMalIngresado)
	}
	filtro.TipoAlimentoDto = Dto.TipoAlimento(tipoAlimento)
	filtro.Nombre = c.Query("nombre")
	nroPaginaStr := c.Query("page")
	erorres := alimentoDto.ValidarFiltroNroPagina(&nroPaginaStr)
	if erorres != nil {
		log.Printf("Valor de número de página no válido en el filtro: %s", Errors.ErrorFiltroNroPaginaMalIngresado)
		c.Error(Errors.ErrorFiltroNroPaginaMalIngresado)
	}
	filtro.NroPagina, _ = strconv.Atoi(nroPaginaStr)
	log.Printf("Iniciando consulta de recetas con filtros: momentosDelDia=%s, tipoAlimento=%s, nombre=%s, usuario=%s", filtro.MomentoDelDiaDto, filtro.TipoAlimentoDto, filtro.Nombre, userInfo)
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
func (handler *RecetaHandler) ContarRecetasPorMomento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	recetas, err := handler.RecetaService.ContarRecetasPorMomento(&userInfo.Codigo)
	if err != nil {
		log.Printf("Error: %v\n", err)
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, recetas)
}
func (handler *RecetaHandler) ContarCantidadDeRecetasPorTipoAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	log.Printf("Iniciando conteo de recetas por tipo de alimento, usuario: %s", userInfo)
	recetas, err := handler.RecetaService.ContarCantidadDeRecetasPorTipoAlimento(&userInfo.Codigo)
	if err != nil {
		log.Printf("Error: %v\n", err)
		c.Error(err)
		return
	}
	log.Printf("Recetas contadas: %v\n", recetas)
	c.JSON(http.StatusOK, recetas)
}

// Mover a utils
func convertirAListaDeMomentos(momentoDelDia []string) []Dto.Momento {
	var momentos []Dto.Momento
	for _, momento := range momentoDelDia {
		momentos = append(momentos, Dto.Momento(momento))
	}
	return momentos
}
