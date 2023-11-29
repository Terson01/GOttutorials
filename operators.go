package main

import (
	f "fmt"
)

func familyName(name string, age int, DOB int, occupation string){
	fmt.Println("Hello", name, "you are", age, "years old and you are born in", DOB,"and you are a" ,occupation)
}
func main(){
	familymembers:=[]string{}
	familyages:=[]int{}
	familyoccupation:=[]string{}
	for{
    var name string
    var age int
    var DOB int
	var occupation string
    
    f.Println("Enter your name: ")
    f.Scan(&name)
    
    f.Println("Enter your age: ")
    f.Scan(&age)
    
    f.Println("Enter your date of birth: ")
    f.Scan(&DOB)
   
	fmt.Println("Enter your Occupation: ")
    fmt.Scan(&occupation)
    
    familyName(name, age, DOB, occupation)

	familymembers=append(familymembers, name)
	familyages=append(familyages, age)
	familyoccupation=append(familyoccupation, occupation)

	fmt.Printf("list of the family members %v\n", familymembers)
	fmt.Printf("list of the family ages %v\n", familyages)
	fmt.Printf("list of the family occupation %v\n", familyoccupation)
	
	}
}