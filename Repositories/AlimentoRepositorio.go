package Repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
	"supercook/Errors"
	"supercook/Models"
	"supercook/Utils"
)

type AlimentoRepositorioInterface interface {
	ObtenerAlimentos(filtro *[4]string, idUsuario *string) ([]Models.Alimento, *Errors.ErrorCodigo)
	ObtenerAlimentoPorID(idAlimento *string, idUsuario *string) (Models.Alimento, *Errors.ErrorCodigo)
	CrearAlimento(alimento *Models.Alimento) (*mongo.InsertOneResult, *Errors.ErrorCodigo)
	ActualizarAlimento(alimento *Models.Alimento) (*mongo.UpdateResult, *Errors.ErrorCodigo)
	EliminarAlimento(id *string, idUsuario *string) (*mongo.DeleteResult, *Errors.ErrorCodigo)
}

type AlimentoRepositorio struct {
	db DB
}

func NuevoAlimentoRepositorio(db DB) *AlimentoRepositorio {
	return &AlimentoRepositorio{
		db: db,
	}
}

func (repositorio *AlimentoRepositorio) ObtenerAlimentos(filtro *[4]string, idUsuario *string) ([]Models.Alimento, *Errors.ErrorCodigo) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	filtros := []bson.M{}
	filtros = append(filtros, bson.M{"idUsuario": *idUsuario})
	if filtro[0] != "" {
		momentos := strings.Split(filtro[0], ",")
		filtros = append(filtros, bson.M{"momentoDelDia": bson.M{"$in": momentos}})
	}
	if filtro[1] != "" {
		filtros = append(filtros, bson.M{"tipoAlimento": bson.M{"$regex": filtro[1], "$options": "i"}})
	}
	if filtro[2] != "" {
		filtros = append(filtros, bson.M{"nombre": bson.M{"$regex": filtro[2], "$options": "i"}})
	}
	if filtro[3] == "True" {
		filtros = append(filtros, bson.M{"$expr": bson.M{"$lt": []string{"$stock", "$cantMininaStock"}}})
	}
	var filtroBson bson.M
	if len(filtros) > 0 {
		filtroBson = bson.M{"$and": filtros}
	} else {
		filtroBson = bson.M{}
	}
	cursor, err := coleccion.Find(context.TODO(), filtroBson)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	defer cursor.Close(context.Background())
	var alimentos []Models.Alimento
	for cursor.Next(context.Background()) {
		var alimento Models.Alimento
		err := cursor.Decode(&alimento)
		if err != nil {
			log.Printf("Error: %v\n", Errors.ErrorDecodificarReceta)
			return nil, Errors.ErrorDecodificarReceta
		}
		alimentos = append(alimentos, alimento)
	}
	if len(alimentos) == 0 {
		return nil, Errors.ErrorListaVaciaDeAlimentos
	}
	return alimentos, nil
}

func (repositorio *AlimentoRepositorio) ObtenerAlimentoPorID(idAlimento *string, idUsuario *string) (Models.Alimento, *Errors.ErrorCodigo) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	IdObjeto := Utils.GetObjectIDFromStringID(*idAlimento)
	filtro := bson.M{"_id": IdObjeto}

	var alimento Models.Alimento
	err := coleccion.FindOne(context.TODO(), filtro).Decode(&alimento)
	if alimento.IDUsuario != *idUsuario {
		return Models.Alimento{}, Errors.ErrorUsuarioNoAutenticado
	}
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("Error: %v\n", Errors.ErrorAlimentoNoEncontrado)
			return Models.Alimento{}, Errors.ErrorAlimentoNoEncontrado
		}
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return Models.Alimento{}, Errors.ErrorConectarBD
	}

	return alimento, nil
}

func (repositorio *AlimentoRepositorio) CrearAlimento(alimento *Models.Alimento) (*mongo.InsertOneResult, *Errors.ErrorCodigo) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	resultado, err := coleccion.InsertOne(context.TODO(), alimento)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	log.Printf("Alimento creado con éxito: %v\n", resultado.InsertedID)
	return resultado, nil
}

func (repositorio *AlimentoRepositorio) ActualizarAlimento(alimento *Models.Alimento) (*mongo.UpdateResult, *Errors.ErrorCodigo) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	filtro := bson.M{"_id": alimento.ID}
	filtro["idUsuario"] = alimento.IDUsuario
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
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	if resultado.MatchedCount == 0 {
		log.Printf("Error: %v\n", Errors.ErrorAlimentoNoEncontradoActualizar)
		return nil, Errors.ErrorAlimentoNoEncontradoActualizar
	}
	return resultado, nil
}

func (repositorio *AlimentoRepositorio) EliminarAlimento(idAlimento *string, idUsuario *string) (*mongo.DeleteResult, *Errors.ErrorCodigo) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	IdObjeto := Utils.GetObjectIDFromStringID(*idAlimento)
	filtro := bson.M{"_id": IdObjeto}
	filtro["idUsuario"] = *idUsuario
	resultado, err := coleccion.DeleteOne(context.TODO(), filtro)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD
	}
	if resultado.DeletedCount == 0 {
		log.Printf("Error: %v\n", Errors.ErrorAlimentoNoEncontradoEliminar)
		return nil, Errors.ErrorAlimentoNoEncontradoEliminar
	}
	return resultado, nil
}
