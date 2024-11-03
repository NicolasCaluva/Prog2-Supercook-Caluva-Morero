package Dto

import "supercook/Errors"

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

func (a *AlimentoDto) ValidarAlimentoDto() error {
	if a.Nombre == "" {
		return Errors.ErrorAlimentoNombreMalIngresado
	}
	if a.PrecioUnitario <= 0 {
		return Errors.ErrorAlimentoPrecioUnitarioMalIngresado
	}
	if a.Stock < 0 {
		return Errors.ErrorAlimentoStockMalIngresado
	}
	if a.CantMinimaStock < 0 {
		return Errors.ErrorAlimentoCantMinimaStockMalIngresado
	}
	if a.TipoAlimento == "" {
		return Errors.ErrorAlimentoTipoAlimentoMalIngresado
	}
	if len(a.MomentoDelDia) == 0 {
		return Errors.ErrorAlimentoMomentoDelDiaMalIngresado
	}
	return nil
}
