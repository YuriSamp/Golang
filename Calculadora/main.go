package main

import (
	"fmt"
)

func main() {
	var num1, num2, resposta int

	menu()
	fmt.Scanln(&resposta)
	switch resposta {
	case 1:
		fmt.Println("a operação escolhida foi a 1")
		num1, num2 = reconheceNumero()
		var soma = num1 + num2
		fmt.Println(soma)
	case 2:
		fmt.Println("a operação escolhida foi a 2")
		num1, num2 = reconheceNumero()
		var subtracao int = num1 - num2
		fmt.Println(subtracao)
	case 3:
		fmt.Println("a operação escolhida foi a 3")
		num1, num2 = reconheceNumero()
		var multiplicação int = num1 * num2
		fmt.Println(multiplicação)
	case 4:
		fmt.Println("a operação escolhida foi a 4")
		num1, num2 = reconheceNumero()
		var divisão int = num1 / num2
		fmt.Println(divisão)
	default:
		fmt.Println("Número inválido!")
	}
}

func reconheceNumero() (int, int) {
	var num1 int
	var num2 int

	fmt.Println("Digite o valor do primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Println("Digite o valor do segundo número: ")
	fmt.Scanln(&num2)

	return num1, num2
}

func menu() {
	fmt.Println("Seja bem vindo viajante, qual operação deseja realizar?")
	fmt.Println("1) Soma? ")
	fmt.Println("2) Subtração? ")
	fmt.Println("3) Multiplicação")
	fmt.Println("4) Divisão")
}
