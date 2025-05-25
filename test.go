package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var student1 string = "John"
	var student2 = "jane"
	x := 2
	name := "Muxsin Yusuf"
	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)
	fmt.Println(name)

	var (
		r int

		t int = 3

		e string = "Hello"
	)
	fmt.Println(r)
	fmt.Println(t)
	fmt.Println(e)

	// var a string
	// var b int
	// var c bool
	// fmt.Println((a))
	// fmt.Println(b)
	// fmt.Println(c)

	var greet string
	greet = "Hello"
	fmt.Println(greet)

	y := 4
	fmt.Println(y)

	boss := "Nice one"
	fmt.Println(boss)

	k := 7
	j := 10
	z := k + j
	fmt.Println(z)

	// Multiple varaiable declaration
	var m, n, b, v int = 1, 3, 4, 2
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(b)
	fmt.Println(v)

	// declaring multiple variable with different data type
	var c, l = 8, "Dog"

	h, g := 7, "Cow"
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(c)
	fmt.Println(l)

	// constants
	//If a variable should have a fixed value that cannot be changed, you can use the const keyword.
	// it means it is unchangable
	// syntaxt = const CONSTNAME type = value
	const PI = 3.14
	fmt.Println(PI)

	q := 10
	u := 2
	fmt.Println(q / u)

	/// ARAYS
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

	//slices
	// syntax slice_name := []datatype{values}
	var scores = []int{13, 34, 54, 65}
	fmt.Println(scores)
	scores[1] = 100
	fmt.Println((scores))
	scores = append(scores, 87)
	fmt.Println(scores)

	//slice ranges
	rangeOne := scores[1:]
	fmt.Println(rangeOne)

}

// Variables
// int stores integers
// float32 stores floating numbers
//string stores texts
// bool true or false
// there are two ways to declare variables
//1 var variablename type = value
//2 with the := sighn
// variablename := value
