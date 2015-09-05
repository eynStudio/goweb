package goweb

import (
	"fmt"
	"io"
	"net/http"
	"path"
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
		ctx = NewHttpContext(r, w)
	)
	Filters[0](ctx, Filters[1:])

	url := path.Clean(r.URL.Path)
	fmt.Println(url)
	route, params := MyRouter.FindRoute(url)
	if route == nil {
		fmt.Println("route not found")
	} else {
		ctrl := MyRouter.FindController(route, params)

		ctx.Result = ctrl.MethodByName("Get").Call(nil)[0].Interface().(Result)
		fmt.Println(ctx.Result)
		fmt.Println(ctrl.Elem())

	}

	if ctx.Result != nil {
		ctx.Result.Apply(ctx)
	} else {
		fmt.Fprintf(w, "Hi, This is an example of http service in golang!")
	}

	if w, ok := ctx.Resp.(io.Closer); ok {
		w.Close()
	}
}
