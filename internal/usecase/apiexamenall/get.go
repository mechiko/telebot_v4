package apiexamenall

import (
	"fmt"
	"strings"

	"github.com/mechiko/telebot_v4/internal/entity"
)

// собираем все командировки по массивам
// если указан мастер то фильтруем по нему
// если срок меньше нуля надо очищать сделаем отдельную процедуру
func (am *apiExamenAll) GetAll() (*entity.ApiExamenAll, error) {
	if allEx, err := am.app.GetRepo().GetViewExamens().GetAll(); err != nil {
		return am.ApiExamenAll, fmt.Errorf("apiexamenall:getall %w", err)
	} else {
		if len(allEx.Items) > 0 {
			for _, m := range allEx.Items {
				if m.BeforeDateDays == 0 {
					if am.master != "" {
						if strings.Contains(m.Masters, am.master) {
							am.ApiExamenAll.Day0 = append(am.ApiExamenAll.Day0, m)
						}
						continue
					}
					am.ApiExamenAll.Day0 = append(am.ApiExamenAll.Day0, m)
					continue
				}
				if m.BeforeDateDays == 3 {
					if am.master != "" {
						if strings.Contains(m.Masters, am.master) {
							am.ApiExamenAll.Day3 = append(am.ApiExamenAll.Day3, m)
						}
						continue
					}
					am.ApiExamenAll.Day3 = append(am.ApiExamenAll.Day3, m)
					continue
				}
				if m.BeforeDateDays == 30 {
					if am.master != "" {
						if strings.Contains(m.Masters, am.master) {
							am.ApiExamenAll.Month = append(am.ApiExamenAll.Month, m)
						}
						continue
					}
					am.ApiExamenAll.Month = append(am.ApiExamenAll.Month, m)
					continue
				}
				if m.BeforeDateDays > 0 {
					if am.master != "" {
						if strings.Contains(m.Masters, am.master) {
							am.ApiExamenAll.InFuture = append(am.ApiExamenAll.InFuture, m)
						}
						continue
					}
					am.ApiExamenAll.InFuture = append(am.ApiExamenAll.InFuture, m)
					continue
				}
			}
		}
		return am.ApiExamenAll, nil
	}
}
