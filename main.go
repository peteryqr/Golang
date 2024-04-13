package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	/*
		fmt.Println("Hello World")

		// variables
		var name string = "Peter"
		fmt.Println("My name is", name)
		var birthYear = 10
		fmt.Print("My birth year is ", birthYear, "\n")
		month := "January"
		var date int8 = 17
		// int8 limit the number to [-127, 128]
		fmt.Print("Month and date were ", month, " ", date, "\n")
		var gpa float32 = 4.0
		fmt.Printf("The quoted name %q and type of variable %T and regular %v and float %0.1f\n", name, birthYear, month, gpa)
	*/

	// arrays
	var ages [3]int = [3]int{30, 25, 30}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}
	names[0] = "luigi"
	// array size is fixed after declaration
	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	//slices has dynamic length: we could insert and delete
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	//slice ranges
	rangeONe := names[1:3]
	fmt.Println(rangeONe, len(rangeONe))
	// we get the second and third element, not the fourth

	// string package:
	// sort package:
	greeting := "hello there friends!"

	fmt.Println(strings.Contains(greeting, "hello"))
	fmt.Println(strings.Contains(greeting, "bye"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))
	fmt.Println(strings.Split(greeting, " "))

	// the original value is unchanged
	fmt.Println("original string value =", greeting)

	ages1 := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages1)
	fmt.Println(ages1)

	index := sort.SearchInts(ages1, 30)
	fmt.Println(index)

	names1 := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names1)
	fmt.Println(names1)

	fmt.Println(sort.SearchStrings(names1, "bowser"))
}
