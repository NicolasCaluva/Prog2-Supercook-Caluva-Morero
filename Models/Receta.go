package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Receta struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	IDUsuario          string             `bson:"idUsuario"`
	Nombre             string             `bson:"nombre"`
	Alimentos          []AlimentoReceta   `bson:"alimentos"`
	Momento            Momento            `bson:"momento"`
	FechaCreacion      time.Time          `bson:"fechaCreacion"`
	FechaActualizacion time.Time          `bson:"fechaActualizacion"`
}
