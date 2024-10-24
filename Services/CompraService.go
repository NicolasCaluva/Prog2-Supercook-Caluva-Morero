package Services

import (
	"log"
	"supercook/Dto"
	"supercook/Models"
	"supercook/Repositories"
	"time"
)

type CompraInterfaz interface {
	AgregarCompra(compra *Dto.CompraDto) *Dto.Resultado
	ObtenerListaAlimentosStockMenorStockMinimo(idUsuario *string) []*Dto.AlimentoDto
}

type CompraService struct {
	CompraRepositorio Repositories.CompraRepositorioInterfaz
	AlimentoService   AlimentoInterface
}

func NuevoCompraService(compraRepositorio Repositories.CompraRepositorioInterfaz, alimentoService AlimentoInterface) *CompraService {
	return &CompraService{
		CompraRepositorio: compraRepositorio,
		AlimentoService:   alimentoService,
	}
}
func (service *CompraService) ObtenerListaAlimentosStockMenorStockMinimo(idUsuario *string) []*Dto.AlimentoDto {
	log.Printf("ObtenerListaAlimentosStockMenorStockMinimo")
	var alimentos []*Dto.AlimentoDto
	alimentos = service.AlimentoService.ObtenerAlimentosConMenosStockQueCantidadMinima(idUsuario)
	log.Printf("Alimentos pasa por el service: %v", alimentos)
	var alimentoDTO []*Dto.AlimentoDto
	for _, alimento := range alimentos {
		alimentoDTO = append(alimentoDTO, alimento)
	}
	return alimentoDTO
}
func (service *CompraService) AgregarCompra(compra *Dto.CompraDto) *Dto.Resultado {
	resultado := Dto.Resultado{}
	mensaje := compra.ValidarListaAlimentos()
	resultado.ListaMensaje = append(resultado.ListaMensaje, mensaje)
	var compraModel *Models.Compra
	if resultado.ListaMensaje[0] != "" {
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
					compraModel.MontoTotal = compraModel.MontoTotal + (alimentoRecibido.PrecioUnitario * float64(alimento.CantComprada))
					compraModel.Alimentos = append(compraModel.Alimentos, Models.ElementoComprado{IDAlimento: alimento.IDAlimento, CantComprada: alimento.CantComprada})
					compraModel.FechaCreacion = time.Now()
				}
			}
		}
	}
	_, err := service.CompraRepositorio.AgregarCompra(compraModel)
	if err != nil {
		resultado.BoolResultado = false
		resultado.ListaMensaje = append(resultado.ListaMensaje, "Error al agregar compra.")
	} else {
		resultado.ListaMensaje = append(resultado.ListaMensaje, "Compra agregada correctamente.")
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
