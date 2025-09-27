package userstate

import (
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
	"gopkg.in/telebot.v4"
)

// проверяем состояния по ключам из таблицы не просто присутствие любого
func (m *stateMission) checkAllStates(um *entity.UserStatesMaps) bool {
	defer m.App.GetRecovery().RecoverLog("userstate:checkAllStates")
	if um == nil {
		return false
	}
	result := true
	examens := m.App.GetRepo().GetKeyStates().GetExamenKeys()
	for _, ex := range examens {
		if val, ok := um.Examens[ex]; ok {
			if !m.checkStateDate(val) {
				um.Examens[ex] += " неправильная дата!"
				result = false
			}
		} else {
			um.Examens[ex] += " отсутствует поле настройки"
			result = false
		}
	}
	intros := m.App.GetRepo().GetKeyStates().GetIntroKeys()
	for _, ex := range intros {
		if val, ok := um.Intro[ex]; ok {
			if !m.checkIntro(val) {
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
func (m *stateMission) checkStateDate(dt string) bool {
	defer m.App.GetRecovery().RecoverLog("userstate:checkStateDate")
	l := m.App.GetConfiguration().Layouts.TimeLayoutDay
	nowDate := time.Now().Local()
	if date, err := time.Parse(l, dt); err != nil {
		return false
	} else {
		days := int(date.Sub(nowDate).Hours() / 24)
		if days < 0 {
			return false
		}
	}
	return true
}

// проверка что это дата и она в будущем
func (m *stateMission) checkIntro(val string) bool {
	return val != ""
}

func (m *stateMission) copyUserStates(um *entity.UserStatesMaps) *entity.UserStatesMaps {
	m1 := &entity.UserStatesMaps{
		UserId:  um.UserId,
		Intro:   make(map[string]string),
		Examens: make(map[string]string),
	}
	for k, v := range um.Intro {
		m1.Intro[k] = v
	}
	for k, v := range um.Examens {
		m1.Examens[k] = v
	}
	return m1
}

// проверяем старые меню если присылают коллбэки
func (m *stateMission) checkMsgId(c telebot.Context) error {
	// если ID не соответствует сохраненном, то удалим это сообщение
	msg := c.Message()
	if m.State.UserStateMenuMsg.MessageID != int64(msg.ID) {
		c.Bot().Delete(msg)
	}
	return nil
}
