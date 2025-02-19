package main

import "fmt"

func main() {

	x := 2

	if x == 2 || x == 3 {
		fmt.Println("Sim, x é 2 ou 3!")
	} else {
		fmt.Println("Não, x não é 2 nem 3!")
	}

}

/* Tabela importante de Operadores logicos. 

AND lógico (&&) → Retorna true se ambos os operandos forem true, caso contrário, retorna false.

true && true // true 
true && false // false 
false && true  // false
false && // false

OR lógico (||) → Retorna true se pelo menos um dos operandos for true.

true || true // true 
true || false // true 
false || true  // true 
false || false // false

NOT lógico (!) → Inverte o valor booleano.

!true  // false
!false // true */

