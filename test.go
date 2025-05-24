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
