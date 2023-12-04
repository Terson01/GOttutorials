package main

import "fmt"

func

func main(){
	n := 5 
	f := n
	
	for i := 1; i<5; i++{
		f*=(n-1)
		fmt.Println(f)
	}
}

