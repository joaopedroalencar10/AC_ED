package funcoes

import (
	"faculdade/contato"
	"fmt"
)

func ExcluiContato(a [5]contato.Contato) [5]contato.Contato {
	if (a[0] == contato.Contato{}) {
		fmt.Println("Lista de contatos vazia!")
		return a
	}

	for ind, contatob := range a {
		if (contatob == contato.Contato{}) {
			a[ind-1] = contato.Contato{}
		}
	}

	return a
}
