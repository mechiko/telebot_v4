package application

import (
	"fmt"
	"path"

	"github.com/mechiko/telebot_v4/pkg/sqlite3"
)

func (a *applicationType) initDb(p string) error {
	cfg := a.configuration
	dbName := cfg.Database.DbName
	if dbName == "" {
		return fmt.Errorf("initDb Database.DbName empty")
	}
	if p != "" {
		dbName = path.Join(p, dbName)
	}
	dbs, err := sqlite3.NewDbService(dbName, sqlite3.RwModeWithCreate, sqlite3.Version)
	if err != nil {
		return fmt.Errorf("initDb %w", err)
	}
	a.db = dbs
	// return fmt.Errorf("test error")
	return nil
}
