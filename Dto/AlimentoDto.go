package Dto

type AlimentoDto struct {
	IdAlimento      string
	IDUsuario       string
	Nombre          string
	PrecioUnitario  float64
	Stock           int
	CantMinimaStock int
	TipoAlimento    TipoAlimento
	MomentoDelDia   []Momento
}

func (a *AlimentoDto) ValidarAlimentoDto() []string {
	var mensajes []string

	if a.Nombre == "" {
		mensajes = append(mensajes, "El nombre no puede ser vacío ni nulo.")
	}
	if a.PrecioUnitario <= 0 {
		mensajes = append(mensajes, "El precio unitario debe ser mayor que 0.")
	}
	if a.Stock < 0 {
		mensajes = append(mensajes, "El stock no puede ser negativo.")
	}
	if a.CantMinimaStock < 0 {
		mensajes = append(mensajes, "La cantidad mínima de stock no puede ser negativa.")
	}
	if len(a.MomentoDelDia) == 0 {
		mensajes = append(mensajes, "Debe haber al menos un momento del día.")
	}

	return mensajes
}
