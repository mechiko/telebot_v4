package userstates

import (
	"fmt"
)

func (r *userstates) Get(userid int64, key string) (string, error) {
	value := ""
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return value, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select value from user_states where user_id = ? and key = ?;`

	err = db.Get(&value, query, userid, key)
	if err != nil {
		return value, fmt.Errorf("db.Select() %w", err)
	}
	return value, nil
}
