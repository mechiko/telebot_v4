package entity

type ApiMissionAll struct {
	InProgress []*MissionUser
	InFuture   []*MissionUser
	InPast     []*MissionUser
	Day0       []*MissionUser
	Day3       []*MissionUser
	Month      []*MissionUser
}

// Day0 экзамены сегодня
type ApiExamenAll struct {
	Day0     []*ExamenUser
	InFuture []*ExamenUser
	Day3     []*ExamenUser
	Month    []*ExamenUser
}
