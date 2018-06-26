package main

import (
	"fmt"
)

type pessoas struct {
	nome     string
	endereço string
	idade    int
}

func main() {
	p := pessoas{nome: "Junior", endereço: "rua dos Go", idade: 39}
	fmt.Println(p)
	fmt.Println(p.nome)
	fmt.Println(p.endereço)
	fmt.Println(p.idade)
}
