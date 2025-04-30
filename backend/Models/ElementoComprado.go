package Models

type ElementoComprado struct {
	IDAlimento   string `bson:"_id,omitempty"`
	CantComprada int    `bson:"cantComprada"`
}
