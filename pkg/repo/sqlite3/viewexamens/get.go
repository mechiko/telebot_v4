package viewexamens

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (r *viewExamens) GetAll() (*entity.ExamenUserList, error) {
	ml := &entity.ExamenUserList{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return ml, fmt.Errorf("viewexamens:getall %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()
	query := `select * from examen_users;`

	err = db.Select(&ml.Items, query)
	if err != nil {
		return ml, fmt.Errorf("viewexamens:getall %w", err)
	}
	return ml, nil
}
