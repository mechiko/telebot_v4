package usecase

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (uc *usecase) ApiGetTelebotUsers() (*entity.TelebotUserList, error) {
	defer uc.App.GetRecovery().RecoverLog("ApiGetTelebotUsers")
	if list, err := uc.App.GetRepo().GetTelebotUsers().GetList(); err != nil {
		return nil, fmt.Errorf("%w", err)
	} else {
		return list, nil
	}
}
