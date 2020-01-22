package main

import "fmt"

func main() {
	var option string
	var imput string
	var sliceWithElements []string

	for {
		fmt.Println(sliceWithElements)

		lengt, err := fmt.Scanf("%s", &option)

		if lengt == 1 && err == nil {
			_, err := fmt.Scanf("%s\n", &imput)

			if err == nil {
				fmt.Println(imput)
				sliceWithElements = append(sliceWithElements, imput)
			}

		} else {
			fmt.Println(err)
			fmt.Println(lengt)
		}
	}
}
