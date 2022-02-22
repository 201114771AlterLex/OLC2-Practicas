// Code generated from /home/alterlex/go_projects/src/github.com/OLC2-Practicas/Practica3/Practica3copy.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 10, 45, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3,
	5, 3, 6, 3, 6, 3, 7, 6, 7, 31, 10, 7, 13, 7, 14, 7, 32, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 9, 6, 9, 40, 10, 9, 13, 9, 14, 9, 41, 3, 9, 3, 9, 2, 2, 10,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 3, 2, 4, 3, 2, 50,
	59, 5, 2, 11, 12, 15, 15, 34, 34, 2, 46, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 3, 19, 3, 2, 2, 2, 5, 21,
	3, 2, 2, 2, 7, 23, 3, 2, 2, 2, 9, 25, 3, 2, 2, 2, 11, 27, 3, 2, 2, 2, 13,
	30, 3, 2, 2, 2, 15, 34, 3, 2, 2, 2, 17, 39, 3, 2, 2, 2, 19, 20, 7, 43,
	2, 2, 20, 4, 3, 2, 2, 2, 21, 22, 7, 42, 2, 2, 22, 6, 3, 2, 2, 2, 23, 24,
	7, 45, 2, 2, 24, 8, 3, 2, 2, 2, 25, 26, 7, 44, 2, 2, 26, 10, 3, 2, 2, 2,
	27, 28, 7, 48, 2, 2, 28, 12, 3, 2, 2, 2, 29, 31, 9, 2, 2, 2, 30, 29, 3,
	2, 2, 2, 31, 32, 3, 2, 2, 2, 32, 30, 3, 2, 2, 2, 32, 33, 3, 2, 2, 2, 33,
	14, 3, 2, 2, 2, 34, 35, 5, 13, 7, 2, 35, 36, 5, 11, 6, 2, 36, 37, 5, 13,
	7, 2, 37, 16, 3, 2, 2, 2, 38, 40, 9, 3, 2, 2, 39, 38, 3, 2, 2, 2, 40, 41,
	3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2,
	43, 44, 8, 9, 2, 2, 44, 18, 3, 2, 2, 2, 5, 2, 32, 41, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "')'", "'('", "'+'", "'*'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "PRight", "PLeft", "Plus", "Times", "Point", "Numbers", "Decimal",
	"Whitespace",
}

var lexerRuleNames = []string{
	"PRight", "PLeft", "Plus", "Times", "Point", "Numbers", "Decimal", "Whitespace",
}

type Practica3copyLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewPractica3copyLexer(input antlr.CharStream) *Practica3copyLexer {

	l := new(Practica3copyLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Practica3copy.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Practica3copyLexer tokens.
const (
	Practica3copyLexerPRight     = 1
	Practica3copyLexerPLeft      = 2
	Practica3copyLexerPlus       = 3
	Practica3copyLexerTimes      = 4
	Practica3copyLexerPoint      = 5
	Practica3copyLexerNumbers    = 6
	Practica3copyLexerDecimal    = 7
	Practica3copyLexerWhitespace = 8
)
