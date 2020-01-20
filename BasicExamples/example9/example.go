/*When you declare a variable in Go with no initialization that variable will hold the zero-value of the specified type. For example, in the case of `someInt`, its zero-value will be `0`. For a `bool` the zero-value will be `false`, strings will be `""`.
If you want to know more about why Go uses the format `<name> <type>` for declaring variables and parameters you can checkout [Go's declaration syntax](https://blog.golang.org/gos-declaration-syntax)

Type casting in Go is done by enclosing the variable or value you want to cast with parenthesis, preceded by the type you want to cast it to.
For example([GoPlay](https://goplay.space/#rAUMn61amkI)):
```golang*/
package main

import "fmt"

func main() {
	someInt := 2
	castedFloat := float64(someInt)
	fmt.Printf("type of someint: %T\ntype of castedFloat: %T\n", someInt, castedFloat)
	// will result in:
	// type of someint: int
	// type of castedFloat: float64
}
