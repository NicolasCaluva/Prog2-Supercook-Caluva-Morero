package Services

import (
	"supercook/Dto"
	"supercook/Models"
	"supercook/Repositories"
)

type CompraInterfaz interface {
	AgregarCompra(compra *Dto.CompraDto) *Dto.Resultado
	ObtenerListaAlimentosStockMenorStockMinimo(idUsuario *string) []*Dto.AlimentoDto
}

type CompraService struct {
	CompraRepositorio Repositories.CompraRepositorioInterfaz
	AlimentoService   AlimentoInterface
}

func NuevoCompraService(compraRepositorio Repositories.CompraRepositorioInterfaz) *CompraService {
	return &CompraService{
		CompraRepositorio: compraRepositorio,
	}
}
func (service *CompraService) ObtenerListaAlimentosStockMenorStockMinimo(idUsuario *string) []*Dto.AlimentoDto {
	alimentos := service.AlimentoService.ObtenerAlimentosConMenosStockQueCantidadMinima(idUsuario)
	var alimentoDTO []*Dto.AlimentoDto
	for _, alimento := range alimentos {
		alimentoDTO = append(alimentoDTO, alimento)
	}
	return alimentoDTO
}
func (service *CompraService) AgregarCompra(compra *Dto.CompraDto) *Dto.Resultado {
	resultado := Dto.Resultado{}
	resultado.ListaMensaje[0] = compra.ValidarListaAlimentos()
	var compraModel *Models.Compra
	if len(resultado.ListaMensaje) > 0 {
		resultado.BoolResultado = false
		return &resultado
	} else {
		primerIteracion := true
		for _, alimento := range compra.Alimentos {
			alimentoRecibido := service.AlimentoService.ObtenerAlimentoPorID(&alimento.IDAlimento, &compra.IDUsuario)
			if alimentoRecibido != nil {
				alimentoRecibido.Stock = alimentoRecibido.Stock + alimento.CantComprada
				resultadoAlimentos := service.AlimentoService.ActualizarAlimento(alimentoRecibido)
				resultado.ListaMensaje = append(resultado.ListaMensaje, resultadoAlimentos.ListaMensaje...)
				if resultadoAlimentos.BoolResultado {
					if primerIteracion {
						primerIteracion = false
						compraModel = convertirCompra(compra)
					}
					compraModel.Alimentos = append(compraModel.Alimentos, Models.ElementoComprado{
						IDAlimento:   alimento.IDAlimento,
						CantComprada: alimento.CantComprada,
					})
				}
			}
		}
	}
	resultadoCompra, err := service.CompraRepositorio.AgregarCompra(compraModel)
	if err != nil {
		resultado.BoolResultado = false
		resultado.ListaMensaje = append(resultado.ListaMensaje, "Error al agregar compra.")
	} else {
		resultado.ListaMensaje = append(resultado.ListaMensaje, "ERR X MONGO"+resultadoCompra.InsertedID.(string))
	}

	resultado.BoolResultado = true
	return &resultado
}

// haceme una funcion para cambiar de dto a model
func convertirCompra(compra *Dto.CompraDto) *Models.Compra {
	return &Models.Compra{
		IDCompra:      compra.IDCompra,
		IDUsuario:     compra.IDUsuario,
		Alimentos:     convertirElementoComprado(compra.Alimentos),
		FechaCreacion: compra.FechaCreacion,
	}
}
func convertirElementoComprado(alimentos []Dto.ElementoCompradoDto) []Models.ElementoComprado {
	var alimentosModel []Models.ElementoComprado
	for _, alimento := range alimentos {
		alimentosModel = append(alimentosModel, Models.ElementoComprado{
			IDAlimento:   alimento.IDAlimento,
			CantComprada: alimento.CantComprada,
		})
	}
	return alimentosModel
}
