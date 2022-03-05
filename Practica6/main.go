/*
	go mod init Compilador
	go get github.com/antlr/antlr4/runtime/Go/antlr
	antlr -o parser Gramatica.g4
	go run main.go
*/

package main

import (
	"Compilador/parser"
	"bufio"
	"log"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Listener struct {
	*parser.BaseGramaticaListener
}

func analizar(input string) {
	// create input stream
	stream := antlr.NewInputStream(input)
	// create lexer
	lexer := parser.NewGramaticaLexer(stream)
	// create tokenStream
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	// create parser
	parser := parser.NewGramaticaParser(tokenStream)
	// get tree
	tree := parser.Start()

	// activate listener
	listen := Listener{}
	antlr.ParseTreeWalkerDefault.Walk(&listen, tree)
}

func leerEntrada(name string) string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileScanner := bufio.NewScanner(file)
	text := ""
	for fileScanner.Scan() {
		text += fileScanner.Text()
	}
	return text
}

func main() {
	input := leerEntrada("entrada.txt")
	analizar(input)
}
