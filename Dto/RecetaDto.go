package Dto

type RecetaDto struct {
	ID        string
	IDUsuario string
	Nombre    string
	Alimentos []AlimentoRecetaDto
	Momento   Momento
}

func (receta *RecetaDto) ValidarRecetaDto() string {
	if receta.IDUsuario == "" {
		return "El ID del usuario no puede estar vacío"
	}
	if receta.Nombre == "" {
		return "El nombre de la receta no puede estar vacío"
	}
	if len(receta.Alimentos) == 0 {
		return "Debe haber al menos un alimento en la receta"
	}
	if receta.Momento == "" {
		return "Debe haber un momento del día en la receta"
	}
	return ""
}
