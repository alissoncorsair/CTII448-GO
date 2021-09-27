package horistabll

import (
	"POO/erro"
	"POO/horista"
	"fmt"
)

func ValidaDados(e *erro.Erro, h horista.Horista) {
	e.SetErro(false)

	if h.GetQtd() <= 0 && h.GetValor() <= 0 {
		e.SetErro(true)
		e.SetMens("O campo BASE e de ALTURA devem ser maiores que 0")
	} else {
		if h.GetQtd() <= 0 {
			e.SetErro(true)
			e.SetMens("O campo BASE deve ser maior que 0")
		} else if h.GetValor() <= 0 {
			e.SetErro(true)
			e.SetMens("O campo ALTURA deve ser maior que 0")
		}
	}
	fmt.Println(e.GetErro(), e.GetMens())
}
