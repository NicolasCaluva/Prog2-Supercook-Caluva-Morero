package Handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

	if userInfo == nil {
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}

	momentosDelDia := c.QueryArray("momentoDelDia")
	tipoAlimento := c.Query("tipoAlimento")
	nombre := c.Query("nombre")
	StockMenorCantidadMinima := c.Query("StockMenorCantidadMinima")
	momentosDelDiaStr := strings.Join(momentosDelDia, ",")
	filtro := [4]string{momentosDelDiaStr, tipoAlimento, nombre, StockMenorCantidadMinima}

	alimentos, err := handler.AlimentoService.ObtenerAlimentos(&filtro, &userInfo.Codigo)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, alimentos)
}

func (handler *AlimentoHandler) ObtenerAlimentoPorID(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}
	id := c.Param("id")
	alimento, error := handler.AlimentoService.ObtenerAlimentoPorID(&id, &userInfo.Codigo)
	if error != nil {
		c.Error(error)
		return
	}
	c.JSON(http.StatusOK, alimento)
}

func (handler *AlimentoHandler) CrearAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}
	var alimentoDto Dto.AlimentoDto

	if err := c.BindJSON(&alimentoDto); err != nil {
		c.Error(Errors.ErrorJsonInvalidoAlimento)
		return
	}
	alimentoDto.IDUsuario = userInfo.Codigo
	error := handler.AlimentoService.CrearAlimento(&alimentoDto)
	if error != nil {
		c.Error(error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Alimento creado con éxito"})
}

func (handler *AlimentoHandler) ActualizarAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}
	var alimentoDto Dto.AlimentoDto
	c.BindJSON(&alimentoDto)
	alimentoDto.IDUsuario = userInfo.Codigo
	log.Printf("AlimentoDto: %v", alimentoDto)
	error := handler.AlimentoService.ActualizarAlimento(&alimentoDto)
	if error != nil {
		c.Error(error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Alimento actualizado con éxito"})
}

func (handler *AlimentoHandler) EliminarAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.Error(Errors.ErrorUsuarioNoAutenticado)
		return
	}
	id := c.Param("id")
	error := handler.AlimentoService.EliminarAlimento(&id, &userInfo.Codigo)
	if error != nil {
		c.Error(error)
		return
	}
	c.JSON(http.StatusOK, gin.H{"mensaje": "Alimento eliminado con éxito"})
}
