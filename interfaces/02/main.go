package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int      { return len(p) }
func (p people) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }
func main() {

	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	s := []string{"Zeno", "John", "Al", "Jenny"}
	//sort.StringSlice(s).Sort()
	fmt.Println("before:",s)

	sort.Sort(sort.StringSlice(s))
	fmt.Println("after:",s)


	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println("before:",n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))

	fmt.Println("after:",n)

	m := []int{4,1,2}
	sort.Ints(m)
	fmt.Print(m)

}
