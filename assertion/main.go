package main

import "fmt"

//assertion
func main() {
	var name interface{} = 12
	str, status := name.(string)
	if status {
		fmt.Println(str)
	} else {
		fmt.Println("errror")
	}

	fmt.Println(7 + name.(int))

}
