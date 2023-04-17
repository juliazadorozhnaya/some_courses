package main

import (
	"fmt"
	"time"
)

//Sample Input:
//
//1h30m
//Sample Output:
//
//1h30m = 90 min

func main() {

	var s string
	fmt.Scan(&s)
	d, _ := time.ParseDuration(s)
	fmt.Printf("%s = %v min\n", s, d.Minutes())
}
