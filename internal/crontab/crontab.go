package crontab

import (
	"time"

	"github.com/mechiko/telebot_v4/internal/entity"
)

// everyDay ставится true когда cronClock выполнен потом сбрасывается при смене дня
// CheckDay()
type crontab struct {
	app          entity.Application
	run          entity.CronTabFunc
	timeDuration time.Duration // Seconds
	chBreak      chan bool
	timeHour     int
	timeMinute   int
	timeDay      int
	everyDay     bool
}

// Interface assertions
var _ entity.CronTab = (*crontab)(nil)

func NewDuration(app entity.Application, td time.Duration, f entity.CronTabFunc) entity.CronTab {
	return &crontab{
		app:          app,
		run:          f,
		timeDuration: td,
		chBreak:      make(chan bool, 1),
	}
}

func NewClock(app entity.Application, th, tm int, f entity.CronTabFunc) entity.CronTab {
	return &crontab{
		app:        app,
		run:        f,
		timeHour:   th,
		timeMinute: tm,
		timeDay:    time.Now().Day(),
		everyDay:   false,
		chBreak:    make(chan bool, 1),
	}
}
