package main

import "fmt"

func main() {
	var a int
	var b bool
	var c string
	//:= declares and assigns, = just assigns
	d := "test"
	e := 1
	f := false
	g := `this is new string`

	fmt.Printf("%T is %v\n", a, a)
	fmt.Printf("%T is %v\n", b, b)
	fmt.Printf("%T is %v\n", c, c)

	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)

}
