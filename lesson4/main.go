package main

import "fmt"

func main() {
	word := "съешь ещё этих мягких французских булок, да выпей чаю"

	charMap := make(map[int32]int)

	for _, v := range word {
		charMap[v] = charMap[v] + 1
	}

	fmt.Println(charMap)
}
