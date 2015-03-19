package main

import (
	"fmt"
	"github.com/wackonline/goway"
	"os"
)

func main() {
	gm := goway.Bootstrap()

	gm.Get("/", func() string {
		fmt.Println("hello test...")
		return "hello,write"
	})
	gm.Get("/hi", func() string {
		fmt.Println("say hi test...")
		return "say hi,write"
	})
	pwd, _ := os.Getwd()
	gm.Static("/public", pwd)

	gm.Get("/say", func() string {
		fmt.Println("say image test...")
		return "<html><head><title>test loading</title></head><body><img src='/public/loading.gif'></body></html>"
	})

	gm.RunOnAddr("0.0.0.0:8080")
}
