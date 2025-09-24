package userstate

import (
	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/template"
	"github.com/mechiko/telebot_v4/internal/usecase"
	"gopkg.in/telebot.v4"
)

func (m *stateMission) Mode(c telebot.Context) error {
	defer m.App.GetRecovery().RecoverLog("userstate:mode")

	m.ClearDialogMessages(c)
	if userState, err := usecase.New(m.App).GetUserStatesMap(c.Chat().ID); err != nil {
		m.App.GetLogger().Errorf("userstate:mode %s", err.Error())
	} else {
		m.State.UserState = userState
	}
	tempUserStates := m.copyUserStates(m.State.UserState)
	// can_exit := m.checkAllStates(tempUserStates)

	isCallback := false
	if c.Callback() != nil {
		isCallback = true
	}
	// преобразуем тект в режим по команде если экзамены все в норме
	// и можно выйти из режима
	// if can_exit {
	if isCallback {
		// если прилетело из старого сообщения то удалим его
		// сам коллбэк кнопка обрабатывается как обычно, просто мусор уберем из чата
		m.checkMsgId(c)
		m.routerCallbackToMode(c)
		// после обработки команд если изменится режим то обрабатываем
		switch m.State.Mode {
		case entity.ModeStart:
			c.Delete()
			m.ClearAll(c)
			return m.R(c, m.State)
		case entity.ModeEditUserStates:
		case entity.ModeEditUserStatesCallback:
			switch m.State.Stage {
			case entity.StageEmpty:
			case entity.StageUserStateExam:
				message := "Введите дату в формате DD.MM.YYYY для " + m.State.CallbackData
				m.sendMessage(c, message)
			case entity.StageUserStateIntro:
				message := "Введите значение для " + m.State.CallbackData
				m.sendMessage(c, message)
			}
		}
		// callback отработали и выходим
		return nil
	} else {
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
		case entity.ModeEditUserStates:
		case entity.ModeEditUserStatesCallback:
			switch m.State.Stage {
			case entity.StageUserStateIntro:
				if err := m.App.GetRepo().GetUserStates().Set(m.State.UserState.UserId, m.State.CallbackData, m.State.Text); err != nil {
					m.App.GetLogger().Errorf("state save error %s", err.Error())
				}
				m.State.UserState.Intro[m.State.CallbackData] = m.State.Text
				m.addContextMessage(c)
				tempUserStates = m.copyUserStates(m.State.UserState)
			case entity.StageUserStateExam:
				if err := m.App.GetRepo().GetUserStates().Set(m.State.UserState.UserId, m.State.CallbackData, m.State.Text); err != nil {
					m.App.GetLogger().Errorf("state save error %s", err.Error())
				}
				m.State.UserState.Examens[m.State.CallbackData] = m.State.Text
				// удаляем текущее сообщение
				m.addContextMessage(c)
				tempUserStates = m.copyUserStates(m.State.UserState)
			}
		}
		// если сообщение пришло и не обработано ни в одном режиме удалим
		c.Delete()
	}
	// зачищаем сообщения диалогов режима
	m.ClearDialogMessages(c)
	// фиксируем режим и стадию в начало редактора состояний
	m.State.Mode = entity.ModeEditUserStates
	m.State.Stage = entity.StageEmpty
	m.State.UserStateDialogMsgs = make([]*entity.StoredMessage, 0)
	m.checkAllStates(tempUserStates)
	userStatesHtml := ""

	if msgText, err := template.NewTemplate(m.App, "").UserStates(tempUserStates); err != nil {
		m.App.GetLogger().Errorf("userstate:template %s", err.Error())
	} else {
		userStatesHtml = msgText
	}
	return m.SendMenu(c, userStatesHtml)
}
