package entity

type UseCase interface {
	GetUsersStatesMap() (UsersStatesMap, error)
	GetUserStatesMap(userid int64) (*UserStatesMaps, error)
	GetUserActiveMission(uid int64) (*Mission, error)
	SetUserActiveMission(mission *Mission) error
	GetUserInfo(state *StateDialog) (string, error)
	// for http
	ApiGetAllMissions(master string) (*ApiMissionAll, error)
	ApiGetAllExamens(master string) (*ApiExamenAll, error)
	ApiGetUsers() (*UserList, error)
	ApiClearAll() error
	ApiGetTelebotUsers() (*TelebotUserList, error)
	UsersReminder() error
	UsersReminderCheck() error
	MasterReminder() error
	ClearUserActiveMission(mission *Mission) error
}
