package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ba11breaker/gomarri/framework"
)

func FooControllerHandler(ctx *framework.Context) error {
	durationContext, cancel := context.WithTimeout(ctx.BaseContext(), time.Duration(1*time.Second))
	// Tell durationCtx's children ccontexts to end
	defer cancel()

	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	go func() {

		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		time.Sleep(10 * time.Second)
		ctx.Json(200, "ok")

		finish <- struct{}{}
	}()

	select {
	case p := <-panicChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		log.Println(p)
		ctx.Json(500, "panic")
	case <-finish:
		fmt.Println("finish")
	case <-durationContext.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(500, "time out")
		ctx.SetHasTimeout()
	}
	return nil
}
