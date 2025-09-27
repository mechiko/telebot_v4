package sqlite3

import (
	_ "embed"
	"fmt"
	// этот драйвер не зависит от CGO поэтому не проблема для 64 бит
	// _ "modernc.org/sqlite"
	// _ "github.com/mattn/go-sqlite3"
)

// create() создаем БД сначала таблицу для опций, номер версии туда, и скрипт s.CreateDb
func (r *repository) create() error {
	db, err := r.App.GetDbService().Db()
	if err != nil {
		return fmt.Errorf("GetDbService().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.App.GetDbService().Close()
	}()
	if _, err = db.Exec(createVersion); err != nil {
		return fmt.Errorf("db.Exec(createVersion) %w", err)
	}
	if _, err = db.Exec(setVersion, ExeVersion); err != nil {
		return fmt.Errorf("db.Exec(setVersion, ExeVersion) %w", err)
	}
	if _, err = db.Exec(CreateDB); err != nil {
		return fmt.Errorf("db.Exec(CreateDB) %w", err)
	}
	if _, err = db.Exec(IndexDB); err != nil {
		return fmt.Errorf("db.Exec(IndexDB) %w", err)
	}
	if _, err = db.Exec(InsertDB); err != nil {
		return fmt.Errorf("db.Exec(InsertDB) %w", err)
	}
	return err
}
