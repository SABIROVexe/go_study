//Дано число. Проверьте, отрицательное оно или нет. Выведите об этом информацию в консоль.

package main

import "fmt"

func main() {
	var a int

	fmt.Print("Введите число: ")
	fmt.Scan(&a)

	if a < 0 {
		fmt.Print("число ", a, " отрицательное")
	} else {
		fmt.Print("число ", a, " положительное")
	}
}
