package Services

import (
	"supercook/Dto"
	"supercook/Errors"
	"supercook/Models"
	"supercook/Repositories"
	"supercook/Utils"
	"time"
)

type AlimentoInterface interface {
	ObtenerAlimentos(filtro *[4]string, idUsuario *string) ([]*Dto.AlimentoDto, *Errors.ErrorCodigo)
	ObtenerAlimentoPorID(id *string, idUsuario *string) (*Dto.AlimentoDto, *Errors.ErrorCodigo)
	CrearAlimento(alimento *Dto.AlimentoDto) *Errors.ErrorCodigo
	ActualizarAlimento(alimento *Dto.AlimentoDto) *Errors.ErrorCodigo
	EliminarAlimento(id *string, idUsuario *string) *Errors.ErrorCodigo
}

type AlimentoService struct {
	AlimentoRepositorio Repositories.AlimentoRepositorioInterface
	RecetaRepositorio   Repositories.RecetaRepositorioInterface
}

func NuevoAlimentoService(alimentoRepositorio Repositories.AlimentoRepositorioInterface, recetaRepositorio Repositories.RecetaRepositorioInterface) *AlimentoService {
	return &AlimentoService{
		AlimentoRepositorio: alimentoRepositorio,
		RecetaRepositorio:   recetaRepositorio,
	}
}

func (service *AlimentoService) ObtenerAlimentos(filtro *[4]string, idUsuario *string) ([]*Dto.AlimentoDto, *Errors.ErrorCodigo) {
	alimentos, error := service.AlimentoRepositorio.ObtenerAlimentos(filtro, idUsuario)
	if error != nil {
		return nil, error
	}
	var alimentosDto []*Dto.AlimentoDto
	for _, alimento := range alimentos {
		alimentosDto = append(alimentosDto, convertirAlimento(alimento))
	}
	return alimentosDto, nil
}

func (service *AlimentoService) ObtenerAlimentoPorID(idAlimento *string, idUsuario *string) (*Dto.AlimentoDto, *Errors.ErrorCodigo) {
	alimento, err := service.AlimentoRepositorio.ObtenerAlimentoPorID(idAlimento, idUsuario)
	if err != nil {
		return nil, err
	}
	return convertirAlimento(alimento), nil
}

func (service *AlimentoService) CrearAlimento(alimento *Dto.AlimentoDto) *Errors.ErrorCodigo {
	error := alimento.ValidarAlimentoDto()
	if error != nil {
		return error
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
			return err
		}
		return nil
	}
}

func (service *AlimentoService) ActualizarAlimento(alimento *Dto.AlimentoDto) *Errors.ErrorCodigo {
	error := alimento.ValidarAlimentoDto()
	if error != nil {
		return error
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
			return err
		}
		return nil
	}
}

func (service *AlimentoService) EliminarAlimento(idAlimento *string, idUsuario *string) *Errors.ErrorCodigo {
	errorALimento := service.RecetaRepositorio.VerificarAlimentoExistente(*idAlimento)
	if errorALimento != nil {
		return errorALimento
	}
	_, err := service.AlimentoRepositorio.EliminarAlimento(idAlimento, idUsuario)

	if err != nil {
		return err
	}
	return nil
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
