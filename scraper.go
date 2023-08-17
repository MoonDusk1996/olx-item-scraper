package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func scraper(item string, country string) {

	//recebe o item e o estado do item a procurar e logo em seguida cria um url
	baseURL := "https://www.olx.com.br/estado-"
	url := baseURL + country + "?q=" + item

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

	// Seletor CSS para o titulo da busca que o usuário deseja extrair
	searchTitle := "h1[data-ds-component='DS-Text']"

	doc.Find(searchTitle).Each(func(index int, element *goquery.Selection) {
		text := strings.TrimSpace(element.Text())
		fmt.Println("Resultado da busca:", text)
	})

	// Seletor CSS para os elementos que você deseja extrair
	searchItens := "h2[data-ds-component='DS-Text']"

	doc.Find(searchItens).Each(func(index int, element *goquery.Selection) {
		text := strings.TrimSpace(element.Text())
		fmt.Printf("item %d: %s\n", index, text)
	})

}

// <h2 data-ds-component="DS-Text" class="sc-jrQzAO jhFuLz horizontal title">Kindle 10° geração preto / Parcelo no cartão </h2>
