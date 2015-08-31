package goweb

import (
	"fmt"
)

func Run() {
	app := NewApp()
	app.UseHook(func() {
		fmt.Println("I'm a hook")
	})
	fmt.Println("I'm Start...")
	app.Run()
	fmt.Println("I'm Stop")
}
