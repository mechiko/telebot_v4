package entity

var YamlConfig = []byte(`
hostname:	"localhost"
htmlport:  "3550"
hostport: "3650"
debug: true
token: ""
groupid: -1001876139554
channelid: 1001893512470
adminid: 68255313
telebot: true

layouts:
    timelayout: "02.01.2006 15:04:05"
    timelayoutday: "02.01.2006"

database:
  timeout: 2
  driver: "sqlite3"
  connectionuri: "?cache=shared&mode=rw"
  dbname: "telebot.db"
`)
