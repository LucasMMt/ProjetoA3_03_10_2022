package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
)

// uso de struct
type Aluno struct {
	Nome string
	Nota []string
}

func cliente(aluno Aluno) {
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = gob.NewEncoder(c).Encode(aluno)
	if err != nil {
		fmt.Println(err)
	}
	c.Write([]byte(aluno.Nome))
	c.Close()
}

func soma(aluno Aluno) {
	s1 := aluno.Nota[0]
	s2 := aluno.Nota[1]
	s3 := aluno.Nota[2]

	f1, err := strconv.ParseFloat(s1, 2)

	f2, err := strconv.ParseFloat(s2, 2)

	f3, err := strconv.ParseFloat(s3, 2)
	fmt.Println("Media das 3 notas :")
	fmt.Println((f1 + f2 + f3) / 3)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	aluno := Aluno{
		Nome: "Lucas Mauricio",
		Nota: []string{
			"10",
			"9.3",
			"5.5",
		},
	}
	fmt.Println("Cliente enviou dados do Aluno para o Servidor")
	go soma(aluno)
	go cliente(aluno)

	//CTRL + C para sair
	var input string
	fmt.Scanln(&input)
}
