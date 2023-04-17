package main

import (
	"fmt"
	"strings"
)

func main() {
	var source string
	var times int
	fmt.Scan(&source, &times)

	result := strings.Repeat(source, times)

	fmt.Println(result)
}
