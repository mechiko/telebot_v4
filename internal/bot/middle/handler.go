package middle

import (
	"fmt"
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/usecase"
	tele "gopkg.in/telebot.v4"
)

// обработчик всех входящих событий бота
// формируем состояние диалога пользователя или новое или из сохраненного
// проверяем условия принудительного состояни диалога экзаменов
//
// какие то сообщения события прерываются возвратом ошибки или nil
// return nil - обработка прерывается
// return next(c) - обработка продолжается
func Handler(app entity.Application) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			defer app.GetRecovery().RecoverLog("middleware:handler ")
			var state *entity.StateDialog
			app.GetLogger().Debugf("middleware:handler start")

			// update := c.Update()
			chat := c.Chat()

			// получаем или создаем состояние диалога
			// берем состояние или создаем его и назначаем режим
			if state = app.GetBot().GetStates().GetState(c.Chat().ID); state == nil {
				// если нет состояния для пользователя создаем его
				app.GetLogger().Infof("middleware new state create user found %v %s", chat.ID, chat.Username)
				if user, err := app.GetRepo().GetTelebotUsers().GetById(chat.ID); user != nil {
					state = &entity.StateDialog{
						ChatId:               chat.ID,
						User:                 user,
						Text:                 c.Text(),
						ErrorText:            "",
						Mode:                 entity.ModeStart,
						Mission:              &entity.Mission{},
						Stage:                entity.StageEmpty,
						UserStateDialogMsgs:  make([]*entity.StoredMessage, 0),
						MissionDialogMsgs:    make([]*entity.StoredMessage, 0),
						MissionAddDialogMsgs: make([]*entity.StoredMessage, 0),
					}
					app.GetBot().GetStates().In <- state
					// тут бы надо дождаться как состояние проскочит...
					time.Sleep(10 * time.Millisecond)
				} else {
					return fmt.Errorf("middleware:handler GetUsers %w", err)
				}
				if userState, err := usecase.New(app).GetUserStatesMap(chat.ID); err != nil {
					return fmt.Errorf("middleware:handler GetUserStatesMap %w", err)
				} else {
					// по задумке тут всегда будет мап
					state.UserState = userState
				}
				// по наличию заполняем активную командировку
				if mission, err := usecase.New(app).GetUserActiveMission(state.ChatId); err != nil {
					app.GetLogger().Errorf("middleware:handler get active mission error %s", err.Error())
				} else {
					state.Mission = mission
				}
				// при необходимости заполнить экзамены делаем это принудительно
				if !checkAllStates(state.UserState, app) {
					state.Mode = entity.ModeEditUserStates
					state.Text = ""
					return next(c)
				}
				//  else {
				// 	// если экзамены заполнены то очищаем клавиатуру
				// 	// очищаем клавиатуру при создании нового состояни для пользователя
				// 	// сеанса не было еще и будем сбрасывать периодически пока не решил
				// 	clearKeyboard(c)
				// }
			} else {
				// state != nil
				if user, err := app.GetRepo().GetTelebotUsers().GetById(chat.ID); err != nil {
					return fmt.Errorf("middleware:handler GetUser %w", err)
				} else {
					state.User = user
				}
				state.Text = c.Text()
				state.ErrorText = ""
				if userState, err := usecase.New(app).GetUserStatesMap(chat.ID); err != nil {
					return fmt.Errorf("middleware:GetUserStatesMap %w", err)
				} else {
					// по задумке тут всегда будет мап
					state.UserState = userState
				}
				// по наличию заполняем активную командировку
				if mission, err := usecase.New(app).GetUserActiveMission(state.ChatId); err != nil {
					app.GetLogger().Errorf("middleware:handler get active mission error %s", err.Error())
				} else {
					state.Mission = mission
				}
				if state.Mode != entity.ModeEditUserStates && state.Mode != entity.ModeEditUserStatesCallback {
					// для всех режимов не равных редактору будем проверять состояние и если ошибочно
					// принудительно ставим режим и запускаем редактор
					if !checkAllStates(state.UserState, app) {
						state.Mode = entity.ModeEditUserStates
						state.Text = ""
					}
				}

			}
			app.GetLogger().Debugf("handler state %s %s", state.Mode, state.Stage)
			return next(c)
		}
	}
}
