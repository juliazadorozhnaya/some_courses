package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)
	counter := make(map[int]int)

	for _, digit := range strconv.Itoa(number) {
		d, _ := strconv.Atoi(string(digit))
		counter[d]++ //добавляем в map элемент значение d, увеличиваем значение ключа на 1
	}

	printCounter(counter)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// printCounter печатает карту в формате
// key1:val1 key2:val2 ... keyN:valN
func printCounter(counter map[int]int) {
	digits := make([]int, 0)
	for d := range counter {
		digits = append(digits, d)
	}
	sort.Ints(digits)
	for idx, digit := range digits {
		fmt.Printf("%d:%d", digit, counter[digit])
		if idx < len(digits)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}
