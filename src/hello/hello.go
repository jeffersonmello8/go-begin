package main

import (
	"fmt"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
		os.Exit(0)
	default:
		fmt.Println("Comando não reconhecido!")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	var nome string
	versao := 1.1
	fmt.Println("Olá, qual é o seu nome?")
	fmt.Scan(&nome)
	fmt.Println(nome, "este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("O que você deseja fazer?")
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}
