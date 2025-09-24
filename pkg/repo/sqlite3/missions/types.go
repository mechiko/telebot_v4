package missions

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type missions struct {
	Repo  entity.Repo
	Items []*entity.Mission
}

func New(r entity.Repo) entity.Missions {
	return &missions{
		Repo: r,
	}
}

func (c *missions) GetItems() []*entity.Mission {
	return c.Items
}
