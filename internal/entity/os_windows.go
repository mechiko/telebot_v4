//go:build windows

package entity

var (
	ConfigPath       = ""
	DbPath           = ""
	LogPath          = "./logs"
	Supported        = true
	Windows          = true
	Linux            = false
	PosixUserUIDGUID = 1002
	PosixChownPath   = 0755
	PosixChownFile   = 0644
)
