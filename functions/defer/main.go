package main

import "fmt"

func world(){
	fmt.Println("world")
}

func hello(){
	fmt.Println("hello")
}

func main(){
	defer world()
	hello()
}
