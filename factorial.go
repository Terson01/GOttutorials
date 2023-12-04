package main

import "fmt"

func factorial(n int) int {
	f:=n 
	for i := 1; i<5; i++{
		f*=(n-1)
	}	
	return f }

func main() {
	num:= 0
	fmt.Println("enter your value")
	fmt.Scan(&num)
	fmt.Println(factorial(num))
}

