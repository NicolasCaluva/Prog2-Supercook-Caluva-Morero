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
	var filtro Dto.FiltroAlimentoDto
	momentosDelDia := c.Query("momento")
	if momentosDelDia != "" {
		filtro.MomentoDelDiaDto = make([]Dto.Momento, 0, 1)
		switch momentosDelDia {
		case string(Dto.Desayuno), string(Dto.Almuerzo), string(Dto.Merienda), string(Dto.Cena):
			filtro.MomentoDelDiaDto = append(filtro.MomentoDelDiaDto, Dto.Momento(momentosDelDia))
		default:
			log.Printf("Valor de momento no válido en el filtro: %s", Errors.ErrorFiltroMomentoInvalido)
			c.Error(Errors.ErrorFiltroMomentoInvalido)
			return
		}
	}
	tipoAlimento := c.Query("tipoAlimento")
	if tipoAlimento != "" {
		switch tipoAlimento {
		case string(Dto.Verdura), string(Dto.Fruta), string(Dto.Lacteo), string(Dto.Carne):
			filtro.TipoAlimentoDto = Dto.TipoAlimento(tipoAlimento)
		default:
			log.Printf("Valor de tipo de alimento no válido en el filtro: %s", Errors.ErrorFiltroAlimentoTipoAlimentoMalIngresado)
			c.Error(Errors.ErrorFiltroAlimentoTipoAlimentoMalIngresado)
			return
		}
	}
	filtro.Nombre = c.Query("nombre")
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
