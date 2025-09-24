package apimissionall

import (
	"fmt"
	"strings"

	"github.com/mechiko/telebot_v4/internal/entity"
)

// собираем все командировки по массивам
// если указан мастер то фильтруем по нему
func (am *apiMissionAll) GetAll() (*entity.ApiMissionAll, error) {
	if allMiss, err := am.app.GetRepo().GetViewMissions().GetAll(); err != nil {
		return am.ApiMissionAll, fmt.Errorf("apimissionall:getall %w", err)
	} else {
		if len(allMiss.Items) > 0 {
			for _, m := range allMiss.Items {
				if m.BeforeStartDays > 0 {
					if am.master != "" {
						if strings.Contains(m.Masters, am.master) {
							am.ApiMissionAll.InFuture = append(am.ApiMissionAll.InFuture, m)
						}
						continue
					}
					am.ApiMissionAll.InFuture = append(am.ApiMissionAll.InFuture, m)
					continue
				}
				if m.BeforeEndDays > 0 {
					if am.master != "" {
						if strings.Contains(m.Masters, am.master) {
							am.ApiMissionAll.InProgress = append(am.ApiMissionAll.InProgress, m)
						}
						continue
					}
					am.ApiMissionAll.InProgress = append(am.ApiMissionAll.InProgress, m)
					continue
				}
				if am.master != "" {
					if strings.Contains(m.Masters, am.master) {
						am.ApiMissionAll.InPast = append(am.ApiMissionAll.InPast, m)
					}
					continue
				}
				am.ApiMissionAll.InPast = append(am.ApiMissionAll.InPast, m)
			}
		}
		return am.ApiMissionAll, nil
	}
}
