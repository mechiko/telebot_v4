package entity

import "context"

type CronTab interface {
	RunDuration(context.Context) error
	RunClock(context.Context) error
	Shutdown()
}

type CronTabFunc func() error
