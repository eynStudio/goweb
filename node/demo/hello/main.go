package main

import (
	"log"

	"github.com/eynstudio/goweb/node"
)

func main() {
	log.Println("Hello Start...")
	app := node.NewAppWithCfg(&node.Cfg{Port: 80})

	app.Root.AddNode(NewHome())
	app.Root.AddNode(NewApi())
	app.Start()

}
