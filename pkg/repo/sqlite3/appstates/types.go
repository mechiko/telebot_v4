package appstates

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type appstates struct {
	Repo entity.Repo
}

func New(r entity.Repo) entity.AppStates {
	return &appstates{
		Repo: r,
	}
}
