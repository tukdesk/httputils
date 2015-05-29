package gojimiddleware

import (
	"regexp"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

// RegisterSubroute Example:
// package main

// import (
//     "fmt"
//     "net/http"

//     "github.com/tukdesk/httputils/gojimiddleware"
//     "github.com/zenazn/goji/web"
// )

// func hello1(c web.C, w http.ResponseWriter, r *http.Request) {
//     fmt.Fprint(w, "hello1 _:", c.URLParams)
//     return
// }

// func hello2(c web.C, w http.ResponseWriter, r *http.Request) {
//     fmt.Fprint(w, "hello2 _:", c.URLParams)
//     return
// }

// func hello3(c web.C, w http.ResponseWriter, r *http.Request) {
//     fmt.Fprint(w, "hello3 _:", c.URLParams)
//     return
// }

// func main() {
//     m := web.New()
//     m.Handle("", hello1)
//     m.Handle("/abc/*", hello2)
//     m.Handle("/*", hello3)

//     app := web.New()

//     gojimiddleware.RegisterSubroute("/hello", app, m)
//     http.ListenAndServe(":50001", app)
// }

func RegisterSubroute(prefix string, app, mod *web.Mux) {
	str := "^" + prefix + "(?P<_>(/.*)?)$"
	mod.Use(middleware.SubRouter)
	app.Handle(regexp.MustCompile(str), mod)
	return
}
