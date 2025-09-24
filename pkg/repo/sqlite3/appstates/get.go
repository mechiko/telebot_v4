package appstates

import (
	"fmt"
)

//	Set(module string, key string, value string) error

func (r *appstates) Get(module string, key string) (string, error) {
	value := ""
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return value, fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select value from app_state where module = ? and key = ?;`

	err = db.Get(&value, query, module, key)
	if err != nil {
		return value, fmt.Errorf("db.Select() %w", err)
	}
	return value, nil
}
