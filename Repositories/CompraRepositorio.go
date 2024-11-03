package Repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"supercook/Errors"
	"supercook/Models"
)

type CompraRepositorioInterfaz interface {
	AgregarCompra(compra *Models.Compra) (*mongo.InsertOneResult, error)
}
type CompraRepositorio struct {
	db DB
}

func NuevoCompraRepositorio(db DB) *CompraRepositorio {
	return &CompraRepositorio{
		db: db,
	}
}
func (compraRepositorio *CompraRepositorio) AgregarCompra(compra *Models.Compra) (*mongo.InsertOneResult, error) {
	coleccion := compraRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("compra")
	resultado, err := coleccion.InsertOne(context.Background(), compra)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	return resultado, nil
}
