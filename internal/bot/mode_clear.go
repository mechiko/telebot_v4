package bot

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/bot/missionadd"
	"github.com/mechiko/telebot_v4/internal/bot/userstate"
	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
)

// удаление состояния диалога пользователя
// для зависших состояний
func (b Bot) modeClear(c tele.Context) error {
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		b.App.GetLogger().Errorf("onClear state not found")
		return fmt.Errorf("onClear state not found")
	}
	b.routerTextCommandToMode(c)
	switch state.Mode {
	case entity.ModeStart:
		return b.modeStart(c)
	// case entity.ModeMission:
	// 	return mission.New(b.App, state, b.L, b.routerMode).Mode(c)
	case entity.ModeMissionAdd:
		return missionadd.New(b.App, state, b.L, b.routerMode).Mode(c)
	case entity.ModeEditUserStates:
		return userstate.New(b.App, state, b.L, b.routerMode).Mode(c)
	case entity.ModeEditUserStatesCallback:
		return userstate.New(b.App, state, b.L, b.routerMode).Mode(c)
	}
	b.clearAllStateMessage(c)
	return nil
}
