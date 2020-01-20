/*#### If
If statements, like for, does not used parenthesis. We can start if statements with a statement to execute before the condition, for example:
```golang
if err := funcReturnsError(); err != nil {
	// very important stuff
}
```
We also have an `else` and can concatenate it with another if like `} else if <condition> {`.

#### Switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression. Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. This means we don't need that ugly break statement at the end of each case like in the mentioned languages. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers. For example([GoPlay](https://goplay.space/#SIsdHJKlgxe))
```golang*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}
