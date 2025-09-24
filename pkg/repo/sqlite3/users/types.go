package users

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type users struct {
	Repo  entity.Repo
	Items []*entity.User
}

func New(r entity.Repo) entity.Users {
	return &users{
		Repo: r,
	}
}

func (c *users) GetItems() []*entity.User {
	return c.Items
}
