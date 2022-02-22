package main

import (
	"fmt"

	"github.com/OLC2-Practicas/Practica3/listener"
	"github.com/OLC2-Practicas/Practica3/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func analizar(input string) {
	// create input stream
	stream := antlr.NewInputStream(input)
	// create lexer
	lexer := parser.NewPractica3Lexer(stream)
	// create tokenStream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// create parser
	parser := parser.NewPractica3Parser(tokenStream)
	// get tree
	tree := parser.Start()

	// activate listener

	listen := listener.Listener{}
	antlr.ParseTreeWalkerDefault.Walk(&listen, tree)

	// print value
	fmt.Println("Entrada:", input)
	fmt.Println("Salida: Tipo:", listen.Type, " Valor: ", listen.Value)
}

func main() {
	Test := "5+8*6+10.2"
	input := ""
	Option := 0
	for {
		fmt.Println("Desea analizar cadena de prueba o ingresa una o salir (1/2/3)")
		fmt.Scan(&Option)
		if Option == 1 {
			analizar(Test)
		} else if Option == 2 {
			fmt.Println("Ingrese cadena")
			fmt.Scan(&input)
			analizar(input)
			input = ""
		} else if Option == 3 {
			break
		}
	}
}
