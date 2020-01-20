/*### Packages
In Go, programs are structured using packages. The main function of a package should be in `package main`. If you create a project that does not have a package main then it means you are creating a library. Go's convetion says that the name of the package should match the last element of the import path. An import path is basically the path to the project from within your Go workspace.
Now that we are talking about imports, go gives us the statement `import "<path>"` to import a specific package either from the standard library or some package in our go workspace. Example([GoPlay](https://goplay.space/#bB6DC_CV-bF)):
```golang*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Make the zero value useful.")

	fmt.Printf("Square root of 8: %v\n", math.Sqrt(8))
}