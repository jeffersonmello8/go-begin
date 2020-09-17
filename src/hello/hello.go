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

const tempoDeEsperaEmSegundos = 2
const monitoramentos = 2

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
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
	versao := 1.0
	fmt.Println("Olá, qual é o seu nome?")
	fmt.Scan(&nome)
	pulaUmaLinhaNoConsole()
	fmt.Println(nome+", "+"este programa está na versão", versao)
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

		sites := leSitesDoArquivo()

		for indice, site := range sites {
			fmt.Println("Testando site", indice, ":", site)
			testaSite(site)
			time.Sleep(tempoDeEsperaEmSegundos * time.Second)
			pulaUmaLinhaNoConsole()
		}
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site:", site, "está funcionando perfeitamente!")
		registraLog(site, true)
	} else {
		fmt.Println("O site:", site, "tem algum problema. Código retornado:", resp.StatusCode)
		registraLog(site, false)
	}
}

func pulaUmaLinhaNoConsole() {
	println("")
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("./sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			pulaUmaLinhaNoConsole()
			break
		}
		if err != nil {
			fmt.Println("Ocorreu um erro:", err)
		}
	}

	arquivo.Close()

	return sites

}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + "Endereço: " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println("Exibindo Logs..." + "\n")

	fmt.Println(string(arquivo))
}
