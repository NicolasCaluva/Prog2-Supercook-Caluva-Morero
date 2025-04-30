package Models

type AlimentoReceta struct {
	IDAlimento string `bson:"idAlimento"`
	Cantidad   int    `bson:"cantidad"`
}
