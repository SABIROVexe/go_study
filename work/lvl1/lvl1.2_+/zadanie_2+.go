//Дано целое число. Выведите в консоль последнюю цифру этого числа.

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
	runs := []rune(str) // [13, 34, 56, 76]

	for i := 0; i <= len(runs); i++ {
		if i == len(runs) {
			fmt.Printf("%c", runs[i-1])
		}
	}

}
