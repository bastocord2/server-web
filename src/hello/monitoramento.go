package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	entradaPrincipal()
	//registraLog("site-falso", false)

	for {

		//_, idade := devolveNomeeIdade()
		//fmt.Println("eu tenho a idade", idade)

		exibeMenu()

		//var comando int
		comando := menu()

		/*fmt.Println(" O comando escolhido foi", comando)

		if comando == 1 {
			fmt.Println("Iniciar Monitoramento....")
		} else if comando == 2 {
			fmt.Println("Exibindo Log......")

		} else if comando == 3 {
			fmt.Println("Saindo do Sistema.....")

		} else {
			fmt.Println("Comando Inválido!")
		}
		*/

		switch comando {
		case 1:
			fmt.Println("Iniciar Monitoramento....")
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Log......")
		case 3:
			fmt.Println("Saindo do Sistema.....")
			os.Exit(0)
		default:
			fmt.Println("Comando Inválido!")
			os.Exit(-1)
		}
	}
}
func entradaPrincipal() {

	var nome string = "Eduardo"
	var versao float32 = 1.1

	fmt.Println("Olá, sr", nome)
	fmt.Println("Este programa esta na versão", versao)

}
func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exiber Log")
	fmt.Println("3 - Sair do sistema")

}

/*func devolveNomeeIdade() (string, int) {
	nome := "Edu"
	idade := 41
	return nome, idade

}*/
func menu() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Comando lido foi", comandoLido)
	return comandoLido
}
func iniciarMonitoramento() {
	fmt.Printf("Monitoramendo...")
	fmt.Println("")
	//sites := []string{"http://google.com", "http://terra.com.br", "https://www.amazon.com.br/"}
	//fmt.Println(sites)

	sites := leSitesdoArquivo()
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("posicao ", i, "site ", site)
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
		fmt.Println("Erro ao ler site: ", err)
	}
	//fmt.Println(resp)
	//resp, err := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!!!")
		registraLog(site, true)
	} else {
		fmt.Println("O Site", site, "esta co problemas . Status Code", resp.StatusCode)
		registraLog(site, false)
	}

}
func leSitesdoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	//arquivo, err := ioutil.ReadFile("sites.txt")
	if err != nil {
		fmt.Println("Erro ao ler arquivo: ", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		fmt.Println(linha)
		if err == io.EOF {
			break
		}

	}
	arquivo.Close()
	//fmt.Println(string(arquivo))
	return sites
}
func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(arquivo)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "- online:" + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}
