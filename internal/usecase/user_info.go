package usecase

import (
	"fmt"
	"sort"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/template"
)

func (uc *usecase) GetUserInfo(state *entity.StateDialog) (string, error) {
	defer uc.App.GetRecovery().RecoverLog("usecase:GetUserInfo")
	if state.User != nil {
		return uc.getUserInfoEmploye(state)
	}
	return "", nil
}

func (uc *usecase) getUserInfoEmploye(state *entity.StateDialog) (string, error) {
	defer uc.App.GetRecovery().RecoverLog("usecase:getUserInfoEmploye")
	result := ""
	mission := state.Mission
	l := uc.App.GetConfiguration().Layouts.TimeLayoutDay
	//init the loc
	loc, _ := time.LoadLocation("Europe/Moscow")
	t := time.Now().In(loc)
	nowDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	startMission, err := time.ParseInLocation(l, mission.Start, loc)
	if err != nil {
		return result, fmt.Errorf("%w", err)
	}
	endMission, err := time.ParseInLocation(l, mission.End, loc)
	if err != nil {
		return result, fmt.Errorf("%w", err)
	}
	if endMission.Before(nowDate) {
		// командировка завершилась, отправляем в архив
		result = `Командировка завершилась и переведена в архив`
		if err := uc.App.GetRepo().GetMissions().UpdateClearActive(mission); err != nil {
			return result, fmt.Errorf("%w", err)
		}
	} else {
		if startMission.Before(nowDate) {
			// командировка сейчас идет
			days := int(endMission.Sub(nowDate).Hours() / 24)
			result = fmt.Sprintf("Вы сейчас в командировке, она закончится через %v дней %s числа", days, mission.End)
		} else {
			// командировка еще не началась
			days := int(startMission.Sub(nowDate).Hours() / 24)
			switch days {
			case 0:
				result = fmt.Sprintf("Командировка начнется сегодня %s по %s", mission.Start, mission.End)
			case 1:
				result = fmt.Sprintf("Командировка начнется завтра %s по %s", mission.Start, mission.End)
			default:
				result = fmt.Sprintf("Командировка начнется с %s по %s", mission.Start, mission.End)
			}
		}
	}

	// берем мапу экзаменов преврещаем ее в массив структур и сортируем
	// sort.Slice(planets, func(i, j int) bool {
	// 	return planets[i].Axis < planets[j].Axis
	// })

	arr := make([]*entity.Examen, 0)
	empty := make([]*entity.Examen, 0)
	bad := make([]*entity.Examen, 0)
	for k, v := range state.UserState.Examens {
		if v == "" {
			empty = append(empty, &entity.Examen{Name: k, Date: time.Now().Local(), DateStr: v})
			continue
		}
		l := uc.App.GetConfiguration().Layouts.TimeLayoutDay
		loc, _ := time.LoadLocation("Europe/Moscow")
		date, err := time.ParseInLocation(l, v, loc)
		if err != nil {
			uc.App.GetLogger().Errorf("usecase:userinfo %s", err.Error())
			bad = append(bad, &entity.Examen{Name: k, Date: time.Now().Local(), DateStr: v})
			continue
		}
		arr = append(arr, &entity.Examen{Name: k, Date: date, DateStr: v})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Date.Before(arr[j].Date)
	})
	arrBefore := make([]*entity.Examen, 0)
	arrIn := make([]*entity.Examen, 0)
	for _, v := range arr {
		if v.Date.Before(startMission) {
			arrBefore = append(arrBefore, v)
		} else if !v.Date.After(endMission) { //uc.dateInMission(v.Date, startMission, endMission) {
			arrIn = append(arrIn, v)
		}
	}
	userInfo := &entity.UserInfo{
		Mission:         result,
		BeforeExamens:   arrBefore,
		ConflictExamens: arrIn,
		Empty:           empty,
		Bad:             bad,
	}
	if strTemplate, err := template.NewTemplate(uc.App, "").UserInfo(userInfo); err != nil {
		return result, fmt.Errorf("%w", err)
	} else {
		return strTemplate, nil
	}
}
