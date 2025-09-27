package missionadd

import (
	"fmt"

	"github.com/mechiko/telebot_v4/internal/usecase"
	tele "gopkg.in/telebot.v4"
)

func (m *stateMission) save(c tele.Context) error {
	defer m.App.GetRecovery().RecoverLog("saveMissionAdd")
	// if !m.checkMission(c) {
	// 	m.State.ErrorText = "сохранение невозможно, есть ошибки"
	// 	return fmt.Errorf("missadd:save ошибка сохранения миссии")
	// }
	m.State.Mission.RecepientId = c.Chat().ID
	if err := usecase.New(m.App).SetUserActiveMission(m.State.Mission); err != nil {
		return fmt.Errorf("missadd:save %w", err)
	}
	return nil
}

func (m *stateMission) del(c tele.Context) error {
	defer m.App.GetRecovery().RecoverLog("delMissionAdd")
	m.State.Mission.RecepientId = c.Chat().ID
	if err := usecase.New(m.App).ClearUserActiveMission(m.State.Mission); err != nil {
		return fmt.Errorf("missadd:save %w", err)
	}
	return nil
}
