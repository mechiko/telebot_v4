package keystates

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type keystates struct {
	Repo entity.Repo
}

func New(r entity.Repo) entity.KeyStates {
	return &keystates{
		Repo: r,
	}
}
