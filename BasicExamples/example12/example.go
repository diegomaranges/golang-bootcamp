/*### Structs
To group fields together Go defines `structs`, this structs are quite similar to those defined in the C language. They are defined with `type name struct`, each field it's going to have its own type and they can be initilized in different ways. For example([GoPlay](https://goplay.space/#x4njww69awl)):
```golang*/
package main

import "fmt"

type someStruct struct {
	intField  int
	boolField bool
}

func main() {
	// this will have the zero-value for all its fields
	s := someStruct{}
	// fields can get initialized by order
	s = someStruct{2, true}
	// we can name the fields we are initializing
	l := someStruct{
		boolField: false,
		intField:  2,
	}

	// We can then access the fields of the struct using the
	// dot notation:
	fmt.Printf("l bool field: %v\n", l.boolField)
	fmt.Printf("s int field field: %v\n", s.intField)
} /*
```
If you declare a pointer to a struct one could think that in order to access the fields we would have to use the pointer notation. But Go is nice and we don't have to type those 3 extra characters, we can simply do the following([GoPlay](https://goplay.space/#wKurxTznpp5)):

__NOTE:__ notice that in the above example the struct attributes are private because the first letter is in lower case.

```golang
package main

type someStruct struct {
	intField int
	boolField bool
}

func main() {
	p := &someStruct{2, false}
	fmt.Printf("doing p.intField works: %v\n", p.intField)
	fmt.Printf("but also we can do (*p).boolField: %v\n", (*p).boolField)
}
```
Anonymous structs are something that Go gives us that comes in quite handy when we are testing. Basically we can group together fields without specifying a type for it. For example([GoPlay](https://goplay.space/#H8CneBlRDfh)):
```golang
package main

func main() {
	// see that it has no name that can be used to reference
	// the "type" of this structure
	s := struct {
		x int
		y int
	}{
		x: 2,
		y: 3,
   	}
	fmt.Printf("x: %v y: %v", s.x, s.y)
}*/
