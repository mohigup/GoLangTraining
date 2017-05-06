package main

import "fmt"

const p1 string = " i am constant 1"
const int1 = " i am int 1"


const (
	PI       = 3.14
	Language = `GO`
	_ = iota
	KB = 1 << (iota * 10)
)

func main() {

	const p2 string = "i am constant 2"
	const int2 = 34

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println()
	fmt.Println(int1)
	fmt.Println(int2)
	fmt.Println()
	fmt.Println(PI)
	fmt.Println(Language)
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
}
