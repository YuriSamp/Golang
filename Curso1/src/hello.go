package main

import (
	"fmt"

	// "reflect"

	"os"
)

func main() {
	exibeIntrodução()
	exibeMenu()
	comando := leComando()

	// if comando == 1 {
	// 	fmt.Println("Monitorando...")
	// } else if comando == 2 {
	// 	fmt.Println("exibindo logs...")
	// } else if comando == 3 {
	// 	fmt.Println("Saindo do programa...")
	// }

	switch comando {
	case 1:
		inicarMonitoramento()
	case 2:
		fmt.Println("exibindo logs...")
	case 3:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando")
	}

	// fmt.Println("O tipo do dinheiro é", reflect.TypeOf(dinheiro))
}

func exibeIntrodução() {
	nome := "Yuri"
	fmt.Println("Ola caro viajante: ", nome)
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func exibeMenu() {
	fmt.Println("1 -iniciar Monitoramento")
	fmt.Println("2 -Exibir logs")
	fmt.Println("3 -Fechar programa")
}

func inicarMonitoramento() {
	fmt.Println("Monitorando...")
	// site := "https://https://www.youtube.com"
	// resp := http.Get(site)
}
