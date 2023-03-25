package framework

import (
	"log"
	"net/http"
	"strings"
)

// The core structure of framework
type Core struct {
	router map[string]*Tree
}

// Initialize the core structure
func NewCore() *Core {
	router := map[string]*Tree{
		"GET":    NewTree(),
		"POST":   NewTree(),
		"PUT":    NewTree(),
		"DELETE": NewTree(),
	}

	return &Core{
		router: router,
	}
}

func (c *Core) Get(url string, handler ControllerHandler) {
	if err := c.router["GET"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Post(url string, handler ControllerHandler) {
	if err := c.router["POST"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Put(url string, handler ControllerHandler) {
	if err := c.router["PUT"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Delete(url string, handler ControllerHandler) {
	if err := c.router["DELETE"].AddRouter(url, handler); err != nil {
		log.Fatal("add router error: ", err)
	}
}

func (c *Core) Group(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// Find the route by request, if none, return nul
func (c *Core) FindRouteByRequest(request *http.Request) ControllerHandler {
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
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
