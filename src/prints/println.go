package main

import "fmt"

const ebulicaoF = 212.0

func main() {

	var F = ebulicaoF
	var C = (F - 32) * 5 / 9

	//Print Line - Adiciona uma nova linha automaticamente: Cada chamada de fmt.Println insere uma quebra de linha (\n) ao final, movendo o cursor para a próxima linha.
	fmt.Println("Ponto de ebulição da água em F é:", F)
	fmt.Println("Ponto de ebulição da água em C é:", C)
}
