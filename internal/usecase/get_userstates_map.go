package usecase

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

// всегда возвращает мап даже при ошибке пустой
// по всем пользователям
func (uc *usecase) GetUsersStatesMap() (entity.UsersStatesMap, error) {
	m := make(entity.UsersStatesMap)
	if l, err := uc.App.GetRepo().GetUserStates().GetUserStateList(); err != nil {
		return m, fmt.Errorf("%w", err)
	} else {
		for _, v := range l.Items {
			if userMap, ok := m[v.UserId]; !ok {
				userMap = entity.UserStatesMaps{
					UserId:  v.UserId,
					Intro:   make(map[string]string),
					Examens: make(map[string]string),
				}
				keysExam := uc.App.GetRepo().GetKeyStates().GetExamenKeys()
				for _, ex := range keysExam {
					userMap.Examens[ex] = ""
				}
				keysIntro := uc.App.GetRepo().GetKeyStates().GetIntroKeys()
				for _, ex := range keysIntro {
					userMap.Intro[ex] = ""
				}
				if v.IsExamen {
					userMap.Examens[v.Key] = v.Value
				}
				if v.IsIntro {
					userMap.Intro[v.Key] = v.Value
				}
				m[v.UserId] = userMap
			} else {
				if v.IsExamen {
					userMap.Examens[v.Key] = v.Value
				}
				if v.IsIntro {
					userMap.Intro[v.Key] = v.Value
				}
				m[v.UserId] = userMap
			}
		}
	}
	return m, nil
}

// всегда возвращает мап даже при ошибке пустой
func (uc *usecase) GetUserStatesMap(userid int64) (*entity.UserStatesMaps, error) {
	m := &entity.UserStatesMaps{
		UserId:  userid,
		Intro:   make(map[string]string),
		Examens: make(map[string]string),
	}
	if l, err := uc.App.GetRepo().GetUserStates().GetUserStateListByUser(userid); err != nil {
		return m, fmt.Errorf("%w", err)
	} else {
		keysExam := uc.App.GetRepo().GetKeyStates().GetExamenKeys()
		for _, ex := range keysExam {
			m.Examens[ex] = ""
		}
		keysIntro := uc.App.GetRepo().GetKeyStates().GetIntroKeys()
		for _, ex := range keysIntro {
			m.Intro[ex] = ""
		}

		for _, v := range l.Items {
			if v.IsExamen {
				m.Examens[v.Key] = v.Value
			}
			if v.IsIntro {
				m.Intro[v.Key] = v.Value
			}
		}
	}
	return m, nil
}
