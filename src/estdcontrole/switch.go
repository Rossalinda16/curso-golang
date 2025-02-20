package main 

import "fmt"

func main() {

	for i := 0; i <=10; i++ {

		switch i{
		case  0: fmt.Println("Cero")
		case  1: fmt.Println("Un")
		case  2: fmt.Println("Dos")
		case  3: fmt.Println("Tres")
		case  4: fmt.Println("Cuatro")
		case  5: fmt.Println("Cinco")

		}
		
	}
}

//O switch em Go é uma estrutura de controle usada para comparar uma variável contra múltiplos valores de maneira mais limpa e organizada do que usar vários if-else.