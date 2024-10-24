package Services

import (
	"fmt"
	"log"
	"supercook/Dto"
	"supercook/Models"
	"supercook/Repositories"
	"supercook/Utils"
	"time"
)

type AlimentoInterface interface {
	ObtenerAlimentos(filtro *[3]string, idUsuario *string) []*Dto.AlimentoDto
	ObtenerAlimentoPorID(id *string, idUsuario *string) *Dto.AlimentoDto
	CrearAlimento(alimento *Dto.AlimentoDto) *Dto.Resultado
	ActualizarAlimento(alimento *Dto.AlimentoDto) *Dto.Resultado
	EliminarAlimento(id *string, idUsuario *string) *Dto.Resultado
	ObtenerAlimentosConMenosStockQueCantidadMinima(idUsuario *string) []*Dto.AlimentoDto
}

type AlimentoService struct {
	AlimentoRepositorio Repositories.AlimentoRepositorioInterface
}

func NuevoAlimentoService(alimentoRepositorio Repositories.AlimentoRepositorioInterface) *AlimentoService {
	return &AlimentoService{
		AlimentoRepositorio: alimentoRepositorio,
	}
}

func (service *AlimentoService) ObtenerAlimentos(filtro *[3]string, idUsuario *string) []*Dto.AlimentoDto {
	alimentos, _ := service.AlimentoRepositorio.ObtenerAlimentos(filtro, idUsuario)
	log.Printf("Alimentos pasa por el service: %v", alimentos)
	var alimentosDto []*Dto.AlimentoDto
	for _, alimento := range alimentos {
		alimentosDto = append(alimentosDto, convertirAlimento(alimento))
	}
	return alimentosDto
}

func (service *AlimentoService) ObtenerAlimentoPorID(idAlimento *string, idUsuario *string) *Dto.AlimentoDto {
	alimento, _ := service.AlimentoRepositorio.ObtenerAlimentoPorID(idAlimento, idUsuario)
	return convertirAlimento(alimento)
}

func (service *AlimentoService) CrearAlimento(alimento *Dto.AlimentoDto) *Dto.Resultado {
	resultado := Dto.Resultado{}
	resultado.ListaMensaje = alimento.ValidarAlimentoDto()
	if len(resultado.ListaMensaje) > 0 {
		resultado.BoolResultado = false
		return &resultado
	} else {
		alimentoModel := Models.Alimento{
			Nombre:          alimento.Nombre,
			IDUsuario:       alimento.IDUsuario,
			PrecioUnitario:  alimento.PrecioUnitario,
			Stock:           alimento.Stock,
			CantMininaStock: alimento.CantMinimaStock,
			TipoAlimento:    Models.TipoAlimento(alimento.TipoAlimento),
			MomentoDelDia:   convertirMomentoaModel(alimento.MomentoDelDia),
			FechaCreacion:   time.Now(),
		}
		_, err := service.AlimentoRepositorio.CrearAlimento(&alimentoModel)
		if err != nil {
			resultado.BoolResultado = false
			resultado.ListaMensaje = append(resultado.ListaMensaje, "Error al crear alimento.")
		} else {
			resultado.BoolResultado = true
			resultado.ListaMensaje = append(resultado.ListaMensaje, fmt.Sprintf("Alimento creado con éxito con las siguientes características: Nombre = %s, Precio Unitario = %.2f, Stock = %d, Cantidad Mínima de Stock = %d, Tipo de Alimento = %d, Momento del Día = %v", alimento.Nombre, alimento.PrecioUnitario, alimento.Stock, alimento.CantMinimaStock, alimento.TipoAlimento, alimento.MomentoDelDia))
		}
	}
	return &resultado
}

func (service *AlimentoService) ActualizarAlimento(alimento *Dto.AlimentoDto) *Dto.Resultado {
	resultado := Dto.Resultado{}
	resultado.ListaMensaje = alimento.ValidarAlimentoDto()
	if len(resultado.ListaMensaje) > 0 {
		resultado.BoolResultado = false
		return &resultado
	} else {
		alimentoModel := Models.Alimento{
			ID:                 Utils.GetObjectIDFromStringID(alimento.IdAlimento),
			Nombre:             alimento.Nombre,
			IDUsuario:          alimento.IDUsuario,
			PrecioUnitario:     alimento.PrecioUnitario,
			Stock:              alimento.Stock,
			CantMininaStock:    alimento.CantMinimaStock,
			TipoAlimento:       Models.TipoAlimento(alimento.TipoAlimento),
			MomentoDelDia:      convertirMomentoaModel(alimento.MomentoDelDia),
			FechaActualizacion: time.Now(),
		}
		_, err := service.AlimentoRepositorio.ActualizarAlimento(&alimentoModel)
		if err != nil {
			resultado.BoolResultado = false
			resultado.ListaMensaje = append(resultado.ListaMensaje, "Error al actualizar alimento.")
		} else {
			resultado.BoolResultado = true
			resultado.ListaMensaje = append(resultado.ListaMensaje, fmt.Sprintf("Alimento actualizado con éxito con las siguientes características: Nombre = %s, Precio Unitario = %.2f, Stock = %d, Cantidad Mínima de Stock = %d, Tipo de Alimento = %d, Momento del Día = %v", alimento.Nombre, alimento.PrecioUnitario, alimento.Stock, alimento.CantMinimaStock, alimento.TipoAlimento, alimento.MomentoDelDia))
		}
	}
	return &resultado
}

func (service *AlimentoService) EliminarAlimento(idAlimento *string, idUsuario *string) *Dto.Resultado {
	resultado := Dto.Resultado{}
	//idAlimentoConvertido:= Utils.GetObjectIDFromStringID(*idAlimento)
	_, err := service.AlimentoRepositorio.EliminarAlimento(idAlimento, idUsuario)
	if err != nil {
		resultado.BoolResultado = false
		resultado.ListaMensaje = append(resultado.ListaMensaje, "Error al eliminar alimento.")
	} else {
		resultado.BoolResultado = true
		resultado.ListaMensaje = append(resultado.ListaMensaje, "Alimento eliminado con éxito.")
	}
	return &resultado
}
func (service *AlimentoService) ObtenerAlimentosConMenosStockQueCantidadMinima(idUsuario *string) []*Dto.AlimentoDto {
	alimentos, _ := service.AlimentoRepositorio.ObtenerAlimentosConStockMenorAlMinimo(idUsuario)
	var alimentosDto []*Dto.AlimentoDto
	for _, alimento := range alimentos {
		alimentosDto = append(alimentosDto, convertirAlimento(alimento))
	}
	return alimentosDto
}

func convertirAlimento(alimento Models.Alimento) *Dto.AlimentoDto {
	alimentoDto := Dto.AlimentoDto{
		IdAlimento:      alimento.ID.Hex(),
		Nombre:          alimento.Nombre,
		IDUsuario:       alimento.IDUsuario,
		PrecioUnitario:  alimento.PrecioUnitario,
		Stock:           alimento.Stock,
		CantMinimaStock: alimento.CantMininaStock,
		TipoAlimento:    Dto.TipoAlimento(alimento.TipoAlimento),
		MomentoDelDia:   convertirMomentosADto(alimento.MomentoDelDia),
	}
	return &alimentoDto
}

func convertirMomentosADto(momentoLista []Models.Momento) []Dto.Momento {
	var dtoMomentoLista []Dto.Momento
	for _, momento := range momentoLista {
		dtoMomentoLista = append(dtoMomentoLista, Dto.Momento(momento))
	}
	return dtoMomentoLista
}

func convertirMomentoaModel(momentoLista []Dto.Momento) []Models.Momento {
	var modelMomentoLista []Models.Momento
	for _, momento := range momentoLista {
		modelMomentoLista = append(modelMomentoLista, Models.Momento(momento))
	}
	return modelMomentoLista
}
