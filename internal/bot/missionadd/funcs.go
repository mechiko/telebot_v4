package missionadd

import (
	"github.com/mechiko/telebot_v4/internal/entity"
	"gopkg.in/telebot.v4"
)

func (m *stateMission) addContextMessage(c telebot.Context) error {
	if msg := c.Message(); msg != nil {
		m.State.MissionAddDialogMsgs = append(m.State.MissionAddDialogMsgs, &entity.StoredMessage{
			MessageID: int64(msg.ID),
			ChatID:    m.State.ChatId,
		})
	}
	return nil
}
