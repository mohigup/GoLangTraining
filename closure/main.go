package main

import "fmt"


func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	//type 1
	x := 0
	increment1 := func() int {
		x++
		return x
	}
	fmt.Println(increment1())
	fmt.Println(increment1())
	//type 2
	increment2 := wrapper()
	fmt.Println(increment2())
	fmt.Println(increment2())

}
