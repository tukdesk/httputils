package gojimiddleware

import (
	"fmt"
	"log"

	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
)

var (
	ErrAppClosed = fmt.Errorf("app closed")
)

type App struct {
	m *web.Mux
}

func NewApp() *App {
	return NewAppWithMux(nil)
}

func NewAppWithMux(m *web.Mux) *App {
	if m == nil {
		m = web.New()
	}
	return &App{
		m: m,
	}
}

func (this *App) Mux() *web.Mux {
	return this.m
}

// see goji.Serve()
func (this *App) Run(addr string) error {
	listener := bind.Socket(addr)
	log.Println("App listen on ", listener.Addr())
	graceful.HandleSignals()

	// support einhorn?
	bind.Ready()

	graceful.PreHook(func() { log.Println("App received signal, gracefully stopping") })
	graceful.PostHook(func() { log.Println("App stopped") })

	if err := graceful.Serve(listener, this.m); err != nil {
		return err
	}

	graceful.Wait()
	return ErrAppClosed
}

func (this *App) Close() {
	graceful.Shutdown()
}

func (this *App) CloseNow() {
	graceful.ShutdownNow()
}
