package Repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"supercook/Models"
	"supercook/Utils"
)

type RecetaRepositorioInterface interface {
	CrearReceta(receta *Models.Receta) (*mongo.InsertOneResult, error)
	ObtenerRecetas(filtro *[3]string, idUsuario *string) ([]Models.Receta, error)
	ObtenerRecetaPorID(idReceta *string, idUsuario *string) (Models.Receta, error)
	ActualizarReceta(receta *Models.Receta) (*mongo.UpdateResult, error)
	EliminarReceta(idReceta *string, idUsuario *string) (*mongo.DeleteResult, error)
}

type RecetaRepositorio struct {
	db DB
}

func NuevoRecetaRepositorio(db DB) *RecetaRepositorio {
	return &RecetaRepositorio{
		db: db,
	}
}

func (recetaRepositorio *RecetaRepositorio) CrearReceta(receta *Models.Receta) (*mongo.InsertOneResult, error) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	resultado, err := coleccion.InsertOne(context.Background(), receta)
	if err != nil {
		return nil, err
	}
	return resultado, nil
}

func (recetaRepositorio *RecetaRepositorio) ObtenerRecetas(filtro *[3]string, idUsuario *string) ([]Models.Receta, error) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")

	filtroBson := bson.M{}

	filtroBson["idUsuario"] = *idUsuario
	if filtro[0] != "" {
		filtroBson["momento"] = bson.M{"$regex": filtro[0], "$options": "i"}
	}
	if filtro[1] != "" {
		filtroBson["nombre"] = bson.M{"$regex": filtro[1], "$options": "i"}
	}

	cursor, err := coleccion.Find(context.TODO(), filtroBson)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var recetas []Models.Receta
	for cursor.Next(context.Background()) {
		var receta Models.Receta
		err := cursor.Decode(&receta)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		recetas = append(recetas, receta)
	}

	return recetas, nil
}

func (recetaRepositorio *RecetaRepositorio) ObtenerRecetaPorID(idReceta *string, idUsuario *string) (Models.Receta, error) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	IdObjeto := Utils.GetObjectIDFromStringID(*idReceta)
	filtro := bson.M{"_id": IdObjeto}
	filtro["idUsuario"] = *idUsuario
	cursor, err := coleccion.Find(context.TODO(), filtro)
	defer cursor.Close(context.Background())
	var receta Models.Receta
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&receta)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
	}
	return receta, err
}

func (recetaRepositorio *RecetaRepositorio) ActualizarReceta(receta *Models.Receta) (*mongo.UpdateResult, error) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	filtro := bson.M{"_id": receta.ID}
	resultado, err := coleccion.ReplaceOne(context.TODO(), filtro, receta)
	return resultado, err
}

func (recetaRepositorio *RecetaRepositorio) EliminarReceta(idReceta *string, idUsuario *string) (*mongo.DeleteResult, error) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	IdObjeto := Utils.GetObjectIDFromStringID(*idReceta)
	filtro := bson.M{"_id": IdObjeto}
	filtro["idUsuario"] = *idUsuario
	resultado, err := coleccion.DeleteOne(context.TODO(), filtro)
	return resultado, err
}
