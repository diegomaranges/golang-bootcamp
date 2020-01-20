/*### Pointers
Go has pointers. A pointer in Go will hold the memory address of a value. The zero-value of pointers it's `nil`.
To declare a pointer to a type we use a syntax similar to C:
```golang
var somePointer *int
```
In this case `somePointer` is a pointer to an int. To access the value a pointer points to or the memory address of a given variable when can do this([GoPlay](https://goplay.space/#5_eTXtJ1h43)):
```golang*/
package main

import "fmt"

func main() {
	var someInt *int
	otherInt := 2
	pointer := &otherInt
	fmt.Printf("zero-value of someInt: %v\n", someInt)
	fmt.Printf("value pointer points to: %v\n", *pointer)
}
