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

	// Conditions
	fmt.Println("Conditions in Golang")

	// the if statement
	// Syntax
	// if condition { code to be executed }
	fmt.Println("If Conditions")
	x1 := 20
	y1 := 11
	if x1 > y1 {
		fmt.Println("True")
	}

	// if else Condtion
	// syntax
	// if condition {code execution
	//}else {
	//	code execution
	//}
	time := 20
	if time > 12 {
		fmt.Println("greater Time")
	} else {
		fmt.Println("Not Great")
	}

	temp := 14
	if temp > 15 {
		fmt.Println("It is Warm")

	} else {
		fmt.Println("its is cold")
	}

	// the else if statetment
	// syntaxt
	// if conditon1{execute code}
	// else if condition2{execute code blck}
	//else{execute code block}

	temp2 := 9
	if temp2 < 18 {
		fmt.Println("It is hot")
	} else if temp2 > 20 {
		fmt.Println("Moderate temp")
	} else {
		fmt.Println("it is cold")
	}

	// The Nested if statements
	// we can have if statements inside if statements
	// syntax
	// if condition1{
	// code to be executed if condition1 is true
	//if condition2{
	//code to be executed if both conditio1 and 2 are true
	//}
	//}

	num := 20
	if num > 15 {
		fmt.Println("Num is greater than 15")
		if num >= 10 {
			fmt.Println("Num is greater or equal to 10")

		} else {
			fmt.Println("Num is less than 15")
		}

		//Go Switch Statements
		// Syntax
		// switch expression{
		// case x:
		//code block
		// case y:
		// code block
		// case z:
		// code block
		// default
		// code block
		//}

		day := 8
		switch day {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		case 5:
			fmt.Println("Friday")
		case 6:
			fmt.Println("Satyrday")
		case 7:
			fmt.Println("Sunday")
		default:
			fmt.Println("No such Day")

		}
	}

	day1 := 2

	switch day1 {
	case 1, 3, 5:
		fmt.Println("Odd Weekdays")
	case 2, 4:
		fmt.Println("even weekdays")
	case 6, 7:
		fmt.Println("weekend")
	default:
		fmt.Println("No such day")

	}

	// For loops
	fmt.Println("For Loops in Go")
	// the for loop can take up to three statements
	// syntaxt
	// for statement1; statement2; statement3{// code to be executed for each iteration}

	// 	statement1 Initializes the loop counter value.

	// statement2 Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. If it evaluates to FALSE, the loop ends.

	// statement3 Increases the loop counter value.

	// initialization , condition , iteration
	for i := 0; i < 5; i++ {
		// fmt.Println(i)
	}

	for x := 0; x <= 100; x += 10 {
		// fmt.Println(x)
	}

	// continue statements in loops
	// continue statement is used to skip on or mor iterations

	for y := 0; y <= 10; y++ {
		if y == 3 {
			continue

		}
		fmt.Println(y)
	}

	// Break statemnt in loops
	// breaks th eloop on that condition
	for k := 0; k < 7; k++ {
		if k == 5 {
			break
		}
		fmt.Println(k)
	}

	// Nested for loops
	// we place a loop inside another loop
	// the inner loop gets excecuted for each iteration of the outher loop

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"orange", "mango", "apple"}

	for j := 0; j < len(adj); j++ {
		for k := 0; k < len(fruits); k++ {
			fmt.Println(adj[j], fruits[k])
		}
	}

	// range key word
	// range keyword returnd both the index and the value
	// syntax
	// for index, value := range array|slice|map{// code to be executed}

	names2 := []string{"Muhsin", "Yusuf", "Hassan"}

	//Range
	for index, value := range names2 {
		fmt.Printf("the value at index %v is %v\n", index, value)
	}
	// wiyhout index we use _
	for _, value := range names2 {
		fmt.Printf("the value is %v\n", value)
	}

	// Verb	Description
	// %v	Prints the value in the default format
	// %#v	Prints the value in Go-syntax format
	// %T	Prints the type of the value
	// %%	Prints the % sign

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
