package viewmissions

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (r *viewMissions) GetAll() (*entity.MissionUserList, error) {
	ml := &entity.MissionUserList{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return ml, fmt.Errorf("viewmissions:getall %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()
	query := `select * from mission_users;`

	err = db.Select(&ml.Items, query)
	if err != nil {
		return ml, fmt.Errorf("viewmissions:getall %w", err)
	}
	return ml, nil
}

func (r *viewMissions) GetAllActive() (*entity.MissionUserList, error) {
	ml := &entity.MissionUserList{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return ml, fmt.Errorf("viewmissions:getall %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()
	query := `select * from mission_users where active = 1;`

	err = db.Select(&ml.Items, query)
	if err != nil {
		return ml, fmt.Errorf("viewmissions:getall %w", err)
	}
	return ml, nil
}
