package examenended

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (r *examenended) Insert(ex *entity.ExamenEnded) error {
	query := `INSERT INTO examen_ended (user_id, key, date) values (?, ?, ?);`
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("examenended:insert %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, ex.UserId, ex.Examen, ex.Date); err != nil {
		return fmt.Errorf("examenended:insert %w", err)
	}
	return nil
}
