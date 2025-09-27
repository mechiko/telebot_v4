package sqlite3

import (
	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/appstates"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/chats"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/examenended"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/keystates"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/missions"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/telebotusers"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/updates"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/users"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/userstates"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/viewexamens"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/viewmasters"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/viewmissions"
	"github.com/mechiko/telebot_v4/pkg/repo/sqlite3/views"
)

func (r *repository) DbService() entity.DbService {
	return r.Db
}

func (r *repository) GetTelebotUsers() entity.TelebotUsers {
	return telebotusers.New(r)
}

func (r *repository) GetChats() entity.Chats {
	return chats.New(r)
}

func (r *repository) GetUpdates() entity.Updates {
	return updates.New(r)
}

func (r *repository) GetMissions() entity.Missions {
	return missions.New(r)
}

func (r *repository) GetAppStates() entity.AppStates {
	return appstates.New(r)
}

func (r *repository) GetUserStates() entity.UserStates {
	return userstates.New(r)
}

func (r *repository) GetKeyStates() entity.KeyStates {
	return keystates.New(r)
}

func (r *repository) GetApplication() entity.Application {
	return r.App
}

func (r *repository) GetUsers() entity.Users {
	return users.New(r)
}

func (r *repository) GetViews() entity.Views {
	return views.New(r)
}

func (r *repository) GetViewMissions() entity.MissionUsers {
	return viewmissions.New(r)
}

func (r *repository) GetViewExamens() entity.ExamenUsers {
	return viewexamens.New(r)
}

func (r *repository) GetExamenEndeds() entity.ExamenEndeds {
	return examenended.New(r)
}

func (r *repository) GetMasters() entity.MasterUsers {
	return viewmasters.New(r)
}
