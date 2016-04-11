package node

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	. "github.com/eynstudio/gobreak"
	"github.com/eynstudio/gobreak/conf"
	"github.com/eynstudio/gobreak/di"
)

type App struct {
	Root INode
	di.Container
	Name   string
	Cfg    *Config
	Server *http.Server
	Router
}

func NewApp(name string) *App {
	var cfg Config
	conf.MustLoadJsonCfg(&cfg, "conf/"+name+".json")
	return NewAppWithCfg(&cfg)
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
	ctx := p.NewCtx(r, w)
	if !ctx.ServeFile(p.Cfg) {
		p.Route(p.Root, ctx)
		log.Println(ctx.Scope)
		if !ctx.Handled {
			ctx.NotFound()
		}
	}

	if w, ok := ctx.Resp.ResponseWriter.(io.Closer); ok {
		w.Close()
	}
}
