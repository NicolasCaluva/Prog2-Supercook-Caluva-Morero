package Services

import (
	"supercook/Dto"
	"supercook/Errors"
	"supercook/Models"
	"supercook/Repositories"
	"time"
)

type CompraInterfaz interface {
	AgregarCompra(compra *Dto.CompraDto) *Errors.ErrorCodigo
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
func (service *CompraService) AgregarCompra(compra *Dto.CompraDto) *Errors.ErrorCodigo {
	error := compra.ValidarListaAlimentos()
	var compraModel *Models.Compra
	if error != nil {
		return error
	} else {
		primerIteracion := true
		for _, alimento := range compra.Alimentos {
			alimentoRecibido, err := service.AlimentoService.ObtenerAlimentoPorID(&alimento.IDAlimento, &compra.IDUsuario)
			if err == nil {
				alimentoRecibido.Stock = alimentoRecibido.Stock + alimento.CantComprada
				service.AlimentoService.ActualizarAlimento(alimentoRecibido)
				if primerIteracion {
					primerIteracion = false
					compraModel = convertirCompra(compra)
				}
				compraModel.MontoTotal = compraModel.MontoTotal + (alimentoRecibido.PrecioUnitario * float64(alimento.CantComprada))
				compraModel.Alimentos = append(compraModel.Alimentos, Models.ElementoComprado{IDAlimento: alimento.IDAlimento, CantComprada: alimento.CantComprada})
				compraModel.FechaCreacion = time.Now()
			} else {
				return err
			}
		}
	}
	_, err1 := service.CompraRepositorio.AgregarCompra(compraModel)
	if err1 != nil {
		return err1
	}
	return nil
}
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
