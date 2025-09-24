telebot v4
systemctl list-timers
systemctl daemon-reload
systemctl stop YOUR_TIMER_NAME.timer
systemctl disable YOUR_TIMER_NAME.timer
remove from /etc/systemd/system/
systemctl list-units --type=service --state=running
