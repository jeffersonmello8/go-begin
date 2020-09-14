package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const tempoDeEsperaEmSegundos = 2
const monitoramentos = 10

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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
}

func exibeIntroducao() {
	var nome string
	versao := 1.1
	fmt.Println("Olá, qual é o seu nome?")
	fmt.Scan(&nome)
	pulaUmaLinhaNoConsole()
	fmt.Println(nome, "este programa está na versão", versao)
	pulaUmaLinhaNoConsole()
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

func iniciarMonitoramento() {
	for i := 0; i < monitoramentos; i++ {
		fmt.Println("Monitorando...")
		pulaUmaLinhaNoConsole()
		sites := []string{"https://translate.google.com/",
			"https://www.google.com.br/",
			"https://random-status-code.herokuapp.com/"}
		for indice, site := range sites {
			fmt.Println("Testando site", indice, ":", site)
			testaSite(site)
			time.Sleep(tempoDeEsperaEmSegundos * time.Second)
			pulaUmaLinhaNoConsole()
		}
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "está funcionando perfeitamente!")
	} else {
		fmt.Println("O site:", site, "tem algum problema. Código retornado:", resp.StatusCode)
	}
}

func pulaUmaLinhaNoConsole() {
	println("")
}
