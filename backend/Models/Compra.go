package Models

import "time"

type Compra struct {
	IDCompra      string             `bson:"_id,omitempty"`
	IDUsuario     string             `bson:"idUsuario"`
	Alimentos     []ElementoComprado `bson:"alimentos"`
	FechaCreacion time.Time          `bson:"fecha"`
	MontoTotal    float64            `bson:"montoTotal"`
}
