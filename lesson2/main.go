package main

import "fmt"

func HasDuplicate(nums []int) bool {
	numMap := make(map[int]int)

	for _, v := range nums {
		if numMap[v] != 0 {
			return true
		} else {
			numMap[v] = v
		}
	}
	return false
}

func main() {
	nums := []int{10, 20, 30, 40, 10}
	fmt.Println(HasDuplicate(nums))
}
