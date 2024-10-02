package Repositories

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type DB interface {
	Conectar() error
	Desconectar() error
	ObtenerCliente() *mongo.Client
}
