package main

import "fmt"

func main() {
	fmt.Println("Мое портфолио Go-разработчика")

	skills := []string{
		"Go (основы)",
		"PostgreSQL (basic queries)",
		"REST API(basic queries)",
		"HTTP протокол",
		"Docker basics",
		"Git и GitHub",
	}

	fmt.Println("\nMy skills:")
	for i, skill := range skills {
		fmt.Printf("%d. %s\n", i+1, skill)
	}
}
