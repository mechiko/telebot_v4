package userstate

import (
	"fmt"
	"strings"

	"github.com/mechiko/telebot_v4/internal/entity"
	"gopkg.in/telebot.v4"
)

func (m *stateMission) routerCallbackToMode(c telebot.Context) error {
	defer m.App.GetRecovery().RecoverLog("missadd:routerCallbackToMode")

	can_exit := m.checkAllStates(m.State.UserState)
	callback := c.Callback()
	if callback == nil {
		return fmt.Errorf("routerCallbackToMode callback is nil")
	}
	// callback data \f + unique + | + data
	data := callback.Data
	data, _ = strings.CutPrefix(data, "\f")
	dataStrings := strings.Split(data, "|")
	unique := ""
	command := ""
	if len(dataStrings) >= 2 {
		unique = dataStrings[0]
		command = dataStrings[1]
	} else {
		return fmt.Errorf("callback data wrong %s", callback.Data)
	}
	m.State.CallbackData = command

	switch unique {
	case "exitState":
		if can_exit {
			m.State.Mode = entity.ModeStart
			m.State.Stage = entity.StageEmpty
		}
	case "intro":
		m.State.Mode = entity.ModeEditUserStatesCallback
		m.State.Stage = entity.StageUserStateIntro
	case "exam1":
		m.State.Mode = entity.ModeEditUserStatesCallback
		m.State.Stage = entity.StageUserStateExam
	case "exam2":
		m.State.Mode = entity.ModeEditUserStatesCallback
		m.State.Stage = entity.StageUserStateExam
	case "exam3":
		m.State.Mode = entity.ModeEditUserStatesCallback
		m.State.Stage = entity.StageUserStateExam
	case "exam4":
		m.State.Mode = entity.ModeEditUserStatesCallback
		m.State.Stage = entity.StageUserStateExam
	case "exam5":
		m.State.Mode = entity.ModeEditUserStatesCallback
		m.State.Stage = entity.StageUserStateExam
	case "exam6":
		m.State.Mode = entity.ModeEditUserStatesCallback
		m.State.Stage = entity.StageUserStateExam
	case "exam7":
		m.State.Mode = entity.ModeEditUserStatesCallback
		m.State.Stage = entity.StageUserStateExam
	}
	return nil
}
