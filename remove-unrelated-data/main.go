package main

import "fmt"

func removeUnrelated(dataMap map[string]any, key string) map[string]any {
	delete(dataMap, key)
	return dataMap
}

func main() {
	m := make(map[string]any)
	m["haha"] = 10
	m["hihi"] = 9
	m["heh"] = 8

	m = removeUnrelated(m, "heh")
	fmt.Println(m)
}
