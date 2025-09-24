package usecase

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

// всегда возвращает командировку даже при ошибке пустую с ID 0 и RecepientId = chat.ID
func (uc *usecase) GetUserActiveMission(uid int64) (*entity.Mission, error) {
	defer uc.App.GetRecovery().RecoverLog("usecase:GetUserActiveMission")
	mission := &entity.Mission{ID: 0, RecepientId: uid}
	if miss, err := uc.App.GetRepo().GetMissions().GetActiveByUid(uid); err != nil {
		return mission, fmt.Errorf("%w", err)
	} else {
		return miss, nil
	}
}

// весь объект командировки нужен для удобства, а так берет из него только RecepientId
func (uc *usecase) SetUserActiveMission(mission *entity.Mission) error {
	defer uc.App.GetRecovery().RecoverLog("usecase:SetUserActiveMission")
	// сбрасываем во всех командировках флаг Активности
	if err := uc.App.GetRepo().GetMissions().UpdateClearActive(mission); err != nil {
		return fmt.Errorf("%w", err)
	}
	mission.Active = true
	if mission.ID == 0 {
		if err := uc.App.GetRepo().GetMissions().Insert(mission); err != nil {
			return fmt.Errorf("%w", err)
		}
	} else if err := uc.App.GetRepo().GetMissions().Update(mission); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// весь объект командировки нужен для удобства, а так берет из него только RecepientId
func (uc *usecase) ClearUserActiveMission(mission *entity.Mission) error {
	defer uc.App.GetRecovery().RecoverLog("usecase:ClearUserActiveMission")
	// сбрасываем во всех командировках флаг Активности
	if err := uc.App.GetRepo().GetMissions().UpdateClearActive(mission); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
