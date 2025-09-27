package telebotusers

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (c *users) Delete(ci *entity.TelebotUser) error {
	query := `DELETE FROM telebotusers WHERE id = ?;`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, ci.ID); err != nil {
		return fmt.Errorf("db.Exec() %w", err)
	}
	return nil
}
