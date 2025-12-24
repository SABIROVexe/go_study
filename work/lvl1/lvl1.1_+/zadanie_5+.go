// Даны два слова. Проверьте, что первые буквы этих слов совпадают.

package main

import "fmt"

func main() {
	var str1 string
	var str2 string

	fmt.Print("str1:> ")
	fmt.Scanln(&str1)

	fmt.Print("str2:> ")
	fmt.Scanln(&str2)

	runs1 := []rune(str1)
	runs2 := []rune(str2)

	if runs1[0] == runs2[0] {
		fmt.Println("первые буквы слов совпадают")
		fmt.Println("str1:", runs1)
		fmt.Println("str2:", runs2)
	} else {
		fmt.Println("первые буквы слов не совпадают")
	}
}
