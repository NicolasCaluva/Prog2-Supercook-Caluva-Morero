package Dto

import "supercook/Errors"

type AlimentoRecetaDto struct {
	IDAlimento string
	Nombre     string
	Cantidad   int
}

func (a *AlimentoRecetaDto) ValidarAlimentoRecetaDto() *Errors.ErrorCodigo {

	if a.IDAlimento == "" {
		return Errors.ErrorAlimentoRecetaIDAlimentoMalIngresado
	}
	if a.Cantidad <= 0 {
		return Errors.ErrorCantidadMenorACero
	}

	return nil
}
