package Services

import (
	"supercook/Dto"
	"supercook/Models"
	"supercook/Repositories"
	"supercook/Utils"
	"time"
)

type RecetaInterface interface {
	ObtenerRecetas(filtro *[3]string, idUsuario *string) []*Dto.RecetaDto
	ObtenerRecetaPorID(idReceta *string, idUsuario *string) *Dto.RecetaDto
	CrearReceta(receta *Dto.RecetaDto) *Dto.Resultado
	ActualizarReceta(receta *Dto.RecetaDto) *Dto.Resultado
	EliminarReceta(idReceta *string, idUsuario *string) *Dto.Resultado
}

type RecetaService struct {
	RecetaRepositorio Repositories.RecetaRepositorioInterface
}

func NuevoRecetaService(recetaRepositorio Repositories.RecetaRepositorioInterface) *RecetaService {
	return &RecetaService{
		RecetaRepositorio: recetaRepositorio,
	}
}

func (service *RecetaService) ObtenerRecetas(filtro *[3]string, idUsuario *string) []*Dto.RecetaDto {
	recetas, _ := service.RecetaRepositorio.ObtenerRecetas(filtro, idUsuario)
	var recetasDto []*Dto.RecetaDto
	for _, receta := range recetas {
		recetasDto = append(recetasDto, convertirReceta(receta))
	}
	return recetasDto
}

func (service *RecetaService) ObtenerRecetaPorID(idReceta *string, idUsuario *string) *Dto.RecetaDto {
	receta, _ := service.RecetaRepositorio.ObtenerRecetaPorID(idReceta, idUsuario)
	return convertirReceta(receta)
}

func (service *RecetaService) CrearReceta(receta *Dto.RecetaDto) *Dto.Resultado {
	resultado := Dto.Resultado{}
	resultado.ListaMensaje = append(resultado.ListaMensaje, receta.ValidarRecetaDto())
	if len(resultado.ListaMensaje) > 0 {
		resultado.BoolResultado = false
		return &resultado
	} else {
		var listaAlimentosReceta []Models.AlimentoReceta
		for _, alimentoReceta := range receta.Alimentos {
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
			resultado.BoolResultado = false
			resultado.ListaMensaje = append(resultado.ListaMensaje, err.Error())
		} else {
			resultado.BoolResultado = true
		}
		return &resultado
	}
}

func (service *RecetaService) ActualizarReceta(receta *Dto.RecetaDto) *Dto.Resultado {
	resultado := Dto.Resultado{}
	resultado.ListaMensaje = append(resultado.ListaMensaje, receta.ValidarRecetaDto())
	if len(resultado.ListaMensaje) > 0 {
		resultado.BoolResultado = false
		return &resultado
	} else {
		var listaAlimentosReceta []Models.AlimentoReceta
		for _, alimentoReceta := range receta.Alimentos {
			listaAlimentosReceta = append(listaAlimentosReceta, convertirAlimentoRecetaAModel(alimentoReceta))
		}

		recetaModel := Models.Receta{
			ID:                 Utils.GetObjectIDFromStringID(receta.ID),
			IDUsuario:          receta.IDUsuario,
			Nombre:             receta.Nombre,
			Alimentos:          listaAlimentosReceta,
			Momento:            Models.Momento(receta.Momento),
			FechaActualizacion: time.Now(),
		}
		_, err := service.RecetaRepositorio.ActualizarReceta(&recetaModel)
		if err != nil {
			resultado.BoolResultado = false
			resultado.ListaMensaje = append(resultado.ListaMensaje, err.Error())
		} else {
			resultado.BoolResultado = true
		}
		return &resultado
	}
}

func (service *RecetaService) EliminarReceta(idReceta *string, idUsuario *string) *Dto.Resultado {
	resultado := Dto.Resultado{}
	_, err := service.RecetaRepositorio.EliminarReceta(idReceta, idUsuario)
	if err != nil {
		resultado.BoolResultado = false
		resultado.ListaMensaje = append(resultado.ListaMensaje, err.Error())
	} else {
		resultado.BoolResultado = true
	}
	return &resultado
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
