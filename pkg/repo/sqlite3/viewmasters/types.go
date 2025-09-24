package viewmasters

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type viewMasters struct {
	Repo  entity.Repo
	Items []*entity.MasterUser
}

// asserts
var _ entity.MasterUsers = (*viewMasters)(nil)

func New(r entity.Repo) entity.MasterUsers {
	return &viewMasters{
		Repo: r,
	}
}
