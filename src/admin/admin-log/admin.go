package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = "admin_a"
	fmt.Println(a)

	const HEIGHT int = 5
	const WIDTH = 6
	var area = HEIGHT * WIDTH
	fmt.Printf("面积为: %d\n", area)
	fmt.Println(HEIGHT)

	const (
		O = "abc"
		P = len(O)
		Q = unsafe.Sizeof(O)
	)
	fmt.Println(O, P, Q)

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)

	var b, c = "admin_b", "admin_c"
	fmt.Println(b, c)
	fmt.Printf("Admin fmt")
}
