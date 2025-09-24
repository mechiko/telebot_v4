package usecase

import (
	"fmt"
	"sort"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/template"
)

func (uc *usecase) UsersReminder() error {
	defer uc.App.GetRecovery().RecoverLog("usecase:UserReminder")
	// if err := uc.ApiClearAll(); err != nil {
	// 	uc.App.GetLogger().Errorf("usecase:UsersReminder %s", err.Error())
	// }
	if listUsers, err := uc.App.GetRepo().GetTelebotUsers().GetList(); err != nil {
		uc.App.GetLogger().Errorf("usecase:UserReminder %w", err)
	} else {
		if len(listUsers.Items) > 0 {
			for _, tu := range listUsers.Items {
				if info, err := uc.getTelebotUserReminderInfo(tu.ID); err != nil {
					uc.App.GetLogger().Errorf("usecase:UserReminder %w", err)
					continue
				} else {
					if info.MissionDay0 == "" && info.MissionDay3 == "" && info.MissionDay30 == "" && len(info.ExamenDay0) == 0 && len(info.ExamenDay3) == 0 && len(info.ExamenDay30) == 0 && len(info.ExamenEmpty) == 0 && len(info.ExamenBad) == 0 {
						continue
					}
					info.UserIdent = tu.Ident
					if strTemplate, err := template.NewTemplate(uc.App, "").ReminderUserInfo(info); err != nil {
						uc.App.GetLogger().Errorf("usecase:UserReminder %w", err)
						continue
					} else {
						if err := uc.App.GetBot().SendMessageAdmin(strTemplate, "HTML"); err != nil {
							uc.App.GetLogger().Errorf("usecase:UserReminder %w", err)
						}
						if err := uc.App.GetBot().SendMessageUser(info.UserId, strTemplate, "HTML"); err != nil {
							uc.App.GetLogger().Errorf("usecase:UserReminder %w", err)
						}
					}
				}
			}
		}
	}
	return nil
}

func (uc *usecase) getTelebotUserReminderInfo(uid int64) (*entity.TelebotUserReminderInfo, error) {
	defer uc.App.GetRecovery().RecoverLog("usecase:getTelebotUserReminderInfo")
	result := &entity.TelebotUserReminderInfo{
		UserId:        uid,
		ExamenDay0:    make([]*entity.Examen, 0),
		ExamenDay3:    make([]*entity.Examen, 0),
		ExamenDay30:   make([]*entity.Examen, 0),
		ExamenDayOt60: make([]*entity.Examen, 0),
		ExamenEmpty:   make([]*entity.Examen, 0),
		ExamenBad:     make([]*entity.Examen, 0),
	}
	l := uc.App.GetConfiguration().Layouts.TimeLayoutDay
	loc, _ := time.LoadLocation("Europe/Moscow")
	t := time.Now().In(loc)
	nowDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if mission, err := uc.App.GetRepo().GetMissions().GetActiveByUid(uid); err != nil {
		uc.App.GetLogger().Errorf("usecase:getTelebotUserReminderInfo %s", err)
	} else if mission.Start != "" {
		if startMission, err := time.ParseInLocation(l, mission.Start, loc); err != nil {
			uc.App.GetLogger().Errorf("usecase:getTelebotUserReminderInfo %s", err)
		} else {
			days := int(startMission.Sub(nowDate).Hours() / 24)
			switch days {
			case 0:
				result.MissionDay0 = fmt.Sprintf("командировка в %s %s", mission.Place, mission.Start)
			case 3:
				result.MissionDay3 = fmt.Sprintf("командировка в %s %s", mission.Place, mission.Start)
			case 30:
				result.MissionDay30 = fmt.Sprintf("командировка в %s %s", mission.Place, mission.Start)
			}
		}
	}
	// берем мапу экзаменов преврещаем ее в массив структур и сортируем
	if userState, err := uc.GetUserStatesMap(uid); err != nil {
		uc.App.GetLogger().Errorf("usecase:getTelebotUserReminderInfo %s", err)
	} else {
		// по задумке тут всегда будет мап
		arr := make([]*entity.Examen, 0)
		for k, v := range userState.Examens {
			// если дата пустая
			if v == "" {
				result.ExamenEmpty = append(result.ExamenEmpty, &entity.Examen{Name: k, Date: time.Now().Local(), DateStr: v})
				continue
			}
			date, err := time.ParseInLocation(l, v, loc)
			if err != nil {
				uc.App.GetLogger().Errorf("usecase:getTelebotUserReminderInfo %s", err)
				result.ExamenBad = append(result.ExamenBad, &entity.Examen{Name: k, Date: time.Now().Local(), DateStr: v})
				continue
			}
			arr = append(arr, &entity.Examen{Name: k, Date: date, DateStr: v})
		}
		sort.Slice(arr, func(i, j int) bool {
			return arr[i].Date.Before(arr[j].Date)
		})
		for _, v := range arr {
			days := int(v.Date.Sub(nowDate).Hours() / 24)
			switch days {
			case 0:
				result.ExamenDay0 = append(result.ExamenDay0, v)
			case 3:
				result.ExamenDay3 = append(result.ExamenDay3, v)
			case 30:
				result.ExamenDay30 = append(result.ExamenDay30, v)
			case 60:
				if v.Name == "ОТ" {
					result.ExamenDayOt60 = append(result.ExamenDayOt60, v)
				}
			}
		}
	}
	return result, nil
}
