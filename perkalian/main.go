package main

import "fmt"

func main() {
	fmt.Println("perkalian(0, 2)")
	fmt.Printf("=> %v\n", perkalian(0, 2))

	fmt.Println("perkalian(1, 14)")
	fmt.Printf("=> %v\n", perkalian(1, 14))

	fmt.Println("perkalian(12, 6)")
	fmt.Printf("=> %v\n", perkalian(12, 6))

}

func perkalian(x int, y int) int {
	var result, base, multiply int = 0, 0, 0
	if x > y {
		base = x
		multiply = y
	} else {
		base = y
		multiply = x
	}

	for i := 0; i < multiply; i++ {
		result += base
	}
	return result
}
