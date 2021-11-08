package livrobll

import (
	"POO/erro"
	"POO/livro"
	"POO/livrodal"
)

func Conecta() {
	livrodal.Conecta()
}

func Desconecta() {
	livrodal.Desconecta()
}

func ValidaDados(e *erro.Erro, l livro.Livro, op string) {
	e.SetErro(false)

	if l.GetCodigo() == "" {
		e.SetMens("O campo código é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetTitulo() == "" {
		e.SetMens("O campo titulo é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetAutor() == "" {
		e.SetMens("O campo autor é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetEditora() == "" {
		e.SetMens("O campo editora é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if l.GetAno() <= 0 {
		e.SetMens("O campo ano tem de ser maior do que 0")
		e.SetErro(true)
		return
	}

	if op == "i" {
		valorErro, valorMens := livrodal.InserirLivro(l)
		e.SetErro(valorErro)
		e.SetMens(valorMens)
	} else {
		valorErro, valorMens := livrodal.AtualizaUmLivro(l)
		e.SetErro(valorErro)
		e.SetMens(valorMens)
	}
}

func ValidaCodigo(e *erro.Erro, l *livro.Livro, op string) {
	e.SetErro(false)
	if l.GetCodigo() == "" {
		e.SetMens("O campo código é de preenchimento obrigatório")
		e.SetErro(true)
		return
	}

	if op == "c" {
		livro, valorErro, valorMens := livrodal.ConsultaUmLivro(*l)
		e.SetErro(valorErro)
		e.SetMens(valorMens)

		if valorErro == false {
			l.SetCodigo(livro.GetCodigo())
			l.SetTitulo(livro.GetTitulo())
			l.SetAutor(livro.GetAutor())
			l.SetEditora(livro.GetEditora())
			l.SetAno(livro.GetAno())
		}
	} else {
		valorErro, valorMens := livrodal.ExcluiUmLivro(*l)
		e.SetErro(valorErro)
		e.SetMens(valorMens)
	}
}
