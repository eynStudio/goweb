package goweb

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type App struct {
	Name       string
	Config     *Config
	Server     *http.Server
	SetupHooks []func()
}

func NewApp(name string) *App {
	app := &App{}
	app.Init(name)
	return app
}

func (this *App) Init(name string) *App {
	this.Name = name
	this.Config = LoadConfig(name)

	this.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", this.Config.Port),
		Handler:      http.HandlerFunc(handler),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	return this
}
func (this *App) Start() {
	this.runSetupHooks()

	if this.Config.Tls {
		err := this.Server.ListenAndServeTLS(this.Config.CertFile, this.Config.KeyFile)
		if err != nil {
			panic(err)
		}
	} else {
		err := this.Server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}
}
func (this *App) UseHook(f func()) *App {
	this.SetupHooks = append(this.SetupHooks, f)
	return this
}

func (this *App) runSetupHooks() {
	for _, hook := range this.SetupHooks {
		hook()
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = NewHttpContext(r, w)
	)
	Filters[0](ctx, Filters[1:])

	if ctx.Result != nil {
		ctx.Result.Apply(ctx)
	} else {
		fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
	}

	if w, ok := ctx.Resp.(io.Closer); ok {
		w.Close()
	}
}
