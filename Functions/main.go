package main

import (
	"fmt"
)

func plus(a, b int) int {
	return a + b
}

func vals() (int, int) {
	return 3, 7
}

func main() {

	finalvalue := plus(2, 3)
	fmt.Println(finalvalue)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

}
