package goweb

import (
	"fmt"
	"github.com/eynstudio/gobreak/di"
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
	handlers   []Handler
	Router
}

func NewApp(name string) *App {
	r := NewRouter()
	di.Root.MapAs(r, (*Router)(nil))
	c := LoadConfig(name)
	app := &App{Container: di.Root, Name: name, Config: c, Router: r}
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

func (this *App) Use(h Handler) *App {
	checkHandler(h)
	this.handlers = append(this.handlers, h)
	return this
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
	req := Req{r}
	resp := &resp{rw}
	c := &context{di.New(), req, resp, this.handlers, make(map[string]string), nil}
	c.MapAs(c, (*Context)(nil))
	c.MapAs(resp, (*Resp)(nil))
	c.Map(req)
	c.SetParent(this)
	return c
}
func (this *App) handler(w http.ResponseWriter, r *http.Request) {
	ctx := this.NewContext(r, w)
	ctx.exec()

	if ctx.Result != nil {
		ctx.Result.Apply(ctx)
	}
	if w, ok := ctx.Resp.(io.Closer); ok {
		w.Close()
	}
}
