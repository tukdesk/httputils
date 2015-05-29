package gojimiddleware

import (
	"regexp"

	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func RegisterSubroute(prefix string, app, mod *web.Mux) {
	str := "^" + prefix + "(?P<_>(/.*)?)$"
	mod.Use(middleware.SubRouter)
	app.Handle(regexp.MustCompile(str), mod)
	return
}
