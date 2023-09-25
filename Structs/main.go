package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	person1 := Person{
		name: "Alice",
		age:  20,
	}

	fmt.Println(person1)

}
