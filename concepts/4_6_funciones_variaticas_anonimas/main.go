package main

import "fmt"

func main() {

	// fmt.Println(sum(2))
	// fmt.Println(sum(2, 3))
	// fmt.Println(sum(2, 3, 12))
	// fmt.Println(sum(2, 3, 12, 1, 24))

	// greet := func() {
	// 	fmt.Println("🖐️ Hello")
	// }

	// greet()

	func(name string) {
		fmt.Println("🖐️ Hello", name)
	}("Diego")
}

// func sum(nums ...int) int {
// 	var total int

// 	for _, num := range nums {
// 		total += num
// 	}

// 	return total
// }

// func sum(nums ...int) (total int) {

// 	for _, num := range nums {
// 		total += num
// 	}
// 	return
// }
