package Dto

type FiltroAlimentoDto struct {
	MomentoDelDiaDto         []Momento
	TipoAlimentoDto          TipoAlimento
	Nombre                   string
	StockMenorCantidadMinima bool
	NroPagina                int
}
