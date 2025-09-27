package crontab

import (
	"context"
	"fmt"
	"time"
)

func (ct *crontab) RunClock(ctx context.Context) error {
	var err error
	if ct.timeDuration != 0 {
		return fmt.Errorf("cron not clock")
	}
	ok := true
	select {
	case _, ok = <-ct.chBreak:
	default:
	}
	if !ok {
		return fmt.Errorf("chanel runclock is closed")
	}
	// проверяем если запустили програму после времени часов, чтобы не сработало сразу
	// а ожидало следующего дня
	if ct.timeHour <= time.Now().Hour() {
		ct.everyDay = true
	}
	tickerEveryMinute := time.NewTicker(time.Minute * 1)
	go func() {
		for range tickerEveryMinute.C {
			// ct.app.GetLogger().Debugf("crontab runClock ticker minute %s", time.Now())
			ct.CheckDay()
			if ct.CheckClock() {
				if err = ct.run(); err != nil {
					return
				}
				// задача выполнена на текущий день
				ct.everyDay = true
			}
		}
	}()

	select {
	case <-ctx.Done():
		// если контект внешне прерван сюда приходит сигнал в канал
		ct.app.GetLogger().Debugf("crontab.Clock run recieve <-ctx.Done() %s", time.Now())
		tickerEveryMinute.Stop()
		return nil
	case <-ct.chBreak:
		// если функция объекта вернула ошибку
		ct.app.GetLogger().Debugf("crontab.Clock run recieve <-ct.chBreak %s", time.Now())
		tickerEveryMinute.Stop()
		return err
	}
}

// если день поменялся сбрасываем признак выполненности ежедневного задания
func (ct *crontab) CheckDay() {
	if ct.timeDay != time.Now().Day() {
		ct.timeDay = time.Now().Day()
		ct.everyDay = false
	}
}

// если день поменялся сбрасываем признак выполненности ежедневного задания
// истина если установленный час меньше текущего и установленная минута меньше текущей минуты
// то есть при запуске скажем в 20:30 и установке 07:00 будет истина потому что текущее время прошло черту...
// один раз запускаем при запуске чтобы проверить
func (ct *crontab) CheckClock() bool {
	if ct.everyDay {
		return false
	}
	if ct.timeHour <= time.Now().Hour() {
		if ct.timeMinute <= time.Now().Minute() {
			return true
		}
	}
	return false
}
