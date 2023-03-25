package main

import "github.com/ba11breaker/gomarri/framework"

// register the router rules
func registerRouter(core *framework.Core) {
	// HTTP methods + static path matching
	core.Get("/user/login", nil)

	subjectApi := core.Group("/subject")
	{
		// HTTP methods + dynamic path matching
		subjectApi.Delete("/:id", nil)
		subjectApi.Get("/:id", nil)
		subjectApi.Put("/:id", nil)
		subjectApi.Get("/list/all", nil)
	}
}
