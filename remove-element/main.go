package main

import "fmt"

func removeLeftRight(arr []any, position string) []any {
	if position == "left" {
		return arr[1:]
	}

	return arr[:len(arr)-1]
}

func main() {

	arr := []any{1, 2, 3}
	fmt.Println(removeLeftRight(arr, "left")...)
	fmt.Println(removeLeftRight(arr, "right")...)

}
