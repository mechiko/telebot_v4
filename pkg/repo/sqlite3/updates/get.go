package updates

import (
	"fmt"
)

func (r *updates) Get() error {
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select * from updates;`

	err = db.Select(&r.Items, query)
	if err != nil {
		return fmt.Errorf("db.Select() %w", err)
	}
	return nil
}
