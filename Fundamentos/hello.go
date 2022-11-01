package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntrodução()

	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			inicarMonitoramento()
		case 2:
			fmt.Println("exibindo logs...")
			imprimeLogs()
		case 3:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando")
		}
	}

}

func exibeIntrodução() {
	nome := "Yuri"
	fmt.Println("Ola caro viajante: ", nome)
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)
	fmt.Println("")
	return comando
}

func exibeMenu() {
	fmt.Println("1 -iniciar Monitoramento")
	fmt.Println("2 -Exibir logs")
	fmt.Println("3 -Fechar programa")
}

func inicarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leArquivoExterno()

	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites {
			fmt.Println("Testando site", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("o site", site, "foi carregado com sucesso")
		escreveLogs(site, true)
	} else {
		fmt.Println("deu merda chefia, tenta de novo, Status code:", resp.StatusCode)
		escreveLogs(site, false)
	}
}

func leArquivoExterno() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
		sites = append(sites, linha)

		if err == io.EOF {
			break
		}

	}
	arquivo.Close()
	return sites
}

func escreveLogs(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
