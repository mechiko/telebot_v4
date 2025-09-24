package entity

type MissionUser struct {
	MissionId       int64  `db:"mission_id" json:"missionId"`
	UserName        string `db:"name" json:"userName"`
	UserId          int64  `db:"user_id" json:"userId"`
	UserIdent       string `db:"ident" json:"userIdent"`
	BeforeStartDays int64  `db:"before_start_days" json:"beforeStartDays"`
	BeforeEndDays   int64  `db:"before_end_days" json:"beforeEndDays"`
	StartOrdered    string `db:"start_ordered" json:"startOrdered"`
	EndOrdered      string `db:"end_ordered" json:"endOrdered"`
	Place           string `db:"place" json:"place"`
	Start           string `db:"start" json:"start"`
	End             string `db:"end" json:"end"`
	Active          bool   `db:"active" json:"active"`
	Masters         string `db:"masters_list" json:"masters"`
	Conflicts       string `db:"conflicts" json:"conflicts"`
	InProgress      bool   `db:"in_progress" json:"inProgress"`
	InFuture        bool   `db:"in_future" json:"inFuture"`
	InPast          bool   `db:"in_past" json:"inPast"`
}

type MissionUserList struct {
	Items []*MissionUser `json:"rows"`
}

type MissionUsers interface {
	GetAll() (*MissionUserList, error)
	GetAllActive() (*MissionUserList, error)
	// GetAllEnded() (*MissionUserList, error)
	// GetByMasterActive() (*MissionUserList, error)
	// GetByUserActive() (*MissionUserList, error)
	// GetByMaster() (*MissionUserList, error)
	// GetByUser() (*MissionUserList, error)
}
