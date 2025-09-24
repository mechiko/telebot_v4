package sqlite3

import (
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (r *repository) InsertUsers() error {
	db, err := r.App.GetDbService().Db()
	if err != nil {
		return fmt.Errorf("InsertUsers().Db() %w", err)
	}
	defer func() {
		db.Close()
		r.App.GetDbService().Close()
	}()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("APremote14!"), 8)
	passwd := hex.EncodeToString(hashedPassword)
	query := `INSERT OR REPLACE INTO users(login, passwd, name, email, active, is_admin, rem) VALUES 
	('kbprime@mail.ru', ?, 'mikl', 'kbprime@mail.ru', 1, 1, 'author'),
	('a.kuleshov.m@gmail.com', ?, 'Admin', 'a.kuleshov.m@gmail.com', 1, 1, 'na4alnik'),
        ('n91n91@mail.ru', ?, 'nastya', 'n91n91@mail.ru', 1, 1, 'info');`

	if _, err = db.Exec(query, passwd, passwd, passwd); err != nil {
		return fmt.Errorf("db.Exec() %w", err)
	}
	return nil
}
