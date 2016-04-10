package main

import (
	"log"

	"github.com/eynstudio/goweb/node"
)

func main() {
	log.Println("Hello Start...")
	app := node.NewAppWithCfg(&node.Config{Port: 80})

	home := &Home{node.NewNode("")}
	app.Root.AddNode(home)
	my := app.Root.NewNode("my")
	n2 := node.NewParamNode("{id}")
	my.AddNode(n2)
	n2.NewNode("hi")
	app.Start()

}
