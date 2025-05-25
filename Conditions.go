package main

import (
	"fmt"
)

func main() {
	// Conditions
	fmt.Println("Conditions in Golang")

	// the if statement
	// Syntax
	// if condition { code to be executed }
	fmt.Println("If Conditions")
	x := 20
	y := 11
	if x > y {
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

}
