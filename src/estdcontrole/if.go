package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "Este número é Par")
		} else {
			fmt.Println(i, "Este número é ímpar")
		}

	}

}

//Se o número é par ou ímpar
//Resto é representado por: %
