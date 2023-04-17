package main

import (
	"fmt"
)

/*
Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:
843

Sample Output:
348
*/

package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%1d%1d%1d", &a, &b, &c)
	fmt.Print(c*100 + b*10 + a)

}
