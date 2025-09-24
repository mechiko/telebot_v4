package checkdbg

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/internal/usecase"
)

type Checks struct {
	app entity.Application
}

func NewChecks(app entity.Application) *Checks {
	return &Checks{
		app: app,
	}
}

func (c *Checks) Run() error {
	defer c.app.GetRecovery().RecoverLog("check:Run()")
	// c.app.GetLogger().Debugf("CHECK!!! Start only in debug mode")

	// if err := c.ReminderMaster(); err != nil {
	// 	return err
	// }
	// if err := c.ReminderUser(); err != nil {
	// 	return err
	// }

	c.app.GetLogger().Debugf("CHECK!!! Finished")
	return nil
}

func (c *Checks) ReminderUser() error {
	defer c.app.GetRecovery().RecoverLog("check:ReminderUser")
	if err := usecase.New(c.app).UsersReminder(); err != nil {
		return fmt.Errorf("check:ReminderUser %w", err)
	}
	return nil
}

func (c *Checks) ReminderMaster() error {
	defer c.app.GetRecovery().RecoverLog("check:ReminderMaster")
	if err := usecase.New(c.app).MasterReminder(); err != nil {
		return fmt.Errorf("check:ReminderUser %w", err)
	}
	return nil
}
