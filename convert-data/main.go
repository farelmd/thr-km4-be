package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	splitted := strings.Split(data, ",")
	age, _ := strconv.Atoi(splitted[1])
	var newUser User = User{
		Name:    splitted[0],
		Age:     age,
		Address: splitted[2]}

	return newUser
}

func main() {

	fmt.Println(ConvertData("Edo,20,Jakarta"))
}
