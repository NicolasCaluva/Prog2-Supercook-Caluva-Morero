package Services

import (
	"log"
	"supercook/Dto"
	"supercook/Errors"
	"supercook/Models"
	"supercook/Repositories"
	"time"
)

type CompraInterfaz interface {
	AgregarCompra(compra *Dto.CompraDto) *Errors.ErrorCodigo
	SumarMontoTotalDeComprasEntreDosFechas(fechaInicio, fechaFin string, idUsuario string) (map[string]float64, *Errors.ErrorCodigo)
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
func (service *CompraService) SumarMontoTotalDeComprasEntreDosFechas(fechaInicio, fechaFin string, idUsuario string) (map[string]float64, *Errors.ErrorCodigo) {
	fechaInicial := time.Time{}
	fechaFinal := time.Time{}
	if fechaInicio != "" && fechaFin != "" {
		fechaInicial1, err1 := time.Parse("2006-01-02", fechaInicio)
		if err1 != nil {
			return nil, Errors.ErrorFechasInvalidas
		}
		fechaFinal1, err2 := time.Parse("2006-01-02", fechaFin)
		if err2 != nil {
			return nil, Errors.ErrorFechasInvalidas
		}
		if fechaFinal.Before(fechaInicial) {
			return nil, Errors.ErrorFechasInvalidas
		}
		fechaInicial = fechaInicial1
		fechaFinal = fechaFinal1
	}
	montoTotal, err := service.CompraRepositorio.SumarMontoTotalDeComprasEntreDosFechasDividoPorMes(fechaInicial, fechaFinal, idUsuario)
	log.Printf("Monto total de compras entre %s y %s: %v\n", fechaInicial, fechaFinal, montoTotal)
	if err != nil {
		return nil, err
	}
	return montoTotal, nil
}
