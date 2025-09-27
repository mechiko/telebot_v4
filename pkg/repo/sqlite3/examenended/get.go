package examenended

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (r *examenended) Get() (*entity.ExamenEndedList, error) {
	value := &entity.ExamenEndedList{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return value, fmt.Errorf("examenended:get %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select * from examen_ended;`

	err = db.Select(&value.Items, query)
	if err != nil {
		return value, fmt.Errorf("examenended:get %w", err)
	}
	return value, nil
}
