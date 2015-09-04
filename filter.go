package goweb

import (
	"fmt"
)

type Filter func(c *Controller, filterChain []Filter)

var Filters = []Filter{testFilter, RouterFilter, testFilter2}

func testFilter(c *Controller, fc []Filter) {
	fmt.Println("just test")
	fc[0](c, fc[1:])
}

func testFilter2(c *Controller, fc []Filter) {
	fmt.Println("just test 2")

}
