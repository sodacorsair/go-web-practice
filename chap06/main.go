package main

import (
	"net/http"
	"time"
)

type Cookie struct {
	Name string
	Value string
	Path string
	Domain string
	Expires time.Time
	RawExpires string

	MaxAge int
	Secure bool
	HttpOnly bool
	Raw  string
	Unparsed []string
}

func main() {
    expiration := time.Now()
    expiration = expiration.AddDate(1, 0, 0)
    cookie := http.Cookie{Name: "username", Value: "astaxie", Expires: expiration}
    http.SetCookie(w, &cookie)
}


