package goweb

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type App struct {
	Server     *http.Server
	SetupHooks []func()
}

func NewApp() *App {
	app := &App{Server: &http.Server{
		Addr:         ":80",
		Handler:      http.HandlerFunc(handler),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}}
	return app
}

func (this *App) Run() {
	this.runSetupHooks()

	this.Server.ListenAndServe()
}

func (this *App) UseHook(f func()) {
	this.SetupHooks = append(this.SetupHooks, f)
}

func (this *App) runSetupHooks() {
	for _, hook := range this.SetupHooks {
		hook()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {

	var (
		req  = NewRequest(r)
		resp = NewResponse(w)
		c    = NewController(req, resp)
	)
	Filters[0](c, Filters[1:])

	if c.Result != nil {
		c.Result.Apply(req, resp)
	} else if c.Response.Status != 0 {
		c.Response.Out.WriteHeader(c.Response.Status)
	} else {
		fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
	}

	if w, ok := resp.Out.(io.Closer); ok {
		w.Close()
	}
}
