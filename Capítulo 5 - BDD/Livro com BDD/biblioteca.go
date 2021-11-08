package main

//importação dos pacotes erro, livro e livro bll. além das bibliotecas lxn/walk e walk/declarative
import (
	"fmt"

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
	livro := new(livro.Livro)
	erro := new(erro.Erro)

	// livrobll.Desconecta()
	livrobll.Conecta()
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

			PushButton{
				ColumnSpan: 2,
				Text:       "Salvar",
				OnClicked: func() {
					livro.SetTitulo(titulo.Text())
					livro.SetAutor(autor.Text())
					livro.SetCodigo(cod.Text())
					livro.SetEditora(editora.Text())
					livro.SetAno(int(ano.Value()))
					livrobll.ValidaDados(erro, *livro, "i")

					if erro.GetErro() {
						walk.MsgBox(form, "Erro!", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {
						walk.MsgBox(form, "Sucesso!", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconInformation))
					}

				},
			},

			PushButton{
				AssignTo:   &consultar,
				ColumnSpan: 2,
				Text:       "Ler",
				OnClicked: func() {
					livro.SetCodigo(cod.Text())
					livrobll.ValidaCodigo(erro, livro, "c")

					if erro.GetErro() {
						walk.MsgBox(form, "Erro!", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {
						walk.MsgBox(form, "Sucesso!", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconInformation))
						cod.SetText(livro.GetCodigo())
						titulo.SetText(livro.GetTitulo())
						autor.SetText(livro.GetAutor())
						editora.SetText(livro.GetEditora())
						ano.SetValue(float64(livro.GetAno()))
					}
				},
			},

			PushButton{
				ColumnSpan: 2,
				Text:       "Alterar",
				OnClicked: func() {
					livro.SetCodigo(cod.Text())
					livro.SetTitulo(titulo.Text())
					livro.SetAutor(autor.Text())
					livro.SetEditora(editora.Text())
					livro.SetAno(int(ano.Value()))
					fmt.Println(livro.GetTitulo(), livro.GetAutor(), livro.GetEditora(), livro.GetAno(), livro.GetCodigo())

					livrobll.ValidaDados(erro, *livro, "a")

					if erro.GetErro() {
						walk.MsgBox(form, "Erro!", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {
						walk.MsgBox(form, "Alterado!", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconInformation))
					}

				},
			},

			// botão excluir
			PushButton{
				ColumnSpan: 2,
				Text:       "Excluir",
				OnClicked: func() {
					livro.SetCodigo(cod.Text())

					livrobll.ValidaCodigo(erro, livro, "e")

					if erro.GetErro() {
						walk.MsgBox(form, "Erro", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {
						walk.MsgBox(form, "Excluído!", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconInformation))
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
