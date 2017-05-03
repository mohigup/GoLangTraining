package main

import "fmt"

func main() {
	//n := average(1, 2, 3, 4)
	data := []float64{43,42,41}
	//n := average(data)
	n := average(data...)
	fmt.Println(n)
}

//func average(sf float64[])
func average(sf ...float64) float64 {
	var total float64
	for _, v := range sf {
		total += v

	}
	return total / float64(len(sf))

}
