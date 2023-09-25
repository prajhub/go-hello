package main

import "fmt"

func main() {

	languages := make(map[string]string)
	kvs := map[string]string{"a": "apple", "b": "banana"}

	languages["JS"] = "Javascrupt"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	kvs["c"] = "cider"

	fmt.Println(kvs)

	//If i want to grab just one value
	fmt.Println(languages["JS"])

	//For deleting
	delete(languages, "RB")

	//looping over maps

	for key, value := range languages {
		fmt.Println("For key %v, value is %v\n", key, value)
	}

}
