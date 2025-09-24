package telebotusers

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (c *users) Insert(u *entity.TelebotUser) error {
	query := `INSERT INTO telebotusers (id, first_name, last_name, username, language_code, is_bot, is_premium, added_to_menu, can_join_groups, can_read_messages,
		supports_inline, is_admin, masters) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?);`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("users:insert %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, u.ID, u.FirstName, u.LastName, u.Username, u.LanguageCode, u.IsBot, u.IsPremium, u.AddedToMenu, u.CanJoinGroups, u.CanReadMessages,
		u.SupportsInline, u.IsAdmin, u.Masters); err != nil {
		return fmt.Errorf("users:insert db.Exec() %w", err)
	}
	return nil
}

func (c *users) Update(u *entity.TelebotUser) error {
	query := `UPDATE telebotusers SET is_admin=?, masters=? WHERE id=?;`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("users:update %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, u.IsAdmin, u.Masters, u.ID); err != nil {
		return fmt.Errorf("users:update db.Exec() %w", err)
	}
	// прописываем Ident
	if u.Ident != "" {
		query1 := `select key from state_key where is_intro = 1;`
		intro := ""
		if err := db.Get(&intro, query1); err != nil {
			return fmt.Errorf("%w", err)
		}
		query2 := `INSERT OR REPLACE INTO user_states (user_id, key, value) values (?, ?, ?);`
		if _, err := db.Exec(query2, u.ID, intro, u.Ident); err != nil {
			return fmt.Errorf("%w", err)
		}
	}
	return nil
}
