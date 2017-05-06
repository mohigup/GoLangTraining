package main

import "fmt"

func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
	for i := 5000; i < 51000; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
}
