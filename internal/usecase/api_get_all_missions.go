package usecase

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/usecase/apimissionall"
)

func (uc *usecase) ApiGetAllMissions(master string) (*entity.ApiMissionAll, error) {
	if err := uc.ApiClearAll(); err != nil {
		uc.App.GetLogger().Errorf("usecase:apiclearall %s", err.Error())
	}
	if apiAll, err := apimissionall.New(uc.App, master).GetAll(); err != nil {
		return apiAll, fmt.Errorf("usecase:ApiGetAllMissions %w", err)
	} else {
		return apiAll, nil
	}
}
