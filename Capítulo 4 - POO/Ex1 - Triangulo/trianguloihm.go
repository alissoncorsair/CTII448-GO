package main

//importação dos pacotes erro, triangulo e triangulo bll. além das bibliotecas lxn/walk e walk/declarative
import (
	"POO/erro"
	"POO/triangulo"
	"POO/triangulobll"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	//definindo valores para os elementos gráficos do programa
	var altura *walk.NumberEdit
	var base *walk.NumberEdit
	var resultado *walk.NumberEdit
	var calcular *walk.PushButton
	var form *walk.MainWindow

	MainWindow{
		AssignTo: &form,
		Title:    "Calculo Área",
		Size:     Size{250, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Base: ",
			},

			NumberEdit{
				AssignTo: &base,
			},

			Label{
				Text: "Altura: ",
			},

			NumberEdit{
				AssignTo: &altura,
			},

			//botão calcular
			PushButton{
				AssignTo: &calcular,
				Text:     "Calcular",
				// assim que for clicado, a função será acionada...
				OnClicked: func() {
					// instanciando uma variável do tipo Triangulo e do tipo Erro.
					// ambas sendo referenciadas por seus packages de origem.
					triangulo := new(triangulo.Triangulo)
					erro := new(erro.Erro)

					// chamando as funções SetBase e SetAltura da variável tipo Triangulo do package triangulo
					// e definindo os valores da base e da altura do triângulo
					// tendo como parâmetro o que foi escrito
					// nas caixas de número que são referenciadas pelas variáveis "base" e "altura"
					triangulo.SetBase(base.Value())
					triangulo.SetAltura(altura.Value())

					// chamando a função ValidaDados do package triangulobll e usando como parâmetro
					// os ponteiros das variáveis tipo Erro e Triangulo, respectivamente.
					triangulobll.ValidaDados(*erro, *triangulo)

					// se a função GetErro retornar true, uma mensagem de erro será exibida na tela.
					// a mensagem exibida vem da função GetMens
					// função esta que recebe essa mensagem através do pacote triangulobll
					// por meio da função SetMens
					if erro.GetErro() {
						walk.MsgBox(form, "Erro!", erro.GetMens(), walk.MsgBoxIconError)
						// caso a função GetErro retorne false, o valor da caixa de número referente à
						// área do triângulo terá como valor o resultado da função GetArea que retorna a área do triângulo criado.
						// além disso, os campos de base e altura serão desativados.
					} else {
						resultado.SetValue(triangulo.GetArea())
						base.SetEnabled(false)
						altura.SetEnabled(false)
					}
				},
			},

			// botão limpar
			PushButton{
				Text: "Limpar",
				// após ser clicado, a função será acionada...
				OnClicked: func() {
					// os campos referentes à altura, base, resultado, e altura terão alterações: valores resetados
					// e botões reativados.
					altura.SetValue(0)
					base.SetValue(0)
					resultado.SetValue(0)
					base.SetEnabled(true)
					altura.SetEnabled(true)
				},
			},

			Label{
				Text: "Área: ",
			},

			// campo onde aparece a área do triângulo.
			// nunca poderá editado e/ou clicado.
			NumberEdit{
				AssignTo: &resultado,
				ReadOnly: true,
				Enabled:  false,
			},
		},
	}.Run()
}
