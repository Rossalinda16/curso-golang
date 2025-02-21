//Exercicio de For - Break e continue

package main

import "fmt"

func main() {

	x := 0

	for {
			if x < 5 {
				x++
				fmt.Println("x < 5")
			} else {
					break
			}
		}
	}


/*package main

import "fmt"

func main() {
	for x := 0; //inicialiação
	x <= 10; x++ { //condição
		fmt.Println(x) // pós - saida
	}

}
Em Go, o único loop existente é o for. Diferente de outras linguagens que possuem while e do-while, o for em Go pode ser usado de diferentes formas para cobrir todas as necessidades de repetição.

Existem varios formatos de usar o For */
