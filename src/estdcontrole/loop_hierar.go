package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d x %d = %d\n", i, j, i*j)
		}
		fmt.Println() // Quebra de linha entre tabelas
	}
}

//Um loop hierárquico (ou loop aninhado) ocorre quando um for está dentro de outro for. Isso é útil para percorrer matrizes, tabelas e resolver problemas que exigem múltiplas iterações.