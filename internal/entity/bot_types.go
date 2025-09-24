package entity

import "context"

type Bot interface {
	// GetTelebot() interface{}
	GetStates() *States
	Start(ctx context.Context) error
	SendMessageUser(int64, string, string) error
	SendMessageAdmin(string, string) error
	SendMessageGroup(string, string) error
	SendMessageChannel(string, string) error
}
