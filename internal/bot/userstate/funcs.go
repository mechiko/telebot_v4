package userstate

import (
	"github.com/mechiko/telebot_v4/internal/entity"
	"gopkg.in/telebot.v4"
)

func (m *stateMission) addContextMessage(c telebot.Context) error {
	if msg := c.Message(); msg != nil {
		m.State.UserStateDialogMsgs = append(m.State.UserStateDialogMsgs, &entity.StoredMessage{
			MessageID: int64(msg.ID),
			ChatID:    m.State.ChatId,
		})
	}
	return nil
}
