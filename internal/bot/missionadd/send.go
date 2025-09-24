package missionadd

import (
	"fmt"
	"strings"

	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
)

// сбрасывает состояние
func (m *stateMission) SendMenu(c tele.Context, message string) error {
	mode := tele.ModeHTML
	keyBoard := m.L.Markup(c, "menuMissionAdd")
	if m.State.MissionAddMenuMsg == nil {
		// переходим к началу формирования сообщения с инлайнкнопками
		m.ClearKeyBoard(c)
		if msg, err := c.Bot().Send(c.Chat(), message, keyBoard, mode); err != nil {
			return fmt.Errorf("missadd:send %w", err)
		} else {
			m.State.MissionAddMenuMsg = &entity.StoredMessage{
				MessageID: int64(msg.ID),
				ChatID:    m.State.ChatId,
			}
		}
	} else {
		if msg, err := c.Bot().Edit(m.State.MissionAddMenuMsg, message, keyBoard, mode); err != nil {
			errDesc := err.Error()
			if !strings.Contains(errDesc, "message is not modified") {
				m.State.MissionAddMenuMsg = nil
				return fmt.Errorf("%w", err)
			}
			return nil
		} else {
			// после редактирования запоминаем заново
			m.State.MissionAddMenuMsg = &entity.StoredMessage{
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
		m.State.MissionAddDialogMsgs = append(m.State.MissionAddDialogMsgs, &entity.StoredMessage{
			MessageID: int64(msg.ID),
			ChatID:    m.State.ChatId,
		})
	}
	return nil
}
