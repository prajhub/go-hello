package main

import (
	"fmt"
	"time"
)

//

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple" // this is shorthand for declaring and initializing a variable var f string = "apple"
	fmt.Println(f)

	const n = "constant"
	fmt.Println(n)

	//For loop is the only loop construct in GO

	i := 1

	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	var num1, num2 = 2, 2

	fmt.Println(num1 % num2)

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// If else statements

	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	//Switch case statemet

	k := 2
	fmt.Print("Write ", k, " as ")
	switch k {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its weekend")
	default:
		fmt.Println("Its weekday")
	}

	//Switch case without an expression
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	//Arrays

	var l [5]int
	fmt.Println("emp: ", l)

	l[4] = 100

	m := [5]int{1, 2, 3, 4, 5}
	fmt.Println(m)

	/*	var twoD [2][3]int
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)

	*/

	s := make([]string, 3)
	//alsos := []int {1, 2, 3}

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "d") //Here what is happening is append returns a new slice so to update s we are reassigning s to the returned value of append.

	fmt.Println("apd:", s)

	exampleArray := []int{1, 2, 3}

	exampleArray = append(exampleArray, 6, 7, 8)
	fmt.Println(exampleArray)

	rangeOne := exampleArray[2:4]
	fmt.Println(rangeOne)

	//So slices are basically arrays under the hood it just that arrays have a fixed length and cant be increased but slices can

}
