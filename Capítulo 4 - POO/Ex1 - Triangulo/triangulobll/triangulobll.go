package triangulobll

import (
	"POO/erro"
	"POO/triangulo"
	"fmt"
)

func ValidaDados(e erro.Erro, t triangulo.Triangulo) {
	e.SetErro(false)
	if t.GetBase() <= 0 && t.GetAltura() <= 0 {
		e.SetErro(true)
		e.SetMens("O campo BASE e de ALTURA devem ser maiores que 0")
		fmt.Println(t.GetBase(), e.GetMens(), e.GetErro())
	} else {
		if t.GetBase() <= 0 {
			e.SetErro(true)
			e.SetMens("O campo BASE deve ser maior que 0")
		} else if t.GetAltura() <= 0 {
			e.SetErro(true)
			e.SetMens("O campo ALTURA deve ser maior que 0")
		}
	}
}
