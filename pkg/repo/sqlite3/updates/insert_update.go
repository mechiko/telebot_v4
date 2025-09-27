package updates

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (c *updates) Insert(u *entity.Update) error {
	query := `INSERT INTO updates (id, message, sender_id, chat_id, recepient, 'update') VALUES(?,?,?,?,?,?);`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("updates:insert %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, u.ID, u.Message, u.Sender, u.Chat, u.Recepient, u.Update); err != nil {
		return fmt.Errorf("updates:insert db.Exec() %w", err)
	}
	return nil
	return nil
}
