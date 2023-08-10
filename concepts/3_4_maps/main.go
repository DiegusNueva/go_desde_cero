package main

import "fmt"

func main() {
	music := make(map[string]string)
	music["guitarra"] = "🎸"
	music["violin"] = "🎻"

	fmt.Println(music)

	tech := map[string]string{
		"computer": "💻",
		"mouse":    "🖱️",
	}

	delete(tech, "computer")
	fmt.Println(tech)

	fmt.Println(music["fake"])

	content, ok := music["guitarra"]

	fmt.Println(content, ok)

}
