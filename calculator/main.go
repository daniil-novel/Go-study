package main

import "fmt"

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += 2
	}

	fmt.Printf("sum = %d\n", sum)
}
