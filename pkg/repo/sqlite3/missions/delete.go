package missions

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (c *missions) Delete(ci *entity.Mission) error {
	query := `DELETE FROM missions WHERE id = ?;`
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

func (c *missions) DeleteById(id int64) error {
	query := `DELETE FROM missions WHERE id = ?;`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, id); err != nil {
		return fmt.Errorf("db.Exec() %w", err)
	}
	return nil
}
