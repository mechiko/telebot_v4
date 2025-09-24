package userstate

import (
	"fmt"
	"strings"

	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
)

func (m *stateMission) SendMenu(c tele.Context, message string) error {
	mode := tele.ModeHTML
	keyBoard := m.L.Markup(c, "menuEditUserStates")
	if m.State.UserStateMenuMsg == nil {
		// переходим к началу формирования сообщения с инлайнкнопками
		if msg, err := c.Bot().Send(c.Chat(), message, keyBoard, mode); err != nil {
			return fmt.Errorf("userstate:sendmenu %w", err)
		} else {
			m.State.UserStateMenuMsg = &entity.StoredMessage{
				MessageID: int64(msg.ID),
				ChatID:    m.State.ChatId,
			}
		}
	} else {
		if msg, err := c.Bot().Edit(m.State.UserStateMenuMsg, message, keyBoard, mode); err != nil {
			errDesc := err.Error()
			if !strings.Contains(errDesc, "message is not modified") {
				m.State.UserStateMenuMsg = nil
				return fmt.Errorf("%w", err)
			}
			return nil
		} else {
			// после редактирования запоминаем заново
			m.State.UserStateMenuMsg = &entity.StoredMessage{
				MessageID: int64(msg.ID),
				ChatID:    m.State.ChatId,
			}
		}
	}
	return nil
}

func (m *stateMission) sendMessage(c tele.Context, message string) error {
	if msg, err := c.Bot().Send(c.Sender(), message); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		m.State.UserStateDialogMsgs = append(m.State.UserStateDialogMsgs, &entity.StoredMessage{
			MessageID: int64(msg.ID),
			ChatID:    m.State.ChatId,
		})
	}
	return nil
}
