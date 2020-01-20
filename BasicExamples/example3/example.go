/*### Functions
Functions in Go can take any number of arguments and return any number of results. The typical model in Go is to return the results you want plus an error, but we'll cover errors later.
The return values of a function may be named, if so then Go will treat them as local variables to the scope of the function.  Example([GoPlay](https://goplay.space/#xPmMJhUXExA)):
```golang*/
package main

import "fmt"

// declare x and y as named return results
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
