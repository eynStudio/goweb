package goweb

import (
	"fmt"
	"github.com/eynStudio/gobreak/di"
	"io"
	"net/http"
	"time"
)

type App struct {
	di.Container
	Name       string
	Config     *Config
	Server     *http.Server
	SetupHooks []func()
	Router
}

func NewApp(name string) *App {
	r := NewRouter()
	di.Root.MapAs(r, (*Router)(nil))
	c := LoadConfig(name)
	app := &App{di.Root, name, c, nil, nil, r}
	app.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Port),
		Handler:      http.HandlerFunc(app.handler),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}

	return app
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
func (this *App) NewContext(r *http.Request, rw http.ResponseWriter) *context {
	c := &context{di.New(), r, rw, nil, nil, make(map[string]string), nil}
	c.SetParent(this)
	return c
}
func (this *App) handler(w http.ResponseWriter, r *http.Request) {
	ctx := this.NewContext(r, w)
	ctx.App = this
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
