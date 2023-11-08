package main

import (
	"fmt"
)

func main() {
	fmt.Println(sumCharCodes("hello world"))
}

func sumCharCodes(n string) int {
	sum := 0
	for k, i := range n {
		fmt.Printf("\n\n%d. %s\n", k, string(i))

		for _, j := range n {
			fmt.Println(string(j))
			sum = sum + int(j)
		}
	}

	return sum
}
