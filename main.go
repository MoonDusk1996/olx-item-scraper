package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// Declarar variáveis para armazenar as configurações e entradas do usuário
var countryOptions = []string{
	"AC", "AL", "AP", "AM", "BA", "CE", "DF", "ES", "GO", "MA", "MT", "MS", "MG",
	"PA", "PB", "PR", "PE", "PI", "RJ", "RN", "RS", "RO", "RR", "SC", "SP", "SE", "TO",
}

func main() {
	var itemToSearch string
	var country string

	//Imprime o campo de item a procurar e armazena o input
	fmt.Print("Item a procurar: ")
	reader := bufio.NewReader(os.Stdin)
	itemToSearch, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		os.Exit(1)
	}
	// Remove a quebra de linha do final
	itemToSearch = strings.TrimSpace(itemToSearch)

	// Cria um prompt usando a biblioteca survey
	prompt := &survey.Select{
		Message: "Estado a procurar: ",
		Options: countryOptions,
	}

	// Pergunta ao usuário o estado em qual buscar o item
	err = survey.AskOne(prompt, &country, survey.WithValidator(survey.Required))
	if err != nil {
		log.Fatal(err)
	}

	//executa o scraper
	scraper(itemToSearch, country)
}
