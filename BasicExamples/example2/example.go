/*#### Exporting names
If you've used Java then you probably already had fun with those endless method definitions, I'm talking about you `public static void main`.
Go has no *specific* access modifiers like public or private. In Go, names are either package-level or public. And you do this by making the first letter of the name be a capital letter. When you import a package you can only refer to those names that start with a capital letter. For example([GoPlay](https://goplay.space/#CrFk7n7AWd6)):
```golang*/
package main

import (
	"fmt"
	"math"
)

func main() {
	// Run this and you will get an error. After that
	// change the pi to be Pi and see what happens.
	fmt.Println(math.Pi)

}
