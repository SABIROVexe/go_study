//Дана строка. Если в этой строке более одного символа, выведите в консоль предпоследний символ этой строки.

package main

import "fmt"

func main() {
	var str string

	fmt.Print("input str:> ")
	fmt.Scanln(&str)

	if len(str) > 1 {
		fmt.Printf("%c", str[len(str)-2])
	} else {
		fmt.Println("error")
	}
}
