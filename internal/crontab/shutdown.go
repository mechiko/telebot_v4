package crontab

func (ct *crontab) Shutdown() {
	ct.chBreak <- true
	close(ct.chBreak)
}
