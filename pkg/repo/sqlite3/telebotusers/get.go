package telebotusers

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (r *users) Get() error {
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select tu.*, ifnull(us.value, '') as ident from telebotusers tu left join user_states us on us.user_id = tu.id and us."key" = 'ИДЕНТ';`

	err = db.Select(&r.Items, query)
	if err != nil {
		return fmt.Errorf("db.Select() %w", err)
	}
	return nil
}

func (r *users) GetList() (*entity.TelebotUserList, error) {
	ul := &entity.TelebotUserList{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return ul, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select tu.*, ifnull(us.value, '') as ident from telebotusers tu left join user_states us on us.user_id = tu.id and us."key" = 'ИДЕНТ';`

	err = db.Select(&ul.Items, query)
	if err != nil {
		return ul, fmt.Errorf("db.Select() %w", err)
	}
	return ul, nil
}

func (r *users) GetById(id int64) (*entity.TelebotUser, error) {
	u := &entity.TelebotUser{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return nil, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select * from telebotusers where id = ?;`

	err = db.Get(u, query, id)
	if err != nil {
		return nil, fmt.Errorf("db.Select() %w", err)
	}
	return u, nil
}
