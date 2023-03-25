package main

import (
	"github.com/ba11breaker/gomarri/framework"
)

func UserLoginController(c *framework.Context) error {
	c.Json(200, "ok, UserLoginController")
	return nil
}
