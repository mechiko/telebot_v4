package missionadd

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/template"
	tele "gopkg.in/telebot.v4"
)

func (m *stateMission) Mode(c tele.Context) error {
	defer m.App.GetRecovery().RecoverLog("missadd:mode")

	m.ClearDialogMessages(c)

	isCallback := false
	if c.Callback() != nil {
		isCallback = true
	}
	// преобразуем тект в режим по команде если экзамены все в норме
	// и можно выйти из режима
	if isCallback {
		// если прилетело из старого сообщения то удалим его
		// сам коллбэк кнопка обрабатывается как обычно, просто мусор уберем из чата
		m.checkMsgId(c)
		m.routerCallbackToMode(c)
		// после обработки команд если изменится режим то обрабатываем
		switch m.State.Mode {
		case entity.ModeStart:
			m.ClearAll(c)
			c.Delete()
			return m.R(c, m.State)
		case entity.ModeMissionAdd:
		case entity.ModeMissionAddCallback:
			switch m.State.Stage {
			case entity.StageEmpty:
			case entity.StageMissionStart:
				message := "Введите дату в формате DD.MM.YYYY для " + m.State.CallbackData
				return m.sendMessage(c, message)
			case entity.StageMissionEnd:
				message := "Введите дату в формате DD.MM.YYYY для " + m.State.CallbackData
				return m.sendMessage(c, message)
			case entity.StageMissionPlace:
				m.State.Mission.Place = m.State.CallbackData
				if err := m.save(c); err != nil {
					m.App.GetLogger().Errorf("missadd:save %s", err.Error())
				}
				m.checkMission(c)
				m.State.Mode = entity.ModeMissionAdd
				m.State.Stage = entity.StageEmpty
			case entity.StageMissionExit:
				c.Delete()
				m.ClearDialogMessages(c)
				m.State.Mode = entity.ModeStart
				m.State.Stage = entity.StageEmpty
				return m.R(c, m.State)
			case entity.StageMissionDel:
				if err := m.del(c); err != nil {
					m.App.GetLogger().Errorf("missadd:del %s", err.Error())
				}
				m.State.Mission = &entity.Mission{}
				// m.checkMission(c)
				m.State.Mode = entity.ModeMissionAdd
				m.State.Stage = entity.StageEmpty
			}
		}
	} else {
		// если пришел текст
		m.Router(c)
		// команда может изменить текущий режим
		switch m.State.Mode {
		case entity.ModeClear:
			m.ClearAll(c)
			c.Delete()
			return m.R(c, m.State)
		case entity.ModeStart:
			m.ClearAll(c)
			c.Delete()
			return m.R(c, m.State)
		// case entity.ModeMission:
		// 	m.ClearAll(c)
		// 	c.Delete()
		// 	return m.R(c, m.State)
		case entity.ModeEditUserStates:
			m.ClearAll(c)
			c.Delete()
			return m.R(c, m.State)
		case entity.ModeMissionAdd:
		case entity.ModeMissionAddCallback:
			switch m.State.Stage {
			case entity.StageMissionStart:
				m.State.Mission.Start = m.State.Text
				if err := m.save(c); err != nil {
					m.App.GetLogger().Errorf("missadd:save %s", err.Error())
				}
				m.addContextMessage(c)
				m.checkMission(c)
				m.State.Mode = entity.ModeMissionAdd
				m.State.Stage = entity.StageEmpty
			case entity.StageMissionEnd:
				m.State.Mission.End = m.State.Text
				if err := m.save(c); err != nil {
					m.App.GetLogger().Errorf("missadd:save %s", err.Error())
				}
				m.addContextMessage(c)
				m.checkMission(c)
				m.State.Mode = entity.ModeMissionAdd
				m.State.Stage = entity.StageEmpty
			}
		}
		// если сообщение пришло и не обработано ни в одном режиме удалим
		c.Delete()
	}

	// m.checkMission(c)
	if m.State.ErrorText != "" {
		m.State.Mission.ErrorTxt = fmt.Sprintf("%s\n%s", m.State.ErrorText, m.State.Mission.ErrorTxt)
	}
	missionHtml := ""
	if msgText, err := template.NewTemplate(m.App, "").MissionMessage(m.State.Mission); err != nil {
		m.App.GetLogger().Errorf("missadd:template %s", err.Error())
	} else {
		missionHtml = msgText
	}

	return m.SendMenu(c, missionHtml)
}
