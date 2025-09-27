package bot

import (
	"fmt"
	"time"

	"github.com/mechiko/telebot_v4/internal/bot/missionadd"
	"github.com/mechiko/telebot_v4/internal/bot/userstate"
	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/usecase"
	"gopkg.in/telebot.v4"
)

// сюда приходят команда старт и все команды меню плюс текст отправлемый просто случайно в этом режиме
func (b Bot) modeStart(c telebot.Context) error {
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		return fmt.Errorf("bot:start state not found")
	}
	b.routerTextCommandToMode(c)
	state.Text = ""
	switch state.Mode {
	case entity.ModeClear:
		return b.modeClear(c)
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
	// message := b.L.Text(c, "start", c.Chat())
	state.MissionDialogMsgs = make([]*entity.StoredMessage, 0)
	// b.clearKeyboard(c, message)
	if b.checkMissionInfo(c) {
		if str, err := usecase.New(b.App).GetUserInfo(state); err != nil {
			return fmt.Errorf("%w", err)
		} else {
			b.sendMessage(c, str)
		}
	} else {
		b.sendMessage(c, "командировка отсутствует или данные некорректны")
	}
	return nil
}

// проверяем миссию в состоянии и прописываем ошибку в mission.ErrorText
func (b *Bot) checkMissionInfo(c telebot.Context) bool {
	defer b.App.GetRecovery().RecoverLog("checkOnSave")
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		return false
	}
	var startDate, endDate time.Time
	var err error
	result := true
	if state.Mission == nil {
		state.Mission = &entity.Mission{}
	}
	state.Mission.ErrorTxt = ""
	l := b.App.GetConfiguration().Layouts.TimeLayoutDay
	// nowDate := time.Now().Local()
	if state.Mission.Place == "" {
		result = false
		state.Mission.ErrorTxt += "место командировки не должно быть пустым \n"
	}
	if startDate, err = time.Parse(l, state.Mission.Start); err != nil {
		result = false
		state.Mission.ErrorTxt += "ошибка формата даты начала \n"
	}
	if endDate, err = time.Parse(l, state.Mission.End); err != nil {
		result = false
		state.Mission.ErrorTxt += "ошибка формата даты окончания"
	}
	if !result {
		return false
	}
	// if startDate.Before(nowDate) {
	// 	result = false
	// 	state.Mission.ErrorTxt = "дата начала не должна быть меньше текущей даты \n"
	// }
	if endDate.Before(startDate) {
		result = false
		state.Mission.ErrorTxt += "дата окончания меньше даты начала "
	}
	return result
}
