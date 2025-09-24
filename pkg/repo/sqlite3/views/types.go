package views

import (
	_ "embed"
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

type views struct {
	Repo entity.Repo
}

//go:embed examen_users.sql
var examenUsersSql string

//go:embed mission_users.sql
var missionUsersSql string

//go:embed masters_users.sql
var masterUsersSql string

func New(r entity.Repo) entity.Views {
	return &views{
		Repo: r,
	}
}

func (v *views) Create() error {
	db, err := v.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("DbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		v.Repo.DbService().Close()
	}()
	if _, err = db.Exec(examenUsersSql); err != nil {
		return fmt.Errorf("views:create %w", err)
	}
	if _, err = db.Exec(missionUsersSql); err != nil {
		return fmt.Errorf("views:create %w", err)
	}
	if _, err = db.Exec(masterUsersSql); err != nil {
		return fmt.Errorf("views:create %w", err)
	}
	return nil
}
