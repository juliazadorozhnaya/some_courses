package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)
	if len(text) == width || len(text) < width {
		fmt.Println(text)
	} else {
		bytes := []byte(text[:width])
		fmt.Printf("%s...", string(bytes))
	}

}
