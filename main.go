package main

import (
	"github.com/gin-gonic/gin"
	"supercook/Handlers"
	"supercook/Middlewares"
	"supercook/Repositories"
	"supercook/Services"
	"supercook/clientes"
)

var (
	alimentoHandler *Handlers.AlimentoHandler
	router          *gin.Engine
)

func main() {
	router = gin.Default()
	router.Use(Middlewares.CorsMiddleware())
	dependencias()
	rutas()

	router.Run(":8080")
}

func rutas() {

	var authClient clients.AuthClientInterface
	authClient = clients.NewAuthClient()
	authMiddleware := Middlewares.NewAuthMiddleware(authClient)
	router.Use(authMiddleware.ValidateToken)

	group := router.Group("/alimentos")
	group.Use(authMiddleware.ValidateToken)
	group.GET("/", alimentoHandler.ObtenerAlimentos)
	group.GET("/:id/", alimentoHandler.ObtenerAlimentoPorID)
	group.POST("/", alimentoHandler.CrearAlimento)
	group.PUT("/:id/", alimentoHandler.ActualizarAlimento)
	group.DELETE("/:id/", alimentoHandler.EliminarAlimento)
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
