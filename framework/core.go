package framework

import (
	"log"
	"net/http"
)

// The core structure of framework
type Core struct {
	router map[string]ControllerHandler
}

// Initialize the core structure
func NewCore() *Core {
	return &Core{
		router: map[string]ControllerHandler{},
	}
}

// Hook Get route with controller
func (c *Core) Get(url string, handler ControllerHandler) {
	c.router[url] = handler
}

// To implement Handler interface in Core struct
func (c *Core) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	log.Println("core.serveHttp")
	ctx := NewContext(request, responseWriter)

	router := c.router["foo"]
	if router == nil {
		return
	}
	log.Println("core.router")
	router(ctx)
}
