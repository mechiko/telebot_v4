package sqlite3

import (
	"fmt"
)

func (r *repository) InsertStateKeys() error {
	db, err := r.App.GetDbService().Db()
	if err != nil {
		return fmt.Errorf("InsertStates().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.App.GetDbService().Close()
	}()
	query := `INSERT OR REPLACE INTO state_key(key, is_examen, is_intro, description) VALUES
('ОТ', 1, 0 ,'ОТ'),
('ППБ', 1, 0 ,'ППБ'),
('ОПЭ АС, ДИ', 1, 0 ,'ОПЭ АС, ДИ'),
('ФНП', 1, 0 ,'ФНП'),
('ПРБ', 1, 0 ,'ПРБ'),
('ЭБ', 1, 0 ,'ЭБ'),
('Медосмотр', 1, 0 ,'Медосмотр'),
('ИДЕНТ', 0, 1 ,'идентификатор');`

	if _, err = db.Exec(query); err != nil {
		return fmt.Errorf("db.Exec() %w", err)
	}
	return nil
}
