package entity

import "gopkg.in/telebot.v4"

// состояние диалога пользователя с командировками
type StateMode interface {
	// DbMissions() (*MissionList, error)
	Mode(c telebot.Context) error
	// Clear(c telebot.Context) error
	// Send(c telebot.Context, message string, keyBoard *telebot.ReplyMarkup, mode string) error
}
