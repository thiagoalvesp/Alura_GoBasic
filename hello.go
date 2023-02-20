package main

import (
	"fmt"
	"os/user"
)

func main() {
	exibeIntroducao()
}

func iniciarMonitoramento() {
	fmt.Printf("Monitorando...")
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
