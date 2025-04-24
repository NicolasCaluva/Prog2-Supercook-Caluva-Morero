package Repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"math"
	"supercook/Dto"
	"supercook/Errors"
	"supercook/Models"
	"supercook/Utils"
)

type RecetaRepositorioInterface interface {
	CrearReceta(receta *Models.Receta) (*mongo.InsertOneResult, *Errors.ErrorCodigo)
	ObtenerRecetas(filtro *Dto.FiltroAlimentoDto, idUsuario *string) ([]Models.Receta, *Errors.ErrorCodigo, *Dto.PaginadoRecetasDto)
	ObtenerRecetaPorID(idReceta *string, idUsuario *string) (Models.Receta, *Errors.ErrorCodigo)
	EliminarReceta(idReceta *string, idUsuario *string) (*mongo.DeleteResult, *Errors.ErrorCodigo)
	VerificarAlimentoExistente(idAlimento string) *Errors.ErrorCodigo
	ContarRecetasPorMomento(idUsuario *string) (map[string]int, *Errors.ErrorCodigo)
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

func (recetaRepositorio *RecetaRepositorio) ObtenerRecetas(filtro *Dto.FiltroAlimentoDto, idUsuario *string) ([]Models.Receta, *Errors.ErrorCodigo, *Dto.PaginadoRecetasDto) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")

	filtros := []bson.M{}
	filtros = append(filtros, bson.M{"idUsuario": *idUsuario})
	if filtro.MomentoDelDiaDto[0] != "" {
		filtros = append(filtros, bson.M{"momento": filtro.MomentoDelDiaDto[0]})
	}
	if filtro.Nombre != "" {
		filtros = append(filtros, bson.M{"nombre": bson.M{"$regex": filtro.Nombre, "$options": "i"}})
	}
	nroRegistrosPorPagina := 10
	var opcionesConsulta options.FindOptions
	if filtro.NroPagina > 0 {
		opcionesConsulta.SetSkip(int64((filtro.NroPagina - 1) * nroRegistrosPorPagina))
		opcionesConsulta.SetLimit(int64(nroRegistrosPorPagina))
	}
	var filtroBson bson.M
	if len(filtros) > 0 {
		filtroBson = bson.M{"$and": filtros}
	} else {
		filtroBson = bson.M{}
	}
	totalRegistros, err := coleccion.CountDocuments(context.TODO(), filtroBson)
	if err != nil {
		log.Printf("Error al contar documentos: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD, nil
	}
	cursor, err := coleccion.Find(context.TODO(), filtroBson)
	if err != nil {
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return nil, Errors.ErrorConectarBD, nil
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
		return nil, Errors.ErrorListaVaciaDeRecetas, nil
	}
	var nroPaginaDto Dto.PaginadoRecetasDto
	if filtro.NroPagina > 0 {
		totalPaginas := int(math.Ceil(float64(totalRegistros) / float64(nroRegistrosPorPagina)))
		nroPaginaDto.PaginasTotales = totalPaginas
		nroPaginaDto.NroPagina = filtro.NroPagina
	} else {
		nroPaginaDto.PaginasTotales = 0
		nroPaginaDto.NroPagina = 0
	}
	return recetas, nil, &nroPaginaDto
}

func (recetaRepositorio *RecetaRepositorio) ObtenerRecetaPorID(idReceta *string, idUsuario *string) (Models.Receta, *Errors.ErrorCodigo) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	IdObjeto := Utils.GetObjectIDFromStringID(*idReceta)
	filtro := bson.M{"_id": IdObjeto, "idUsuario": *idUsuario}
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
func (recetaRepositorio *RecetaRepositorio) VerificarAlimentoExistente(idAlimento string) *Errors.ErrorCodigo {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	filtro := bson.M{"alimentos.idAlimento": idAlimento}
	err := coleccion.FindOne(context.TODO(), filtro).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		log.Printf("Error: %v\n", Errors.ErrorConectarBD)
		return Errors.ErrorConectarBD
	}
	return Errors.ErrorNoSePuedeEliminarAlimentoPerteneceaReceta
}
func (recetaRepositorio *RecetaRepositorio) ContarRecetasPorMomento(idUsuario *string) (map[string]int, *Errors.ErrorCodigo) {
	coleccion := recetaRepositorio.db.ObtenerCliente().Database("mongodb-SuperCook").Collection("receta")
	consulta := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{{Key: "idUsuario", Value: *idUsuario}}}},
		{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$momento"},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}}},
		{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 0},
			{Key: "momento", Value: "$_id"},
			{Key: "count", Value: 1},
		}}},
	}
	cursor, err := coleccion.Aggregate(context.Background(), consulta)
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v\n", err)
		return nil, Errors.ErrorConectarBD
	}
	defer cursor.Close(context.Background())
	resultado := make(map[string]int)
	for cursor.Next(context.Background()) {
		var cantRecetasPorMomento struct {
			Momento  string `bson:"momento"`
			Contador int    `bson:"count"`
		}
		err := cursor.Decode(&cantRecetasPorMomento)
		if err != nil {
			log.Printf("Error al decodificar el documento: %v\n", err)
			return nil, Errors.ErrorDecodificarAlimento
		}
		resultado[cantRecetasPorMomento.Momento] = cantRecetasPorMomento.Contador
	}
	if err := cursor.Err(); err != nil {
		log.Printf("Error en el cursor: %v\n", err)
		return nil, Errors.ErrorConectarBD
	}
	return resultado, nil
}
