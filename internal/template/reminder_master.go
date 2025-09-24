package template

import (
	"bytes"
	_ "embed"
	"path"
	"text/template"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/utility"
	// "reflect"
)

//go:embed tmplReminderMasterInfo.html
var tmplReminderMasterInfo string

func (tt *templateString) ReminderMasterInfo(ui *entity.MasterReminderInfo) (string, error) {
	defer tt.app.GetRecovery().RecoverLog("ReminderMasterInfo")

	var err error
	var buf bytes.Buffer
	var result string = ""

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
	}

	ftmpl_src := "E:/src/goproj/!telegram_bot_telebot/telebot/internal/template/tmplReminderMasterInfo.html"
	ftmpl_local := "template/tmplReminderMasterInfo.html"
	name := path.Base(ftmpl_src)
	if utility.PathOrFileExists(ftmpl_src) {
		t := template.Must(template.New(name).Funcs(funcMap).ParseFiles(ftmpl_src))
		err = t.ExecuteTemplate(&buf, name, ui)
		if err != nil {
			return result, err
		}
	} else if utility.PathOrFileExists(ftmpl_local) {
		t := template.Must(template.New(name).Funcs(funcMap).ParseFiles(ftmpl_local))
		err = t.ExecuteTemplate(&buf, name, ui)
		if err != nil {
			return result, err
		}
	} else {
		t := template.Must(template.New(name).Funcs(funcMap).Parse(tmplReminderMasterInfo))
		err = t.ExecuteTemplate(&buf, name, ui)
		if err != nil {
			return result, err
		}
	}
	result = buf.String()
	return result, err
}
