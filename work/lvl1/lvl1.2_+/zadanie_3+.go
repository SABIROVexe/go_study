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

	first_num := rune(str[0])
	last_num := rune(str[len(str)-1])

	firstDigit := int(first_num - '0')
	lastDigit := int(last_num - '0')

	sum := firstDigit + lastDigit

	fmt.Printf("Первая цифра: %d\n", firstDigit)
	fmt.Printf("Последняя цифра: %d\n", lastDigit)
	fmt.Printf("Сумма: %d\n", sum)

}
