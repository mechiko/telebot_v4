package missionadd

import (
	"fmt"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
	"gopkg.in/telebot.v4"
)

func (m *stateMission) Check(c telebot.Context) bool {
	return false
}

// проверяем дату в строке
func (m *stateMission) checkDate(d string) (string, bool) {
	defer m.App.GetRecovery().RecoverLog("checkDate")
	l := m.App.GetConfiguration().Layouts.TimeLayoutDay
	nowDate := time.Now().Local()
	if date, err := time.Parse(l, d); err != nil {
		return fmt.Sprintf("%s %s", d, err.Error()), false
	} else {
		fmtDate := date.Format(l)
		if date.Before(nowDate) {
			return fmt.Sprintf("%s дата не может быть раньше текущей даты %s", d, nowDate.Format(l)), false
		}
		return fmtDate, true
	}
}

// проверяем миссию в состоянии и прописываем ошибку в mission.ErrorText
func (m *stateMission) checkMission(c telebot.Context) bool {
	defer m.App.GetRecovery().RecoverLog("checkOnSave")
	var startDate, endDate time.Time
	var err error
	result := true
	if m.State.Mission == nil {
		m.State.Mission = &entity.Mission{}
	}
	m.State.Mission.ErrorTxt = ""
	l := m.App.GetConfiguration().Layouts.TimeLayoutDay
	nowDate := time.Now().Local()
	if m.State.Mission.Place == "" {
		result = false
		m.State.Mission.ErrorTxt += "место командировки не должно быть пустым \n"
	}
	if startDate, err = time.Parse(l, m.State.Mission.Start); err != nil {
		result = false
		m.State.Mission.ErrorTxt += "ошибка формата даты начала \n"
	}
	if endDate, err = time.Parse(l, m.State.Mission.End); err != nil {
		result = false
		m.State.Mission.ErrorTxt += "ошибка формата даты окончания"
	}
	if !result {
		return false
	}
	if startDate.Before(nowDate) {
		result = false
		m.State.Mission.ErrorTxt = "дата начала не должна быть меньше текущей даты \n"
	}
	if endDate.Before(startDate) {
		result = false
		m.State.Mission.ErrorTxt += "дата окончания меньше даты начала "
	}
	return result
}

// проверяем старые меню если присылают коллбэки
func (m *stateMission) checkMsgId(c telebot.Context) error {
	// если ID не соответствует сохраненном, то удалим это сообщение
	msg := c.Message()
	if m.State.MissionAddMenuMsg.MessageID != int64(msg.ID) {
		c.Bot().Delete(msg)
	}
	return nil
}
