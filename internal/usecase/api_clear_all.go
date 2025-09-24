package usecase

import (
	"fmt"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (uc *usecase) ApiClearAll() error {
	if allMiss, err := uc.App.GetRepo().GetViewMissions().GetAll(); err != nil {
		return fmt.Errorf("usecase:ApiClearAll %w", err)
	} else {
		if len(allMiss.Items) > 0 {
			for _, m := range allMiss.Items {
				if m.BeforeEndDays < 0 {
					// командировка закончилась
					uc.App.GetLogger().Infof("missions ended now[%v] user %s[%s] palce %s start %s end %s", time.Now().Local(), m.UserName, m.UserIdent, m.Place, m.StartOrdered, m.EndOrdered)
					if err := uc.App.GetRepo().GetMissions().UpdateClearActiveById(m.MissionId); err != nil {
						uc.App.GetLogger().Errorf("usecase:ApiClearAll %w", err)
					}
				}
			}
		}
	}
	if allEx, err := uc.App.GetRepo().GetViewExamens().GetAll(); err != nil {
		return fmt.Errorf("usecase:ApiClearAll %w", err)
	} else {
		if len(allEx.Items) > 0 {
			for _, m := range allEx.Items {
				if m.BeforeDateDays < 0 {
					// экзамен закончился делаем дату пустой
					examenEnded := &entity.ExamenEnded{UserId: m.UserId, Examen: m.ExamenKey, Date: m.Date}
					if err := uc.App.GetRepo().GetExamenEndeds().Insert(examenEnded); err != nil {
						uc.App.GetLogger().Errorf("usecase:ApiClearAll %w", err)
					}
					uc.App.GetLogger().Infof("examens ended user %s[%s] key %s date %s", m.UserName, m.UserIdent, m.ExamenKey, m.Date)
					if err := uc.App.GetRepo().GetUserStates().Set(m.UserId, m.ExamenKey, ""); err != nil {
						uc.App.GetLogger().Errorf("usecase:ApiClearAll %w", err)
					}
				}
			}
		}
	}
	return nil
}
