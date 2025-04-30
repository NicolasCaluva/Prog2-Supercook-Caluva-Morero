package Repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"supercook/Errors"
	"supercook/Models"
	"time"
)

type CompraRepositorioInterfaz interface {
	AgregarCompra(compra *Models.Compra) (*mongo.InsertOneResult, *Errors.ErrorCodigo)
	SumarMontoTotalDeComprasEntreDosFechasDividoPorMes(fechaInicio, fechaFin time.Time, idUsuario string) (map[string]float64, *Errors.ErrorCodigo)
}
type CompraRepositorio struct {
	db DB
}

func NuevoCompraRepositorio(db DB) *CompraRepositorio {
	return &CompraRepositorio{
		db: db,
	}
}
func (compraRepositorio *CompraRepositorio) AgregarCompra(compra *Models.Compra) (*mongo.InsertOneResult, *Errors.ErrorCodigo) {
	coleccion := compraRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("compra")
	resultado, err := coleccion.InsertOne(context.Background(), compra)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	return resultado, nil
}
func (CompraRepositorio *CompraRepositorio) SumarMontoTotalDeComprasEntreDosFechasDividoPorMes(fechaInicio, fechaFin time.Time, idUsuario string) (map[string]float64, *Errors.ErrorCodigo) {
	coleccion := CompraRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("compra")
	var filtro bson.M
	if fechaInicio.IsZero() || fechaFin.IsZero() {
		filtro = bson.M{"idUsuario": idUsuario}
	} else {
		filtro = bson.M{"idUsuario": idUsuario, "fecha": bson.M{"$gte": fechaInicio, "$lte": fechaFin}}
	}
	cursor, err := coleccion.Find(context.Background(), filtro)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	var compras []Models.Compra
	if err = cursor.All(context.Background(), &compras); err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	montoTotalPorMes := make(map[string]float64)
	for _, compra := range compras {
		mes := compra.FechaCreacion.Month().String()
		montoTotalPorMes[mes] = montoTotalPorMes[mes] + compra.MontoTotal
	}
	return montoTotalPorMes, nil
}
