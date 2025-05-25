package main

import (
	"fmt"
)

func main() {
	fmt.Println("Excercise for printing name, age, location")

	fmt.Println("My name is mucsin am 23 years old and am in nairobi")

	add(4, 8)
	fmt.Println(greet("mucsin", "Hello"))
	evenOrOdd(7)
	count()

	// slice looping excersice
	products := []string{"Gas", "Fruits", "Aniamls"}
	for idx, val := range products {
		fmt.Printf("at %v is %v\n", idx, val)
	}
}

func add(a int, b int) {
	fmt.Println(a + b)

}

func greet(name string, greet string) (string, string) {
	return greet, name

}

func evenOrOdd(x int) {
	if x%2 == 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("Number is odd")
	}
}

func count() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
