package Models

type TipoAlimento string
type Momento string

const (
	Verdura TipoAlimento = "verdura"
	Fruta   TipoAlimento = "fruta"
	Lacteo  TipoAlimento = "lacteo"
	Carne   TipoAlimento = "carne"
)
const (
	Desayuno Momento = "desayuno"
	Almuerzo Momento = "almuerzo"
	Merienda Momento = "merienda"
	Cena     Momento = "cena"
)
