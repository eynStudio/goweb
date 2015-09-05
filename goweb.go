package goweb

import (
	"fmt"
)

func Run() {

	NewApp("conf").UseHook(func() {
		fmt.Println("I'm a hook")
	}).Start()

}
