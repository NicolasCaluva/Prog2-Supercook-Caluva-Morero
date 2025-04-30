package Dto

import (
	"log"
	"strconv"
	"supercook/Errors"
)

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

func (a *AlimentoDto) ValidarAlimentoDto() *Errors.ErrorCodigo {
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
func (a *AlimentoDto) ValidarFiltroMomentoDelDia(momentosDelDia *[]string) *Errors.ErrorCodigo {
	var filtro FiltroAlimentoDto
	if len(*momentosDelDia) > 0 && (*momentosDelDia)[0] != "" {
		filtro.MomentoDelDiaDto = make([]Momento, 0, len(*momentosDelDia))
		for _, momento := range *momentosDelDia {
			switch momento {
			case string(Desayuno), string(Almuerzo), string(Merienda), string(Cena):
				filtro.MomentoDelDiaDto = append(filtro.MomentoDelDiaDto, Momento(momento))
			default:
				log.Printf("Valor de momento no válido en el filtro: %s", Errors.ErrorFiltroMomentoInvalido)
				return Errors.ErrorFiltroMomentoInvalido
			}
		}
	}
	return nil
}
func (a *AlimentoDto) ValidarFiltroTipoAlimento(tipoAlimento *string) *Errors.ErrorCodigo {
	var filtro FiltroAlimentoDto
	if *tipoAlimento != "" {
		switch *tipoAlimento {
		case string(Verdura), string(Fruta), string(Lacteo), string(Carne):
			filtro.TipoAlimentoDto = TipoAlimento(*tipoAlimento)
		default:
			log.Printf("Valor de tipo de alimento no válido en el filtro: %s", Errors.ErrorFiltroAlimentoTipoAlimentoMalIngresado)
			return Errors.ErrorFiltroAlimentoTipoAlimentoMalIngresado
		}
	}
	return nil
}
func (a *AlimentoDto) ValidarFiltroNroPagina(nroPagina *string) *Errors.ErrorCodigo {
	if *nroPagina != "" {
		nroPaginaInt, err := strconv.Atoi(*nroPagina)
		if err != nil {
			log.Printf("El número de página no puede ser menor a 0: %s", Errors.ErrorFiltroNroPaginaMalIngresado)
			return Errors.ErrorFiltroNroPaginaMalIngresado
		}
		if nroPaginaInt < 0 {
			log.Printf("El número de página no puede ser menor a 0: %s", Errors.ErrorFiltroNroPaginaMalIngresado)
			return Errors.ErrorFiltroNroPaginaMalIngresado
		}
	}
	return nil
}
