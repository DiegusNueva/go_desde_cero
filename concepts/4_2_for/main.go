package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {

	// 	fmt.Println(i)

	// }

	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }

	// i := 1
	// for {
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// 	i++
	// }

	// for i, v := range []string{"🍕", "🍔", "🍎", "🌭"} {
	// 	fmt.Println("indice:", i, "valor:", v)
	// }

	// numbers := []uint8{2, 4, 6, 8}

	// for i := range numbers {
	// 	numbers[i] *= 2
	// }

	// fmt.Println(numbers)

	// food := map[string]string{
	// 	"pizza":     "🍕",
	// 	"hamburger": "🍔",
	// 	"apple":     "🍎",
	// 	"hotdog":    "🌭",
	// }

	// for key, value := range food {
	// 	fmt.Println("key:", key, "value:", value)
	// }

	for i, v := range "EDteam" {
		fmt.Println("indice:", i, "valor:", string(v))
	}

}
