package viewmissions

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type viewMissions struct {
	Repo  entity.Repo
	Items []*entity.MissionUser
}

// asserts
var _ entity.MissionUsers = (*viewMissions)(nil)

func New(r entity.Repo) entity.MissionUsers {
	return &viewMissions{
		Repo: r,
	}
}
