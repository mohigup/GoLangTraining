package main

import (
	"fmt"
)

func main() {

	data := []int{43, 32, 21, 55}
	fmt.Println(findMax(data...))

}

func findMax(input ...int) int {

	var output int
	for _, v := range input {
		if v > output {
			output = v
		}
	}
	return output
}
