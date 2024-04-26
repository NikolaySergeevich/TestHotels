package main

import "fmt"

func main() {
	fmt.Println(find([]int{42, 12, 18}))
}

func find(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	// Находим наименьшее число из массива
	minNum := nums[0]
	for _, num := range nums[1:] {
		if num < minNum {
			minNum = num
		}
	}

	// Находим общие делители
	comDiv := make([]int, 0)
	for i := 1; i <= minNum; i++ {
		isComDiv := true
		for _, num := range nums {
			if num%i != 0 {
				isComDiv = false
				break
			}
		}
		if isComDiv {
			comDiv = append(comDiv, i)
		}
	}

	return comDiv
}