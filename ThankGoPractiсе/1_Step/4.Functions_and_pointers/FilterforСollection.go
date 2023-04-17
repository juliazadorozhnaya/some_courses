package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func filter(predicate func(int) bool, iterable []int) []int {
	var result []int
	for _, n := range iterable {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

func main() {
	src := readInput()
	res := filter(func(n int) bool { return n%2 == 0 }, src)
	fmt.Println(res)
}

func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin) //Сtrl + D
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}
