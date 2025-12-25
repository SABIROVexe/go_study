//Дано целое число. Выведите в консоль первую цифру этого числа.

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

	fmt.Print("первая цифра числа: ")
	fmt.Printf("%c", runs[0])

}
