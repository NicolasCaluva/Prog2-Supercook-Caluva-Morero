package Dto

type CompraDto struct {
	IDCompra      string
	IDUsuario     string ``
	Alimentos     []ElementoCompradoDto
	FechaCreacion string
	MontoTotal    float64
}

func (Compra *CompraDto) ValidarListaAlimentos() string {
	var mensajes string
	if len(Compra.Alimentos) == 0 {
		mensajes = "La lista de alimentos no puede estar vacía"
	}
	return mensajes
}
