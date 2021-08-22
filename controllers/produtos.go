package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/wilker/loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco:", err)
		}

		quantidadeParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade:", err)
		}

		models.CriarNovoProduto(nome, descricao, precoParaFloat, quantidadeParaInt)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	idDoProdutoParaInt, err := strconv.Atoi(idDoProduto)
	if err != nil {
		log.Println("Erro na conversao id:", err)
	}
	models.DeletaProduto(idDoProdutoParaInt)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco:", err)
		}

		quantidadeParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade:", err)
		}

		idParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversao do id:", err)
		}

		models.AtualizaProduto(idParaInt, nome, descricao, precoParaFloat, quantidadeParaInt)
	}

	http.Redirect(w, r, "/", 301)
}
