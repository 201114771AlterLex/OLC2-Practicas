// Code generated from Gramatica.g4 by ANTLR 4.9. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 13, 62, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 6, 10, 45, 10, 10, 13, 10, 14,
	10, 46, 3, 11, 3, 11, 7, 11, 51, 10, 11, 12, 11, 14, 11, 54, 11, 11, 3,
	12, 6, 12, 57, 10, 12, 13, 12, 14, 12, 58, 3, 12, 3, 12, 2, 2, 13, 3, 3,
	5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13,
	3, 2, 6, 3, 2, 50, 59, 5, 2, 67, 92, 97, 97, 99, 124, 6, 2, 50, 59, 67,
	92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34, 2, 64, 2, 3, 3, 2, 2,
	2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2,
	2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2,
	2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 3, 25, 3, 2, 2, 2, 5, 27, 3,
	2, 2, 2, 7, 31, 3, 2, 2, 2, 9, 33, 3, 2, 2, 2, 11, 35, 3, 2, 2, 2, 13,
	37, 3, 2, 2, 2, 15, 39, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2, 19, 44, 3, 2, 2,
	2, 21, 48, 3, 2, 2, 2, 23, 56, 3, 2, 2, 2, 25, 26, 7, 61, 2, 2, 26, 4,
	3, 2, 2, 2, 27, 28, 7, 107, 2, 2, 28, 29, 7, 112, 2, 2, 29, 30, 7, 118,
	2, 2, 30, 6, 3, 2, 2, 2, 31, 32, 7, 45, 2, 2, 32, 8, 3, 2, 2, 2, 33, 34,
	7, 47, 2, 2, 34, 10, 3, 2, 2, 2, 35, 36, 7, 44, 2, 2, 36, 12, 3, 2, 2,
	2, 37, 38, 7, 49, 2, 2, 38, 14, 3, 2, 2, 2, 39, 40, 7, 42, 2, 2, 40, 16,
	3, 2, 2, 2, 41, 42, 7, 43, 2, 2, 42, 18, 3, 2, 2, 2, 43, 45, 9, 2, 2, 2,
	44, 43, 3, 2, 2, 2, 45, 46, 3, 2, 2, 2, 46, 44, 3, 2, 2, 2, 46, 47, 3,
	2, 2, 2, 47, 20, 3, 2, 2, 2, 48, 52, 9, 3, 2, 2, 49, 51, 9, 4, 2, 2, 50,
	49, 3, 2, 2, 2, 51, 54, 3, 2, 2, 2, 52, 50, 3, 2, 2, 2, 52, 53, 3, 2, 2,
	2, 53, 22, 3, 2, 2, 2, 54, 52, 3, 2, 2, 2, 55, 57, 9, 5, 2, 2, 56, 55,
	3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 56, 3, 2, 2, 2, 58, 59, 3, 2, 2, 2,
	59, 60, 3, 2, 2, 2, 60, 61, 8, 12, 2, 2, 61, 24, 3, 2, 2, 2, 6, 2, 46,
	52, 58, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "';'", "'int'", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "MAS", "MEN", "POR", "DIV", "PARI", "PARD", "NUM", "ID", "WHITESPACE",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "MAS", "MEN", "POR", "DIV", "PARI", "PARD", "NUM", "ID",
	"WHITESPACE",
}

type GramaticaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewGramaticaLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *GramaticaLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGramaticaLexer(input antlr.CharStream) *GramaticaLexer {
	l := new(GramaticaLexer)
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
	l.GrammarFileName = "Gramatica.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GramaticaLexer tokens.
const (
	GramaticaLexerT__0       = 1
	GramaticaLexerT__1       = 2
	GramaticaLexerMAS        = 3
	GramaticaLexerMEN        = 4
	GramaticaLexerPOR        = 5
	GramaticaLexerDIV        = 6
	GramaticaLexerPARI       = 7
	GramaticaLexerPARD       = 8
	GramaticaLexerNUM        = 9
	GramaticaLexerID         = 10
	GramaticaLexerWHITESPACE = 11
)
