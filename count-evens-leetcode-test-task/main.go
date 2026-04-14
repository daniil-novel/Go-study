package main

import "fmt"

func countEven1(nums []int) {
	even_count := 0
	odd_count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] % 2 == 0 {
			even_count++
		} else {
			odd_count++
		}
	}

	fmt.Println("Even:", even_count, "Odd: ", odd_count)
}

func countEven2(nums []int) {
	even_count := 0
	odd_count := 0
	for _, num := range(nums) { 
		if num % 2 == 0 {
			even_count++
		} else {
			odd_count++
		}
	}

	fmt.Println("Even:", even_count, "Odd: ", odd_count)
}
func main() {
	nums := []int{1, 2, 3, 4, 6}
	countEven1(nums)
	countEven2(nums)
}
