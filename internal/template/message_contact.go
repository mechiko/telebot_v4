package template

import (
	"bytes"
	_ "embed"
	"path"
	"text/template"

	"github.com/mechiko/utility"
	tele "gopkg.in/telebot.v4"
	// "reflect"
)

//go:embed tmplMessageContact.txt
var tmplMessageContact string

func (tt *templateString) MessageContact(ml *tele.Contact) (string, error) {
	defer tt.app.GetRecovery().RecoverLog("MessageLocation")

	var err error
	var buf bytes.Buffer
	var result string = ""

	funcMap := template.FuncMap{
		// The name "inc" is what the function will be called in the template text.
		"inc": func(i int) int {
			return i + 1
		},
	}

	ftmpl_src := "E:/src/goproj/!telegram_bot_telebot/telebot/internal/template/tmplMessageContact.txt"
	ftmpl_local := "template/tmplMessageContact.txt"
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
		t := template.Must(template.New(name).Funcs(funcMap).Parse(tmplMessageContact))
		err = t.ExecuteTemplate(&buf, name, ml)
		if err != nil {
			return result, err
		}
	}
	result = buf.String()
	return result, err
}
