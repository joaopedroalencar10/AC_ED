package main

import (
	"bufio"
	"faculdade/contato"
	"faculdade/funcoes"
	"fmt"
	"os"
)

func main() {
	var contatos [5]contato.Contato
	var nome, email, opcao string

	leitor := bufio.NewReader(os.Stdin)

	fmt.Println("Lista de contatos!")
	for {
		fmt.Print("Digite (1) para adicionar, (2)excluir contato, (3)mostra todos os contatos: ")
		fmt.Scanln(&opcao)

		switch opcao {
		case "1":
			fmt.Print("Informe o seu nome: ")
			nome, _ = leitor.ReadString('\n')

			fmt.Print("Informe o seu email: ")
			fmt.Scanln(&email)

			contatos = funcoes.AdicionaContato(contato.Contato{Nome: nome, Email: email}, contatos)
		case "2":
			contatos = funcoes.ExcluiContato(contatos)
			fmt.Println(contatos)

		case "3":
			fmt.Println(contatos)

		default:
			return
		}

	}

}
