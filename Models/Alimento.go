package Models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Alimento struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty"`
	IDUsuario          int                `bson:"idUsuario"`
	Nombre             string             `bson:"nombre"`
	PrecioUnitario     float64            `bson:"precioUnitario"`
	Stock              int                `bson:"stock"`
	CantMininaStock    int                `bson:"cantMininaStock"`
	TipoAlimento       TipoAlimento       `bson:"tipoAlimento"`
	MomentoDelDia      []Momento          `bson:"momentoDelDia"`
	FechaCreacion      time.Time          `bson:"fechaCreacion"`
	FechaActualizacion time.Time          `bson:"fechaActualizacion"`
}
