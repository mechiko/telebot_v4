package bot

import (
	"net/mail"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
)

// проверяем заполненность полей в состоянии пользователя
// true если все заполнено и правильно по датам
//
//	func (b *Bot) checkUserStateMod(c tele.Context) bool {
//		var state *entity.StateDialog
//		chat := c.Chat()
//		if state = b.States.GetState(chat.ID); state != nil {
//			switch state.Mode {
//			case entity.ModeEditUserStates:
//				return true
//			}
//		}
//		return false
//	}
//
// проверяем состояния по ключам из таблицы не просто присутствие любого
func (b Bot) checkAllStates(um *entity.UserStatesMaps) bool {
	defer b.App.GetRecovery().RecoverLog("bot:checkAllStates")
	if um == nil {
		return false
	}
	result := true
	examens := b.App.GetRepo().GetKeyStates().GetExamenKeys()
	for _, ex := range examens {
		if val, ok := um.Examens[ex]; ok {
			if !b.checkStateDate(val) {
				um.Examens[ex] += " неправильная дата!"
				result = false
			}
		} else {
			um.Examens[ex] += " отсутствует поле настройки"
			result = false
		}
	}
	intros := b.App.GetRepo().GetKeyStates().GetIntroKeys()
	for _, ex := range intros {
		if val, ok := um.Intro[ex]; ok {
			if !b.checkIntro(val) {
				um.Intro[ex] += " пустой идентификатор!"
				result = false
			}
		} else {
			um.Intro[ex] += " отсутствует поле настройки"
			result = false
		}
	}
	return result
}

// проверка что это дата и она в будущем
func (b Bot) checkStateDate(dt string) bool {
	defer b.App.GetRecovery().RecoverLog("bot:checkStateDate")
	l := b.App.GetConfiguration().Layouts.TimeLayoutDay
	nowDate := time.Now().Local()
	if date, err := time.Parse(l, dt); err != nil {
		return false
	} else {
		if date.Before(nowDate) {
			return false
		}
	}
	return true
}

// проверка что это дата и она в будущем
func (b Bot) checkEmail(email string) bool {
	if adr, err := mail.ParseAddress(email); err != nil {
		return false
	} else {
		b.App.GetLogger().Debugf("email %v", adr)
	}
	return true
}

// проверка что это дата и она в будущем
func (b Bot) checkIntro(val string) bool {
	return val != ""
}

func (b Bot) copyUserStates(um *entity.UserStatesMaps) *entity.UserStatesMaps {
	m := &entity.UserStatesMaps{
		UserId:  um.UserId,
		Intro:   make(map[string]string),
		Examens: make(map[string]string),
	}
	for k, v := range um.Intro {
		m.Intro[k] = v
	}
	for k, v := range um.Examens {
		m.Examens[k] = v
	}
	return m
}
