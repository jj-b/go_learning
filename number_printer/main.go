package main

import"fmt"

func main() {
	//initialise an array
	intArray := [100]int{}
	//Use a for loop to populate the array
	for count := 0; count <= 99; count++ {
		intArray[count] = count+1
	}
	//For debugging purposes, print the array.
	fmt.Println(intArray)
/*Iterate over the array and evaluate each number.
If the number is not divisible by 3, 5, or 3 and 5, print the number.
If it meets one of the above conditions print the appropriate string*/
	for count := 0; count <= 99; count++ {
	if intArray[count] % 3 == 0 && intArray[count] % 5 == 0{
fmt.Println("foo", "bar")
} else if intArray[count] % 3 == 0{
fmt.Println("foo")
} else if intArray[count] % 5 == 0{
fmt.Println("bar")
} else {
fmt.Println(intArray.Get[count])
}
//Close loop
	}
}
