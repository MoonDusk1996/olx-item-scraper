package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

// Declarar variáveis para armazenar as entradas do usuário
var itemToSearch string

func main() {

	//Imprime o campo de item a procurar e armazena o input
	fmt.Print("Item a procurar: ")
	_, err := fmt.Scan(&itemToSearch)
	if err != nil {
		fmt.Println("Erro ao ler a entrada:", err)
		os.Exit(1)
	}

	// Defina as opções disponíveis
	options := []string{"AM", "RJ", "SP", "BA"}

	// Crie um prompt usando a biblioteca survey
	var country string
	prompt := &survey.Select{
		Message: "Estado a procurar: ",
		Options: options,
	}

	// Pergunte ao usuário
	err = survey.AskOne(prompt, &country, survey.WithValidator(survey.Required))
	if err != nil {
		log.Fatal(err)
	}

	//executa o scraper
	scraper(itemToSearch, country)
}
