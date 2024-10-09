package Repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"supercook/Models"
	"supercook/Utils"
)

type AlimentoRepositorioInterface interface {
	ObtenerAlimentos(filtro [3]string) ([]Models.Alimento, error)
	ObtenerAlimentoPorID(id string) (Models.Alimento, error)
	CrearAlimento(alimento Models.Alimento) (*mongo.InsertOneResult, error)
	ActualizarAlimento(id string, alimento Models.Alimento) (*mongo.UpdateResult, error)
	EliminarAlimento(id string) (*mongo.DeleteResult, error)
}

type AlimentoRepositorio struct {
	db DB
}

func NuevoAlimentoRepositorio(db DB) *AlimentoRepositorio {
	return &AlimentoRepositorio{
		db: db,
	}
}

func (repositorio AlimentoRepositorio) ObtenerAlimentos(filtro [3]string) ([]Models.Alimento, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")

	filtroBson := bson.M{}
	if filtro[0] != "" {
		filtroBson["momentoDelDia"] = bson.M{"$regex": filtro[0], "$options": "i"}
	}
	if filtro[1] != "" {
		filtroBson["tipoAlimento"] = bson.M{"$regex": filtro[1], "$options": "i"}
	}
	if filtro[2] != "" {
		filtroBson["nombre"] = bson.M{"$regex": filtro[2], "$options": "i"}
	}

	cursor, err := coleccion.Find(context.TODO(), filtroBson)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var alimentos []Models.Alimento
	for cursor.Next(context.Background()) {
		var alimento Models.Alimento
		err := cursor.Decode(&alimento)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		alimentos = append(alimentos, alimento)
	}
	return alimentos, err
}

func (repositorio AlimentoRepositorio) ObtenerAlimentoPorID(id string) (Models.Alimento, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	IdObjeto := Utils.GetObjectIDFromStringID(id)
	filtro := bson.M{"_id": IdObjeto}
	cursor, err := coleccion.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
	var alimento Models.Alimento
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&alimento)
		if err != nil {

			fmt.Printf("Error: %v\n", err)
		}
	}
	return alimento, err
}

func (repositorio AlimentoRepositorio) CrearAlimento(alimento Models.Alimento) (*mongo.InsertOneResult, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	resultado, err := coleccion.InsertOne(context.TODO(), alimento)
	return resultado, err
}

func (repositorio AlimentoRepositorio) ActualizarAlimento(id string, alimento Models.Alimento) (*mongo.UpdateResult, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	filtro := bson.M{"_id": alimento.ID}
	entidad := bson.M{
		"$set": bson.M{
			"nombre":             alimento.Nombre,
			"precioUnitario":     alimento.PrecioUnitario,
			"stock":              alimento.Stock,
			"cantMininaStock":    alimento.CantMininaStock,
			"tipoAlimento":       alimento.TipoAlimento,
			"momentoDelDia":      alimento.MomentoDelDia,
			"fechaActualizacion": alimento.FechaActualizacion,
			"fechaCreacion":      alimento.FechaCreacion,
		},
	}
	resultado, err := coleccion.UpdateOne(context.TODO(), filtro, entidad)
	return resultado, err
}

func (repositorio AlimentoRepositorio) EliminarAlimento(id string) (*mongo.DeleteResult, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	filtro := bson.M{"_id": id}
	resultado, err := coleccion.DeleteOne(context.TODO(), filtro)
	return resultado, err
}
