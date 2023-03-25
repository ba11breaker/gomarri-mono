package framework

import (
	"net/http"
	"strings"
)

// The core structure of framework
type Core struct {
	router map[string]map[string]ControllerHandler
}

// Initialize the core structure
func NewCore() *Core {
	getRouter := map[string]ControllerHandler{}
	postRouter := map[string]ControllerHandler{}
	putRouter := map[string]ControllerHandler{}
	deleteRouter := map[string]ControllerHandler{}

	router := map[string]map[string]ControllerHandler{
		"GET":    getRouter,
		"POST":   postRouter,
		"PUT":    putRouter,
		"DELETE": deleteRouter,
	}

	return &Core{
		router: router,
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["GET"][upperUrl] = handler
}

func (c *Core) Post(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["POST"][upperUrl] = handler
}

func (c *Core) Put(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["PUT"][upperUrl] = handler
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	upperUrl := strings.ToUpper(url)
	c.router["DELETE"][upperUrl] = handler
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// Find the route by request, if none, return nul
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)
	upperUri := strings.ToUpper(uri)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		if handler, ok := methodHandlers[upperUri]; ok {
			return handler
		}
	}
	return nil
}

// To implement Handler interface in Core struct
func (c *Core) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// To wrap the context
	ctx := NewContext(request, responseWriter)

	// Find the route by request
	router := c.FindRouteByRequest(request)
	if router == nil {
		ctx.Json(404, "Not Found")
		return
	}

	// Execute the route, if it returns err, it means there is internal error
	// and we should return 500
	if err := router(ctx); err != nil {
		ctx.Json(500, "Internal Error")
		return
	}
}
