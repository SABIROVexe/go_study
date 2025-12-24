// Дано число. Проверьте, четное оно или нет.

package main

import "fmt"

func main() {

	var number int

	fmt.Println("ПРОВЕРКА ЧЕТНОСТИ")
	fmt.Print("input number:> ")

	fmt.Scanln(&number)

	c := number % 2

	fmt.Println("--------------------")

	if c == 0 {
		fmt.Println("число:", number, "четное")
	} else {
		fmt.Println("число:", number, "не четное")
	}

}
