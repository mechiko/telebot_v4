package usecase

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/usecase/apiexamenall"
)

func (uc *usecase) ApiGetAllExamens(master string) (*entity.ApiExamenAll, error) {
	if err := uc.ApiClearAll(); err != nil {
		uc.App.GetLogger().Errorf("usecase:apiclearall %s", err.Error())
	}
	if apiAll, err := apiexamenall.New(uc.App, master).GetAll(); err != nil {
		return apiAll, fmt.Errorf("usecase:ApiGetAllExamens %w", err)
	} else {
		return apiAll, nil
	}
}
