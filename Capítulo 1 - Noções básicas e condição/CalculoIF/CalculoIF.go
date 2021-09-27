package main

import "fmt"

func main() {
	// Declaração de variaveis.
	var vlr1, vlr2 float64
	var op string

	// Entrada do usuario com os 2 valores.
	fmt.Print("1° Valor: ")
	fmt.Scanln(&vlr1)
	fmt.Print("2° Valor: ")
	fmt.Scanln(&vlr2)
	println("")

	// Entrada do usuario com um dos operadores matemáticos (+, -, *, /).
	fmt.Print(" *  ->  multiplicar\n /  ->  dividir\n +  ->  somar\n -  ->  subtrair\n\nInsira o operador lógico: ")
	fmt.Scanln(&op)

	// Estrutura condicional com if/else para executar o calculo com base no operador matemático escolhido.
	if op == "+" {
		fmt.Print(vlr1, " + ", vlr2, " = ", vlr1+vlr2, "\n")
	} else if op == "-" {
		fmt.Print(vlr1, " - ", vlr2, " = ", vlr1-vlr2, "\n")
	} else if op == "*" {
		fmt.Print(vlr1, " * ", vlr2, " = ", vlr1*vlr2, "\n")
	} else if op == "/" {
		fmt.Print(vlr1, " / ", vlr2, " = ", vlr1/vlr2, "\n")
	} else {
		fmt.Print("Erro! Você digitou o caracter incorreto.\n")
	}
}
