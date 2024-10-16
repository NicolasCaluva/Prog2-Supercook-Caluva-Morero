package Handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"supercook/Dto"
	"supercook/Services"
	"supercook/Utils"
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
	log.Print("ObtenerAlimentos")
	userInfo := Utils.GetUserInfoFromContext(c)
	filtro := [3]string{
		c.Query("momentoDelDia"),
		c.Query("tipoAlimento"),
		c.Query("nombre"),
	}

	log.Println("ACA ESTA EL CODIGO", userInfo.Codigo)

	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	userId, err := strconv.Atoi(userInfo.Codigo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("userId", userId)

	var alimentos = handler.AlimentoService.ObtenerAlimentos(&filtro, &userId)
	c.JSON(http.StatusOK, alimentos)
}

func (handler *AlimentoHandler) ObtenerAlimentoPorID(c *gin.Context) {
	id := c.Param("id")
	idUsuario, err := strconv.Atoi(c.Query("idUsuario"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	alimento := handler.AlimentoService.ObtenerAlimentoPorID(&id, &idUsuario)
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
	resultado := handler.AlimentoService.ActualizarAlimento(&id, &alimentoDto)
	c.JSON(http.StatusOK, resultado)
}

func (handler *AlimentoHandler) EliminarAlimento(c *gin.Context) {
	id := c.Param("id")
	resultado := handler.AlimentoService.EliminarAlimento(&id)
	c.JSON(http.StatusOK, resultado)
}
