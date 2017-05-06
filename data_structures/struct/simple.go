package main

import "fmt"

type foo int
type person struct {
	name string
	age  int
	male bool
}

type child struct {
	person
	hobby string
	name  string
}

func (p person) getName() {
	fmt.Println("I am the Parent")
}

func (c child) getName() {
	fmt.Println("I am the child")
}

func main() {
	// custom type
	var age foo
	fmt.Printf("%T %v \n", age, age)

	var age1 int
	//static- typing
	fmt.Println(age1 + int(age))

	// person type
	p1 := person{"john", 4, true}
	fmt.Println(p1.age)
	// embeded type
	child := child{
		person: person{
			name: "mohit",
			age:  29,
			male: true,
		},
		hobby: "cricket",
		name:  "mohit gupta",
	}
	fmt.Println("overiding fields")
	fmt.Println(child)
	fmt.Println("child name overrides parent:\n",child.name)
	fmt.Println("explicit pull child name:\n",child.person.name)
	fmt.Println("overiding methods")
	child.getName()

}
