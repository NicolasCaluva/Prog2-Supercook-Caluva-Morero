package Repositories

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"supercook/Models"
	"supercook/Utils"
)

type AlimentoRepositoryInterface interface {
	ObtenerAlimentos() ([]Models.Alimento, error)
}

type AlimentoRepository struct {
	db DB
}

func NuevoAlimentoRepository(db DB) *AlimentoRepository {
	return &AlimentoRepository{
		db: db,
	}
}
func (repository AlimentoRepository) ObtenerAlimentos() ([]Models.Alimento, error) {
	coleccion := repository.db.GetClient().Database("mongodb-SuperCook").Collection("alimento")
	filtro := bson.M{}
	cursor, err := coleccion.Find(context.TODO(), filtro)
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
func (repository AlimentoRepository) ObtenerAlimentoPorID(id string) (Models.Alimento, error) {
	coleccion := repository.db.GetClient().Database("mongodb-SuperCook").Collection("alimento")
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
func (repository AlimentoRepository) CrearAlimento(alimento Models.Alimento) (*mongo.InsertOneResult, error) {
	coleccion := repository.db.GetClient().Database("mongodb-SuperCook").Collection("alimento")
	resultado, err := coleccion.InsertOne(context.TODO(), alimento)
	return resultado, err
}
func (repository AlimentoRepository) ActualizarAlimento(id string, alimento Models.Alimento) (*mongo.UpdateResult, error) {
	coleccion := repository.db.GetClient().Database("mongodb-SuperCook").Collection("alimento")
	filtro := bson.M{"_id": alimento.ID}
	entidad := bson.M{
		"$set": bson.M{
			"nombre":          alimento.Nombre,
			"precioUnitario":  alimento.PrecioUnitario,
			"stock":           alimento.Stock,
			"cantMininaStock": alimento.CantMininaStock,
			"tipoAlimento":    alimento.TipoAlimento,
			"momentoDelDia":   alimento.MomentoDelDia,
		},
	}
	resultado, err := coleccion.UpdateOne(context.TODO(), filtro, entidad)
	return resultado, err
}
func (repository AlimentoRepository) EliminarAlimento(id string) (*mongo.DeleteResult, error) {
	coleccion := repository.db.GetClient().Database("mongodb-SuperCook").Collection("alimento")
	filtro := bson.M{"_id": id}
	resultado, err := coleccion.DeleteOne(context.TODO(), filtro)
	return resultado, err
}
