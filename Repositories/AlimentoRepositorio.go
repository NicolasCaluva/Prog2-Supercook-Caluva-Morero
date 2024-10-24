package Repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"supercook/Models"
	"supercook/Utils"
)

type AlimentoRepositorioInterface interface {
	ObtenerAlimentos(filtro *[3]string, idUsuario *string) ([]Models.Alimento, error)
	ObtenerAlimentoPorID(idAlimento *string, idUsuario *string) (Models.Alimento, error)
	CrearAlimento(alimento *Models.Alimento) (*mongo.InsertOneResult, error)
	ActualizarAlimento(alimento *Models.Alimento) (*mongo.UpdateResult, error)
	EliminarAlimento(id *string, idUsuario *string) (*mongo.DeleteResult, error)
	ObtenerAlimentosConStockMenorAlMinimo(idUsuario *string) ([]Models.Alimento, error)
}

type AlimentoRepositorio struct {
	db DB
}

func NuevoAlimentoRepositorio(db DB) *AlimentoRepositorio {
	return &AlimentoRepositorio{
		db: db,
	}
}

func (repositorio AlimentoRepositorio) ObtenerAlimentos(filtro *[3]string, idUsuario *string) ([]Models.Alimento, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")

	filtroBson := bson.M{}

	filtroBson["idUsuario"] = *idUsuario
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

func (repositorio AlimentoRepositorio) ObtenerAlimentoPorID(idAlimento *string, idUsuario *string) (Models.Alimento, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	IdObjeto := Utils.GetObjectIDFromStringID(*idAlimento)
	filtro := bson.M{"_id": IdObjeto}
	filtro["idUsuario"] = *idUsuario
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

func (repositorio AlimentoRepositorio) CrearAlimento(alimento *Models.Alimento) (*mongo.InsertOneResult, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	resultado, err := coleccion.InsertOne(context.TODO(), alimento)
	return resultado, err
}

func (repositorio AlimentoRepositorio) ActualizarAlimento(alimento *Models.Alimento) (*mongo.UpdateResult, error) {
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
		log.Printf("Error al actualizar el alimento: %v\n", err)
		return nil, err
	}
	return resultado, err
}

func (repositorio AlimentoRepositorio) EliminarAlimento(idAlimento *string, idUsuario *string) (*mongo.DeleteResult, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	IdObjeto := Utils.GetObjectIDFromStringID(*idAlimento)
	filtro := bson.M{"_id": IdObjeto}
	filtro["idUsuario"] = *idUsuario
	resultado, err := coleccion.DeleteOne(context.TODO(), filtro)
	return resultado, err
}
func (repositorio AlimentoRepositorio) ObtenerAlimentosConStockMenorAlMinimo(idUsuario *string) ([]Models.Alimento, error) {
	coleccion := repositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("alimento")
	filtro := bson.M{"stock": bson.M{"$lt": "cantMinimaStock"}}
	filtro["idUsuario"] = *idUsuario
	alimentosLista, err := coleccion.Find(context.TODO(), filtro)
	defer alimentosLista.Close(context.Background())
	var alimentos []Models.Alimento
	for alimentosLista.Next(context.Background()) {
		var alimento Models.Alimento
		err := alimentosLista.Decode(&alimento)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		alimentos = append(alimentos, alimento)
	}
	return alimentos, err
}
