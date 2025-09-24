package missionadd

import (
	"github.com/mechiko/telebot_v4/internal/entity"
	tele "gopkg.in/telebot.v4"
)

// получаем команды и меняем состояния, свои команды исполняем
func (m *stateMission) Router(c tele.Context) error {
	switch m.State.Text {
	case `/start`:
		c.Delete()
		m.State.Text = ""
		m.State.Mode = entity.ModeStart
		m.State.Stage = entity.StageEmpty
	case `/clear`:
		c.Delete()
		m.State.Text = ""
		m.State.Mode = entity.ModeClear
		m.State.Stage = entity.StageEmpty
	case `/mission`:
		c.Delete()
		m.State.Text = ""
		m.State.Mode = entity.ModeMissionAdd
		m.State.Stage = entity.StageEmpty
	case `/examens`:
		c.Delete()
		m.State.Text = ""
		m.State.Mode = entity.ModeEditUserStates
		m.State.Stage = entity.StageEmpty
	}
	return nil
}
