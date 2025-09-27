package bot

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
)

func (b Bot) clearAllStateMessage(c tele.Context) {
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		b.App.GetLogger().Errorf("routerText state not found")
	}
	if state.UserStateMenuMsg != nil {
		b.Delete(state.UserStateMenuMsg)
		state.UserStateMenuMsg = nil
	}
	if state.MissionAddMenuMsg != nil {
		b.Delete(state.MissionAddMenuMsg)
		state.MissionAddMenuMsg = nil
	}
	if len(state.UserStateDialogMsgs) > 0 {
		for _, m := range state.UserStateDialogMsgs {
			if m != nil {
				b.Delete(m)
			}
		}
		state.UserStateDialogMsgs = make([]*entity.StoredMessage, 0)
	}
	if len(state.MissionDialogMsgs) > 0 {
		for _, m := range state.MissionDialogMsgs {
			if m != nil {
				b.Delete(m)
			}
		}
		state.MissionDialogMsgs = make([]*entity.StoredMessage, 0)
	}

}

func (b Bot) sendDialogMessage(c tele.Context, message string, keyBoard *tele.ReplyMarkup, mode string) error {
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		return fmt.Errorf("sendDialogMessage state not found")
	}
	if msg, err := b.Send(c.Sender(), message, keyBoard, mode); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		state.MissionDialogMsgs = append(state.MissionDialogMsgs, &entity.StoredMessage{
			MessageID: int64(msg.ID),
			ChatID:    state.ChatId,
		})
	}
	return nil
}

func (b Bot) clearDialogsMessage(c tele.Context) {
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		b.App.GetLogger().Errorf("routerText state not found")
	}
	if len(state.MissionDialogMsgs) > 0 {
		for _, m := range state.MissionDialogMsgs {
			if m != nil {
				b.Delete(m)
			}
		}
		state.MissionDialogMsgs = make([]*entity.StoredMessage, 0)
	}

}

func (b Bot) clearUserStateDialogMessages(c tele.Context) {
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		b.App.GetLogger().Errorf("routerText state not found")
	}
	if len(state.UserStateDialogMsgs) > 0 {
		for _, m := range state.UserStateDialogMsgs {
			if m != nil {
				b.Delete(m)
			}
		}
		state.UserStateDialogMsgs = make([]*entity.StoredMessage, 0)
	}
}

func (b Bot) sendMessage(c tele.Context, message string) error {
	var state *entity.StateDialog
	if state = b.States.GetState(c.Chat().ID); state == nil {
		return fmt.Errorf("sendMessage state not found")
	}
	if msg, err := b.Send(c.Sender(), message); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		state.MissionDialogMsgs = append(state.MissionDialogMsgs, &entity.StoredMessage{
			MessageID: int64(msg.ID),
			ChatID:    state.ChatId,
		})
	}
	return nil
}
