package main

import "fmt"

func main() {

	mySlice1 := []int{1, 2, 3, 4, 4}
	fmt.Printf("%v - %T\n", mySlice1, mySlice1)

	mySlice3 := []int{9,10,11}

	mySlice1 = append(mySlice1,mySlice3...)
	fmt.Println(mySlice1)

	mySlice2 := []string{"a", "b", "c", "d",}
	// n-1th index to nth index
	fmt.Println(mySlice2[1:2])
	fmt.Println(mySlice2[1:3])
	fmt.Println(mySlice2[1:])
	fmt.Println(mySlice2[:3])
	fmt.Println(mySlice2[:])
	fmt.Println(mySlice2[3])

	// use append to add new items to slice
	/*myslice3 := make([]int, 0, 5)
	for i := 0; i < 43; i++ {
		myslice3 = append(myslice3, i)
		fmt.Println("Len: ", len(myslice3), " Capacity: ", cap(myslice3), " Value:", myslice3[i])
	}*/
}
