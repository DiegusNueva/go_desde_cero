package main

import "fmt"

func main() {
	// var flags [3]string
	// flags[0] = "🇻🇮"
	// flags[1] = "🇻🇬"
	// flags[2] = "🇻🇪"

	flags := [3]string{"🇼🇸", "🇻🇺", "🇺🇿"}
	flags2 := [...]string{"🇼🇸", "🇻🇺", "🇺🇿", "🇼🇫"}

	fmt.Println(flags, flags2)
}
