package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go Maps")
	// maps are used to store data values in a key:value pair
	// each element in a map is a key:value pair
	// a map is unordered and challangable collection that does not allow duplicates
	// defualt value of a map is nill
	// go has multiple ways of creating maps

	// 1 Creating maps using var and :=
	// syntax
	// var a = map [keytype]valuetype{key1:value1:, key2:value2,...}
	// b := map[keytype]valuetype{key1:value1, key2:value2,...}

	// example
	var a = map[string]string{"brand": "ford", "model": "Mustang", "year": "1994"}
	// b := map[string]int{"Oslo": 1, "Bergen": 2, "trondheim": 3, "Stavanger": 4}

	// fmt.Printf("a\t%v\n", a)
	// fmt.Printf("b\t%v\n", b)

	// 2 Creating MAps using the make() function:

	// syntax
	// var a = make(map[keytype]valuetype)
	// b:= make(map[keytype]valueType)

	// example
	var x = make(map[string]string)

	x["brand"] = "ford"
	x["model"] = "mustang"
	x["year"] = "1964"

	k := make(map[string]int)

	k["Oslo"] = 1
	k["Bergen"] = 2
	k["Trondheim"] = 3
	k["Stavanger"] = 4

	// fmt.Printf("a\t%v\n", a)
	// fmt.Printf("b\t%v\n", b)

	// Accessing Map Elements
	// syntaxt value = map_name[key]

	fmt.Println(a["brand"])

	// update and add map elements
	// syntax
	// map_name[key]=value

	k["Oslo"] = 6
	a["color"] = "red"

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("k\t%v\n", k)

	// remove element form map
	// removing elements is done usmng delet() function

	// syntax
	// delete(map_name,key)
	// example
	delete(a, "color")
	fmt.Printf("a\t%v\n", a)

	// cheking for specific element
	// syntax
	// val, ok :=map_name[key]

	// example

	val1, ok1 := a["brand"]
	fmt.Println(val1, ok1)

	val2, ok2 := a["year"]
	fmt.Println(val2, ok2)

	_, ok3 := a["model"] // only cheking for exsisting key not its value
	fmt.Println(ok3)

	val4, ok4 := a["color"]
	fmt.Println(val4, ok4)

	// Maps are references
	// changes made to two map varriables of the same hash tableaffects all the maps

	// iterate over maps

	for i, v := range a {
		fmt.Printf("%v : %v ", i, v)
	}

}
