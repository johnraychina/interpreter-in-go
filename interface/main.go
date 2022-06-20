package main

import "fmt"

type human struct {
	name string
	age  int
}

func main() {

	var people []human
	var alice human
	people = append(people, alice)
	fmt.Println(people)
}
