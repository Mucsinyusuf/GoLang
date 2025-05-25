package main

import (
	"fmt"
)

func main() {
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

}
