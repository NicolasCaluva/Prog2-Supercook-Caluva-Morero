package Dto

type PaginadoAlimentoDto struct {
	NroPagina      int
	PaginasTotales int
	AlimentosDto   []AlimentoDto
}
