package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("indes: ", i)
		}
	}
}
