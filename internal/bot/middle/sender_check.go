package middle

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
)

// обработчик всех входящих событий бота
// какие то сообщения события прерываются возвратом ошибки или nil
// return nil - обработка прерывается
// return next(c) - обработка продолжается
func SenderChecker(app entity.Application) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			defer app.GetRecovery().RecoverLog("middleware:sender ")
			chat := c.Chat()
			// если отправитель есть в БД то только для приватных чатов
			if sender := c.Sender(); sender != nil {
				if checkSender(c.Sender().ID, app) {
					if chat.Type == "private" {
						app.GetLogger().Infof("middleware:sender user found %v %s", chat.ID, chat.Username)
					} else {
						// chat.Type != "private"
						// не обрабатываем
						return fmt.Errorf("middleware:sender reject chat %v", chat.ID)
					}
				} else {
					// sender == nil
					// не обрабатываем чужого
					app.GetLogger().Errorf("middleware reject sender %v", sender.Username)
					return fmt.Errorf("reject")
				}
			} else {
				// c.Sender == nil if user is not presented
				return fmt.Errorf("middleware:sender reject sender not present")
			}
			return next(c)
		}
	}
}
