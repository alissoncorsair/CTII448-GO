package livrodal

import (
	"POO/livro"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var erroBD error
var valorMens string
var valorErro bool

func Conecta() {
	db, erroBD = sql.Open("mysql", "root:@tcp(localhost:3306)/cadastrolivro")
	if erroBD != nil {
		panic(erroBD.Error())
	}

}

func Desconecta() {
	defer db.Close()
}

func InserirLivro(livro livro.Livro) (bool, string) {
	var pesquisa = "insert into livros (codigo,titulo,autor,editora,ano) values (?,?,?,?,?)"

	stmt, erroPesquisa := db.Prepare(pesquisa)

	if erroPesquisa != nil {
		valorErro = true
		valorMens = "erro ao preparar a consulta ao banco de dados!"
		return valorErro, valorMens
	}

	respostaInserir, erroInserir := stmt.Exec(livro.GetCodigo(), livro.GetTitulo(), livro.GetAutor(), livro.GetEditora(), livro.GetAno())

	if erroInserir != nil {
		valorErro = true
		valorMens = "erro ao inserir o livro ao banco de dados!"
		return valorErro, valorMens
	}

	fmt.Print(respostaInserir)
	valorErro = false
	valorMens = "livro cadastrado com sucesso!"
	return valorErro, valorMens
}

func ExcluiUmLivro(livro livro.Livro) (bool, string) {
	var pesquisa = "delete from livros where codigo = (?)"

	stmt, erroPesquisa := db.Prepare(pesquisa)

	if erroPesquisa != nil {
		valorErro = true
		valorMens = "erro ao preparar a consulta ao banco de dados!"
		return valorErro, valorMens
	}

	respostaExclui, erroExclui := stmt.Exec(livro.GetCodigo())

	if erroExclui != nil {
		valorErro = true
		valorMens = "Houve um erro ao excluir o livro!"
		return valorErro, valorMens
	}
	fmt.Print(respostaExclui)

	valorErro = false
	valorMens = "o livro foi excluído com sucesso"
	return valorErro, valorMens
}

func AtualizaUmLivro(livro livro.Livro) (bool, string) {
	var pesquisa = "update livros set titulo=(?), autor=(?), editora=(?), ano=(?) where codigo = (?)"

	stmt, erroPesquisa := db.Prepare(pesquisa)

	if erroPesquisa != nil {
		valorErro = true
		valorMens = "erro ao preparar a consulta ao banco de dados!"
		return valorErro, valorMens
	}

	respostaAtualiza, erroAtualiza := stmt.Exec(livro.GetTitulo(), livro.GetAutor(), livro.GetEditora(), livro.GetAno(), livro.GetCodigo())

	if erroAtualiza != nil {
		valorErro = true
		valorMens = "não foi possível alterar o livro!"
		return valorErro, valorMens
	}

	fmt.Print(respostaAtualiza)
	valorErro = false
	valorMens = "o livro foi alterado com sucesso!"
	return valorErro, valorMens
}

func ConsultaUmLivro(livro livro.Livro) (livro.Livro, bool, string) {
	var pesquisa = "select * from livros where codigo=(?)"
	var entrou = false
	smtr, erroPesquisa := db.Prepare(pesquisa)

	if erroPesquisa != nil {
		valorErro = true
		valorMens = "erro ao preparar a consulta ao banco de dados!"
		return livro, valorErro, valorMens
	}

	respostaConsulta, erroConsulta := smtr.Query(livro.GetCodigo())
	livro.SetCodigo("")
	livro.SetTitulo("")
	livro.SetAutor("")
	livro.SetEditora("")
	livro.SetAno(0)
	valorErro = true
	valorMens = "erro ao consultar o livro!"

	if erroConsulta != nil {
		valorErro = true
		valorMens = "erro ao consultar o livro!"
		return livro, valorErro, valorMens
	}

	for respostaConsulta.Next() {
		entrou = true
		var codigo string
		var titulo string
		var autor string
		var editora string
		var ano int

		erroLeituraDados := respostaConsulta.Scan(&codigo, &titulo, &autor, &editora, &ano)

		if erroLeituraDados != nil {
			valorErro = true
			valorMens = "houve um erro ao ler os dados da tabela do banco de dados!"
			return livro, valorErro, valorMens
		}
		livro.SetCodigo(codigo)
		livro.SetTitulo(titulo)
		livro.SetAutor(autor)
		livro.SetEditora(editora)
		livro.SetAno(ano)

		fmt.Println(livro.GetCodigo(), livro.GetTitulo(), livro.GetAutor(), livro.GetEditora(), livro.GetAno())
	}

	if entrou {
		valorErro = false
		valorMens = "livro consultado com sucesso"
	}
	return livro, valorErro, valorMens
}
