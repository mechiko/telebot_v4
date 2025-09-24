package userstates

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type userstates struct {
	Repo entity.Repo
}

func New(r entity.Repo) entity.UserStates {
	return &userstates{
		Repo: r,
	}
}
