package Repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Cliente *mongo.Client
}

func NuevaMongoDB() *MongoDB {
	instancia := &MongoDB{}
	instancia.Conectar()

	return instancia
}

func (mongoDB *MongoDB) ObtenerCliente() *mongo.Client {
	return mongoDB.Cliente
}

func (mongoDB *MongoDB) Conectar() error {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	mongoDB.Cliente = client

	return nil
}

func (mongoDB *MongoDB) Desconectar() error {
	return mongoDB.Cliente.Disconnect(context.Background())
}
