/*#### For
Go gives us only one looping construct, the `for` loop. This means we don't have a while, or a repeat until or anything like that.
For loops have a basic structure similar to the one used at C, except we don't use parenthesis, they will actually be a compilation error:
```golang
for i := 0; i < 2; i++ {
	// code
}
```
The first and last part of the for can be optional, meaning that we can only use the evaluation part and we basically have a while like so:
```golang
for i != 4 {
	// do some stuff
}
```
But you can also omit everything and you'll have an endless loop.

**range** iterates over elements in a variety of data structures. Let’s see how to use range with some of the data structures. ([GoPlay](https://play.golang.org/p/ChWJFN-Zaoy))

```golang*/

package main

import "fmt"

func main() {

	// Here we use `range` to sum the numbers in a slice.
	// Arrays work like this too.
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// `range` on arrays and slices provides both the
	// index and value for each entry. Above we didn't
	// need the index, so we ignored it with the
	// blank identifier `_`. Sometimes we actually want
	// the indexes though.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// `range` on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` can also iterate over just the keys of a map.
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// `range` on strings iterates over Unicode code
	// points. The first value is the starting byte index
	// of the `rune` and the second the `rune` itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
