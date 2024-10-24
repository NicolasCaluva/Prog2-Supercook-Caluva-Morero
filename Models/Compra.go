package Models

type Compra struct {
	IDCompra      string             `bson:"_id,omitempty"`
	IDUsuario     string             `bson:"idUsuario"`
	Alimentos     []ElementoComprado `bson:"alimentos"`
	FechaCreacion string             `bson:"fecha"`
}
