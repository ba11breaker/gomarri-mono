package main

import "github.com/ba11breaker/gomarri/framework"

// register the router rules
func registerRouter(core *framework.Core) {
	// HTTP methods + static path matching
	core.Get("/user/login", UserLoginController)

	subjectApi := core.Group("/subject")
	{
		// HTTP methods + dynamic path matching
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Get("/:id", SubjectGetController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/list/all", SubjectListController)
	}
}
