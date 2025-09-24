//go:build linux || darwin || freebsd

package entity

// import "github.com/mechiko/telebot_v4/internal/entity"
// if !entity.supported
var (
	DbPath               = "/var/local/telebot"
	LogPath              = "/var/log/telebot"
	ConfigPath           = "/etc/telebot"
	Supported            = true
	Linux                = true
	Windows              = false
	PosixUserUIDGUID int = 1002
	PosixChownPath   int = 0755
	PosixChownFile   int = 0644
)
