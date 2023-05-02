package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {

	if position == "up" {
		newSlice := []int{newData}

		return append(newSlice, data...)
	}

	return append(data, newData)
}

func main() {

	data := []int{1, 2, 3}
	fmt.Println(AddElement(data, 4, "up"))
	fmt.Println(AddElement(data, 4, "down"))

}
