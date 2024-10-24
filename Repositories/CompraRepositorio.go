package Repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
		return nil, err
	}
	return resultado, nil
}
