//Дано целое число. Выведите количество цифр в этом числе.

package main

import (
	"fmt"
	"strconv"
)

func main() {

	var number int

	fmt.Println("input number:> ")
	fmt.Scanln(&number)

	str := strconv.Itoa(number)

	x := len(str)

	fmt.Println(x)
}
