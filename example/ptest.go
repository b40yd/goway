package main

import (
	"fmt"
)

const (
	A = 0
	B = 1
	C = 2
	D = 3
	E = 4
)

var all int

func testMain() {
	all = A | B | C | D
	fmt.Printf("x0:%x\n", "E_ALL")
	fmt.Println("ALL: ", all)

	isa := all & A
	fmt.Println("A: ", isa)

	isb := all & B
	fmt.Println("B: ", isb)

	isc := all & C
	fmt.Println("C: ", isc)

	isd := all & D
	fmt.Println("D: ", isd)

	ise := all & E
	fmt.Println("E: ", ise)

}
