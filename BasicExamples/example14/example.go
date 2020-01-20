/*### Arrays, slices and maps
In Go there is something that is often confusing to newcomers and that is the difference between arrays and slices.
Arrays are like you would think, a fixed size list of indexed values that all share the same type.
Slices are something a bit more interesting. The first difference would be that arrays have a fixed size, while slices are dynamically allocated. In practice you will probably always see slices instead of arrays.
Something that is worth noting is that slices themselves don't hold the array of the data. They are simply `structs` that contain three fields: Len, Cap and Data. Len is going to be the length of the array, Cap is the maximum capacity of the array and Data is a pointer to the backing array. [Here](https://golang.org/pkg/reflect/#SliceHeader) you can see the struct of the slice.
Slices and arrays are initializaed differently and they let us do different type of operations. Below is a code that will explain all this in more detail([GoPlay](https://goplay.space/#NPh97D1qgEY)):
```golang*/
package main

import "fmt"

func main() {
	// This is how you would declare an array
	// var someArray [3]int
	// Slices can be declared in many different ways.
	// This slice will have its zero value that is going to be nil, since
	// no array has been alocated yet.
	// var names []string
	// slice of strings with an initial size of 2 and unlimited capacity
	// otherNames := make([]string, 2)
	// slice of strings with an initial size of 2 and maximum capacity of 4
	// capacity := make([]string, 2, 4)
	// We can initialize with values
	numbers := []int{1, 2, 3, 4, 5}
	// Slices let us do operations using the indices
	oneToThree := numbers[0:2]
	fmt.Println(oneToThree)
	// We can omit one of the indices and it will go to the last or the first
	threeToFive := numbers[2:]
	fmt.Println(threeToFive)
	fmt.Println(numbers)

	// Incrementing the size of slice.
	// If we want to append values to already declared slice we can use the append function
	oneToThree = append(oneToThree, 4)
	fmt.Println(oneToThree)
	// append doesn't care about the receiver, so we could declare a new variable with append:
	fiveToSix := append(threeToFive[len(threeToFive)-1:], 6)
	fmt.Println(fiveToSix)
}

/*```
To learn more about the usage of slices and how they work internally refer to [this blog post](https://blog.golang.org/go-slices-usage-and-internals).
**Excercise:** Go to [Go Tour slices excercise](https://tour.golang.org/moretypes/18) and implement what it's requested there.

Maps are basically like dictonaries in python or HashMaps in java. They map a key to a given value and let us access those values using the specified keys. The zero value of a map is going to be nil as with slices. Both keys and values can be of any given type, from structs to basic types to interfaces(which we will cover later). Much like slices, we can use `make` to create maps. Examples([GoPlay](https://goplay.space/#xjeZ48dknCc)):
```golang
package main

import "fmt"

type someStruct struct {
	intField  int
	boolField bool
}

func main() {
	// this a nil map since it's the zero-value
	// var m map[string]someStruct
	// declaring and initializing maps
	m := map[string]someStruct{
		"key1": someStruct{2, false},
		"key2": someStruct{3, true},
	}
	// We can reference the fields of the struct by accessing the value of the map
	fmt.Printf("m[\"key1\"].intField=%v\n", m["key1"].intField)
	// We can allocate new key-value pairs
	m["key3"] = someStruct{4, true}
	fmt.Printf("m[\"key3\"].boolField=%v\n", m["key3"].boolField)
}
```
**Exercise:** go to [Go tour maps excercise](https://tour.golang.org/moretypes/23) and implement what it's requested.*/
