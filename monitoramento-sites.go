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

const monitoramento = 1
const delay = 1

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
			fmt.Print("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Print("Não reconheço este comando!")
			os.Exit(-1)
		}

	}
}

func exibeIntroducao() {
	nome := "Juliano"
	versao := 1.1
	idade := 30
	fmt.Println("Olá sr.", nome, ". Sua idade é", idade)
	fmt.Println("Este programa está na versão: ", versao)
}

func exibeMenu() {
	fmt.Println("1- iniciar monitoramento: ")
	fmt.Println("2- exibir logs ")
	fmt.Println("0- sair do programa ")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := lendoSitesDoArquivo()

	for i := 0; i < monitoramento; i++ {
		for _, site := range sites {
			testeSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func testeSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "está funcionando corretamente!")
		registrLog(site, true)
	} else {
		fmt.Println("O site", site, "está com problemas. StatusCode:", resp.StatusCode)
		registrLog(site, false)
	}
}

func lendoSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao encontrar o arquivo:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		if err == io.EOF {
			break
		}
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}

	arquivo.Close()

	fmt.Println(sites)

	return sites
}

func registrLog(site string, online bool) {

	arquivo, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " " + site + " - online: " + strconv.FormatBool(online) + "\n")

	arquivo.Close()
}

func exibirLogs() {

	//POde ser feito de 2 maneiras
	/* 	fmt.Print("Exibindo os Logs..." + "\n")

	   	arquivo, err := os.Open("logs.txt")

	   	if err != nil {
	   		fmt.Println(err)
	   	}

	   	leitor := bufio.NewReader(arquivo)

	   	for {
	   		linha, err := leitor.ReadString('\n')

	   		if err == io.EOF {
	   			break
	   		}

	   		linha = strings.TrimSpace(linha)
	   		fmt.Println(linha)
	   	}

		   arquivo.Close() */

	arquivo, err := ioutil.ReadFile("logs.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))
}
