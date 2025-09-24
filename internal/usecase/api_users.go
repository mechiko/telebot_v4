package usecase

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (uc *usecase) ApiGetUsers() (*entity.UserList, error) {
	defer uc.App.GetRecovery().RecoverLog("GetUsers")
	appList := &entity.UserList{}
	apps := uc.App.GetRepo().GetUsers()
	if err := apps.Get(); err != nil {
		return appList, fmt.Errorf("%w", err)
	}
	appList.Items = apps.GetItems()
	return appList, nil
}
