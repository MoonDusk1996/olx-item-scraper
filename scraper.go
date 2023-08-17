package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var totalItems int

func scraper(itemToSearch string, country string) {

	//formata os inputs do usuario
	formatedItemToSearch := strings.ReplaceAll(itemToSearch, " ", "%20")
	formatedCountry := strings.ToLower(country)

	//recebe o item e o estado do item a procurar e logo em seguida cria um url
	baseURL := "https://www.olx.com.br/estado-"
	url := baseURL + formatedCountry + "?q=" + formatedItemToSearch

	// Faz a requisição HTTP para o URL e espera obter o conteúdo da página
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Carrega o conteúdo da página no goquery.Document
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Seletor CSS para o titulo da busca
	searchTitle := "h1[data-ds-component='DS-Text']"

	doc.Find(searchTitle).Each(func(index int, element *goquery.Selection) {
		text := strings.TrimSpace(element.Text())
		fmt.Printf("Resultado da busca:%s\n", text)
		fmt.Println(" ")
	})

	// Seletor CSS para os items referente a busca
	searchItens := "h2[data-ds-component='DS-Text']"

	doc.Find(searchItens).Each(func(index int, element *goquery.Selection) {
		text := strings.TrimSpace(element.Text())
		totalItems++
		fmt.Printf("item %d: %s\n", index, text)
	})
	if totalItems <= 3 {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Println("Itens não carregados corretamente, refazendo a busca...")
		scraper(itemToSearch, country)
	}
}
