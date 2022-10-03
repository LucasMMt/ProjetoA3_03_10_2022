package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

type Aluno struct {
	Nome string
	Nota []string
}

func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	//loop for para Aceitar conexão
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	var aluno Aluno
	err := gob.NewDecoder(c).Decode(&aluno)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Servidor recebeu Dados de Cliente \nNome Do Aluno: ", aluno.Nome+"\n"+
			"Primeira Nota: "+aluno.Nota[0]+"\n"+
			"Segunda Nota: "+aluno.Nota[1]+"\n"+
			"Terceira Nota: "+aluno.Nota[2])
	}

}

func main() {

	go servidor()

	var input string
	fmt.Scanln(&input)
}
