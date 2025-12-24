//Дана строка. Выведите в консоль последний символ строки.

package main

import "fmt"

func main() {
	var str string = "aboba"

	runs := []rune(str) // [a, b, o, b, a]

	for i := 0; i <= len(str); i++ {
		if i == len(str) {
			fmt.Printf("%c", runs[i-1])
		}
	}
}
