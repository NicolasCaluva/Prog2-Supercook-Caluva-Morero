package main

import (
	"github.com/gin-gonic/gin"
	"supercook/Handlers"
	"supercook/Middlewares"
	"supercook/Repositories"
	"supercook/Services"
)

var (
	alimentoHandler *Handlers.AlimentoHandler
	router          *gin.Engine
)

func main() {
	router = gin.Default()
	dependencias()
	rutas()
	router.Run(":8080")
}

func rutas() {
	router.Use(Middlewares.CorsMiddleware())

	router.GET("/alimentos", alimentoHandler.ObtenerAlimentos)
	router.GET("/alimentos/:id", alimentoHandler.ObtenerAlimentoPorID)
	router.POST("/alimentos", alimentoHandler.CrearAlimento)
	router.PUT("/alimentos/:id", alimentoHandler.ActualizarAlimento)
	router.DELETE("/alimentos/:id", alimentoHandler.EliminarAlimento)
}

func dependencias() {
	var database Repositories.DB
	var alimentoRepository Repositories.AlimentoRepositorioInterface
	var alimentoService Services.AlimentoInteface

	database = Repositories.NuevaMongoDB()
	alimentoRepository = Repositories.NuevoAlimentoRepositorio(database)
	alimentoService = Services.NuevoAlimentoService(alimentoRepository)
	alimentoHandler = Handlers.NuevoAlimentoHandler(alimentoService)
}
