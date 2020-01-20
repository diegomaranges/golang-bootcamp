/*### Variables and types
Go has almost all the typical type values:
```
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```
This types can be manipulated with different operators. Strings can be added together with `+` as can integers and floats.
Booleans have all the boolean operators as expected.

This types are used when declaring variables. In Go you can declare variables in three different ways that will be shown in the code below.
Go also lets you declare constants by preceding the keyword `const` before the name. The type of the constant gets inferred.
[Run it online](https://goplay.space/#cmQu-3Uf58J)
```golang*/
package main

import "fmt"

func main() {
	// Declare with no initialization, go will give the default value
	// which in this case is 0.
	var someInt int
	// Giving it a specific initial value
	var someBool, hi = true, "hi"
	// Inferring the type from the right side of the expression
	hello := "hello world" // hello will be of type string

	// declaring constants. Type gets inferred
	const number = 2
	const str = "some string"

	fmt.Printf("someInt: %v\n", someInt)
	fmt.Printf("someBool: %v\nhi: %v\n", someBool, hi)
	fmt.Printf("helloWorld: %v\n", hello)
	fmt.Printf("number: %v\n", number)
	fmt.Printf("str: %v\n", str)
}
