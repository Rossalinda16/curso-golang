/*package main

import "fmt"

func main() {

	var nome string = "Rosa Linda"
	var idade int = 25
	var versao float32 = 3.2
	var cidade string = "São Paulo"
	fmt.Println("Meu nome é", nome, "e minha idade é", idade, "anos")
	fmt.Println("Estou usando o programa de versão:", versao)
	fmt.Println("Moro na cidade de:", cidade)

}*/

package main

import "fmt"

func main() {

	var nome = "Rosa Linda"
	var idade = 25
	var versao = 3.2
	var cidade = "São Paulo"
	fmt.Println("Meu nome é", nome, "e minha idade é", idade, "anos")
	fmt.Println("Estou usando o programa de versão:", versao)
	fmt.Println("Moro na cidade de:", cidade)

}

//A inferência de tipos em Go é a capacidade do compilador de determinar o tipo de uma variável com base no valor atribuído a ela, sem que você precise especificar o tipo explicitamente.
