package main

import (
	"fmt"
)

func main() {
	var num [5]int = [5]int{23, 43, 12, 32, 4}
	var num2 = [...]int{1, 3, 5}

	fmt.Println(num)
	fmt.Println(num[0])
	fmt.Println(num2)
	fmt.Println(num2[0])

}
