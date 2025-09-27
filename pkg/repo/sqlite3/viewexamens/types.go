package viewexamens

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type viewExamens struct {
	Repo  entity.Repo
	Items []*entity.ExamenUser
}

// asserts
var _ entity.ExamenUsers = (*viewExamens)(nil)

func New(r entity.Repo) entity.ExamenUsers {
	return &viewExamens{
		Repo: r,
	}
}
