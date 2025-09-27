package viewmasters

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (r *viewMasters) GetAll() (*entity.MasterUserList, error) {
	ml := &entity.MasterUserList{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return ml, fmt.Errorf("viewmasters:getall %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()
	query := `select * from master_users;`

	err = db.Select(&ml.Items, query)
	if err != nil {
		return ml, fmt.Errorf("viewmasters:getall %w", err)
	}
	return ml, nil
}
