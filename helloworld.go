package main

import (
	"fmt"
)

func printMyInfo(name string, age int) (result int) {
	fmt.Printf("Hi, my name is %v and I am %v years old!\n", name, age)
	result = 444
	return
}

func main() {
	fmt.Println("Hello World!") // semi-colon optional

	var firstName string = "John" // Explicit typing
	var lastName = "Doe"          // Inferred type
	x := 2                        // inferred short-hand

	fmt.Println(firstName + " " + lastName)
	fmt.Println("Student ID:", x)

	var a string
	var b int
	var c bool

	// uninitialized variables are set to defaults
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	const PI = 3.14

	fmt.Printf("PI has value %v and type %T\n", PI, PI)

	var arr1 = [3]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(len(arr1))

	mySlice := arr1[1:3]
	mySlice = append(mySlice, 4)
	fmt.Println(mySlice)

	var mySlice2 = []int{5, 6, 7}
	mySlice = append(mySlice, mySlice2...)
	fmt.Println(mySlice)

	for idx, val := range mySlice {
		fmt.Println(val, "is at index", idx)
	}

	var myName = "Francoise"
	var myAge = 26
	fmt.Println(printMyInfo(myName, myAge))

	type Person struct {
		name string
		age  int
		id   int
	}

	var pers1 Person
	pers1.name = "Francoise"
	pers1.age = 26
	pers1.id = 44

	fmt.Println(pers1)

	var myMap = make(map[string]string)
	myMap["Brand"] = "Ford"
	myMap["Model"] = "Mustang"
	myMap["Year"] = "1964"

	fmt.Printf("myMap\t%v\n", myMap)
	var isYearPresent bool
	_, isYearPresent = myMap["Year"]
	fmt.Println(isYearPresent)

	delete(myMap, "Year")

	fmt.Printf("myMap\t%v\n", myMap)
	_, isYearPresent = myMap["Year"]
	fmt.Println(isYearPresent)
}
