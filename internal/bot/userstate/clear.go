package userstate

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
	if m.State.UserStateMenuMsg != nil {
		c.Bot().Delete(m.State.UserStateMenuMsg)
		m.State.UserStateMenuMsg = nil
	}
	if len(m.State.UserStateDialogMsgs) > 0 {
		for _, m := range m.State.UserStateDialogMsgs {
			if m != nil {
				c.Bot().Delete(m)
			}
		}
		m.State.UserStateDialogMsgs = make([]*entity.StoredMessage, 0)
	}
	return nil
}

func (m *stateMission) ClearDialogMessages(c telebot.Context) error {
	if len(m.State.UserStateDialogMsgs) > 0 {
		for _, m := range m.State.UserStateDialogMsgs {
			if m != nil {
				c.Bot().Delete(m)
			}
		}
		m.State.UserStateDialogMsgs = make([]*entity.StoredMessage, 0)
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
