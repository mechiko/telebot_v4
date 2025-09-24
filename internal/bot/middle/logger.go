package middle

import (
	"encoding/json"
	"errors"

	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
	"modernc.org/sqlite"
)

// Logger returns a middleware that logs incoming updates.
// при ошибках обработка не прерывается
func Logger(app entity.Application) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			defer app.GetRecovery().RecoverLog("middleware:logger ")
			update := c.Update()
			updateBot := &entity.Update{
				ID:      int64(update.ID),
				Message: c.Text(),
			}
			if chat := c.Chat(); chat != nil {
				updateBot.Chat = chat.ID
			}
			if sender := c.Sender(); sender != nil {
				updateBot.Sender = sender.ID
			}
			if recepient := c.Recipient(); recepient != nil {
				recepientData, _ := json.Marshal(recepient)
				updateBot.Recepient = string(recepientData)
			}
			// updateData, _ := json.Marshal(update)
			// updateBot.Update = string(updateData)
			if err := app.GetRepo().GetUpdates().Insert(updateBot); err != nil {
				e := errors.Unwrap(err)
				if er, ok := e.(*sqlite.Error); ok {
					// 1555 - code:1555 msg:constraint failed: UNIQUE constraint failed: chats.id (1555)
					if er.Code() != 1555 {
						app.GetLogger().Errorf("middleware:logger db insert %s", err.Error())
					}
				} else {
					app.GetLogger().Errorf("middleware:logger db insert %s", err.Error())
				}
			}
			return next(c)
		}
	}
}
