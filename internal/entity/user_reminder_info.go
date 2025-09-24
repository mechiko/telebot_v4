package entity

type TelebotUserReminderInfo struct {
	UserIdent     string
	UserId        int64
	MissionDay0   string // началась и активна
	MissionDay3   string
	MissionDay30  string
	ExamenDay0    []*Examen
	ExamenDay3    []*Examen
	ExamenDay30   []*Examen
	ExamenDayOt60 []*Examen
	ExamenEmpty   []*Examen
	ExamenBad     []*Examen
}

type MasterReminderInfo struct {
	UserIdent     string
	UserId        int64
	MissionDay0   []*MissionUser
	MissionDay3   []*MissionUser
	MissionDay30  []*MissionUser
	MissionBad    []*MissionUser
	ExamenDay0    []*ExamenUser
	ExamenDay3    []*ExamenUser
	ExamenDay30   []*ExamenUser
	ExamenDayOt60 []*ExamenUser
	ExamenEmpty   []*ExamenUser
	ExamenBad     []*ExamenUser
}
