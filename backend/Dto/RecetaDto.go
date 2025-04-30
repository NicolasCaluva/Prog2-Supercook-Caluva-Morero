package Dto

import "supercook/Errors"

type RecetaDto struct {
	ID        string
	IDUsuario string
	Nombre    string
	Alimentos []AlimentoRecetaDto
	Momento   Momento
}

func (receta *RecetaDto) ValidarRecetaDto() *Errors.ErrorCodigo {
	if receta.Nombre == "" {
		return Errors.ErrorRecetaNombreMalIngresado
	}
	if len(receta.Alimentos) == 0 {
		return Errors.ErrorRecetaAlimentosMalIngresados
	}
	if receta.Momento == "" {
		return Errors.ErrorRecetaMomentoDelDiaMalIngresado
	}
	return nil
}
