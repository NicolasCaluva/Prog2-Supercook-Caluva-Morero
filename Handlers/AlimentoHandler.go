package Handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
	"supercook/Dto"
	"supercook/Errors"
	"supercook/Services"
	"supercook/Utils"
)

type AlimentoHandler struct {
	AlimentoService Services.AlimentoInterface
}

func NuevoAlimentoHandler(alimentoService Services.AlimentoInterface) *AlimentoHandler {
	return &AlimentoHandler{
		AlimentoService: alimentoService,
	}
}
func (handler *AlimentoHandler) ObtenerAlimentos(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	var alimento Dto.AlimentoDto
	var filtro Dto.FiltroAlimentoDto
	momentosDelDia := strings.Split(c.Query("momentoDelDia"), ",")
	errores := alimento.ValidarFiltroMomentoDelDia(&momentosDelDia)
	if errores != nil {
		log.Printf("Valor de momento no válido en el filtro: %s", Errors.ErrorFiltroMomentoInvalido)
		c.Error(errores)
	}
	tipoAlimento := c.Query("tipoAlimento")
	errores = alimento.ValidarFiltroTipoAlimento(&tipoAlimento)
	if errores != nil {
		log.Printf("Valor de tipo de alimento no válido en el filtro: %s", Errors.ErrorFiltroAlimentoTipoAlimentoMalIngresado)
		c.Error(Errors.ErrorFiltroAlimentoTipoAlimentoMalIngresado)
	}
	filtro.Nombre = c.Query("nombre")
	filtro.StockMenorCantidadMinima = c.Query("StockMenorCantidadMinima") == "true"
	strNroPagina := c.Query("page")
	errores = alimento.ValidarFiltroNroPagina(&strNroPagina)
	filtro.NroPagina, _ = strconv.Atoi(strNroPagina)
	log.Printf("Iniciando consulta de alimentos con filtros: momentosDelDia=%s, tipoAlimento=%s, nombre=%s, StockMenorCantidadMinima=%s, usuario=%s, nroPagina=%d", momentosDelDia, tipoAlimento, filtro.Nombre, filtro.StockMenorCantidadMinima, userInfo, filtro.NroPagina)
	alimentos, err := handler.AlimentoService.ObtenerAlimentos(&filtro, &userInfo.Codigo)
	if err != nil {
		log.Printf("Error: %v\n", err)
		c.Error(err)
		return
	}
	log.Printf("Alimentos obtenidos: %v\n", alimentos)
	c.JSON(http.StatusOK, alimentos)
}

func (handler *AlimentoHandler) ObtenerAlimentoPorID(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	id := c.Param("id")
	log.Printf("Iniciando consulta de alimento con ID: %s, usuario: %s", id, userInfo)
	alimento, error := handler.AlimentoService.ObtenerAlimentoPorID(&id, &userInfo.Codigo)
	if error != nil {
		log.Printf("Error: %v\n", error)
		c.Error(error)
		return
	}
	log.Printf("Alimento obtenido: %v\n", alimento)
	c.JSON(http.StatusOK, alimento)
}

func (handler *AlimentoHandler) CrearAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	var alimentoDto Dto.AlimentoDto

	if err := c.BindJSON(&alimentoDto); err != nil {
		c.Error(Errors.ErrorJsonInvalidoAlimento)
		return
	}
	alimentoDto.IDUsuario = userInfo.Codigo
	log.Printf("iniciando creación de alimento: %s, usuario: %s ", alimentoDto, userInfo)
	error := handler.AlimentoService.CrearAlimento(&alimentoDto)
	if error != nil {
		log.Printf("Error: %v\n", error)
		c.Error(error)
		return
	}
	log.Printf("Alimento creado con éxito: %v\n", alimentoDto)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Alimento creado con éxito"})
}

func (handler *AlimentoHandler) ActualizarAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	var alimentoDto Dto.AlimentoDto
	c.BindJSON(&alimentoDto)
	alimentoDto.IDUsuario = userInfo.Codigo
	log.Printf("AlimentoDto: %v", alimentoDto)
	log.Printf("Iniciando actualización de alimento: %s, usuario: %s", alimentoDto, userInfo)
	error := handler.AlimentoService.ActualizarAlimento(&alimentoDto)
	if error != nil {
		log.Printf("Error: %v\n", error)
		c.Error(error)
		return
	}
	log.Printf("Alimento actualizado con éxito: %v\n", alimentoDto)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Alimento actualizado con éxito"})
}

func (handler *AlimentoHandler) EliminarAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	id := c.Param("id")
	log.Printf("Iniciando eliminación de alimento con ID: %s, usuario: %s", id, userInfo)
	error := handler.AlimentoService.EliminarAlimento(&id, &userInfo.Codigo)
	if error != nil {
		log.Printf("Error: %v\n", error)
		c.Error(error)
		return
	}
	log.Printf("Alimento eliminado con éxito: %s\n", id)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Alimento eliminado con éxito"})
}
