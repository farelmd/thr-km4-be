package main

import "fmt"

func howManyElements(data []any) int {

	return len(data)
}

func main() {

	slc := []any{1, 2, "heh"}
	fmt.Println(howManyElements(slc))
}
