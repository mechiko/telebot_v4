package chats

import (
	"fmt"
)

func (r *chats) Get() error {
	db, err := r.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.Repo.DbService().Close()
	}()

	query := `select * from chats;`

	err = db.Select(&r.Items, query)
	if err != nil {
		return fmt.Errorf("db.Select() %w", err)
	}
	return nil
}
