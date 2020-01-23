package main

import "fmt"

func main() {
	//initialise an array
	intArray := [100]int{}
	//Use a for loop to populate the array
	for count := 0; count <= 99; count++ {
		intArray[count] = count+1
	}
	//For debugging purposes, print the array.
	fmt.Println(intArray)

}
