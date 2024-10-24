package Handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"supercook/Dto"
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
	log.Print("ObtenerAlimentos")
	userInfo := Utils.GetUserInfoFromContext(c)

	log.Println("ACA ESTA EL CODIGO", userInfo.Codigo)

	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}

	momentoDelDia := c.Query("momentoDelDia")
	tipoAlimento := c.Query("tipoAlimento")
	nombre := c.Query("nombre")

	filtro := [3]string{momentoDelDia, tipoAlimento, nombre}

	var alimentos = handler.AlimentoService.ObtenerAlimentos(&filtro, &userInfo.Codigo)
	c.JSON(http.StatusOK, alimentos)
}

func (handler *AlimentoHandler) ObtenerAlimentoPorID(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	id := c.Param("id")
	alimento := handler.AlimentoService.ObtenerAlimentoPorID(&id, &userInfo.Codigo)
	c.JSON(http.StatusOK, alimento)
}

func (handler *AlimentoHandler) CrearAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	var alimentoDto Dto.AlimentoDto

	if err := c.BindJSON(&alimentoDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	alimentoDto.IDUsuario = userInfo.Codigo
	resultado := handler.AlimentoService.CrearAlimento(&alimentoDto)
	if !resultado.BoolResultado {
		c.JSON(http.StatusBadRequest, resultado)
		return
	}
	c.JSON(http.StatusOK, resultado)
}

func (handler *AlimentoHandler) ActualizarAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	var alimentoDto Dto.AlimentoDto
	c.BindJSON(&alimentoDto)
	alimentoDto.IDUsuario = userInfo.Codigo
	log.Printf("AlimentoDto: %v", alimentoDto)
	resultado := handler.AlimentoService.ActualizarAlimento(&alimentoDto)
	if !resultado.BoolResultado {
		c.JSON(http.StatusBadRequest, resultado)
		return
	}
	c.JSON(http.StatusOK, resultado)
}

func (handler *AlimentoHandler) EliminarAlimento(c *gin.Context) {
	userInfo := Utils.GetUserInfoFromContext(c)
	if userInfo == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario no autenticado"})
		return
	}
	id := c.Param("id")
	resultado := handler.AlimentoService.EliminarAlimento(&id, &userInfo.Codigo)
	if !resultado.BoolResultado {
		c.JSON(http.StatusBadRequest, resultado)
		return
	}
	c.JSON(http.StatusOK, resultado)
}
