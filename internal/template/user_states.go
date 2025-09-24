package template

import (
	"bytes"
	_ "embed"
	"html/template"
	"path"

	"github.com/mechiko/telebot_v4/internal/entity"
	"github.com/mechiko/utility"
	// "reflect"
)

//go:embed tmplUserStates.html
var tmplUserStates string

func (tt *templateString) UserStates(ml *entity.UserStatesMaps) (string, error) {
	defer tt.app.GetRecovery().RecoverLog("MissionMessage")
	// так как в шаблоне используем Must то ловим Panic() через defer

	var err error
	var buf bytes.Buffer
	var result string = ""

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
	}

	ftmpl_src := "E:/src/goproj/!telegram_bot_telebot/telebot/internal/template/tmplUserStates.html"
	ftmpl_local := "template/tmplUserStates.html"
	name := path.Base(ftmpl_src)
	if utility.PathOrFileExists(ftmpl_src) {
		t := template.Must(template.New(name).Funcs(funcMap).ParseFiles(ftmpl_src))
		err = t.ExecuteTemplate(&buf, name, ml)
		if err != nil {
			return result, err
		}
	} else if utility.PathOrFileExists(ftmpl_local) {
		t := template.Must(template.New(name).Funcs(funcMap).ParseFiles(ftmpl_local))
		err = t.ExecuteTemplate(&buf, name, ml)
		if err != nil {
			return result, err
		}
	} else {
		t := template.Must(template.New(name).Funcs(funcMap).Parse(tmplUserStates))
		err = t.ExecuteTemplate(&buf, name, ml)
		if err != nil {
			return result, err
		}
	}
	result = buf.String()
	return result, err
}
