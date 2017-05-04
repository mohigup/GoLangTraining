package main

import (
	"fmt"
)

func main() {

	map1 := make(map[string]int)
	map1["a"] = 1
	map1["b"] = 2

	fmt.Println(map1)
	delete(map1, "a")
	_, prs := map1["b"]
	fmt.Println(prs)

	fmt.Println("Example demo for comma ok Idiom MAP")
	map2 := map[string]int{"foo":1,"bar":2}
	if val, exists := map2["foo"]; exists {
		fmt.Println("found")
		fmt.Println(exists)
		fmt.Println(val)
	} else{
		fmt.Println("not found")
	}
	fmt.Println("Example demo for range MAP")
	for key, value := range map2 {
		fmt.Println(key)
		fmt.Println(value)
	}

}
