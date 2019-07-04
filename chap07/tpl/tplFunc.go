package main

import (
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Person struct {
	Username string
	Emails   []string
	Friends  []*Friend
}

type Friend struct {
	Fname string
}

func EmailDealWith(args ...interface{}) string {
	ok := false
	var s string
	if len(args) == 1 {
		s, ok = args[0].(string)
	}
	if !ok {
		s = fmt.Sprint(args...)
	}
	substrs := strings.Split(s, "@")
	if len(substrs) != 2 {
		return s
	}

	return (substrs[0] + " at " + substrs[1])
}

func main() {
	f1 := Friend{Fname: "mm"}
	f2 := Friend{Fname: "ssss"}
	t := template.New("fieldname example")
	t = t.Funcs(template.FuncMap{"emailDeal": EmailDealWith})
	t, _ = t.Parse(`hello {{.Username}}
								{{range .Emails}}
									an emails {{.|emailDeal}}
								{{end}}}
								{{with .Friends}}
								{{range .}}
										my friend name is {{.Fname}}
								{{end}}
								{{end}}
								`)
	p := Person{Username: "Astaxie",
		Emails:  []string{"sssss@ffs.com", "fdfdfdA@ff.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
