package template

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type templateString struct {
	app    entity.Application
	layout string
}

var _ entity.BotTemplate = &templateString{}

func NewTemplate(app entity.Application, layout string) entity.BotTemplate {
	if layout == "" {
		layout = app.GetConfiguration().Layouts.TimeLayoutDay
	}
	return &templateString{
		app:    app,
		layout: layout,
	}
}
