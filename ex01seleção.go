package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Print("Digite o primeiro número:")
	fmt.Scan(&num1)

	fmt.Print("Digite o primeiro número:")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Print("O maior número é", num1)
	} else {
		fmt.Print("O maior número é", num2)

	}
}
