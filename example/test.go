package main

import (
	"github.com/7ym0n/goway"
	"os"
)

func main() {
	gm := goway.Bootstrap()
	gm.Get("/", func() string {
		a := "args"
		b := 2
		gm.Logs.Notice("hello test... %v --- %d", a, b)
		return "hello,write"
	})
	gm.Get("/hi/:id", func() string {
		gm.Logs.Notice("say hi test... %v",goway.Params)
		return "say hi,write"
	})
	gm.Get("/hi2/:id/:page", func() string {
		return "say hi,write"
	})

	pwd, _ := os.Getwd()
	gm.Static("/public", pwd)

	gm.NotFound(func() string {
		return "this not found match router!!!"
	})

	gm.Get("/say", func() string {
		return "<html><head><title>test loading</title></head><body><img src='/public/loading.gif'></body></html>"
	})

	gm.Run()
}
