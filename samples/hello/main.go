package main

import (
	"fmt"
	"github.com/eynstudio/goweb"
)

func main() {

	goweb.MyRouter.Route("/api/{ctrl}-{eyn}/{id}")
	//	goweb.Run()
}

func Perror(err error) {
	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}
}
