package users

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (r *users) Get() error {
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select login, name, email, active, is_admin, rem from users  order by login;`

	err = db.Select(&r.Items, query)
	if err != nil {
		return fmt.Errorf("db.Select() %w", err)
	}
	return nil
}

func (r *users) GetByLogin(login string) (*entity.User, error) {
	rl := &entity.User{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return nil, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()
	sql := "select * from users where active = 1 and login = ? limit 1;"
	err = db.Get(rl, sql, login)
	if err != nil {
		return nil, fmt.Errorf("db.Select() %w", err)
	}
	return rl, nil
}

func (r *users) GetByLoginAdmin(login string) (*entity.User, error) {
	rl := &entity.User{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return nil, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()
	// client_id = 1 это НЕВАКОД пока так в тестовых данных
	sql := "select * from users where active = 1 and is_admin = 1 and login = ? limit 1;"
	err = db.Get(rl, sql, login)
	if err != nil {
		return nil, fmt.Errorf("db.Select() %w", err)
	}
	return rl, nil
}
