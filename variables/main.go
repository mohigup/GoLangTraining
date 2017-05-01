package main

import "fmt"

func main() {
	var a int
	var b bool
	var c string
	d := "test"
	e := 1
	f := false

	fmt.Printf("%T is %v\n", a, a)
	fmt.Printf("%T is %v\n", b, b)
	fmt.Printf("%T is %v\n", c, c)

	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)

}
