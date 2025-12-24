//Дано целое число. Выведите в консоль сумму первой и последней цифры этого числа.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var number int
	fmt.Print("input number:> ")
	fmt.Scanln(&number)

	str := strconv.Itoa(number)
	runs := []rune(str)

	number_first := runs[0]
	number_second := 0

	for i := 0; i <= len(runs); i++ {
		if i == len(runs) {
			number_second = runs[i-1]
		}
	}

	fmt.Println(number_first, number_second)
}
