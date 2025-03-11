package main

import (
	"github.com/gin-gonic/gin"
	"supercook/Handlers"
	"supercook/Middlewares"
	"supercook/Repositories"
	"supercook/Services"
	clients "supercook/clientes"
)

var (
	alimentoHandler *Handlers.AlimentoHandler
	compraHandler   *Handlers.CompraHandler
	recetaHandler   *Handlers.RecetaHandler
	router          *gin.Engine
)

func main() {
	router = gin.Default()
	router.Use(Middlewares.CORSMiddleware())
	router.Use(Middlewares.ErrorMiddleware)
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
	group.Use(Middlewares.CORSMiddleware())
	group.GET("/", alimentoHandler.ObtenerAlimentos)
	group.GET("/:id/", alimentoHandler.ObtenerAlimentoPorID)
	group.POST("/", alimentoHandler.CrearAlimento)
	group.PUT("/", alimentoHandler.ActualizarAlimento)
	group.DELETE("/:id/", alimentoHandler.EliminarAlimento)

	groupCompra := router.Group("/compras")
	groupCompra.POST("/", compraHandler.CrearCompra)

	groupReceta := router.Group("/recetas")
	groupReceta.GET("/", recetaHandler.ObtenerRecetas)
	groupReceta.GET("/:id/", recetaHandler.ObtenerRecetaPorID)
	groupReceta.POST("/", recetaHandler.CrearReceta)
	groupReceta.DELETE("/:id/", recetaHandler.EliminarReceta)
	groupReceta.GET("/contarRecetasPorMomento/", recetaHandler.ContarRecetasPorMomento)
	groupReceta.GET("/contarRecetasPorTipoAlimento/", recetaHandler.ContarCantidadDeRecetasPorTipoAlimento)
}

func dependencias() {
	var database Repositories.DB
	var alimentoRepository Repositories.AlimentoRepositorioInterface
	var alimentoService Services.AlimentoInterface
	var compraService Services.CompraInterfaz
	var compraRepositorio Repositories.CompraRepositorioInterfaz
	var recetaRepository Repositories.RecetaRepositorioInterface
	var recetaService Services.RecetaInterface
	database = Repositories.NuevaMongoDB()
	alimentoRepository = Repositories.NuevoAlimentoRepositorio(database)
	recetaRepository = Repositories.NuevoRecetaRepositorio(database)
	compraRepositorio = Repositories.NuevoCompraRepositorio(database)
	alimentoService = Services.NuevoAlimentoService(alimentoRepository, recetaRepository)
	alimentoHandler = Handlers.NuevoAlimentoHandler(alimentoService)
	compraService = Services.NuevoCompraService(compraRepositorio, alimentoService)
	compraHandler = Handlers.NuevoCompraHandler(compraService)
	recetaService = Services.NuevoRecetaService(recetaRepository, alimentoService)
	recetaHandler = Handlers.NuevoRecetaHandler(recetaService)
}
