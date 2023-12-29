package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// userNames := []string{} // create an empty array
	// userNames[0] = "Max" // this will throw an error because the array is empty, it has no slots

	// Make create an array with 2 slots with null values
	// and now this is value -> userNames[0] = "Max"
	userNames := make([]string, 2, 5) // create an empty array with a length of 2

	userNames[0] = "Max"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRating := make(floatMap, 3)

	courseRating["go"] = 4.7
	courseRating["react"] = 4.8
	courseRating["angular"] = 4.7

	courseRating.output()

	// for arrays, slices
	for idx, value := range userNames {
		fmt.Println(idx, value)
	}

	// for maps
	for key, value := range courseRating {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
