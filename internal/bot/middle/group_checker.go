package middle

import (
	"strings"

	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
)

// обработчик всех входящих событий бота
// return nil - обработка прерывается
// return next(c) - обработка продолжается
func GroupChecker(app entity.Application) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			defer app.GetRecovery().RecoverLog("middleware:groupChecker ")
			// app.GetLogger().Debugf("middleware:groupChecker start")
			chat := c.Chat()
			// обработка сообщения в группе для прописки пользователя в таблицу доступа
			if chat.ID == app.GetConfiguration().AdminGroupId { // группа андрея начальства
				// if c.Text() == "/start@Meshiko_bot" {
				if strings.HasPrefix(c.Text(), "/start@ctai_bot") {
					insertSender(c, app)
					insertChat(c, app)
					if _, err := c.Bot().Send(c.Sender(), "Вы авторизованы"); err != nil {
						app.GetLogger().Errorf("middleware:groupChecker send error %s", err.Error())
					}
					return nil
				}
			}
			// if chat.ID == app.GetConfiguration().GroupId { // группа пока нет
			// 	if c.Text() == "/start@Meshiko_bot" {
			// if strings.HasPrefix(c.Text(), "/start@ctai_bot") {
			// 		insertSender(c, app)
			// 		insertChat(c, app)
			// 		if _, err := c.Bot().Send(c.Sender(), "вас услышали"); err != nil {
			// 			app.GetLogger().Errorf("middleware:groupChecker send error %s", err.Error())
			// 		}
			// 		return nil
			// 	}
			// }
			if chat.ID == app.GetConfiguration().TestGroupId { // группа бота
				// if c.Text() == "/start@Meshiko_bot" {
				if strings.HasPrefix(c.Text(), "/start@ctai_bot") {
					insertSender(c, app)
					insertChat(c, app)
					if _, err := c.Bot().Send(c.Sender(), "Вы авторизованы"); err != nil {
						app.GetLogger().Errorf("middleware:groupChecker send error %s", err.Error())
					}
					return nil
				}
			}
			return next(c)
		}
	}
}
