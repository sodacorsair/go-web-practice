package main

import (
	"html/template"
	"os"
)

//type Person struct {
//	Username string
//	Emails   []string
//	Friends  []*Friend
//}
//
//type Friend struct {
//	Fname string
//}

func main() {
	//t := template.New("filename example")
	//t, _ = t.Parse("hello {{.Username}}!")
	//p := Person{Username: "Astaxie"}
	//t.Execute(os.Stdout, p)

	f1 := Friend{Fname: "minux.ma"}
	f2 := Friend{Fname: "sssss"}
	t := template.New("fieldname example")
	t, _ = t.Parse(`hello {{.Username}}!
					{{range .Emails}}
							an email {{.}}
					{{end}}
					{{with .Friends}}
					{{range .}}
						my friend name is {{.Fname}}
					{{end}}
					{{end}}
					`)
	p := Person{Username: "Astaxie",
		Emails:  []string{"astaxie@gmail.com", "astaxie@beego.me"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)
}
