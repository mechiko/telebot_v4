package userstate

import (
	"github.com/mechiko/telebot_v4/internal/entity"
	"gopkg.in/telebot.v4"
	"gopkg.in/telebot.v4/layout"
)

type RouteFunc func(c telebot.Context, state *entity.StateDialog) error

// создаем прокси на структуру состояния диалога, чтобы организовать в новом модуле все действия
// и избежать перекрестных ссылок когда импортировать будем тип StateDialog то там то там
type stateMission struct {
	App   entity.Application
	State *entity.StateDialog
	// States *entity.States
	// Context telebot.Context
	// Bot     telebot.Bot
	L *layout.Layout
	R RouteFunc
}

func New(app entity.Application, state *entity.StateDialog, l *layout.Layout, f RouteFunc) entity.StateMode {
	return &stateMission{
		App:   app,
		State: state,
		// Bot:     bot,
		// Context: c,
		L: l,
		R: f,
	}
}
