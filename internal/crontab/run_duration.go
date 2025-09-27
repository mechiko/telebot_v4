package crontab

import (
	"context"
	"fmt"
	"time"
)

func (ct *crontab) RunDuration(ctx context.Context) error {
	var err error

	if ct.timeDuration == 0 {
		return fmt.Errorf("cron not clock")
	}
	ok := true
	select {
	case _, ok = <-ct.chBreak:
	default:
	}
	if !ok {
		return fmt.Errorf("chanel duration is closed")
	}
	ticker := time.NewTicker(ct.timeDuration)
	go func() {
		for range ticker.C {
			// ct.app.GetLogger().Debugf("crontab run ticket %s", time.Now())
			if err = ct.run(); err != nil {
				return
			}
		}
	}()

	select {
	case <-ctx.Done():
		// если контект внешне прерван сюда приходит сигнал в канал
		ct.app.GetLogger().Debugf("crontab.Duration run recieve <-ctx.Done() %s", time.Now())
		ticker.Stop()
		return nil
	case <-ct.chBreak:
		// если функция объекта вернула ошибку
		ct.app.GetLogger().Debugf("crontab.Duration run recieve <-ct.chBreak %s", time.Now())
		ticker.Stop()
		return err
	}
}
