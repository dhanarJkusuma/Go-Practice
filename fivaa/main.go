package main

import "fmt"

func main() {
	fmt.Println("hello")
	fivaa(5)
}

func fivaa(input int) {
	for i := input; i > 0; i-- {
		for j := 0; j < i+2; j++ {
			if j < 2 {
				fmt.Print(i - 1)
			} else {
				fmt.Print(i + 1)
			}
		}
		fmt.Print("\n")
	}
}
