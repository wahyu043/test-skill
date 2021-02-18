package main

import "fmt"

func main() {
	for i := 0 + 1; i < 8; i++ {
		for j := i; j < 8; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
