package models

import "github.com/JoaoValentimDev/shoop/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(
			&id,
			&nome,
			&preco,
			&quantidade,
			&descricao)
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

func BuscarUnicoProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	query := "select * from produtos where id=$1"

	buscarProduto, err := db.Prepare(query)

	if err != nil {
		panic(err.Error())
	}

	result, err := buscarProduto.Query(id)

	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for result.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = result.Scan(
			&id,
			&nome,
			&preco,
			&quantidade,
			&descricao)

		if err != nil {
			panic(err.Error())
		}

		produto.Id = id
		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

	}

	defer db.Close()
	return produto
}

func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	defer db.Close()

}

func DeletaProduto(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()

}

func Update(nome, descricao string, preco float64, quantidade, id int) {
	db := db.ConectaComBancoDeDados()

	sql := "update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5"

	update, err := db.Prepare(sql)

	if err != nil {
		panic(err.Error())
	}

	update.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
