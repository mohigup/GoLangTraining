package main

import "fmt"

func main() {
	userInput := "X"
	switch "D" {
	case "A":
		fmt.Println("Found A")
		fallthrough
	case "B", "C":
		fmt.Println("Found B or C")
	case "D":
		fmt.Println("Found D")

	}
	switch  {
	case len(userInput) > 1 :
		fmt.Println("Found User Input X")
	default:
		fmt.Println("No Match Found, Default case")
	}
}
