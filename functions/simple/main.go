package main

import "fmt"

func main() {

	fmt.Println(greet("Jack", "Harris"))
	fmt.Println(welcome("Jack", "Harris"))
	fmt.Println(hi("Jack", "Harris"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

func welcome(fname, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

func hi(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}


