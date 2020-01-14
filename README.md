# go-web-scrap-simple-example
Pequeno web scrap escrito em Go, que busca os títulos dos posts [deste](https://golangcode.com) blog.

Este script utiliza a biblioteca **net/http** para fazer a request, e a **goquery** para buscar os elementos do DOM com os mesmos métodos do jquery (no caso, utilizei o comando `Find()` e o `Each()`).

## Como utilizar
- Para instalar a biblioteca goquery:
  - `go get github.com/PuerkitoBio/goquery`

- Para executar:
  - `go run main.go`
