package main

import (
	"net/http"

	"github.com/ba11breaker/gomarri/framework"
)

func main() {
	server := &http.Server{
		// The core handler function
		Handler: framework.NewCore(),
		Addr:    ":8000",
	}

	server.ListenAndServe()
}
