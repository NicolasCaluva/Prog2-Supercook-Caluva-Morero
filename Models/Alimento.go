package Models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Alimento struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	Nombre          string             `bson:"nombre"`
	PrecioUnitario  float64            `bson:"precioUnitario"`
	Stock           int                `bson:"stock"`
	CantMininaStock int                `bson:"cantMininaStock"`
	TipoAlimento    TipoAlimento       `bson:"tipoAlimento"`
	MomentoDelDia   []Momento          `bson:"momentoDelDia"`
}
