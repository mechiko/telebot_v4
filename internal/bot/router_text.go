package bot

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/bot/missionadd"
	"github.com/mechiko/telebot_v4/internal/bot/userstate"
	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
)

func (b Bot) Router(c tele.Context) error {
	defer b.App.GetRecovery().RecoverLog("routerText")
	var state *entity.StateDialog
	// cmds := b.L.Commands()
	// if cmds != nil {
	// 	if err := b.SetCommands(cmds); err != nil {
	// 		b.App.GetLogger().Errorf("bot:router set commands error %s", err.Error())
	// 	}
	// }

	if state = b.States.GetState(c.Chat().ID); state == nil {
		b.App.GetLogger().Errorf("routerText state not found")
		return fmt.Errorf("routerText state not found")
	}
	switch state.Mode {
	case entity.ModeClear:
		return b.modeClear(c)
	case entity.ModeStart:
		return b.modeStart(c)
	// case entity.ModeMission:
	// 	return mission.New(b.App, state, b.L, b.routerMode).Mode(c)
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

func (b Bot) routerTextCommandToMode(c tele.Context) error {
	defer b.App.GetRecovery().RecoverLog("routerTextCommandToMode")
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		b.App.GetLogger().Errorf("routerTextCommandToMode state not found")
		return fmt.Errorf("routerTextCommandToMode state not found")
	}
	switch state.Text {
	case `/start`:
		c.Delete()
		state.Text = ""
		state.Mode = entity.ModeStart
		state.Stage = entity.StageEmpty
	case `/clear`:
		c.Delete()
		state.Text = ""
		state.Mode = entity.ModeClear
		state.Stage = entity.StageEmpty
	case `/mission`:
		c.Delete()
		state.Text = ""
		state.Mode = entity.ModeMissionAdd
		state.Stage = entity.StageEmpty
	case `/examens`:
		c.Delete()
		state.Text = ""
		state.Mode = entity.ModeEditUserStates
		state.Stage = entity.StageEmpty
	}
	return nil
}

func (b Bot) routerMode(c tele.Context, state *entity.StateDialog) error {
	switch state.Mode {
	case entity.ModeClear:
		return b.modeClear(c)
	case entity.ModeStart:
		return b.modeStart(c)
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
