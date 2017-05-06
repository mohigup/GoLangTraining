package main

import (
	"fmt"
)

func evenOddChecker(n int) (float64, bool) {
	return float64(n)/2, n%2 == 0
}

func main(){
	half := func (n int) (float64, bool) {
	return float64(n)/2, n%2 == 0
	}

	number, status := half(4)
	fmt.Println(number, status)
}
