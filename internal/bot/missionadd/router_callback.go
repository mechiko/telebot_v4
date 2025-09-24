package missionadd

import (
	"fmt"
	"strings"

	"github.com/mechiko/telebot_v4/internal/entity"
	"gopkg.in/telebot.v4"
)

func (m *stateMission) routerCallbackToMode(c telebot.Context) error {
	defer m.App.GetRecovery().RecoverLog("missadd:routerCallbackToMode")

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
		m.State.Mode = entity.ModeMissionAddCallback
		m.State.Stage = entity.StageMissionExit
	case "mistart":
		m.State.Mode = entity.ModeMissionAddCallback
		m.State.Stage = entity.StageMissionStart
	case "miend":
		m.State.Mode = entity.ModeMissionAddCallback
		m.State.Stage = entity.StageMissionEnd
	case "miplace1":
		m.State.Mode = entity.ModeMissionAddCallback
		m.State.Stage = entity.StageMissionPlace
	case "miplace2":
		m.State.Mode = entity.ModeMissionAddCallback
		m.State.Stage = entity.StageMissionPlace
	case "miplace3":
		m.State.Mode = entity.ModeMissionAddCallback
		m.State.Stage = entity.StageMissionPlace
	case "miplace4":
		m.State.Mode = entity.ModeMissionAddCallback
		m.State.Stage = entity.StageMissionPlace
	case "misave":
		m.State.Mode = entity.ModeMissionAddCallback
		m.State.Stage = entity.StageMissionSave
	case "miclear":
		m.State.Mode = entity.ModeMissionAddCallback
		m.State.Stage = entity.StageMissionDel
	}
	return nil
}
