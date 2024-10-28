package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Receta struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	IDUsuario          string             `bson:"id_usuario"`
	Nombre             string             `bson:"nombre"`
	Alimentos          []AlimentoReceta   `bson:"alimentos"`
	Momento            Momento            `bson:"momento"`
	FechaCreacion      time.Time          `bson:"fecha_creacion"`
	FechaActualizacion time.Time          `bson:"fecha_actualizacion"`
}
