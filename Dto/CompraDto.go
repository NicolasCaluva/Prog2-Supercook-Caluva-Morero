package Dto

import (
	"supercook/Errors"
	"time"
)

type CompraDto struct {
	IDCompra      string
	IDUsuario     string
	Alimentos     []ElementoCompradoDto
	FechaCreacion time.Time
	MontoTotal    float64
}

func (Compra *CompraDto) ValidarListaAlimentos() error {
	if len(Compra.Alimentos) == 0 {
		return Errors.ErrorListaVaciaDeCompras
	}
	return nil
}
