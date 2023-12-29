package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial((num - 1))
}

// recursion will help us to solve the factorial problem and other problem
// func factorial(num int) int {
// 	result := 1

// 	for i := 1; i <= num; i++ {
// 		result = result * i
// 	}

// 	return result
// }
