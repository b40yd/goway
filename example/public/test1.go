package main

import (
	"fmt"
	"reflect"
)

type Say func(str string)

var t map[reflect.Type]reflect.Value

func test() Say {
	return func(str string) {
		fmt.Println("say:", str)
	}
}

func main() {
	t = make(map[reflect.Type]reflect.Value)
	t[reflect.TypeOf(test())] = reflect.ValueOf(test())
	ev := t[reflect.TypeOf(Say(nil))]
	s := ev.Interface().(Say)
	s("test")

	//
	v := reflect.ValueOf(test())

	vv := v.Interface().(Say)
	vv("vvvv")
}
