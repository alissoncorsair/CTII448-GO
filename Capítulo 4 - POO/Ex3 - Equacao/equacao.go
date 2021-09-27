package main

//importação dos pacotes erro, equacao e equacao bll. além das bibliotecas lxn/walk e walk/declarative
import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"

	"POO/equacao"
	"POO/equacaobll"
	"POO/erro"
)

func main() {
	//definindo valores para os elementos gráficos do programa
	var vlrA, vlrB, vlrC, x1, x2 *walk.NumberEdit
	var form *walk.MainWindow

	MainWindow{
		AssignTo: &form,
		Title:    "Equação do segundo grau",
		Size:     Size{400, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Valor de a:",
			},

			NumberEdit{
				AssignTo: &vlrA,
			},

			Label{
				Text: "Valor de b:",
			},

			NumberEdit{
				AssignTo: &vlrB,
			},

			Label{
				Text: "Valor de c:",
			},

			NumberEdit{
				AssignTo: &vlrC,
			},

			//botão calcular
			PushButton{
				Text: "Calcular",
				// assim que for clicado, a função será acionada...
				OnClicked: func() {
					// instanciando uma variável do tipo Equacao e do tipo Erro.
					// ambas sendo referenciadas por seus packages de origem.
					erro := new(erro.Erro)
					equacao := new(equacao.Equacao)

					// chamando as funções SetA, SetB e SetC da variável tipo Equacao do package equacao
					// e definindo os valores de A, B e C
					// tendo como parâmetro o que foi escrito
					// nas caixas de número que são referenciadas pelas variáveis "vlrA", "vlrB" e "vlrC"
					// e validando os dados através da função ValidaDados (origem no package equacaobll)
					// tendo como parâmetros a variável erro e o ponteiro da variável equacao
					equacao.SetA(int(vlrA.Value()))
					equacao.SetB(int(vlrB.Value()))
					equacao.SetC(int(vlrC.Value()))
					equacaobll.ValidaDados(erro, *equacao)

					// se GetErro retornar true, uma mensagem de erro aparecerá
					// a mensagem exibida vem da função GetMens
					// função esta que recebe essa mensagem através do pacote equacaobll
					// por meio da função SetMens
					if erro.GetErro() {
						walk.MsgBox(form, "Erro", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
					} else {

						// caso a função GetErro retorne false, o valor das caixas de número referentes ao
						// valor x1 e x2 terão como valor o resultado das funções X1 e X2 que retornam x1 e x2 respectivamente.
						// além disso, os campos de A, B e C serão desativados.

						x1.SetValue(equacao.X1())
						x2.SetValue(equacao.X2())
						vlrA.SetEnabled(false)
						vlrB.SetEnabled(false)
						vlrC.SetEnabled(false)
					}
				},
			},

			// botão limpar
			PushButton{
				Text: "Limpar",
				OnClicked: func() {
					// os campos referentes aos valores de A, B, C, x1 e x2 terão alterações: valores resetados
					// e A, B e C também terão seus botões reativados.
					vlrA.SetValue(0)
					vlrB.SetValue(0)
					vlrC.SetValue(0)
					x1.SetValue(0)
					x2.SetValue(0)
					vlrA.SetEnabled(true)
					vlrB.SetEnabled(true)
					vlrC.SetEnabled(true)
				},
			},

			Label{
				Text: "x1: ",
			},
			// campo onde aparece o valor de x1.
			// nunca poderá editado e/ou clicado.
			NumberEdit{
				AssignTo: &x1,
				Enabled:  false,
				Decimals: 2,
			},

			Label{
				Text: "x2: ",
			},
			// campo onde aparece o valor de x2.
			// nunca poderá editado e/ou clicado.
			NumberEdit{
				AssignTo: &x2,
				Enabled:  false,
				Decimals: 2,
			},
		},
	}.Run()
}
