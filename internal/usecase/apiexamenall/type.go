package apiexamenall

import "github.com/mechiko/telebot_v4/internal/entity"

type IApiMissionAll interface {
	GetAll() (*entity.ApiExamenAll, error)
}

type apiExamenAll struct {
	*entity.ApiExamenAll
	app    entity.Application
	master string
}

// asserts
// var _ entity.ApiExamenAll = (*useCase)(nil)

// передаем идентификатор мастера для отбора
func New(app entity.Application, master string) *apiExamenAll {
	apiExAll := &entity.ApiExamenAll{
		Day0:     make([]*entity.ExamenUser, 0),
		Day3:     make([]*entity.ExamenUser, 0),
		Month:    make([]*entity.ExamenUser, 0),
		InFuture: make([]*entity.ExamenUser, 0),
	}
	return &apiExamenAll{
		ApiExamenAll: apiExAll,
		app:          app,
		master:       master,
	}
}
