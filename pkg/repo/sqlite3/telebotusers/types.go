package telebotusers

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type users struct {
	Repo  entity.Repo
	Items []*entity.TelebotUser
}

func New(r entity.Repo) entity.TelebotUsers {
	return &users{
		Repo: r,
	}
}

func (c *users) GetItems() []*entity.TelebotUser {
	return c.Items
}
