package main

import (
	"fmt"
	"github.com/wackonline/goway"
	// "log"
	"os"
	// "reflect"
)

func main() {
	gm := goway.Bootstrap()
	fmt.Println(gm.Logs)
	gm.Get("/", func() string {
		a := "args"
		b := 2
		gm.Logs.Notice("hello test... %v --- %d", a, b)
		return "hello,write"
	})
	gm.Get("/hi/:id", func() string {
		//gm.Router.Params()
		//gm.logger.Notice("say hi test... %v %v",p,p["id"])
		return "say hi,write"
	})
	gm.Get("/hi2/:id/:page", func() string {
		//gm.Router.Params()
		//gm.logger.Notice("say hi test... %v %v",p,p["id2"])
		return "say hi,write"
	})

	pwd, _ := os.Getwd()
	gm.Static("/public", pwd)

	gm.NotFound(func() string {
		return "this not found match router!!!"
	})

	gm.Get("/say", func() string {
		// gm.Logger.Printf("say hello!!!")
		// logger := gm.Injector.Get(reflect.TypeOf(log.Logger(nil))).Interface().(*log.Logger)
		// logger.Printf("say image test...")
		return "<html><head><title>test loading</title></head><body><img src='/public/loading.gif'></body></html>"
	})

	gm.Run()
}
