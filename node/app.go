package node

import (
	"fmt"
	"log"
	//	"io"
	"net/http"
	"time"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/di"
)

type App struct {
	Root INode
	di.Container
	Name   string
	Cfg    *Config
	Server *http.Server
	//	SetupHooks []func()
	//	handlers   []Handler
	//	Router
}

func NewApp(name string) *App {
	//	r := NewRouter()
	//	di.Root.MapAs(r, (*Router)(nil))
	c := LoadConfig(name)
	return NewAppWithCfg(c)
}
func NewAppWithCfg(c *Config) *App {
	app := &App{Root: NewNode(""), Container: di.Root, Name: "", Cfg: c}
	app.Server = &http.Server{
		Addr:         fmt.Sprintf(":%d", c.Port),
		Handler:      http.HandlerFunc(app.handler),
		ReadTimeout:  time.Minute,
		WriteTimeout: time.Minute,
	}
	app.Map(c)
	return app
}

func (P *App) Start() {
	if P.Cfg.Tls {
		err := P.Server.ListenAndServeTLS(P.Cfg.CertFile, P.Cfg.KeyFile)
		if err != nil {
			panic(err)
		}
	} else {
		err := P.Server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}
}

//func (this *App) Use(h Handler) *App {
//	checkHandler(h)
//	this.handlers = append(this.handlers, h)
//	return this
//}

//func (this *App) UseHook(f func()) *App {
//	this.SetupHooks = append(this.SetupHooks, f)
//	return this
//}

//func (this *App) runSetupHooks() {
//	for _, hook := range this.SetupHooks {
//		hook()
//	}
//}
func (p *App) NewCtx(r *http.Request, rw http.ResponseWriter) *Ctx {
	req := Req{r}
	resp := &Resp{rw}
	c := &Ctx{Container: di.New(), Req: req, Resp: resp, Scope: M{}}
	c.Map(c) //需要吗？
	c.Map(resp)
	c.Map(req)
	c.SetParent(p)
	c.urlParts = *newUrlParts(req.Url())
	return c
}
func (p *App) handler(w http.ResponseWriter, r *http.Request) {
	log.Println("App.handler", r.URL.Path)

	ctx := p.NewCtx(r, w)
	p.Root.Router(ctx)
	//	ctx.exec()

	//	if ctx.Result != nil {
	//		ctx.Result.Apply(ctx)
	//	}
	//	if w, ok := ctx.Resp.(io.Closer); ok {
	//		w.Close()
	//	}
}
