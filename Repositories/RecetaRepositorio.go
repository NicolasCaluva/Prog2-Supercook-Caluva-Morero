package Repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"supercook/Errors"
	"supercook/Models"
	"supercook/Utils"
)

type RecetaRepositorioInterface interface {
	CrearReceta(receta *Models.Receta) (*mongo.InsertOneResult, *Errors.ErrorCodigo)
	ObtenerRecetas(filtro *[3]string, idUsuario *string) ([]Models.Receta, *Errors.ErrorCodigo)
	ObtenerRecetaPorID(idReceta *string, idUsuario *string) (Models.Receta, *Errors.ErrorCodigo)
	ActualizarReceta(receta *Models.Receta) (*mongo.UpdateResult, *Errors.ErrorCodigo)
	EliminarReceta(idReceta *string, idUsuario *string) (*mongo.DeleteResult, *Errors.ErrorCodigo)
}

type RecetaRepositorio struct {
	db DB
}

func NuevoRecetaRepositorio(db DB) *RecetaRepositorio {
	return &RecetaRepositorio{
		db: db,
	}
}

func (recetaRepositorio *RecetaRepositorio) CrearReceta(receta *Models.Receta) (*mongo.InsertOneResult, *Errors.ErrorCodigo) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	resultado, err := coleccion.InsertOne(context.Background(), receta)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	return resultado, nil
}

func (recetaRepositorio *RecetaRepositorio) ObtenerRecetas(filtro *[3]string, idUsuario *string) ([]Models.Receta, *Errors.ErrorCodigo) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")

	filtroBson := bson.M{}

	filtroBson["idUsuario"] = *idUsuario
	if filtro[0] != "" {
		filtroBson["momento"] = bson.M{"$regex": filtro[0], "$options": "i"}
	}
	if filtro[1] != "" {
		filtroBson["nombre"] = bson.M{"$regex": filtro[1], "$options": "i"}
	}
	if filtro[2] != "" {
		filtroBson["alimentos.tipoAlimento"] = bson.M{"$regex": filtro[2], "$options": "i"}
	}

	cursor, err := coleccion.Find(context.TODO(), filtroBson)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	defer cursor.Close(context.Background())

	var recetas []Models.Receta
	for cursor.Next(context.Background()) {
		var receta Models.Receta
		err := cursor.Decode(&receta)
		if err != nil {
			log.Printf("Error: %v\n", Errors.ErrorDecodificarAlimento)
		}
		recetas = append(recetas, receta)
	}
	if len(recetas) == 0 {
		return nil, Errors.ErrorListaVaciaDeRecetas
	}
	return recetas, nil
}

func (recetaRepositorio *RecetaRepositorio) ObtenerRecetaPorID(idReceta *string, idUsuario *string) (Models.Receta, *Errors.ErrorCodigo) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	IdObjeto := Utils.GetObjectIDFromStringID(*idReceta)
	filtro := bson.M{"_id": IdObjeto}
	filtro["idUsuario"] = *idUsuario
	var receta Models.Receta
	err := coleccion.FindOne(context.TODO(), filtro).Decode(&receta)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Error: %v\n", Errors.ErrorRecetaNoEncontrada)
			return Models.Receta{}, Errors.ErrorRecetaNoEncontrada
		}
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return Models.Receta{}, Errors.ErrorConectarBD
	}

	return receta, nil
}

func (recetaRepositorio *RecetaRepositorio) ActualizarReceta(receta *Models.Receta) (*mongo.UpdateResult, *Errors.ErrorCodigo) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	filtro := bson.M{"_id": receta.ID}
	filtro["idUsuario"] = receta.IDUsuario
	entidad := bson.M{
		"$set": bson.M{
			"nombre":             receta.Nombre,
			"alimentos":          receta.Alimentos,
			"momento":            receta.Momento,
			"fechaActualizacion": receta.FechaActualizacion,
		},
	}
	resultado, err := coleccion.UpdateOne(context.TODO(), filtro, entidad)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	if resultado.MatchedCount == 0 {
		log.Printf("Error: %v\n", Errors.ErrorAlimentoNoEncontradoActualizar)
		return nil, Errors.ErrorRecetaNoEncontradoActualizar

	}
	return resultado, nil
}

func (recetaRepositorio *RecetaRepositorio) EliminarReceta(idReceta *string, idUsuario *string) (*mongo.DeleteResult, *Errors.ErrorCodigo) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	IdObjeto := Utils.GetObjectIDFromStringID(*idReceta)
	filtro := bson.M{"_id": IdObjeto}
	filtro["idUsuario"] = *idUsuario
	resultado, err := coleccion.DeleteOne(context.TODO(), filtro)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	if resultado.DeletedCount == 0 {
		log.Printf("Error: %v\n", Errors.ErrorRecetaNoEncontradoEliminar)
		return nil, Errors.ErrorRecetaNoEncontradoEliminar
	}
	return resultado, nil
}
