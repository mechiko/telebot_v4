package chats

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type chats struct {
	Repo  entity.Repo
	Items []*entity.Chat
}

func New(r entity.Repo) entity.Chats {
	return &chats{
		Repo: r,
	}
}

func (c *chats) GetItems() []*entity.Chat {
	return c.Items
}
