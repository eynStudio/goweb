package main

import (
	"fmt"
	"github.com/eynstudio/goweb"
	"github.com/eynstudio/goweb/samples/hello/controllers"
)

func main() {

	app := goweb.NewApp("conf").UseHook(func() {
		fmt.Println("I'm a hook")
	})
	app.Router.Route("/api/{ctrl}/{id:[0-9]+}")
	app.Router.Register((*controllers.Home)(nil))
	app.Start()

}

func Perror(err error) {
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}
}
