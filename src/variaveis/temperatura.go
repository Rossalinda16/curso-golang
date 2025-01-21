// Declara meu pacote principal
package main

//Importar a função fmt
import "fmt"

//Declarção da variável CONST do ponto de ebulição da água em F.
const ebulicaoF = 212.0

//Função principal
func main() {

	var F = ebulicaoF        //Valor da temperatura em graus F. CONST
	var C = (F - 32) * 5 / 9 //Converção de graus F para graus C
	fmt.Println("Ponto de ebulição da água em F é:", F)
	fmt.Println("Ponto de ebulição da água em C é:", C)
}



