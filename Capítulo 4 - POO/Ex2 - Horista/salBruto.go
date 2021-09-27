package main

//importação dos pacotes erro, horista e horista bll. além das bibliotecas lxn/walk e walk/declarative
import (
	"POO/erro"
	"POO/horista"
	"POO/horistabll"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	//definindo valores para os elementos gráficos do programa
	var qtdH, vlrH, Salario *walk.NumberEdit
	var form *walk.MainWindow

	MainWindow{
		AssignTo: &form,
		Title:    "Calculo de Salário Bruto",
		Size:     Size{400, 200},
		Layout:   Grid{Columns: 2},
		Font:     Font{Family: "Arial", PointSize: 14},
		Children: []Widget{
			Label{
				Text: "Digite a quantidade de horas:",
			},
			NumberEdit{
				AssignTo: &qtdH,
			},

			Label{
				Text: "Digite o valor da hora:",
			},
			NumberEdit{
				AssignTo: &vlrH,
				Suffix:   " BRL",
				Decimals: 2,
			},

			//botão calcular
			PushButton{
				Text: "Calcular",
				// assim que for clicado, a função será acionada...
				OnClicked: func() {
					// instanciando uma variável do tipo Horista e do tipo Erro.
					// ambas sendo referenciadas por seus packages de origem.
					erro := new(erro.Erro)
					horista := new(horista.Horista)

					// chamando as funções SetQtd e SetValor da variável tipo Horista do package horista
					// e definindo os valores da hora e do valor da hora do horista
					// tendo como parâmetro o que foi escrito
					// nas caixas de número que são referenciadas pelas variáveis "qtdH" e "vlrH"
					// e validando os dados através da função ValidaDados (origem no package horistabll)
					// tendo como parâmetros a variável erro e o ponteiro da variável horista
					horista.SetQtd(int(qtdH.Value()))
					horista.SetValor(vlrH.Value())
					horistabll.ValidaDados(erro, *horista)

					// se GetErro retornar true, uma mensagem de erro aparecerá
					// a mensagem exibida vem da função GetMens
					// função esta que recebe essa mensagem através do pacote horistabll
					// por meio da função SetMens
					if erro.GetErro() {
						walk.MsgBox(form, "Erro!", erro.GetMens(), walk.MsgBoxStyle(walk.MsgBoxIconError))
						// caso a função GetErro retorne false, o valor da caixa de número referente ao
						// valor bruto terá como valor o resultado da função SalarioBruto que retorna o salário bruto.
						// além disso, os campos de quantidade de horas e valor por hora serão desativados.
					} else {
						Salario.SetValue(horista.SalarioBruto())
						qtdH.SetEnabled(false)
						vlrH.SetEnabled(false)
					}
				},
			},

			// botão limpar
			PushButton{
				Text: "Limpar",
				OnClicked: func() {
					// os campos referentes à quantidade de horas, valor por hora e salario terão alterações: valores resetados
					// e botões reativados.
					qtdH.SetValue(0)
					vlrH.SetValue(0)
					Salario.SetValue(0)
					vlrH.SetEnabled(true)
					qtdH.SetEnabled(true)
				},
			},

			Label{
				Text: "Salário Bruto: ",
			},

			// campo onde aparece o salário.
			// nunca poderá editado e/ou clicado.
			NumberEdit{
				AssignTo: &Salario,
				Suffix:   " BRL",
				Decimals: 2,
				Enabled:  false,
			},
		},
	}.Run()
}
