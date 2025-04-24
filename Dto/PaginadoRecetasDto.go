package Dto

type PaginadoRecetasDto struct {
	NroPagina      int
	PaginasTotales int
	RecetaDto      []RecetaDto
}
