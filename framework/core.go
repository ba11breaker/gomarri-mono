package framework

import "net/http"

// The core structure of framework
type Core struct {
}

// Initialize the core structure
func NewCore() *Core {
	return &Core{}
}

// To implement Handler interface in Core struct
func (c *Core) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// To do
}
