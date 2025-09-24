package bot

import (
	"fmt"
	"strings"

	"github.com/mechiko/telebot_v4/internal/bot/missionadd"
	"github.com/mechiko/telebot_v4/internal/bot/userstate"
	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
)

func (b Bot) routerCallback(c tele.Context) error {
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		return fmt.Errorf("routerCallback state not found")
	}
	callback := c.Callback()
	if callback == nil {
		return fmt.Errorf("routerCallback callback is nil")
	}

	switch state.Mode {
	case entity.ModeClear:
	case entity.ModeStart:
		return b.modeStart(c)
	// case entity.ModeMission:
	case entity.ModeMissionAdd:
		return missionadd.New(b.App, state, b.L, b.routerMode).Mode(c)
	case entity.ModeMissionAddCallback:
		return missionadd.New(b.App, state, b.L, b.routerMode).Mode(c)
	case entity.ModeEditUserStates:
		return userstate.New(b.App, state, b.L, b.routerMode).Mode(c)
	case entity.ModeEditUserStatesCallback:
		return userstate.New(b.App, state, b.L, b.routerMode).Mode(c)
	}

	return nil
}

func (b Bot) routerCallbackToMode(c tele.Context) error {
	defer b.App.GetRecovery().RecoverLog("routerCallbackToMode")
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		b.App.GetLogger().Errorf("routerCallbackToMode state not found")
		return fmt.Errorf("routerCallbackToMode state not found")
	}
	can_exit := b.checkAllStates(state.UserState)
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
	state.CallbackData = command
	b.App.GetLogger().Infof("routerCallbackToMode callback uniq %s command %s", unique, command)
	switch unique {
	case "exitState":
		if can_exit {
			state.Mode = entity.ModeStart
			state.Stage = entity.StageEmpty
		}
	case "intro":
		state.Mode = entity.ModeEditUserStatesCallback
		state.Stage = entity.StageUserStateIntro
	case "exam1":
		state.Mode = entity.ModeEditUserStatesCallback
		state.Stage = entity.StageUserStateExam
	case "exam2":
		state.Mode = entity.ModeEditUserStatesCallback
		state.Stage = entity.StageUserStateExam
	case "exam3":
		state.Mode = entity.ModeEditUserStatesCallback
		state.Stage = entity.StageUserStateExam
	case "exam4":
		state.Mode = entity.ModeEditUserStatesCallback
		state.Stage = entity.StageUserStateExam
	case "exam5":
		state.Mode = entity.ModeEditUserStatesCallback
		state.Stage = entity.StageUserStateExam
	case "exam6":
		state.Mode = entity.ModeEditUserStatesCallback
		state.Stage = entity.StageUserStateExam
	case "exam7":
		state.Mode = entity.ModeEditUserStatesCallback
		state.Stage = entity.StageUserStateExam
	}
	return nil
}
