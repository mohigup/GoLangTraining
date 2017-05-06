package main

import "fmt"

const PI = 3.14

func main() {
	var input float64
	fmt.Print("Enter value\n")
	fmt.Scan(&input)
	answer := input * PI
	fmt.Println(answer)
}
