package userstates

import (
	"fmt"
)

func (r *userstates) Set(userid int64, key string, value string) error {
	query := `INSERT OR REPLACE INTO user_states (user_id, key, value) values (?, ?, ?);`
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, userid, key, value); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}
