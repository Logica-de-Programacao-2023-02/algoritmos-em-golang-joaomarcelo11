package main

import "fmt"

func main() {
	var n1, n2, n3 int
	fmt.Print("Indique um número")
	fmt.Scan(&n1)
	fmt.Print("Indique um número")
	fmt.Scan(&n2)
	fmt.Print("Indique um número")
	fmt.Scan(&n3)

	if n1 < n2 && n1 < n3 {
		fmt.Print("Seu menor número é:", n1)
	}
	if n2 < n1 && n3 < n2 {
		fmt.Print("Seu menor número é:", n2)
	}
	if n3 < n1 && n3 < n2 {
		fmt.Print("Seu menor número é:", n3)
	}
}
