//Operador Curto

// = é utilizado para atribuir valores a variaveis já existentes

// Gopher := é utilizado para criar novas variaveis, dentro de code blocks. O operador := só pode ser usado dentro de funções, não é permitido no nível de pacote.

//O codigo abaixo, é o mesmo do arquivo temperatura.go, uso ele aqui pra exemplificar.

// Declara meu pacote principal
package main

//Importar a função fmt
import "fmt"

//Declarção da variável CONST do ponto de ebulição da água em F.
const ebulicaoF = 212.0

//Função principal
func main() {

	F := ebulicaoF        //Valor da temperatura em graus F. CONST
	C := (F - 32) * 5 / 9 //Converção de graus F para graus C
	fmt.Println("Ponto de ebulição da água em F é:", F)
	fmt.Println("Ponto de ebulição da água em C é:", C)
}
