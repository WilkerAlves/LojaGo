package models

import "github.com/wilker/loja/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()
	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	produtos := []Produto{}
	p := Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	insereDadoNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values(?, ?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}
	insereDadoNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id int) {
	db := db.ConectaComBancoDeDados()
	deletaDadoNoBanco, err := db.Prepare("delete from produtos where id = ?")
	if err != nil {
		panic(err.Error())
	}
	deletaDadoNoBanco.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()
	produtoDoBanco, err := db.Query("select * from produtos where id = ?", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}
	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}

	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()
	insereDadoNoBanco, err := db.Prepare("update produtos set nome = ?, descricao = ?, preco = ?, quantidade = ? where id = ? ")
	if err != nil {
		panic(err.Error())
	}
	insereDadoNoBanco.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
