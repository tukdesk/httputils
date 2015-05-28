package gojimiddleware

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

var (
	ErrAppClosed = fmt.Errorf("app closed")
)

type App struct {
	m    *web.Mux
	quit chan error
}

func NewApp() *App {
	return NewAppWithMux(nil)
}

func NewAppWithMux(m *web.Mux) *App {
	if m == nil {
		m = web.New()
	}
	return &App{
		m:    m,
		quit: make(chan error, 1),
	}
}

func (this *App) Mux() *web.Mux {
	return this.m
}

func (this *App) run(addr string) {
	this.quit <- http.ListenAndServe(addr, this.m)
}

func (this *App) Run(addr string) error {
	go this.run(addr)
	return <-this.quit
}

func (this *App) Close() {
	this.quit <- ErrAppClosed
}
