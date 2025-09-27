package updates

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type updates struct {
	Repo  entity.Repo
	Items []*entity.Update
}

func New(r entity.Repo) entity.Updates {
	return &updates{
		Repo: r,
	}
}

func (c *updates) GetItems() []*entity.Update {
	return c.Items
}
