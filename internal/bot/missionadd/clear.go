package missionadd

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
	"gopkg.in/telebot.v4"
)

func (m *stateMission) ClearAll(c telebot.Context) error {
	menuClear := &telebot.ReplyMarkup{RemoveKeyboard: true}
	if msg, err := c.Bot().Send(c.Sender(), "с", menuClear, telebot.ModeDefault); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		c.Bot().Delete(msg)
	}
	if m.State.MissionAddMenuMsg != nil {
		c.Bot().Delete(m.State.MissionAddMenuMsg)
		m.State.MissionAddMenuMsg = nil
	}
	if len(m.State.MissionAddDialogMsgs) > 0 {
		for _, m := range m.State.MissionAddDialogMsgs {
			if m != nil {
				c.Bot().Delete(m)
			}
		}
		m.State.MissionAddDialogMsgs = make([]*entity.StoredMessage, 0)
	}
	if len(m.State.MissionDialogMsgs) > 0 {
		for _, m := range m.State.MissionDialogMsgs {
			if m != nil {
				c.Bot().Delete(m)
			}
		}
		m.State.MissionDialogMsgs = make([]*entity.StoredMessage, 0)
	}
	return nil
}

func (m *stateMission) ClearDialogMessages(c telebot.Context) error {
	if len(m.State.MissionAddDialogMsgs) > 0 {
		for _, m := range m.State.MissionAddDialogMsgs {
			if m != nil {
				c.Bot().Delete(m)
			}
		}
		m.State.MissionAddDialogMsgs = make([]*entity.StoredMessage, 0)
	}
	return nil
}

func (m *stateMission) ClearKeyBoard(c telebot.Context) error {
	menuClear := &telebot.ReplyMarkup{RemoveKeyboard: true}
	if msg, err := c.Bot().Send(c.Sender(), "с", menuClear, telebot.ModeDefault); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		c.Bot().Delete(msg)
	}
	return nil
}
