package usecase

import (
	"fmt"
	"sort"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/template"
)

func (uc *usecase) UsersReminderCheck() error {
	defer uc.App.GetRecovery().RecoverLog("usecase:UsersReminderCheck")

	if info, err := uc.getTelebotUserReminderInfoCheck(0); err != nil {
		return fmt.Errorf("usecase:UsersReminderCheck %w", err)
	} else {
		if strTemplate, err := template.NewTemplate(uc.App, "").ReminderUserInfo(info); err != nil {
			return fmt.Errorf("usecase:UsersReminderCheck %w", err)
		} else {
			if err := uc.App.GetBot().SendMessageAdmin(strTemplate, "HTML"); err != nil {
				uc.App.GetLogger().Errorf("usecase:UsersReminderCheck %w", err)
			}
		}
	}
	if info, err := uc.getTelebotUserReminderInfoCheck(3); err != nil {
		return fmt.Errorf("usecase:UsersReminderCheck %w", err)
	} else {
		if strTemplate, err := template.NewTemplate(uc.App, "").ReminderUserInfo(info); err != nil {
			return fmt.Errorf("usecase:UsersReminderCheck %w", err)
		} else {
			if err := uc.App.GetBot().SendMessageAdmin(strTemplate, "HTML"); err != nil {
				uc.App.GetLogger().Errorf("usecase:UsersReminderCheck %w", err)
			}
		}
	}
	if info, err := uc.getTelebotUserReminderInfoCheck(30); err != nil {
		return fmt.Errorf("usecase:UsersReminderCheck %w", err)
	} else {
		if strTemplate, err := template.NewTemplate(uc.App, "").ReminderUserInfo(info); err != nil {
			return fmt.Errorf("usecase:UsersReminderCheck %w", err)
		} else {
			if err := uc.App.GetBot().SendMessageAdmin(strTemplate, "HTML"); err != nil {
				uc.App.GetLogger().Errorf("usecase:UsersReminderCheck %w", err)
			}
		}
	}
	if info, err := uc.getTelebotUserReminderInfoCheck(31); err != nil {
		return fmt.Errorf("usecase:UsersReminderCheck %w", err)
	} else {
		if strTemplate, err := template.NewTemplate(uc.App, "").ReminderUserInfo(info); err != nil {
			return fmt.Errorf("usecase:UsersReminderCheck %w", err)
		} else {
			if err := uc.App.GetBot().SendMessageAdmin(strTemplate, "HTML"); err != nil {
				uc.App.GetLogger().Errorf("usecase:UsersReminderCheck %w", err)
			}
		}
	}
	if info, err := uc.getTelebotUserReminderInfoCheck(99); err != nil {
		return fmt.Errorf("usecase:UsersReminderCheck %w", err)
	} else {
		if strTemplate, err := template.NewTemplate(uc.App, "").ReminderUserInfo(info); err != nil {
			return fmt.Errorf("usecase:UsersReminderCheck %w", err)
		} else {
			if err := uc.App.GetBot().SendMessageAdmin(strTemplate, "HTML"); err != nil {
				uc.App.GetLogger().Errorf("usecase:UsersReminderCheck %w", err)
			}
		}
	}
	return nil
}

func (uc *usecase) getTelebotUserReminderInfoCheck(v int) (*entity.TelebotUserReminderInfo, error) {
	defer uc.App.GetRecovery().RecoverLog("usecase:getTelebotUserReminderInfoCheck")
	var mission *entity.Mission
	result := &entity.TelebotUserReminderInfo{
		UserIdent:   "11111",
		UserId:      0,
		ExamenDay0:  make([]*entity.Examen, 0),
		ExamenDay3:  make([]*entity.Examen, 0),
		ExamenDay30: make([]*entity.Examen, 0),
		ExamenEmpty: make([]*entity.Examen, 0),
		ExamenBad:   make([]*entity.Examen, 0),
	}
	l := uc.App.GetConfiguration().Layouts.TimeLayoutDay
	loc, _ := time.LoadLocation("Europe/Moscow")
	t := time.Now().In(loc)
	nowDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	switch v {
	case 0:
		mission = &entity.Mission{Place: "Test Place", Start: nowDate.Format(l), End: nowDate.AddDate(0, 1, 0).Format(l)}
	case 3:
		mission = &entity.Mission{Place: "Test Place", Start: nowDate.AddDate(0, 0, 3).Format(l), End: nowDate.AddDate(0, 1, 0).Format(l)}
	case 30:
		mission = &entity.Mission{Place: "Test Place", Start: nowDate.AddDate(0, 0, 30).Format(l), End: nowDate.AddDate(0, 1, 0).Format(l)}
	case 31:
		mission = &entity.Mission{Place: "Test Place", Start: nowDate.AddDate(0, 0, 30).Format(l), End: nowDate.AddDate(0, 1, 0).Format(l)}
	default:
		mission = &entity.Mission{}
	}
	if startMission, err := time.ParseInLocation(l, mission.Start, loc); err != nil {
		uc.App.GetLogger().Errorf("usecase:getTelebotUserReminderInfoCheck %s", err)
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
	// берем мапу экзаменов преврещаем ее в массив структур и сортируем
	// по задумке тут всегда будет мап
	userState := &entity.UserStatesMaps{
		UserId:  0,
		Intro:   make(map[string]string),
		Examens: make(map[string]string),
	}
	keysExam := uc.App.GetRepo().GetKeyStates().GetExamenKeys()
	for _, ex := range keysExam {
		userState.Examens[ex] = ""
	}
	keysIntro := uc.App.GetRepo().GetKeyStates().GetIntroKeys()
	for _, ex := range keysIntro {
		userState.Intro[ex] = ""
	}
	switch v {
	case 0:
		userState.Intro["ИДЕНТ"] = "00000"
		userState.Examens["ОТ"] = "09.07.2023"
		userState.Examens["ППБ"] = "09.07.2023"
		userState.Examens["ОАЭ АС, ДИ"] = "09.07.2023"
		userState.Examens["ФНП"] = "09.07.2023"
		userState.Examens["ПРБ"] = "09.07.2023"
		userState.Examens["ЭБ"] = "09.07.2023"
		userState.Examens["Медосмотр"] = "09.07.2023"
	case 3:
		userState.Intro["ИДЕНТ"] = "11111"
		userState.Examens["ОТ"] = "12.07.2023"
		userState.Examens["ППБ"] = "12.07.2023"
		userState.Examens["ОАЭ АС, ДИ"] = "12.07.2023"
		userState.Examens["ФНП"] = "12.07.2023"
		userState.Examens["ПРБ"] = "12.07.2023"
		userState.Examens["ЭБ"] = "12.07.2023"
		userState.Examens["Медосмотр"] = "12.07.2023"
	case 30:
		userState.Intro["ИДЕНТ"] = "11111"
		userState.Examens["ОТ"] = "08.08.2023"
		userState.Examens["ППБ"] = "08.08.2023"
		userState.Examens["ОАЭ АС, ДИ"] = "08.08.2023"
		userState.Examens["ФНП"] = "08.08.2023"
		userState.Examens["ПРБ"] = "08.08.2023"
		userState.Examens["ЭБ"] = "08.08.2023"
		userState.Examens["Медосмотр"] = "08.08.2023"
	case 31:
		userState.Intro["ИДЕНТ"] = "11111"
		userState.Examens["ОТ"] = "09.08.2023"
		userState.Examens["ППБ"] = "09.08.2023"
		userState.Examens["ОАЭ АС, ДИ"] = "09.08.2023"
		userState.Examens["ФНП"] = "09.08.2023"
		userState.Examens["ПРБ"] = "09.08.2023"
		userState.Examens["ЭБ"] = "09.08.2023"
		userState.Examens["Медосмотр"] = "09.08.2023"
	default:
		userState.Intro["ИДЕНТ"] = ""
		userState.Examens["ОТ"] = ""
		userState.Examens["ППБ"] = ""
		userState.Examens["ОАЭ АС, ДИ"] = ""
		userState.Examens["ФНП"] = ""
		userState.Examens["ПРБ"] = ""
		userState.Examens["ЭБ"] = ""
		userState.Examens["Медосмотр"] = ""
	}

	arr := make([]*entity.Examen, 0)
	for k, v := range userState.Examens {
		// если дата пустая
		if v == "" {
			result.ExamenEmpty = append(result.ExamenEmpty, &entity.Examen{Name: k, Date: time.Now().Local(), DateStr: v})
			continue
		}
		date, err := time.ParseInLocation(l, v, loc)
		if err != nil {
			uc.App.GetLogger().Errorf("usecase:getTelebotUserReminderInfoCheck %s", err)
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
		}
	}
	return result, nil
}
