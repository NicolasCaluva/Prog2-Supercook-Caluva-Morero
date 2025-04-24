package Services

import (
	"fmt"
	"supercook/Dto"
	"supercook/Errors"
	"supercook/Models"
	"supercook/Repositories"
	"time"
)

type RecetaInterface interface {
	ObtenerRecetas(filtro *Dto.FiltroAlimentoDto, idUsuario *string) (*Dto.PaginadoRecetasDto, *Errors.ErrorCodigo)
	ObtenerRecetaPorID(idReceta *string, idUsuario *string) (*Dto.RecetaDto, *Errors.ErrorCodigo)
	CrearReceta(receta *Dto.RecetaDto) *Errors.ErrorCodigo
	EliminarReceta(idReceta *string, idUsuario *string) *Errors.ErrorCodigo
	ContarRecetasPorMomento(idUsuario *string) (map[string]int, *Errors.ErrorCodigo)
	ContarCantidadDeRecetasPorTipoAlimento(idUsuario *string) (map[string]int, *Errors.ErrorCodigo)
}

type RecetaService struct {
	RecetaRepositorio Repositories.RecetaRepositorioInterface
	AlimentoService   AlimentoInterface
}

func NuevoRecetaService(recetaRepositorio Repositories.RecetaRepositorioInterface, alimentoService AlimentoInterface) *RecetaService {
	return &RecetaService{
		RecetaRepositorio: recetaRepositorio,
		AlimentoService:   alimentoService,
	}
}

func (service *RecetaService) ObtenerRecetas(filtro *Dto.FiltroAlimentoDto, idUsuario *string) (*Dto.PaginadoRecetasDto, *Errors.ErrorCodigo) {
	recetas, err, paginas := service.RecetaRepositorio.ObtenerRecetas(filtro, idUsuario)
	if err != nil {
		return nil, err
	}
	var recetasDto []*Dto.RecetaDto
	var correcto = ""
	for _, receta := range recetas {
		if filtro.TipoAlimentoDto != "" {
			correcto = "pendiente"
		}
		recetasDto = append(recetasDto, convertirReceta(receta))
		for i, alimento := range receta.Alimentos {
			alimento, err := service.AlimentoService.ObtenerAlimentoPorID(&alimento.IDAlimento, idUsuario)
			if err != nil {
				return nil, err
			}
			recetasDto[len(recetasDto)-1].Alimentos[i].Nombre = alimento.Nombre
			if fmt.Sprintf("%v", alimento.TipoAlimento) == string(filtro.TipoAlimentoDto) {
				correcto = "hecho"
			}
		}
		if correcto == "pendiente" {
			recetasDto = recetasDto[:len(recetasDto)-1]
		}
	}
	if len(recetasDto) == 0 {
		return nil, Errors.ErrorListaVaciaDeRecetas

	}
	paginas.RecetaDto = convertirSliceAPunterosRecetas(recetasDto)
	return paginas, nil
}

func (service *RecetaService) ObtenerRecetaPorID(idReceta *string, idUsuario *string) (*Dto.RecetaDto, *Errors.ErrorCodigo) {
	receta, err := service.RecetaRepositorio.ObtenerRecetaPorID(idReceta, idUsuario)
	if err != nil {
		return nil, err
	}
	return convertirReceta(receta), nil
}

func (service *RecetaService) CrearReceta(receta *Dto.RecetaDto) *Errors.ErrorCodigo {
	resultadoValidacion := receta.ValidarRecetaDto()
	var alimentos []Dto.AlimentoDto
	if resultadoValidacion != nil {
		return resultadoValidacion
	} else {
		var listaAlimentosReceta []Models.AlimentoReceta
		for _, alimentoReceta := range receta.Alimentos {
			err := alimentoReceta.ValidarAlimentoRecetaDto()
			if err != nil {
				return err
			}
			alimentoObtenido, err := service.AlimentoService.ObtenerAlimentoPorID(&alimentoReceta.IDAlimento, &receta.IDUsuario)
			if err != nil {
				return err
			}
			if alimentoObtenido.Stock < alimentoReceta.Cantidad {
				return Errors.ErrorNoHayStock
			}
			alimentoObtenido.Stock = alimentoObtenido.Stock - alimentoReceta.Cantidad
			alimentos = append(alimentos, *alimentoObtenido)

			listaAlimentosReceta = append(listaAlimentosReceta, convertirAlimentoRecetaAModel(alimentoReceta))
		}

		recetaModel := Models.Receta{
			IDUsuario:     receta.IDUsuario,
			Nombre:        receta.Nombre,
			Alimentos:     listaAlimentosReceta,
			Momento:       Models.Momento(receta.Momento),
			FechaCreacion: time.Now(),
		}
		_, err := service.RecetaRepositorio.CrearReceta(&recetaModel)
		if err != nil {
			return err
		}

		for _, alimento := range alimentos {
			err := service.AlimentoService.ActualizarAlimento(&alimento)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func (service *RecetaService) EliminarReceta(idReceta *string, idUsuario *string) *Errors.ErrorCodigo {
	recetaObtenida, err := service.RecetaRepositorio.ObtenerRecetaPorID(idReceta, idUsuario)
	if err != nil {
		return err
	}
	_, err = service.RecetaRepositorio.EliminarReceta(idReceta, idUsuario)
	if err != nil {
		return err
	}
	for _, alimentoReceta := range recetaObtenida.Alimentos {
		alimentoObtenido, err := service.AlimentoService.ObtenerAlimentoPorID(&alimentoReceta.IDAlimento, idUsuario)
		if err != nil {
			return err
		}
		alimentoObtenido.Stock = alimentoObtenido.Stock + alimentoReceta.Cantidad
		err = service.AlimentoService.ActualizarAlimento(alimentoObtenido)
		if err != nil {
			return err
		}
	}
	return nil
}
func (service *RecetaService) ContarRecetasPorMomento(idUsuario *string) (map[string]int, *Errors.ErrorCodigo) {
	recetas, err := service.RecetaRepositorio.ContarRecetasPorMomento(idUsuario)
	if err != nil {
		return nil, err
	}
	return recetas, nil
}
func (service *RecetaService) ContarCantidadDeRecetasPorTipoAlimento(idUsuario *string) (map[string]int, *Errors.ErrorCodigo) {
	resultado := make(map[string]int)
	filtro := &Dto.FiltroAlimentoDto{
		TipoAlimentoDto: "",
	}
	recetas, err, _ := service.RecetaRepositorio.ObtenerRecetas(filtro, idUsuario)
	if err != nil {
		return nil, err
	}
	for _, receta := range recetas {
		banderaFruta := true
		banderaVerdura := true
		banderaLacteo := true
		banderaCarne := true
		for _, alimento := range receta.Alimentos {
			alimentoObtenido, err := service.AlimentoService.ObtenerAlimentoPorID(&alimento.IDAlimento, idUsuario)
			if err != nil {
				return nil, err
			}
			if banderaFruta && alimentoObtenido.TipoAlimento == Dto.Fruta {
				resultado[string(Models.Fruta)]++
				banderaFruta = false
			}
			if banderaVerdura && alimentoObtenido.TipoAlimento == Dto.Verdura {
				resultado[string(Models.Verdura)]++
				banderaVerdura = false
			}
			if banderaLacteo && alimentoObtenido.TipoAlimento == Dto.Lacteo {
				resultado[string(Models.Lacteo)]++
				banderaLacteo = false
			}
			if banderaCarne && alimentoObtenido.TipoAlimento == Dto.Carne {
				resultado[string(Models.Carne)]++
				banderaCarne = false
			}
		}
	}
	return resultado, nil
}
func convertirReceta(receta Models.Receta) *Dto.RecetaDto {
	var listaAlimentosRecetaDto []Dto.AlimentoRecetaDto
	for _, alimentoReceta := range receta.Alimentos {
		listaAlimentosRecetaDto = append(listaAlimentosRecetaDto, *convertirAlimentoRecetaADto(alimentoReceta))
	}

	return &Dto.RecetaDto{
		ID:        receta.ID.Hex(),
		IDUsuario: receta.IDUsuario,
		Nombre:    receta.Nombre,
		Alimentos: listaAlimentosRecetaDto,
		Momento:   convertirMomentoaDto(receta.Momento),
	}
}

func convertirAlimentoRecetaADto(alimentoReceta Models.AlimentoReceta) *Dto.AlimentoRecetaDto {

	return &Dto.AlimentoRecetaDto{
		IDAlimento: alimentoReceta.IDAlimento,
		Cantidad:   alimentoReceta.Cantidad,
	}
}

func convertirAlimentoRecetaAModel(alimentoRecetaDto Dto.AlimentoRecetaDto) Models.AlimentoReceta {
	return Models.AlimentoReceta{
		IDAlimento: alimentoRecetaDto.IDAlimento,
		Cantidad:   alimentoRecetaDto.Cantidad,
	}
}

func convertirMomentoaDto(momento Models.Momento) Dto.Momento {
	return Dto.Momento(momento)
}
func convertirSliceAPunterosRecetas(recetas []*Dto.RecetaDto) []Dto.RecetaDto {
	var recetasDto []Dto.RecetaDto
	for _, receta := range recetas {
		recetasDto = append(recetasDto, *receta)
	}
	return recetasDto
}
