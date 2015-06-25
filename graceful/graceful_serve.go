package graceful

import (
	"net/http"

	"github.com/zenazn/goji/bind"
	"github.com/zenazn/goji/graceful"
)

func Serve(h http.Handler, addr string) error {
	listener := bind.Socket(addr)
	graceful.HandleSignals()

	// support einhorn?
	// bind.Ready()

	if err := graceful.Serve(listener, h); err != nil {
		return err
	}

	graceful.Wait()
	return nil

}
