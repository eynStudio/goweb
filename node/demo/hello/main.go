package main

import (
	"log"

	"github.com/eynstudio/goweb/node"
)

func main() {
	log.Println("Hello Start...")
	app := node.NewAppWithCfg(&node.Config{Port: 80})

	home := NewHome()
	app.Root.AddNode(home)
	home.NewParamNode("id")
	app.Root.AddNode(node.NewParamNode("id"))
	app.Start()

}
