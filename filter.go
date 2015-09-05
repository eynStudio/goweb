package goweb

import (
	"fmt"
)

type Filter func(ctx *HttpContext, filterChain []Filter)

var Filters = []Filter{testFilter, RouterFilter, testFilter2}

func testFilter(ctx *HttpContext, fc []Filter) {
	fmt.Println("just test")
	fc[0](ctx, fc[1:])
}

func testFilter2(ctx *HttpContext, fc []Filter) {
	fmt.Println("just test 2")
}
