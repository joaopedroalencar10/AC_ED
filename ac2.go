package main

import (
	"bufio"
	"fmt"
	"os"
)

type Contato struct {
	Nome  string
	Email string
}

func main() {
	var c Contato
	var email string

	leitor := bufio.NewReader(os.Stdin)
	fmt.Println("informe um nome:")
	nome, _ := leitor.ReadString('\n')

	fmt.Println("Informe um email: ")
	fmt.Scanln(&email)

	c = Contato{
		Nome:  nome,
		Email: email,
	}

	var contatos [5]Contato

	contatos[0] = c

}

/*func adiciona(Contato,contatos){
  for _, contato := range contatos {

    if (contato!= Contato{}) {
			break
		}
}
  return contatos
}
*/
