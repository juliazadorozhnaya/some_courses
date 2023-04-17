package main

import (
	"fmt"
	"strings"
)

func main() {
	m := map[int]int{
		12: 2,
		1:  5,
	}
	fmt.Println(m)      // map[1:5 12:2]
	fmt.Println(len(m)) // 2
	fmt.Println(m[12])  // 2

	fmt.Println()

	for key, value := range m {
		fmt.Printf("for. key: %v, value: %v\n", key, value)
	}

	fmt.Println()

	var value int
	var hasValue bool

	value, hasValue = m[404]
	fmt.Println(value, hasValue) // 0 false

	value, hasValue = m[1]
	fmt.Println(value, hasValue) // 5 true

	fmt.Println()

	delete(m, 12)
	fmt.Println(m) // map[1:5 12:2]

	fmt.Println()

	if value, ok := m[1]; ok {
		fmt.Println("key: 1, value:", value) // 10
	}

	if value, ok := m[404]; ok {
		fmt.Println("key: 404, value:", value) // Условие не выполняется
	}

	fmt.Println()

	cityPopulation := map[string]int{
		"A_10": 10, "B_10": 15, "C_10": 20,
		"A_100": 100, "B_100": 200, "C_100": 400,
		"A_1000": 1001, "B_1000": 2000, "C_1000": 3000,
	}
	for k := range cityPopulation {
		if strings.HasPrefix(k, "C_") {
			delete(cityPopulation, k)
		}
	}
	fmt.Println(cityPopulation)
	// map[A_10:10 A_100:100 A_1000:1001 B_10:15 B_100:200 B_1000:2000]
}
