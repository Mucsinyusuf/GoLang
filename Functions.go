package main

import "fmt"

func main() {

	// we call back the function
	myMessage()

	// a function can be called multiple times
	myMessage()

	familyName("Mucsin")
	familyName("Hamdi")
	familyName("Hakima")
	familyName("Farhan")

	person("mucsin", 23)

	fmt.Println(add(4, 6))

	fmt.Println(greet("Lucy", 56))

	fmt.Println(swap("Hello", "World"))

	textCount(1)

	fmt.Println(fuctorial(8))

}

// Functions is go

// to declare a function in go we start with the func the the name of the fucntion
// follow with a parenthesis()
// inside the curly braces we add the codes that defines the functions
// syntaxt
// func FunctionName() {
// code to be executed
// }
func myMessage() {
	fmt.Println("Functions")

	fmt.Println("Am a function executed")

	// parameters and Arguments
	// syntax
	// func FunctionName(param1 type, param2 type, param3 type) {
	// code to be executed

}

func familyName(fname string) {
	fmt.Println("Hello", fname, "yussuf")
	// Note: When a parameter is passed to the function,
	// it is called an argument. So,
	// from the example above: fname is a parameter,
	// while Liam, Jenny and Anja are arguments.
}

// function with multiple parameteres

func person(name string, age int) {
	fmt.Println("hello", name, "you are now", age)
}

// function Returns

// to return a value from a function , we define the data type of the return value,
// also use return keyword inside the function
// syntax
// func FunctionName(param1 type, param2 type) type {
// code to be executed
//   return output
// }

func add(x int, y int) int {
	return x + y

}

// returning multiple values

func greet(name string, age int) (string, int) {
	return name, age

}

func swap(j string, k string) (text1 string, text2 string) {
	text1 = j

	text2 = k

	return k, j

}

// Recursion fnction
// a function that calls itself

func textCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return textCount(x + 1)

}

func fuctorial(x float64) (y float64) {
	if x > 0 {
		y = x * fuctorial(x-1)

	} else {
		y = 1
	}
	return

}
