package main

import "fmt"

/*
Даны два числа. Определить цифры, входящие в запись как первого,
так и второго числа.

Входные данные
Программа получает на вход два числа. Гарантируется, что цифры в числах
не повторяются. Числа в пределах от 0 до 10000.

Выходные данные
Программа должна вывести цифры, которые имеются в обоих числах, через пробел.
Цифры выводятся в порядке их нахождения в первом числе.

Sample Input:
564 8954

Sample Output:
5 4
*/

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	for _, a1 := range a {
		for _, b1 := range b {
			if a1 == b1 {
				fmt.Print(string(a1) + " ")
			}
		}
	}
}
