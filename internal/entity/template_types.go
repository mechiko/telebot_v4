package entity

import (
	"time"

	tele "gopkg.in/telebot.v4"
)

type Examen struct {
	Name    string
	DateStr string
	Date    time.Time
}
type UserInfo struct {
	Mission         string
	ConflictExamens []*Examen
	BeforeExamens   []*Examen
	Empty           []*Examen
	Bad             []*Examen
}

type BotTemplate interface {
	MissionList(*MissionList) (string, error)
	MissionMessage(*Mission) (string, error)
	MessageLocation(*tele.Location) (string, error)
	MessageContact(*tele.Contact) (string, error)
	UserStates(*UserStatesMaps) (string, error)
	UserInfo(ui *UserInfo) (string, error)
	ReminderUserInfo(ui *TelebotUserReminderInfo) (string, error)
	ReminderMasterInfo(ui *MasterReminderInfo) (string, error)
}
