package apimissionall

import "github.com/mechiko/telebot_v4/internal/entity"

type IApiMissionAll interface {
	GetAll() (*entity.ApiMissionAll, error)
}

type apiMissionAll struct {
	*entity.ApiMissionAll
	app    entity.Application
	master string
}

// asserts
// var _ entity.ApiMissionAll = (*useCase)(nil)

// передаем идентификатор мастера для отбора
func New(app entity.Application, master string) *apiMissionAll {
	apiMissAll := &entity.ApiMissionAll{
		InFuture:   make([]*entity.MissionUser, 0),
		InProgress: make([]*entity.MissionUser, 0),
		InPast:     make([]*entity.MissionUser, 0),
		Day0:       make([]*entity.MissionUser, 0),
		Day3:       make([]*entity.MissionUser, 0),
		Month:      make([]*entity.MissionUser, 0),
	}
	apiAll := &apiMissionAll{
		ApiMissionAll: apiMissAll,
		app:           app,
		master:        master,
	}

	return apiAll
}
