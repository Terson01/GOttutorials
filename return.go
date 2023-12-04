package main

import "fmt"
func main() {
	fmt.Println(Add(1,2))
	fmt.Println(name("anikwe", "peterson"))

}
func Add(x int, y int)int { return x + y }

//function that concatenates two strings

func name(firstName string, secondName string) string {
	fullName := firstName+ " " + secondName
	return fullName
}