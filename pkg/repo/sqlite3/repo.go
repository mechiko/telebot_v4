package sqlite3

import (
	"github.com/mechiko/telebot_v4/internal/entity"
)

type repository struct {
	App entity.Application
	Db  entity.DbService
}

// asserts
var _ entity.Repo = (*repository)(nil)

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewRepository(a entity.Application) entity.Repo {
	return &repository{
		App: a,
		Db:  a.GetDbService(),
	}
}
