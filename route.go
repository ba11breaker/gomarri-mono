package main

import "github.com/ba11breaker/gomarri/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
