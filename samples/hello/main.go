package main

import (
	"fmt"
	"github.com/eynstudio/goweb"
	"github.com/eynstudio/goweb/samples/hello/controllers"
)

func main() {

	goweb.MyRouter.Route("/api/{ctrl}-{eyn}/{id:[0-9]+}")
	goweb.MyRouter.Register((*controllers.Home)(nil))

	goweb.Run()
}

func Perror(err error) {
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}
}
