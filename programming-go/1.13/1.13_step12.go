/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных
продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".

Входные данные
Дано число n (0<n<100).

Выходные данные
Программа должна вывести введенное число n и одно из слов (на латинице):
korov, korova или korovy, например, 1 korova, 2 korovy, 5 korov.
Между числом и словом должен стоять ровно один пробел.

Sample Input:
10

Sample Output:
10 korov
*/

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	switch {
	case a%10 == 1 && a%100 != 11:
		fmt.Println(a, "korova")
	case a%10 == 2 && a%100 != 12:
		fallthrough
	case a%10 == 3 && a%100 != 13:
		fallthrough
	case a%10 == 4 && a%100 != 14:
		fmt.Println(a, "korovy")
	default:
		fmt.Println(a, "korov")
	}
}
