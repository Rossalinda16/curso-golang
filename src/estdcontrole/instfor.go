package main //Exercicio 1

import "fmt"

func main() {
	numero := 5

	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)

	}
}

 /* outra forma de usar o for:

func main() {
	for i:= 1; i <=10; i++{
		fmt.Println(i)
	}


}



i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
*/
