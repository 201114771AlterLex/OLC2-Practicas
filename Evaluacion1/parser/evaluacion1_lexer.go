// Code generated from Evaluacion1.g4 by ANTLR 4.9. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 34, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 5, 4, 22, 10, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 6, 7, 29, 10, 7, 13, 7, 14, 7, 30, 3, 7, 3, 7, 2, 2, 8,
	3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 3, 2, 5, 3, 2, 50, 59, 3, 2, 67,
	72, 5, 2, 11, 12, 15, 15, 34, 34, 2, 35, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2,
	2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3,
	2, 2, 2, 3, 15, 3, 2, 2, 2, 5, 17, 3, 2, 2, 2, 7, 21, 3, 2, 2, 2, 9, 23,
	3, 2, 2, 2, 11, 25, 3, 2, 2, 2, 13, 28, 3, 2, 2, 2, 15, 16, 7, 48, 2, 2,
	16, 4, 3, 2, 2, 2, 17, 18, 7, 46, 2, 2, 18, 6, 3, 2, 2, 2, 19, 22, 5, 9,
	5, 2, 20, 22, 5, 11, 6, 2, 21, 19, 3, 2, 2, 2, 21, 20, 3, 2, 2, 2, 22,
	8, 3, 2, 2, 2, 23, 24, 9, 2, 2, 2, 24, 10, 3, 2, 2, 2, 25, 26, 9, 3, 2,
	2, 26, 12, 3, 2, 2, 2, 27, 29, 9, 4, 2, 2, 28, 27, 3, 2, 2, 2, 29, 30,
	3, 2, 2, 2, 30, 28, 3, 2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2,
	32, 33, 8, 7, 2, 2, 33, 14, 3, 2, 2, 2, 5, 2, 21, 30, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'.'", "','",
}

var lexerSymbolicNames = []string{
	"", "Point", "Colon", "Hex", "Number", "Letter", "Whitespace",
}

var lexerRuleNames = []string{
	"Point", "Colon", "Hex", "Number", "Letter", "Whitespace",
}

type Evaluacion1Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewEvaluacion1Lexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *Evaluacion1Lexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewEvaluacion1Lexer(input antlr.CharStream) *Evaluacion1Lexer {
	l := new(Evaluacion1Lexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Evaluacion1.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Evaluacion1Lexer tokens.
const (
	Evaluacion1LexerPoint      = 1
	Evaluacion1LexerColon      = 2
	Evaluacion1LexerHex        = 3
	Evaluacion1LexerNumber     = 4
	Evaluacion1LexerLetter     = 5
	Evaluacion1LexerWhitespace = 6
)
