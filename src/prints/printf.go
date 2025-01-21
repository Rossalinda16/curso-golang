package main

import "fmt"

const ebulicaoF = 212.0

func main() {

	var F = ebulicaoF
	var C = (F - 32) * 5 / 9

	//Print Format - É uma função do pacote fmt usada para imprimir mensagens formatadas no terminal. Ao contrário de fmt.Println, o fmt.Printf permite personalizar a saída utilizando placeholders para formatar variáveis e valores.

	fmt.Printf("Ponto de ebulição da água em F é: %g e o ponto de ebulição da água em C é: %g.", F, C)
}

/* Principais características do fmt.Printf

O fmt.Printf é muito poderoso por permitir personalização avançada da saída. Aqui estão os principais aspectos do uso:

1. Placeholders para formatação
O fmt.Printf utiliza placeholders para especificar como os valores devem ser exibidos:

Placeholder
%s	Exibe uma string.
%d	Exibe um número inteiro (base decimal).
%f	Exibe um número de ponto flutuante com precisão padrão.
%.nf	Exibe um número de ponto flutuante com n casas decimais.
%t	Exibe valores booleanos (true ou false).
%v	Valor padrão para o tipo da variável (útil para depuração).
%T	Exibe o tipo da variável.
%g	Exibe um número de ponto flutuante no formato mais compacto possível.
%%	Exibe o caractere %. */
