package missions

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func (r *missions) Get() (*entity.MissionList, error) {
	ml := &entity.MissionList{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return ml, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()
	query := `select * from missions;`

	err = db.Select(&ml.Items, query)
	if err != nil {
		return ml, fmt.Errorf("db.Select() %w", err)
	}
	return ml, nil
}

func (r *missions) GetById(id int64) (*entity.Mission, error) {
	ml := &entity.Mission{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return ml, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()
	query := `select * from missions where id = ?;`

	err = db.Get(ml, query, id)
	if err != nil {
		return ml, fmt.Errorf("db.Select() %w", err)
	}
	return ml, nil
}

func (r *missions) GetActiveByUid(uid int64) (*entity.Mission, error) {
	ml := &entity.Mission{}
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return ml, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()
	query := `select * from missions where recepient_id = ? and active = 1;`

	err = db.Get(ml, query, uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// пустой результат
			return ml, nil
		}
		if liteErr, ok := err.(*sqlite.Error); ok {
			code := liteErr.Code()
			if code == sqlite3.SQLITE_CONSTRAINT_PRIMARYKEY {
				// нарушение ключа дубликат
				return ml, fmt.Errorf("primary key %w", err)
			}
		}
		return ml, fmt.Errorf("%w", err)
	}
	return ml, nil
}
