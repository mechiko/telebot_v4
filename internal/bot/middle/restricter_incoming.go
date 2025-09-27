package middle

import (
	"fmt"
	"strings"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/samber/lo"
	"golang.org/x/exp/slices"
	tele "gopkg.in/telebot.v4"
)

// обработчик всех входящих событий бота
// какие то сообщения события прерываются возвратом ошибки или nil
// return nil - обработка прерывается
// return next(c) - обработка продолжается
func RestricterIn(app entity.Application) tele.MiddlewareFunc {
	return func(next tele.HandlerFunc) tele.HandlerFunc {
		return func(c tele.Context) error {
			defer app.GetRecovery().RecoverLog("middleware:handler ")
			var state *entity.StateDialog
			app.GetLogger().Debugf("middleware:RestricterIn start")

			message := c.Update().Message
			callback := c.Update().Callback

			// получаем состояние диалога
			// оно уже есть создается на этапе Handler если его нет в памяти бота
			if state = app.GetBot().GetStates().GetState(c.Chat().ID); state == nil {
				return fmt.Errorf("middleware:RestricterIn state is nil")
			}
			if !checkAllStates(state.UserState, app) {
				// если экзамены не заполнены как надо
				// запрещены любые команды бота
				if strings.HasPrefix(state.Text, `/`) {
					state.Text = ""
				}
				if message != nil {
					switch state.Mode {
					case entity.ModeEditUserStates:
						return next(c)
					case entity.ModeEditUserStatesCallback:
						return next(c)
					}
				}
				if callback != nil {
					switch state.Mode {
					case entity.ModeEditUserStates:
						return next(c)
					case entity.ModeEditUserStatesCallback:
						return next(c)
					}
				}
				return fmt.Errorf("экзамены не заполнены все кроме редактора запрещено")
			}
			if message != nil {
				// text := message.Text
				// для состояния коллбэк и стадий кроме интро и мыла игнорируем текстовые сообщения
				// через выход из обработки
				switch state.Mode {
				case entity.ModeClear:
					// enabledTexts := []string{`/start`, `/clear`, `/mission`, `/examens`}
					// if slices.Contains(enabledTexts, c.Text()) {
					// 	return next(c)
					// }
					return next(c)
				case entity.ModeStart:
					// // stage не важно кроме старт и очистка перечисляем допустимые
					// enabledTexts := []string{`/start`, `/clear`, `/mission`, `/examens`}
					// if slices.Contains(enabledTexts, c.Text()) {
					// 	return next(c)
					// }
					return next(c)
				// case entity.ModeMission:
				// 	switch state.Stage {
				// 	case entity.StageEmpty:
				// 		enabledTexts := []string{`назад`, `Добавить`, `Список`, `Удалить`, `Редактировать`}
				// 		if slices.Contains(enabledTexts, c.Text()) {
				// 			return next(c)
				// 		}
				// 	case entity.StageMissionDel:
				// 		return next(c)
				// 	case entity.StageMissionEdit:
				// 		return next(c)
				// 	}
				case entity.ModeMissionAdd:
					return next(c)
				case entity.ModeMissionAddCallback:
					return next(c)
				case entity.ModeEditUserStates:
					return next(c)
				case entity.ModeEditUserStatesCallback:
					return next(c)
				}
			}
			if callback != nil {
				// callback принимает только когда ждем после отправки меню
				// ожидание callback происходит при выводе основного меню и режиме entity.ModeEditUserStates
				// коллбэки от старых сообщений когда режим не редактора будет косвенно удалять эти сообщения из истории
				// есть ньюанс что если старое и новое висят то не важно какой коллбэк прилетит он будет для нового обрабатываться
				// можно попробовать засечь этот момент проверкой msg.ID в обработке и если он не такой как в сохраненном удалять его
				// TODO
				switch state.Mode {
				case entity.ModeEditUserStates:
					return next(c)
				case entity.ModeEditUserStatesCallback:
					return next(c)
				case entity.ModeMissionAdd:
					return next(c)
				case entity.ModeMissionAddCallback:
					return next(c)
				}
			}
			// все что не разрешено запрещено
			enabledTexts := []string{`/start`, `/clear`, `/mission`, `/examens`}
			if slices.Contains(enabledTexts, c.Text()) {
				return next(c)
			}
			// удаляем сообщение если не обрабатываем
			// при этом если это будет колбэк из старых инлайн кнопок то удалится это сообщение тоже
			c.Delete()
			// app.GetLogger().Debugf("middle:restricter reject state[%v] stage[%v] text[%s] callback[%v]", state.Mode, state.Stage, c.Text(), callback)
			// если режектим что то .. то переведем режим в старт наверное...
			app.GetLogger().Errorf("middle:restricter reject state[%v] stage[%v] text[%s] callback[%v]", state.Mode, state.Stage, lo.Substring(state.Text, 0, 10), callback)
			// state.Stage = entity.StageEmpty
			state.Text = ""
			app.GetLogger().Errorf("middle:restricter changTo state[%v] stage[%v] text[%s] callback[%v]", state.Mode, state.Stage, lo.Substring(state.Text, 0, 10), callback)
			return next(c)
		}
	}
}
