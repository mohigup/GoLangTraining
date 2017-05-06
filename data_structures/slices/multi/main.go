package main

import "fmt"

func main() {

	//var records [][]string
	records := make([][]string, 0)

	student1 := make([]string, 2)
	student1[0] = "jack"
	student1[1] = "sparrow"

	student2 := make([]string, 2)
	student2[0] = "oliver"
	student2[1] = "queen"

	records = append(records,student1)
	records = append(records,student2)
	fmt.Println(records)
}
