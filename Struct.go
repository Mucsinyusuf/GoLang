package main

import (
	"fmt"
)

func main() {

	fmt.Println("GoStructs")

	// a struct is a shortform of structures
	// used to create a collection of members of different data types into a single variable
	// syntax
	// type struct_name struct{
	// memeber1 datatype;
	// member2 datatype;
	//member3 datatype;
	//member4 datatype;
	//}

	// example declare a struct named person
	type person struct {
		name   string
		age    int
		job    string
		salary int
	}

	var pers1 person
	var pers2 person

	pers1.name = "Abdu"
	pers1.age = 23
	pers1.salary = 25000
	pers1.job = "Software dev"

	pers2.name = "kali"
	pers2.age = 34
	pers2.job = "Teacher"
	pers2.salary = 3400

	fmt.Println("Name: ", pers1.name)
	fmt.Println("age: ", pers1.age)
	fmt.Println("job: ", pers1.job)
	fmt.Println("salary: ", pers1.salary)

	fmt.Println("Name: ", pers2.name)
	fmt.Println("age: ", pers2.age)
	fmt.Println("job: ", pers2.job)
	fmt.Println("salary: ", pers2.salary)

	// passing structs as a function argument

}
