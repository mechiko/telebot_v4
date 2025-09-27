package examenended

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type examenended struct {
	Repo entity.Repo
}

func New(r entity.Repo) entity.ExamenEndeds {
	return &examenended{
		Repo: r,
	}
}
