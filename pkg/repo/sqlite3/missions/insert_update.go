package missions

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
)

func (c *missions) Insert(m *entity.Mission) error {
	query := `INSERT INTO missions (recepient_id, start, end, place, departament, target, rem, active)	VALUES(?,?,?,?,?,?,?,?);`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("missions:insert %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, m.RecepientId, m.Start, m.End, m.Place, m.Department, m.Target, m.Rem, m.Active); err != nil {
		return fmt.Errorf("missions:insert db.Exec() %w", err)
	}
	return nil
}

func (c *missions) Update(m *entity.Mission) error {
	query := `UPDATE missions SET recepient_id=?, start=?, end=?, place=?, departament=?, target=?, rem=?, active=? 	WHERE id=?;`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("missions:update %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, m.RecepientId, m.Start, m.End, m.Place, m.Department, m.Target, m.Rem, m.Active, m.ID); err != nil {
		return fmt.Errorf("missions:update db.Exec() %w", err)
	}
	return nil
}

// очищаем активность у всех по пациенту
func (c *missions) UpdateClearActive(m *entity.Mission) error {
	query := `UPDATE missions SET active=0 	WHERE recepient_id=?;`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("missions:update %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, m.RecepientId); err != nil {
		return fmt.Errorf("missions:update db.Exec() %w", err)
	}
	return nil
}

// очищаем активность у всех по пациенту
func (c *missions) UpdateClearActiveById(id int64) error {
	query := `UPDATE missions SET active=0 	WHERE id=?;`
	db, err := c.Repo.DbService().Db()
	if err != nil {
		return fmt.Errorf("missions:UpdateClearActiveById %w", err)
	}
	defer func() {
		db.Close()
		c.Repo.DbService().Close()
	}()

	if _, err := db.Exec(query, id); err != nil {
		return fmt.Errorf("missions:UpdateClearActiveById %w", err)
	}
	return nil
}
