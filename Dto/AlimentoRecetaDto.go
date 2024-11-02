package Dto

type AlimentoRecetaDto struct {
	IDAlimento string
	Nombre     string
	Cantidad   int
}

func (a *AlimentoRecetaDto) ValidarAlimentoRecetaDto() []string {
	var mensajes []string

	if a.IDAlimento == "" {
		mensajes = append(mensajes, "El ID del alimento no puede ser vacío ni nulo.")
	}
	if a.Cantidad <= 0 {
		mensajes = append(mensajes, "La cantidad debe ser mayor que 0.")
	}

	return mensajes
}
