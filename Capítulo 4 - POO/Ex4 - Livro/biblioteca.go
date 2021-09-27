package main

//importação dos pacotes erro, livro e livro bll. além das bibliotecas lxn/walk e walk/declarative
import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"POO/erro"
	"POO/livro"
	"POO/livrobll"
)

func main() {
	//definindo valores para os elementos gráficos do programa
	var cod, titulo, editora, autor *walk.TextEdit
	var consultar *walk.PushButton
	var ano *walk.NumberEdit
	var form *walk.MainWindow

	// instanciando duas variáveis do tipo Livro e uma do tipo Erro.
	// ambas sendo referenciadas por seus packages de origem.
	livroAux := new(livro.Livro)
	livro := new(livro.Livro)
	erroCadastra := new(erro.Erro)

	MainWindow{
		AssignTo: &form,
		Title:    "Cadastro de Livro",
		Size:     Size{300, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Código:",
			},

			TextEdit{
				Text:     "",
				AssignTo: &cod,
			},

			Label{
				Text: "Titulo:",
			},

			TextEdit{
				Text:     "",
				AssignTo: &titulo,
			},

			Label{
				Text: "Autor",
			},

			TextEdit{
				Text:     "",
				AssignTo: &autor,
			},

			Label{
				Text: "Editora",
			},

			TextEdit{
				Text:     "",
				AssignTo: &editora,
			},

			Label{
				Text: "Ano",
			},

			NumberEdit{
				AssignTo:           &ano,
				RightToLeftReading: false,
			},

			//botão cadastrar
			PushButton{
				ColumnSpan: 2,
				Text:       "Cadastrar",
				// assim que for clicado, a função será acionada...
				OnClicked: func() {

					// chamando as funções SetTitulo, SetAutor, SetCodigo, SetEditora e SetAno da variável tipo Equacao do package livro
					// e definindo os valores de titulo, autor, codigo, editora e ano
					// tendo como parâmetro o que foi escrito
					// nas caixas de texto/numero que são referenciadas pelas variáveis "titulo", "autor", "cod", "editora" e "ano"
					// e validando os dados através da função ValidaDados (origem no package livrobll)
					// tendo como parâmetros a variável erroCadastra e o ponteiro da variável livroAux
					livroAux.SetTitulo(titulo.Text())
					livroAux.SetAutor(autor.Text())
					livroAux.SetCodigo(cod.Text())
					livroAux.SetEditora(editora.Text())
					livroAux.SetAno(int(ano.Value()))
					livrobll.ValidaDados(erroCadastra, *livroAux)

					// se GetErro retornar true, uma mensagem de erro aparecerá
					// a mensagem exibida vem da função GetMens
					// função esta que recebe essa mensagem através do pacote livrobll
					// por meio da função SetMens
					if erroCadastra.GetErro() {
						walk.MsgBox(form, "Erro", erroCadastra.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {

						// caso a função GetErro retorne false, o valor das caixas de texto/numero referentes a
						// titulo, autor, codigo, editora e ano terão como valor as informações do livro cadastrado
						// através de livroAux, que, passará os valores para a variável livro e cadastrará no sistema.
						consultar.SetEnabled(true)
						walk.MsgBox(form, "Boa!", "Os dados inseridos com sucesso!", walk.MsgBoxStyle(walk.MsgBoxIconInformation))
						livro.SetTitulo(livroAux.GetTitulo())
						livro.SetAutor(livroAux.GetAutor())
						livro.SetCodigo(livroAux.GetCodigo())
						livro.SetEditora(livroAux.GetEditora())
						livro.SetAno(livroAux.GetAno())
					}
				},
			},

			PushButton{
				AssignTo:   &consultar,
				ColumnSpan: 2,
				Text:       "Consultar",
				Enabled:    false,
				OnClicked: func() {
					// instanciamento de outra variável do tipo Erro
					erro := new(erro.Erro)

					// validando o código através da função ValidaCodigo (origem no package livrobll)
					// tendo como parâmetros a variável erro, o ponteiro da variável livro e o valor escrito na caixa de texto do campo código
					livrobll.ValidaCodigo(erro, *livro, cod.Text())

					// se GetErro retornar true, uma mensagem de erro aparecerá
					// a mensagem exibida vem da função GetMens
					// função esta que recebe essa mensagem através do pacote livrobll
					// por meio da função SetMens
					if erro.GetErro() {
						walk.MsgBox(form, "Erro", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))

						// caso GetErro retornar false, o código do livro foi encontrado na memória, então
						// as informações do livro serão exibidas em seus respectivos campos.
					} else {
						cod.SetText(livro.GetCodigo())
						titulo.SetText(livro.GetTitulo())
						autor.SetText(livro.GetAutor())
						editora.SetText(livro.GetEditora())
						ano.SetValue(float64(livro.GetAno()))
					}
				},
			},

			// botão limpar
			PushButton{
				ColumnSpan: 2,
				Text:       "Limpar",
				OnClicked: func() {
					// após clique no botão limpar, todos os campos serão resetados para valores vazios ou 0.
					cod.SetText("")
					titulo.SetText("")
					autor.SetText("")
					editora.SetText("")
					ano.SetValue(0)
				},
			},
		},
	}.Run()
}
