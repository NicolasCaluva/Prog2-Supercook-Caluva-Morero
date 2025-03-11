package Repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"supercook/Errors"
	"supercook/Models"
)

type CompraRepositorioInterfaz interface {
	AgregarCompra(compra *Models.Compra) (*mongo.InsertOneResult, *Errors.ErrorCodigo)
	SumarMontoTotalDeComprasEntreDosFechas(fechaInicio, fechaFin string) (float64, *Errors.ErrorCodigo)
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
func (CompraRepositorio *CompraRepositorio) SumarMontoTotalDeComprasEntreDosFechas(fechaInicio, fechaFin string) (float64, *Errors.ErrorCodigo) {
	coleccion := CompraRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("compra")
	cursor, err := coleccion.Find(context.Background(), bson.M{"fecha": bson.M{"$gte": fechaInicio, "$lte": fechaFin}})
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return 0, Errors.ErrorConectarBD
	}
	var compras []Models.Compra
	if err = cursor.All(context.Background(), &compras); err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return 0, Errors.ErrorConectarBD
	}
	var montoTotal float64
	for _, compra := range compras {
		montoTotal += compra.MontoTotal
	}
	return montoTotal, nil
}
