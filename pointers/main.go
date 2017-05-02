package main

import "fmt"
func main(){

	a := 42
	b := &a
	//var c *int = &a
	fmt.Printf("B is %v\n",*b)
	*b = 43
	fmt.Printf("A value after B changed is %v\n",a)
}