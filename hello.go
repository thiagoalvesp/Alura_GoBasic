package main

import (
	"fmt"
	"net/http"
	"os"
	"os/user"
)

func main() {
	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			exibirLogs()
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheco este comando")
			os.Exit(-1)
		}
	}
}

func iniciarMonitoramento() {
	fmt.Println()
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br/"
	fmt.Println()
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
	fmt.Println()
}

func exibirLogs() {
	fmt.Println()
	fmt.Println("Exibindo Logs...")
	fmt.Println()
}

func exibeIntroducao() {
	usr, _ := user.Current()
	versao := 1.2
	nome := usr.Name

	fmt.Println("Olá, sr. ", nome)
	fmt.Println("Versão: ", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi ", comandoLido)

	return comandoLido
}
