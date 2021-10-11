package main

import (
	"fmt"
	"os"
)

func main() {

	entradaPrincipal()
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
func entradaPrincipal() {

	var nome string = "Edurdo"
	var versao float32 = 1.1
	casa := "home"
	fmt.Println("Olá, sr", nome)
	fmt.Println("Este programa esta na versão", versao)
	fmt.Println("Este programa esta localizado", casa)

}
func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exiber Log")
	fmt.Println("3 - Sair do sistema")

}
func menu() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("Comando lido foi", comandoLido)
	return comandoLido
}
