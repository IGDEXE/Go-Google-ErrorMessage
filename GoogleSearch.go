// Funcao para pesquisar um erro diretamente no Google
// Ivo Dias

// Define um pacote principal
package main

// Importa os pacotes necessarios
import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func GoogleSearch(mensagemErro string) {
	// Cria a variaveis
	var err error
	var url string

	// Cofigura a pesquisa
	url = "https://www.google.com/search?q=" + mensagemErro

	// Verifica o SO e abre o navegador
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Exemplo
func main() {
	var err string
	switch runtime.GOOS {
	default:
		err = "Golang error unsupported platform"
	}
	if err != "" {
		GoogleSearch(err)
	}
}
