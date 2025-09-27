package users

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

// -- таблица логинов клиентов на всякий пожарный чтобы связь с клиентом была
// -- добавить надо НЕВАКОД главным клиентом
// -- таблица логинов клиентов на всякий пожарный чтобы связь с клиентом была
// -- добавить надо НЕВАКОД главным клиентом
// CREATE TABLE if not exists users (
//
//	login TEXT NOT NULL DEFAULT ('') PRIMARY KEY,
//	client_id  integer not null DEFAULT 0, -- link to clients.id
//	passwd TEXT DEFAULT(''),
//	name TEXT DEFAULT(''),
//	email TEXT DEFAULT(''),
//	active integer not null DEFAULT 0,
//	is_admin integer not null DEFAULT 0,
//	rem TEXT DEFAULT(''),
//	unique(email)
//
// );
func (c *users) Insert(ci *entity.User) error {
	query := `INSERT OR REPLACE INTO users(login, client_id, passwd, name, email, active, is_admin, rem) VALUES(?, ?, ?, ?, ?, ?, ?, ?);`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("users:insert %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, ci.Login, ci.ClientId, ci.Passwd, ci.Name, ci.Email, ci.Active, ci.IsAdmin, ci.Rem); err != nil {
		return fmt.Errorf("users:insert db.Exec() %w", err)
	}
	return nil
}

// что инсерт что апдэйт будет одинаково себя вести, по логину как ключу или создает или обновляет поля с логином связанные, возможно
// исключение при нарушении уникальности адреса почты
func (c *users) Update(ci *entity.User) error {
	query := `INSERT OR REPLACE INTO users(login, client_id, passwd, name, email, active, is_admin, rem) VALUES(?, ?, ?, ?, ?, ?, ?, ?);`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("users:update %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, ci.Login, ci.ClientId, ci.Passwd, ci.Name, ci.Email, ci.Active, ci.IsAdmin, ci.Rem); err != nil {
		return fmt.Errorf("users:update %w", err)
	}
	return nil
}

func (c *users) UpdateActive(ci *entity.User) error {
	query := `UPDATE users SET client_id=?, name=?, email=?, active=?, is_admin=?, rem=? WHERE login=?;`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("users:UpdateActive %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, ci.ClientId, ci.Name, ci.Email, ci.Active, ci.IsAdmin, ci.Rem, ci.Login); err != nil {
		return fmt.Errorf("users:UpdateActive %w", err)
	}
	return nil
}
