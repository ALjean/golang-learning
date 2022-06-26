package main

import "fmt"

func isArraySorted(arr []string) bool {
	for i := range arr {
		if arr[i+1] < arr[i] {
			return false
		}
	}

	return true
}

func main() {
	words := []string{"Rabbit", "Fish", "Dog", "Parrot", "Cat", "Hamster"}
	fmt.Println(isArraySorted(words))
}
