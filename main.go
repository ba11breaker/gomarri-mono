package main

import (
	"net/http"

	"github.com/ba11breaker/gomarri/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		// The core handler function
		Handler: core,
		Addr:    ":8000",
	}

	server.ListenAndServe()
}
