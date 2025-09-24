package entity

type Repository struct {
	App Application
	Db  DbService
}

type Repo interface {
	DbService() DbService
	Start() error
	CheckVersionDb() (bool, error)
	GetTelebotUsers() TelebotUsers
	GetChats() Chats
	GetUpdates() Updates
	GetMissions() Missions
	GetAppStates() AppStates
	GetUserStates() UserStates
	GetKeyStates() KeyStates
	GetApplication() Application
	GetUsers() Users
	GetViews() Views
	GetViewMissions() MissionUsers
	GetViewExamens() ExamenUsers
	GetExamenEndeds() ExamenEndeds
	GetMasters() MasterUsers
}
