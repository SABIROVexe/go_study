//Даны два целых числа. Проверьте, что первые цифры этих чисел совпадают.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var first_num int
	var second_num int

	fmt.Print("intput first_num:> ")
	fmt.Scanln(&first_num)

	fmt.Print("intput second_num:> ")
	fmt.Scanln(&second_num)

	str1 := strconv.Itoa(first_num)
	str2 := strconv.Itoa(second_num)

	if str1[0] == str2[0] {
		fmt.Println("первые цифры чисел совпадают")
	} else {
		fmt.Println("первые цифры чисел не совпадают")
	}

}
