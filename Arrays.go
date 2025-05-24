// arrays
// in go there are two ways to declare an array
// 1 with var keyword
// syntax
// var array_name = [lenght] datatype{values}//definig length
// or
// varr arrayname = [...] datatype{values}// lenght inferred
// 2 with the := sign
// syntax
// array_name := [lenght]datatype{values}
// or
// array_name := [...]datatype{values}// the compiler desides the lenght of the array
package main

import (
	"fmt"
)

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{1, 2, 4, 5, 3}
	fmt.Println(arr1)
	fmt.Println(arr2)

	var numbers = [...]int{3, 4, 5, 6, 7}
	numbers2 := [...]int{4, 3, 2, 1, 1}
	fmt.Println(numbers)
	fmt.Println(numbers2)

	var cars = [4]string{"Mazda", "Benz", "Bmw", "Harreir"}
	fmt.Println((cars))
	names := [...]string{"Moha", "Dennis", "james", "Alison"}
	fmt.Println((names))

	// Accsessing elements of an Array

	// access we use index number starting at [0]
	fmt.Println(cars[0])
	fmt.Println(numbers2[3])
	fmt.Println(arr1[2])

	// Changing Elements of an Array
	// we also use index to change values at the specific index

	cars[3] = "Nissan"
	fmt.Println(cars)

	numbers[0] = 4
	fmt.Println(numbers)

	// Array initialization
	//If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

	//Tip: The default value for int is 0, and the default value for string is "".
	fmt.Println("Initialization")
	dog := [5]int{}
	his := [4]int{6, 4}
	has := [3]int{1, 2, 3}

	fmt.Println(dog)
	fmt.Println(his)
	fmt.Println(has)

	// specific initialization of an array
	fmt.Println("Specific Initializatinon")

	bags := [5]int{1: 10, 2: 9}
	fmt.Println(bags)
	//1:10 means: assign 10 to array index 1 (second element).
	//2:40 means: assign 40 to array index 2 (third element).

	// Finding the lenght of an Array
	fmt.Println("Using len()")
	// The len() function is used to find the length of an array:
	fmt.Println(len(cars))
	fmt.Println(len(arr1))

}
