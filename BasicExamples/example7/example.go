/*#### Defer
Defer is something that gets used a lot in Go. This statements defers the execution of the function until the surrounding function returns. The arguments that it receives are evaluated immediately but the function does not get executed until the surrounding function returns.
Deferred function calls get pushed into a call stack. When a function returns the go runtime will pop each of the deferred functions and execute them(it's a LIFO structure). Example([GoPlay](https://goplay.space/#6hzIT6lIo5F)):
```golang*/
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
