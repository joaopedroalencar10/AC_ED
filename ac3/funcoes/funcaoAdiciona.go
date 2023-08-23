package funcoes

import (
	"faculdade/contato"
)

func AdicionaContato(c contato.Contato, a [5]contato.Contato) [5]contato.Contato {
	for ind, contatoa := range a {
		if (contatoa == contato.Contato{}) {
			a[ind] = c
			break
		}
	}

	return a
}
