// Code generated from /home/alterlex/go_projects/src/github.com/OLC2-Practicas/Practica6/Gramatica.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Gramatica

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "Compilador/gen"
import "Compilador/entorno"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 40, 261,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3, 2, 3,
	2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 66, 10, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 101, 10, 4, 12, 4, 14, 4, 104, 11,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 5, 5, 118, 10, 5, 3, 6, 7, 6, 121, 10, 6, 12, 6, 14, 6, 124, 11, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 133, 10, 7, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 7, 9, 153, 10, 9, 12, 9, 14, 9, 156, 11, 9, 3,
	9, 3, 9, 5, 9, 160, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 173, 10, 10, 3, 10, 3, 10, 6, 10, 177,
	10, 10, 13, 10, 14, 10, 178, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 185, 10,
	10, 5, 10, 187, 10, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 202, 10, 11, 3, 11, 3,
	11, 6, 11, 206, 10, 11, 13, 11, 14, 11, 207, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 5, 11, 215, 10, 11, 3, 11, 3, 11, 5, 11, 219, 10, 11, 3, 11, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 2, 3, 6, 18, 2,
	4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 2, 4, 3, 2, 6,
	8, 3, 2, 3, 4, 2, 279, 2, 34, 3, 2, 2, 2, 4, 37, 3, 2, 2, 2, 6, 65, 3,
	2, 2, 2, 8, 117, 3, 2, 2, 2, 10, 122, 3, 2, 2, 2, 12, 132, 3, 2, 2, 2,
	14, 134, 3, 2, 2, 2, 16, 140, 3, 2, 2, 2, 18, 163, 3, 2, 2, 2, 20, 191,
	3, 2, 2, 2, 22, 223, 3, 2, 2, 2, 24, 230, 3, 2, 2, 2, 26, 238, 3, 2, 2,
	2, 28, 243, 3, 2, 2, 2, 30, 254, 3, 2, 2, 2, 32, 258, 3, 2, 2, 2, 34, 35,
	5, 10, 6, 2, 35, 36, 7, 2, 2, 3, 36, 3, 3, 2, 2, 2, 37, 38, 8, 3, 1, 2,
	38, 5, 3, 2, 2, 2, 39, 40, 8, 4, 1, 2, 40, 41, 7, 3, 2, 2, 41, 42, 5, 6,
	4, 16, 42, 43, 8, 4, 1, 2, 43, 66, 3, 2, 2, 2, 44, 45, 7, 4, 2, 2, 45,
	46, 5, 6, 4, 15, 46, 47, 8, 4, 1, 2, 47, 66, 3, 2, 2, 2, 48, 49, 7, 5,
	2, 2, 49, 50, 5, 6, 4, 14, 50, 51, 8, 4, 1, 2, 51, 66, 3, 2, 2, 2, 52,
	53, 7, 12, 2, 2, 53, 54, 5, 6, 4, 2, 54, 55, 7, 13, 2, 2, 55, 56, 8, 4,
	1, 2, 56, 66, 3, 2, 2, 2, 57, 58, 7, 37, 2, 2, 58, 66, 8, 4, 1, 2, 59,
	60, 7, 38, 2, 2, 60, 66, 8, 4, 1, 2, 61, 62, 7, 14, 2, 2, 62, 66, 8, 4,
	1, 2, 63, 64, 7, 15, 2, 2, 64, 66, 8, 4, 1, 2, 65, 39, 3, 2, 2, 2, 65,
	44, 3, 2, 2, 2, 65, 48, 3, 2, 2, 2, 65, 52, 3, 2, 2, 2, 65, 57, 3, 2, 2,
	2, 65, 59, 3, 2, 2, 2, 65, 61, 3, 2, 2, 2, 65, 63, 3, 2, 2, 2, 66, 102,
	3, 2, 2, 2, 67, 68, 12, 13, 2, 2, 68, 69, 9, 2, 2, 2, 69, 70, 5, 6, 4,
	14, 70, 71, 8, 4, 1, 2, 71, 101, 3, 2, 2, 2, 72, 73, 12, 12, 2, 2, 73,
	74, 9, 3, 2, 2, 74, 75, 5, 6, 4, 13, 75, 76, 8, 4, 1, 2, 76, 101, 3, 2,
	2, 2, 77, 78, 12, 11, 2, 2, 78, 79, 5, 8, 5, 2, 79, 80, 5, 6, 4, 12, 80,
	81, 8, 4, 1, 2, 81, 101, 3, 2, 2, 2, 82, 83, 12, 10, 2, 2, 83, 84, 7, 9,
	2, 2, 84, 85, 5, 4, 3, 2, 85, 86, 5, 6, 4, 11, 86, 87, 8, 4, 1, 2, 87,
	101, 3, 2, 2, 2, 88, 89, 12, 9, 2, 2, 89, 90, 7, 10, 2, 2, 90, 91, 5, 4,
	3, 2, 91, 92, 5, 6, 4, 10, 92, 93, 8, 4, 1, 2, 93, 101, 3, 2, 2, 2, 94,
	95, 12, 8, 2, 2, 95, 96, 7, 11, 2, 2, 96, 97, 5, 4, 3, 2, 97, 98, 5, 6,
	4, 9, 98, 99, 8, 4, 1, 2, 99, 101, 3, 2, 2, 2, 100, 67, 3, 2, 2, 2, 100,
	72, 3, 2, 2, 2, 100, 77, 3, 2, 2, 2, 100, 82, 3, 2, 2, 2, 100, 88, 3, 2,
	2, 2, 100, 94, 3, 2, 2, 2, 101, 104, 3, 2, 2, 2, 102, 100, 3, 2, 2, 2,
	102, 103, 3, 2, 2, 2, 103, 7, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2, 105, 106,
	7, 16, 2, 2, 106, 118, 8, 5, 1, 2, 107, 108, 7, 17, 2, 2, 108, 118, 8,
	5, 1, 2, 109, 110, 7, 18, 2, 2, 110, 118, 8, 5, 1, 2, 111, 112, 7, 19,
	2, 2, 112, 118, 8, 5, 1, 2, 113, 114, 7, 20, 2, 2, 114, 118, 8, 5, 1, 2,
	115, 116, 7, 21, 2, 2, 116, 118, 8, 5, 1, 2, 117, 105, 3, 2, 2, 2, 117,
	107, 3, 2, 2, 2, 117, 109, 3, 2, 2, 2, 117, 111, 3, 2, 2, 2, 117, 113,
	3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 9, 3, 2, 2, 2, 119, 121, 5, 12,
	7, 2, 120, 119, 3, 2, 2, 2, 121, 124, 3, 2, 2, 2, 122, 120, 3, 2, 2, 2,
	122, 123, 3, 2, 2, 2, 123, 11, 3, 2, 2, 2, 124, 122, 3, 2, 2, 2, 125, 133,
	5, 14, 8, 2, 126, 133, 5, 16, 9, 2, 127, 133, 5, 20, 11, 2, 128, 133, 5,
	22, 12, 2, 129, 133, 5, 24, 13, 2, 130, 133, 5, 26, 14, 2, 131, 133, 5,
	28, 15, 2, 132, 125, 3, 2, 2, 2, 132, 126, 3, 2, 2, 2, 132, 127, 3, 2,
	2, 2, 132, 128, 3, 2, 2, 2, 132, 129, 3, 2, 2, 2, 132, 130, 3, 2, 2, 2,
	132, 131, 3, 2, 2, 2, 133, 13, 3, 2, 2, 2, 134, 135, 7, 38, 2, 2, 135,
	136, 7, 22, 2, 2, 136, 137, 5, 6, 4, 2, 137, 138, 7, 23, 2, 2, 138, 139,
	8, 8, 1, 2, 139, 15, 3, 2, 2, 2, 140, 141, 7, 24, 2, 2, 141, 142, 5, 6,
	4, 2, 142, 143, 8, 9, 1, 2, 143, 144, 5, 30, 16, 2, 144, 154, 8, 9, 1,
	2, 145, 146, 7, 25, 2, 2, 146, 147, 7, 24, 2, 2, 147, 148, 5, 6, 4, 2,
	148, 149, 8, 9, 1, 2, 149, 150, 5, 30, 16, 2, 150, 151, 8, 9, 1, 2, 151,
	153, 3, 2, 2, 2, 152, 145, 3, 2, 2, 2, 153, 156, 3, 2, 2, 2, 154, 152,
	3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155, 159, 3, 2, 2, 2, 156, 154, 3, 2,
	2, 2, 157, 158, 7, 25, 2, 2, 158, 160, 5, 30, 16, 2, 159, 157, 3, 2, 2,
	2, 159, 160, 3, 2, 2, 2, 160, 161, 3, 2, 2, 2, 161, 162, 8, 9, 1, 2, 162,
	17, 3, 2, 2, 2, 163, 164, 7, 26, 2, 2, 164, 165, 5, 6, 4, 2, 165, 176,
	7, 27, 2, 2, 166, 167, 7, 28, 2, 2, 167, 168, 5, 6, 4, 2, 168, 169, 7,
	29, 2, 2, 169, 172, 8, 10, 1, 2, 170, 173, 5, 30, 16, 2, 171, 173, 5, 32,
	17, 2, 172, 170, 3, 2, 2, 2, 172, 171, 3, 2, 2, 2, 173, 174, 3, 2, 2, 2,
	174, 175, 8, 10, 1, 2, 175, 177, 3, 2, 2, 2, 176, 166, 3, 2, 2, 2, 177,
	178, 3, 2, 2, 2, 178, 176, 3, 2, 2, 2, 178, 179, 3, 2, 2, 2, 179, 186,
	3, 2, 2, 2, 180, 181, 7, 30, 2, 2, 181, 184, 7, 29, 2, 2, 182, 185, 5,
	30, 16, 2, 183, 185, 5, 32, 17, 2, 184, 182, 3, 2, 2, 2, 184, 183, 3, 2,
	2, 2, 185, 187, 3, 2, 2, 2, 186, 180, 3, 2, 2, 2, 186, 187, 3, 2, 2, 2,
	187, 188, 3, 2, 2, 2, 188, 189, 8, 10, 1, 2, 189, 190, 7, 31, 2, 2, 190,
	19, 3, 2, 2, 2, 191, 192, 7, 26, 2, 2, 192, 193, 5, 6, 4, 2, 193, 194,
	7, 27, 2, 2, 194, 205, 8, 11, 1, 2, 195, 196, 7, 28, 2, 2, 196, 197, 5,
	6, 4, 2, 197, 198, 7, 29, 2, 2, 198, 201, 8, 11, 1, 2, 199, 202, 5, 30,
	16, 2, 200, 202, 5, 32, 17, 2, 201, 199, 3, 2, 2, 2, 201, 200, 3, 2, 2,
	2, 202, 203, 3, 2, 2, 2, 203, 204, 8, 11, 1, 2, 204, 206, 3, 2, 2, 2, 205,
	195, 3, 2, 2, 2, 206, 207, 3, 2, 2, 2, 207, 205, 3, 2, 2, 2, 207, 208,
	3, 2, 2, 2, 208, 218, 3, 2, 2, 2, 209, 210, 7, 30, 2, 2, 210, 211, 7, 29,
	2, 2, 211, 214, 8, 11, 1, 2, 212, 215, 5, 30, 16, 2, 213, 215, 5, 32, 17,
	2, 214, 212, 3, 2, 2, 2, 214, 213, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216,
	217, 8, 11, 1, 2, 217, 219, 3, 2, 2, 2, 218, 209, 3, 2, 2, 2, 218, 219,
	3, 2, 2, 2, 219, 220, 3, 2, 2, 2, 220, 221, 8, 11, 1, 2, 221, 222, 7, 31,
	2, 2, 222, 21, 3, 2, 2, 2, 223, 224, 7, 32, 2, 2, 224, 225, 8, 12, 1, 2,
	225, 226, 5, 6, 4, 2, 226, 227, 8, 12, 1, 2, 227, 228, 5, 30, 16, 2, 228,
	229, 8, 12, 1, 2, 229, 23, 3, 2, 2, 2, 230, 231, 7, 33, 2, 2, 231, 232,
	8, 13, 1, 2, 232, 233, 5, 30, 16, 2, 233, 234, 7, 32, 2, 2, 234, 235, 5,
	6, 4, 2, 235, 236, 7, 23, 2, 2, 236, 237, 8, 13, 1, 2, 237, 25, 3, 2, 2,
	2, 238, 239, 7, 34, 2, 2, 239, 240, 8, 14, 1, 2, 240, 241, 5, 30, 16, 2,
	241, 242, 8, 14, 1, 2, 242, 27, 3, 2, 2, 2, 243, 244, 7, 35, 2, 2, 244,
	245, 7, 38, 2, 2, 245, 246, 7, 22, 2, 2, 246, 247, 5, 6, 4, 2, 247, 248,
	7, 36, 2, 2, 248, 249, 5, 6, 4, 2, 249, 250, 8, 15, 1, 2, 250, 251, 7,
	33, 2, 2, 251, 252, 5, 30, 16, 2, 252, 253, 8, 15, 1, 2, 253, 29, 3, 2,
	2, 2, 254, 255, 7, 27, 2, 2, 255, 256, 5, 10, 6, 2, 256, 257, 7, 31, 2,
	2, 257, 31, 3, 2, 2, 2, 258, 259, 5, 10, 6, 2, 259, 33, 3, 2, 2, 2, 18,
	65, 100, 102, 117, 122, 132, 154, 159, 172, 178, 184, 186, 201, 207, 214,
	218,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'+'", "'-'", "'!'", "'*'", "'/'", "'%'", "'&&'", "'^'", "'||'", "'('",
	"')'", "'true'", "'false'", "'=='", "'!='", "'>'", "'<'", "'>='", "'<='",
	"'='", "';'", "'if'", "'else'", "'switch'", "'{'", "'case'", "':'", "'default'",
	"'}'", "'while'", "'do'", "'loop'", "'for'", "'to'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "NUM",
	"ID", "COMMENT", "WHITESPACE",
}

var ruleNames = []string{
	"start", "marcador", "expresion", "oprel", "instrucciones", "instruccion",
	"inst_asignacion", "inst_if", "inst_switch_propuesta1", "inst_switch_propuesta2",
	"inst_while", "inst_doWhile", "inst_loop", "inst_for", "bloque", "bloqueSinLLaves",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type GramaticaParser struct {
	*antlr.BaseParser
}

func NewGramaticaParser(input antlr.TokenStream) *GramaticaParser {
	this := new(GramaticaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Gramatica.g4"

	return this
}

// posición relativa de un símbolo
var desp int = 0

// entorno actual
var tope *entorno.Entorno

// GramaticaParser tokens.
const (
	GramaticaParserEOF        = antlr.TokenEOF
	GramaticaParserT__0       = 1
	GramaticaParserT__1       = 2
	GramaticaParserT__2       = 3
	GramaticaParserT__3       = 4
	GramaticaParserT__4       = 5
	GramaticaParserT__5       = 6
	GramaticaParserT__6       = 7
	GramaticaParserT__7       = 8
	GramaticaParserT__8       = 9
	GramaticaParserT__9       = 10
	GramaticaParserT__10      = 11
	GramaticaParserT__11      = 12
	GramaticaParserT__12      = 13
	GramaticaParserT__13      = 14
	GramaticaParserT__14      = 15
	GramaticaParserT__15      = 16
	GramaticaParserT__16      = 17
	GramaticaParserT__17      = 18
	GramaticaParserT__18      = 19
	GramaticaParserT__19      = 20
	GramaticaParserT__20      = 21
	GramaticaParserT__21      = 22
	GramaticaParserT__22      = 23
	GramaticaParserT__23      = 24
	GramaticaParserT__24      = 25
	GramaticaParserT__25      = 26
	GramaticaParserT__26      = 27
	GramaticaParserT__27      = 28
	GramaticaParserT__28      = 29
	GramaticaParserT__29      = 30
	GramaticaParserT__30      = 31
	GramaticaParserT__31      = 32
	GramaticaParserT__32      = 33
	GramaticaParserT__33      = 34
	GramaticaParserNUM        = 35
	GramaticaParserID         = 36
	GramaticaParserCOMMENT    = 37
	GramaticaParserWHITESPACE = 38
)

// GramaticaParser rules.
const (
	GramaticaParserRULE_start                  = 0
	GramaticaParserRULE_marcador               = 1
	GramaticaParserRULE_expresion              = 2
	GramaticaParserRULE_oprel                  = 3
	GramaticaParserRULE_instrucciones          = 4
	GramaticaParserRULE_instruccion            = 5
	GramaticaParserRULE_inst_asignacion        = 6
	GramaticaParserRULE_inst_if                = 7
	GramaticaParserRULE_inst_switch_propuesta1 = 8
	GramaticaParserRULE_inst_switch_propuesta2 = 9
	GramaticaParserRULE_inst_while             = 10
	GramaticaParserRULE_inst_doWhile           = 11
	GramaticaParserRULE_inst_loop              = 12
	GramaticaParserRULE_inst_for               = 13
	GramaticaParserRULE_bloque                 = 14
	GramaticaParserRULE_bloqueSinLLaves        = 15
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(GramaticaParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GramaticaParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(32)
		p.Instrucciones()
	}
	{
		p.SetState(33)
		p.Match(GramaticaParserEOF)
	}

	return localctx
}

// IMarcadorContext is an interface to support dynamic dispatch.
type IMarcadorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLista returns the lista attribute.
	GetLista() []string

	// SetLista sets the lista attribute.
	SetLista([]string)

	// IsMarcadorContext differentiates from other interfaces.
	IsMarcadorContext()
}

type MarcadorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	lista  []string
}

func NewEmptyMarcadorContext() *MarcadorContext {
	var p = new(MarcadorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_marcador
	return p
}

func (*MarcadorContext) IsMarcadorContext() {}

func NewMarcadorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, lista []string) *MarcadorContext {
	var p = new(MarcadorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_marcador

	p.lista = lista

	return p
}

func (s *MarcadorContext) GetParser() antlr.Parser { return s.parser }

func (s *MarcadorContext) GetLista() []string { return s.lista }

func (s *MarcadorContext) SetLista(v []string) { s.lista = v }

func (s *MarcadorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MarcadorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Marcador(lista []string) (localctx IMarcadorContext) {
	localctx = NewMarcadorContext(p, p.GetParserRuleContext(), p.GetState(), lista)
	p.EnterRule(localctx, 2, GramaticaParserRULE_marcador)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	gen.Soltar(lista)

	return localctx
}

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOp returns the op token.
	GetOp() antlr.Token

	// GetN returns the n token.
	GetN() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// SetN sets the n token.
	SetN(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// GetOpr returns the opr rule contexts.
	GetOpr() IOprelContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// SetOpr sets the opr rule contexts.
	SetOpr(IOprelContext)

	// GetDir returns the dir attribute.
	GetDir() string

	// GetLv returns the lv attribute.
	GetLv() []string

	// GetLf returns the lf attribute.
	GetLf() []string

	// GetCad returns the cad attribute.
	GetCad() string

	// SetDir sets the dir attribute.
	SetDir(string)

	// SetLv sets the lv attribute.
	SetLv([]string)

	// SetLf sets the lf attribute.
	SetLf([]string)

	// SetCad sets the cad attribute.
	SetCad(string)

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	dir    string
	lv     []string
	lf     []string
	cad    string
	e1     IExpresionContext
	op     antlr.Token
	e      IExpresionContext
	n      antlr.Token
	id     antlr.Token
	e2     IExpresionContext
	opr    IOprelContext
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_expresion
	return p
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) GetOp() antlr.Token { return s.op }

func (s *ExpresionContext) GetN() antlr.Token { return s.n }

func (s *ExpresionContext) GetId() antlr.Token { return s.id }

func (s *ExpresionContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExpresionContext) SetN(v antlr.Token) { s.n = v }

func (s *ExpresionContext) SetId(v antlr.Token) { s.id = v }

func (s *ExpresionContext) GetE1() IExpresionContext { return s.e1 }

func (s *ExpresionContext) GetE() IExpresionContext { return s.e }

func (s *ExpresionContext) GetE2() IExpresionContext { return s.e2 }

func (s *ExpresionContext) GetOpr() IOprelContext { return s.opr }

func (s *ExpresionContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *ExpresionContext) SetE(v IExpresionContext) { s.e = v }

func (s *ExpresionContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *ExpresionContext) SetOpr(v IOprelContext) { s.opr = v }

func (s *ExpresionContext) GetDir() string { return s.dir }

func (s *ExpresionContext) GetLv() []string { return s.lv }

func (s *ExpresionContext) GetLf() []string { return s.lf }

func (s *ExpresionContext) GetCad() string { return s.cad }

func (s *ExpresionContext) SetDir(v string) { s.dir = v }

func (s *ExpresionContext) SetLv(v []string) { s.lv = v }

func (s *ExpresionContext) SetLf(v []string) { s.lf = v }

func (s *ExpresionContext) SetCad(v string) { s.cad = v }

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *ExpresionContext) NUM() antlr.TerminalNode {
	return s.GetToken(GramaticaParserNUM, 0)
}

func (s *ExpresionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *ExpresionContext) Oprel() IOprelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOprelContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOprelContext)
}

func (s *ExpresionContext) Marcador() IMarcadorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMarcadorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMarcadorContext)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *GramaticaParser) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, GramaticaParserRULE_expresion, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__0:
		{
			p.SetState(38)

			var _m = p.Match(GramaticaParserT__0)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(39)

			var _x = p.expresion(14)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.GenOperacionUnaria(localctx.(*ExpresionContext).dir, (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), localctx.(*ExpresionContext).GetE().GetDir())

	case GramaticaParserT__1:
		{
			p.SetState(42)

			var _m = p.Match(GramaticaParserT__1)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(43)

			var _x = p.expresion(13)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).dir = gen.NewTemp()
		gen.GenOperacionUnaria(localctx.(*ExpresionContext).dir, (func() string {
			if localctx.(*ExpresionContext).GetOp() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetOp().GetText()
			}
		}()), localctx.(*ExpresionContext).GetE().GetDir())

	case GramaticaParserT__2:
		{
			p.SetState(46)

			var _m = p.Match(GramaticaParserT__2)

			localctx.(*ExpresionContext).op = _m
		}
		{
			p.SetState(47)

			var _x = p.expresion(12)

			localctx.(*ExpresionContext).e = _x
		}

		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLv()

	case GramaticaParserT__9:
		{
			p.SetState(50)
			p.Match(GramaticaParserT__9)
		}
		{
			p.SetState(51)

			var _x = p.expresion(0)

			localctx.(*ExpresionContext).e = _x
		}
		{
			p.SetState(52)
			p.Match(GramaticaParserT__10)
		}

		localctx.(*ExpresionContext).dir = localctx.(*ExpresionContext).GetE().GetDir()
		localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE().GetLv()
		localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE().GetLf()
		localctx.(*ExpresionContext).cad = localctx.(*ExpresionContext).GetE().GetCad()

	case GramaticaParserNUM:
		{
			p.SetState(55)

			var _m = p.Match(GramaticaParserNUM)

			localctx.(*ExpresionContext).n = _m
		}

		localctx.(*ExpresionContext).dir = (func() string {
			if localctx.(*ExpresionContext).GetN() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetN().GetText()
			}
		}())
		localctx.(*ExpresionContext).cad = ""

	case GramaticaParserID:
		{
			p.SetState(57)

			var _m = p.Match(GramaticaParserID)

			localctx.(*ExpresionContext).id = _m
		}

		localctx.(*ExpresionContext).dir = (func() string {
			if localctx.(*ExpresionContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*ExpresionContext).GetId().GetText()
			}
		}())
		localctx.(*ExpresionContext).cad = ""

	case GramaticaParserT__11:
		{
			p.SetState(59)
			p.Match(GramaticaParserT__11)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewEti())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewEti())
		gen.GenGoto(localctx.(*ExpresionContext).lv[0])

	case GramaticaParserT__12:
		{
			p.SetState(61)
			p.Match(GramaticaParserT__12)
		}

		localctx.(*ExpresionContext).dir = ""
		localctx.(*ExpresionContext).cad = ""
		localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewEti())
		localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewEti())
		gen.GenGoto(localctx.(*ExpresionContext).lf[0])

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(65)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(66)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GramaticaParserT__3)|(1<<GramaticaParserT__4)|(1<<GramaticaParserT__5))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(67)

					var _x = p.expresion(12)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).dir = gen.NewTemp()
				gen.GenOperacion(localctx.(*ExpresionContext).dir, localctx.(*ExpresionContext).GetE1().GetDir(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetE2().GetDir())

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(70)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(71)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*ExpresionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GramaticaParserT__0 || _la == GramaticaParserT__1) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*ExpresionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(72)

					var _x = p.expresion(11)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).dir = gen.NewTemp()
				gen.GenOperacion(localctx.(*ExpresionContext).dir, localctx.(*ExpresionContext).GetE1().GetDir(), (func() string {
					if localctx.(*ExpresionContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*ExpresionContext).GetOp().GetText()
					}
				}()), localctx.(*ExpresionContext).GetE2().GetDir())

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(75)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(76)

					var _x = p.Oprel()

					localctx.(*ExpresionContext).opr = _x
				}
				{
					p.SetState(77)

					var _x = p.expresion(10)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lv = append(localctx.(*ExpresionContext).lv, gen.NewEti())
				localctx.(*ExpresionContext).lf = append(localctx.(*ExpresionContext).lf, gen.NewEti())
				localctx.(*ExpresionContext).cad = gen.GenIf(localctx.(*ExpresionContext).GetE1().GetDir(), localctx.(*ExpresionContext).GetOpr().GetOp(), localctx.(*ExpresionContext).GetE2().GetDir(), localctx.(*ExpresionContext).lv[0])
				gen.GenGoto(localctx.(*ExpresionContext).lf[0])

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(81)

					var _m = p.Match(GramaticaParserT__6)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(82)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLv())
				}
				{
					p.SetState(83)

					var _x = p.expresion(9)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE2().GetLv()
				localctx.(*ExpresionContext).lf = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLf(), localctx.(*ExpresionContext).GetE2().GetLf())

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(87)

					var _m = p.Match(GramaticaParserT__7)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(88)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(89)

					var _x = p.expresion(8)

					localctx.(*ExpresionContext).e2 = _x
				}

				gen.Soltar(localctx.(*ExpresionContext).GetE1().GetLv())
				gen.GenIfCad(localctx.(*ExpresionContext).GetE2().GetCad(), localctx.(*ExpresionContext).GetE2().GetLf()[0])
				gen.GenGoto("goto " + localctx.(*ExpresionContext).GetE2().GetLv()[0])
				localctx.(*ExpresionContext).lv = localctx.(*ExpresionContext).GetE2().GetLv()
				localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE2().GetLf()

			case 6:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				localctx.(*ExpresionContext).e1 = _prevctx
				p.PushNewRecursionContext(localctx, _startState, GramaticaParserRULE_expresion)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(93)

					var _m = p.Match(GramaticaParserT__8)

					localctx.(*ExpresionContext).op = _m
				}
				{
					p.SetState(94)
					p.Marcador(localctx.(*ExpresionContext).GetE1().GetLf())
				}
				{
					p.SetState(95)

					var _x = p.expresion(7)

					localctx.(*ExpresionContext).e2 = _x
				}

				localctx.(*ExpresionContext).lf = localctx.(*ExpresionContext).GetE2().GetLf()
				localctx.(*ExpresionContext).lv = gen.Unir(localctx.(*ExpresionContext).GetE1().GetLv(), localctx.(*ExpresionContext).GetE2().GetLv())

			}

		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

// IOprelContext is an interface to support dynamic dispatch.
type IOprelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOpe returns the ope token.
	GetOpe() antlr.Token

	// SetOpe sets the ope token.
	SetOpe(antlr.Token)

	// GetOp returns the op attribute.
	GetOp() string

	// SetOp sets the op attribute.
	SetOp(string)

	// IsOprelContext differentiates from other interfaces.
	IsOprelContext()
}

type OprelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	op     string
	ope    antlr.Token
}

func NewEmptyOprelContext() *OprelContext {
	var p = new(OprelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_oprel
	return p
}

func (*OprelContext) IsOprelContext() {}

func NewOprelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OprelContext {
	var p = new(OprelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_oprel

	return p
}

func (s *OprelContext) GetParser() antlr.Parser { return s.parser }

func (s *OprelContext) GetOpe() antlr.Token { return s.ope }

func (s *OprelContext) SetOpe(v antlr.Token) { s.ope = v }

func (s *OprelContext) GetOp() string { return s.op }

func (s *OprelContext) SetOp(v string) { s.op = v }

func (s *OprelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OprelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Oprel() (localctx IOprelContext) {
	localctx = NewOprelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GramaticaParserRULE_oprel)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(103)

			var _m = p.Match(GramaticaParserT__13)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)

			var _m = p.Match(GramaticaParserT__14)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__15:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(107)

			var _m = p.Match(GramaticaParserT__15)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__16:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(109)

			var _m = p.Match(GramaticaParserT__16)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__17:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(111)

			var _m = p.Match(GramaticaParserT__17)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	case GramaticaParserT__18:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(113)

			var _m = p.Match(GramaticaParserT__18)

			localctx.(*OprelContext).ope = _m
		}
		localctx.(*OprelContext).op = (func() string {
			if localctx.(*OprelContext).GetOpe() == nil {
				return ""
			} else {
				return localctx.(*OprelContext).GetOpe().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Instrucciones() (localctx IInstruccionesContext) {
	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GramaticaParserRULE_instrucciones)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-22)&-(0x1f+1)) == 0 && ((1<<uint((_la-22)))&((1<<(GramaticaParserT__21-22))|(1<<(GramaticaParserT__23-22))|(1<<(GramaticaParserT__29-22))|(1<<(GramaticaParserT__30-22))|(1<<(GramaticaParserT__31-22))|(1<<(GramaticaParserT__32-22))|(1<<(GramaticaParserID-22)))) != 0 {
		{
			p.SetState(117)
			p.Instruccion()
		}

		p.SetState(122)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Inst_asignacion() IInst_asignacionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_asignacionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_asignacionContext)
}

func (s *InstruccionContext) Inst_if() IInst_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_ifContext)
}

func (s *InstruccionContext) Inst_switch_propuesta2() IInst_switch_propuesta2Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_switch_propuesta2Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_switch_propuesta2Context)
}

func (s *InstruccionContext) Inst_while() IInst_whileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_whileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_whileContext)
}

func (s *InstruccionContext) Inst_doWhile() IInst_doWhileContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_doWhileContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_doWhileContext)
}

func (s *InstruccionContext) Inst_loop() IInst_loopContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_loopContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_loopContext)
}

func (s *InstruccionContext) Inst_for() IInst_forContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInst_forContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInst_forContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Instruccion() (localctx IInstruccionContext) {
	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GramaticaParserRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(130)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GramaticaParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(123)
			p.Inst_asignacion()
		}

	case GramaticaParserT__21:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Inst_if()
		}

	case GramaticaParserT__23:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(125)
			p.Inst_switch_propuesta2()
		}

	case GramaticaParserT__29:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(126)
			p.Inst_while()
		}

	case GramaticaParserT__30:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(127)
			p.Inst_doWhile()
		}

	case GramaticaParserT__31:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(128)
			p.Inst_loop()
		}

	case GramaticaParserT__32:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(129)
			p.Inst_for()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IInst_asignacionContext is an interface to support dynamic dispatch.
type IInst_asignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// IsInst_asignacionContext differentiates from other interfaces.
	IsInst_asignacionContext()
}

type Inst_asignacionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	id     antlr.Token
	e      IExpresionContext
}

func NewEmptyInst_asignacionContext() *Inst_asignacionContext {
	var p = new(Inst_asignacionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_asignacion
	return p
}

func (*Inst_asignacionContext) IsInst_asignacionContext() {}

func NewInst_asignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_asignacionContext {
	var p = new(Inst_asignacionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_asignacion

	return p
}

func (s *Inst_asignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_asignacionContext) GetId() antlr.Token { return s.id }

func (s *Inst_asignacionContext) SetId(v antlr.Token) { s.id = v }

func (s *Inst_asignacionContext) GetE() IExpresionContext { return s.e }

func (s *Inst_asignacionContext) SetE(v IExpresionContext) { s.e = v }

func (s *Inst_asignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *Inst_asignacionContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_asignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_asignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Inst_asignacion() (localctx IInst_asignacionContext) {
	localctx = NewInst_asignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GramaticaParserRULE_inst_asignacion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(132)

		var _m = p.Match(GramaticaParserID)

		localctx.(*Inst_asignacionContext).id = _m
	}
	{
		p.SetState(133)
		p.Match(GramaticaParserT__19)
	}
	{
		p.SetState(134)

		var _x = p.expresion(0)

		localctx.(*Inst_asignacionContext).e = _x
	}
	{
		p.SetState(135)
		p.Match(GramaticaParserT__20)
	}
	gen.GenAsignacion((func() string {
		if localctx.(*Inst_asignacionContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_asignacionContext).GetId().GetText()
		}
	}()), localctx.(*Inst_asignacionContext).GetE().GetDir())

	return localctx
}

// IInst_ifContext is an interface to support dynamic dispatch.
type IInst_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// IsInst_ifContext differentiates from other interfaces.
	IsInst_ifContext()
}

type Inst_ifContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lsalida string
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_ifContext() *Inst_ifContext {
	var p = new(Inst_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_if
	return p
}

func (*Inst_ifContext) IsInst_ifContext() {}

func NewInst_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_ifContext {
	var p = new(Inst_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_if

	return p
}

func (s *Inst_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_ifContext) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_ifContext) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_ifContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_ifContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_ifContext) GetLsalida() string { return s.lsalida }

func (s *Inst_ifContext) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_ifContext) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Inst_ifContext) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_ifContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_ifContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Inst_if() (localctx IInst_ifContext) {
	localctx = NewInst_ifContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GramaticaParserRULE_inst_if)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(138)
		p.Match(GramaticaParserT__21)
	}
	{
		p.SetState(139)

		var _x = p.expresion(0)

		localctx.(*Inst_ifContext).e1 = _x
	}

	localctx.(*Inst_ifContext).lsalida = gen.NewEti()
	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLv())

	{
		p.SetState(141)
		p.Bloque()
	}

	gen.GenGoto(localctx.(*Inst_ifContext).lsalida)
	gen.Soltar(localctx.(*Inst_ifContext).GetE1().GetLf())

	p.SetState(152)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(143)
				p.Match(GramaticaParserT__22)
			}
			{
				p.SetState(144)
				p.Match(GramaticaParserT__21)
			}
			{
				p.SetState(145)

				var _x = p.expresion(0)

				localctx.(*Inst_ifContext).e2 = _x
			}

			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLv())

			{
				p.SetState(147)
				p.Bloque()
			}

			gen.GenGoto(localctx.(*Inst_ifContext).lsalida)
			gen.Soltar(localctx.(*Inst_ifContext).GetE2().GetLf())

		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__22 {
		{
			p.SetState(155)
			p.Match(GramaticaParserT__22)
		}
		{
			p.SetState(156)
			p.Bloque()
		}

	}
	gen.GenDestino(localctx.(*Inst_ifContext).lsalida)

	return localctx
}

// IInst_switch_propuesta1Context is an interface to support dynamic dispatch.
type IInst_switch_propuesta1Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// GetLv returns the lv attribute.
	GetLv() string

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// SetLv sets the lv attribute.
	SetLv(string)

	// IsInst_switch_propuesta1Context differentiates from other interfaces.
	IsInst_switch_propuesta1Context()
}

type Inst_switch_propuesta1Context struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lsalida string
	lv      string
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_switch_propuesta1Context() *Inst_switch_propuesta1Context {
	var p = new(Inst_switch_propuesta1Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_switch_propuesta1
	return p
}

func (*Inst_switch_propuesta1Context) IsInst_switch_propuesta1Context() {}

func NewInst_switch_propuesta1Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_switch_propuesta1Context {
	var p = new(Inst_switch_propuesta1Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_switch_propuesta1

	return p
}

func (s *Inst_switch_propuesta1Context) GetParser() antlr.Parser { return s.parser }

func (s *Inst_switch_propuesta1Context) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_switch_propuesta1Context) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_switch_propuesta1Context) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_switch_propuesta1Context) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_switch_propuesta1Context) GetLsalida() string { return s.lsalida }

func (s *Inst_switch_propuesta1Context) GetLv() string { return s.lv }

func (s *Inst_switch_propuesta1Context) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_switch_propuesta1Context) SetLv(v string) { s.lv = v }

func (s *Inst_switch_propuesta1Context) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta1Context) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_switch_propuesta1Context) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta1Context) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_switch_propuesta1Context) AllBloqueSinLLaves() []IBloqueSinLLavesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueSinLLavesContext)(nil)).Elem())
	var tst = make([]IBloqueSinLLavesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueSinLLavesContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta1Context) BloqueSinLLaves(i int) IBloqueSinLLavesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueSinLLavesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueSinLLavesContext)
}

func (s *Inst_switch_propuesta1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_switch_propuesta1Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Inst_switch_propuesta1() (localctx IInst_switch_propuesta1Context) {
	localctx = NewInst_switch_propuesta1Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GramaticaParserRULE_inst_switch_propuesta1)

	localctx.(*Inst_switch_propuesta1Context).lsalida = gen.NewEti()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(161)
		p.Match(GramaticaParserT__23)
	}
	{
		p.SetState(162)

		var _x = p.expresion(0)

		localctx.(*Inst_switch_propuesta1Context).e1 = _x
	}
	{
		p.SetState(163)
		p.Match(GramaticaParserT__24)
	}
	p.SetState(174)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramaticaParserT__25 {
		{
			p.SetState(164)
			p.Match(GramaticaParserT__25)
		}
		{
			p.SetState(165)

			var _x = p.expresion(0)

			localctx.(*Inst_switch_propuesta1Context).e2 = _x
		}
		{
			p.SetState(166)
			p.Match(GramaticaParserT__26)
		}

		localctx.(*Inst_switch_propuesta1Context).lv = gen.NewEti()
		gen.GenIf(localctx.(*Inst_switch_propuesta1Context).GetE1().GetDir(), "!=", localctx.(*Inst_switch_propuesta1Context).GetE2().GetDir(), localctx.(*Inst_switch_propuesta1Context).lv)

		p.SetState(170)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__24:
			{
				p.SetState(168)
				p.Bloque()
			}

		case GramaticaParserT__21, GramaticaParserT__23, GramaticaParserT__25, GramaticaParserT__27, GramaticaParserT__28, GramaticaParserT__29, GramaticaParserT__30, GramaticaParserT__31, GramaticaParserT__32, GramaticaParserID:
			{
				p.SetState(169)
				p.BloqueSinLLaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		gen.GenGoto(localctx.(*Inst_switch_propuesta1Context).lsalida)
		gen.GenDestino(localctx.(*Inst_switch_propuesta1Context).lv)

		p.SetState(176)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__27 {
		{
			p.SetState(178)
			p.Match(GramaticaParserT__27)
		}
		{
			p.SetState(179)
			p.Match(GramaticaParserT__26)
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__24:
			{
				p.SetState(180)
				p.Bloque()
			}

		case GramaticaParserT__21, GramaticaParserT__23, GramaticaParserT__28, GramaticaParserT__29, GramaticaParserT__30, GramaticaParserT__31, GramaticaParserT__32, GramaticaParserID:
			{
				p.SetState(181)
				p.BloqueSinLLaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

	}
	gen.GenDestino(localctx.(*Inst_switch_propuesta1Context).lsalida)
	{
		p.SetState(187)
		p.Match(GramaticaParserT__28)
	}

	return localctx
}

// IInst_switch_propuesta2Context is an interface to support dynamic dispatch.
type IInst_switch_propuesta2Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetLprueba returns the lprueba attribute.
	GetLprueba() string

	// GetCad returns the cad attribute.
	GetCad() string

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// GetLv returns the lv attribute.
	GetLv() string

	// GetDefecto returns the defecto attribute.
	GetDefecto() bool

	// SetLprueba sets the lprueba attribute.
	SetLprueba(string)

	// SetCad sets the cad attribute.
	SetCad(string)

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// SetLv sets the lv attribute.
	SetLv(string)

	// SetDefecto sets the defecto attribute.
	SetDefecto(bool)

	// IsInst_switch_propuesta2Context differentiates from other interfaces.
	IsInst_switch_propuesta2Context()
}

type Inst_switch_propuesta2Context struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	lprueba string
	cad     string
	lsalida string
	lv      string
	defecto bool
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_switch_propuesta2Context() *Inst_switch_propuesta2Context {
	var p = new(Inst_switch_propuesta2Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_switch_propuesta2
	return p
}

func (*Inst_switch_propuesta2Context) IsInst_switch_propuesta2Context() {}

func NewInst_switch_propuesta2Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_switch_propuesta2Context {
	var p = new(Inst_switch_propuesta2Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_switch_propuesta2

	return p
}

func (s *Inst_switch_propuesta2Context) GetParser() antlr.Parser { return s.parser }

func (s *Inst_switch_propuesta2Context) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_switch_propuesta2Context) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_switch_propuesta2Context) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_switch_propuesta2Context) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_switch_propuesta2Context) GetLprueba() string { return s.lprueba }

func (s *Inst_switch_propuesta2Context) GetCad() string { return s.cad }

func (s *Inst_switch_propuesta2Context) GetLsalida() string { return s.lsalida }

func (s *Inst_switch_propuesta2Context) GetLv() string { return s.lv }

func (s *Inst_switch_propuesta2Context) GetDefecto() bool { return s.defecto }

func (s *Inst_switch_propuesta2Context) SetLprueba(v string) { s.lprueba = v }

func (s *Inst_switch_propuesta2Context) SetCad(v string) { s.cad = v }

func (s *Inst_switch_propuesta2Context) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_switch_propuesta2Context) SetLv(v string) { s.lv = v }

func (s *Inst_switch_propuesta2Context) SetDefecto(v bool) { s.defecto = v }

func (s *Inst_switch_propuesta2Context) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta2Context) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_switch_propuesta2Context) AllBloque() []IBloqueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueContext)(nil)).Elem())
	var tst = make([]IBloqueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta2Context) Bloque(i int) IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_switch_propuesta2Context) AllBloqueSinLLaves() []IBloqueSinLLavesContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBloqueSinLLavesContext)(nil)).Elem())
	var tst = make([]IBloqueSinLLavesContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBloqueSinLLavesContext)
		}
	}

	return tst
}

func (s *Inst_switch_propuesta2Context) BloqueSinLLaves(i int) IBloqueSinLLavesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueSinLLavesContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBloqueSinLLavesContext)
}

func (s *Inst_switch_propuesta2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_switch_propuesta2Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Inst_switch_propuesta2() (localctx IInst_switch_propuesta2Context) {
	localctx = NewInst_switch_propuesta2Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GramaticaParserRULE_inst_switch_propuesta2)

	localctx.(*Inst_switch_propuesta2Context).lprueba = gen.NewEti()
	localctx.(*Inst_switch_propuesta2Context).lsalida = gen.NewEti()
	localctx.(*Inst_switch_propuesta2Context).cad = ""
	localctx.(*Inst_switch_propuesta2Context).defecto = false

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(189)
		p.Match(GramaticaParserT__23)
	}
	{
		p.SetState(190)

		var _x = p.expresion(0)

		localctx.(*Inst_switch_propuesta2Context).e1 = _x
	}
	{
		p.SetState(191)
		p.Match(GramaticaParserT__24)
	}

	gen.GenGoto(localctx.(*Inst_switch_propuesta2Context).lprueba)

	p.SetState(203)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GramaticaParserT__25 {
		{
			p.SetState(193)
			p.Match(GramaticaParserT__25)
		}
		{
			p.SetState(194)

			var _x = p.expresion(0)

			localctx.(*Inst_switch_propuesta2Context).e2 = _x
		}
		{
			p.SetState(195)
			p.Match(GramaticaParserT__26)
		}

		localctx.(*Inst_switch_propuesta2Context).lv = gen.NewEti()
		gen.GenDestino(localctx.(*Inst_switch_propuesta2Context).lv)
		localctx.(*Inst_switch_propuesta2Context).cad += "if " + localctx.(*Inst_switch_propuesta2Context).GetE1().GetDir() + " = " + localctx.(*Inst_switch_propuesta2Context).GetE2().GetDir() + " then goto " + localctx.(*Inst_switch_propuesta2Context).lv + "\n"

		p.SetState(199)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__24:
			{
				p.SetState(197)
				p.Bloque()
			}

		case GramaticaParserT__21, GramaticaParserT__23, GramaticaParserT__25, GramaticaParserT__27, GramaticaParserT__28, GramaticaParserT__29, GramaticaParserT__30, GramaticaParserT__31, GramaticaParserT__32, GramaticaParserID:
			{
				p.SetState(198)
				p.BloqueSinLLaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		gen.GenGoto(localctx.(*Inst_switch_propuesta2Context).lsalida)

		p.SetState(205)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(216)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GramaticaParserT__27 {
		{
			p.SetState(207)
			p.Match(GramaticaParserT__27)
		}
		{
			p.SetState(208)
			p.Match(GramaticaParserT__26)
		}

		localctx.(*Inst_switch_propuesta2Context).lv = gen.NewEti()
		localctx.(*Inst_switch_propuesta2Context).defecto = true
		gen.GenDestino(localctx.(*Inst_switch_propuesta2Context).lv)

		p.SetState(212)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case GramaticaParserT__24:
			{
				p.SetState(210)
				p.Bloque()
			}

		case GramaticaParserT__21, GramaticaParserT__23, GramaticaParserT__28, GramaticaParserT__29, GramaticaParserT__30, GramaticaParserT__31, GramaticaParserT__32, GramaticaParserID:
			{
				p.SetState(211)
				p.BloqueSinLLaves()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		gen.GenGoto(localctx.(*Inst_switch_propuesta2Context).lsalida)

	}

	gen.GenDestino(localctx.(*Inst_switch_propuesta2Context).lprueba)
	gen.Gen(localctx.(*Inst_switch_propuesta2Context).cad)
	if localctx.(*Inst_switch_propuesta2Context).defecto {
		gen.GenGoto(localctx.(*Inst_switch_propuesta2Context).lv)
	}
	gen.GenDestino(localctx.(*Inst_switch_propuesta2Context).lsalida)

	{
		p.SetState(219)
		p.Match(GramaticaParserT__28)
	}

	return localctx
}

// IInst_whileContext is an interface to support dynamic dispatch.
type IInst_whileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// GetLinicio returns the linicio attribute.
	GetLinicio() string

	// SetLinicio sets the linicio attribute.
	SetLinicio(string)

	// IsInst_whileContext differentiates from other interfaces.
	IsInst_whileContext()
}

type Inst_whileContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	linicio string
	e       IExpresionContext
}

func NewEmptyInst_whileContext() *Inst_whileContext {
	var p = new(Inst_whileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_while
	return p
}

func (*Inst_whileContext) IsInst_whileContext() {}

func NewInst_whileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_whileContext {
	var p = new(Inst_whileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_while

	return p
}

func (s *Inst_whileContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_whileContext) GetE() IExpresionContext { return s.e }

func (s *Inst_whileContext) SetE(v IExpresionContext) { s.e = v }

func (s *Inst_whileContext) GetLinicio() string { return s.linicio }

func (s *Inst_whileContext) SetLinicio(v string) { s.linicio = v }

func (s *Inst_whileContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_whileContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_whileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_whileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Inst_while() (localctx IInst_whileContext) {
	localctx = NewInst_whileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GramaticaParserRULE_inst_while)

	localctx.(*Inst_whileContext).linicio = gen.NewEti()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(221)
		p.Match(GramaticaParserT__29)
	}
	gen.GenDestino(localctx.(*Inst_whileContext).linicio)
	{
		p.SetState(223)

		var _x = p.expresion(0)

		localctx.(*Inst_whileContext).e = _x
	}
	gen.Soltar(localctx.(*Inst_whileContext).GetE().GetLv())
	{
		p.SetState(225)
		p.Bloque()
	}

	gen.GenGoto(localctx.(*Inst_whileContext).linicio)
	gen.Soltar(localctx.(*Inst_whileContext).GetE().GetLf())

	return localctx
}

// IInst_doWhileContext is an interface to support dynamic dispatch.
type IInst_doWhileContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetE returns the e rule contexts.
	GetE() IExpresionContext

	// SetE sets the e rule contexts.
	SetE(IExpresionContext)

	// GetLinicio returns the linicio attribute.
	GetLinicio() string

	// SetLinicio sets the linicio attribute.
	SetLinicio(string)

	// IsInst_doWhileContext differentiates from other interfaces.
	IsInst_doWhileContext()
}

type Inst_doWhileContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	linicio string
	e       IExpresionContext
}

func NewEmptyInst_doWhileContext() *Inst_doWhileContext {
	var p = new(Inst_doWhileContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_doWhile
	return p
}

func (*Inst_doWhileContext) IsInst_doWhileContext() {}

func NewInst_doWhileContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_doWhileContext {
	var p = new(Inst_doWhileContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_doWhile

	return p
}

func (s *Inst_doWhileContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_doWhileContext) GetE() IExpresionContext { return s.e }

func (s *Inst_doWhileContext) SetE(v IExpresionContext) { s.e = v }

func (s *Inst_doWhileContext) GetLinicio() string { return s.linicio }

func (s *Inst_doWhileContext) SetLinicio(v string) { s.linicio = v }

func (s *Inst_doWhileContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_doWhileContext) Expresion() IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_doWhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_doWhileContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Inst_doWhile() (localctx IInst_doWhileContext) {
	localctx = NewInst_doWhileContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GramaticaParserRULE_inst_doWhile)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(GramaticaParserT__30)
	}

	localctx.(*Inst_doWhileContext).linicio = gen.NewEti()
	gen.GenDestino(localctx.(*Inst_doWhileContext).linicio)

	{
		p.SetState(230)
		p.Bloque()
	}
	{
		p.SetState(231)
		p.Match(GramaticaParserT__29)
	}
	{
		p.SetState(232)

		var _x = p.expresion(0)

		localctx.(*Inst_doWhileContext).e = _x
	}
	{
		p.SetState(233)
		p.Match(GramaticaParserT__20)
	}

	gen.Soltar(localctx.(*Inst_doWhileContext).GetE().GetLv())
	gen.GenGoto(localctx.(*Inst_doWhileContext).linicio)
	gen.Soltar(localctx.(*Inst_doWhileContext).GetE().GetLf())

	return localctx
}

// IInst_loopContext is an interface to support dynamic dispatch.
type IInst_loopContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLinicio returns the linicio attribute.
	GetLinicio() string

	// SetLinicio sets the linicio attribute.
	SetLinicio(string)

	// IsInst_loopContext differentiates from other interfaces.
	IsInst_loopContext()
}

type Inst_loopContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	linicio string
}

func NewEmptyInst_loopContext() *Inst_loopContext {
	var p = new(Inst_loopContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_loop
	return p
}

func (*Inst_loopContext) IsInst_loopContext() {}

func NewInst_loopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_loopContext {
	var p = new(Inst_loopContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_loop

	return p
}

func (s *Inst_loopContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_loopContext) GetLinicio() string { return s.linicio }

func (s *Inst_loopContext) SetLinicio(v string) { s.linicio = v }

func (s *Inst_loopContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_loopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_loopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Inst_loop() (localctx IInst_loopContext) {
	localctx = NewInst_loopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GramaticaParserRULE_inst_loop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(236)
		p.Match(GramaticaParserT__31)
	}

	localctx.(*Inst_loopContext).linicio = gen.NewEti()
	gen.GenDestino(localctx.(*Inst_loopContext).linicio)

	{
		p.SetState(238)
		p.Bloque()
	}

	gen.GenGoto(localctx.(*Inst_loopContext).linicio)

	return localctx
}

// IInst_forContext is an interface to support dynamic dispatch.
type IInst_forContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetE1 returns the e1 rule contexts.
	GetE1() IExpresionContext

	// GetE2 returns the e2 rule contexts.
	GetE2() IExpresionContext

	// SetE1 sets the e1 rule contexts.
	SetE1(IExpresionContext)

	// SetE2 sets the e2 rule contexts.
	SetE2(IExpresionContext)

	// GetLinicio returns the linicio attribute.
	GetLinicio() string

	// GetLsalida returns the lsalida attribute.
	GetLsalida() string

	// SetLinicio sets the linicio attribute.
	SetLinicio(string)

	// SetLsalida sets the lsalida attribute.
	SetLsalida(string)

	// IsInst_forContext differentiates from other interfaces.
	IsInst_forContext()
}

type Inst_forContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	linicio string
	lsalida string
	id      antlr.Token
	e1      IExpresionContext
	e2      IExpresionContext
}

func NewEmptyInst_forContext() *Inst_forContext {
	var p = new(Inst_forContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_inst_for
	return p
}

func (*Inst_forContext) IsInst_forContext() {}

func NewInst_forContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Inst_forContext {
	var p = new(Inst_forContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_inst_for

	return p
}

func (s *Inst_forContext) GetParser() antlr.Parser { return s.parser }

func (s *Inst_forContext) GetId() antlr.Token { return s.id }

func (s *Inst_forContext) SetId(v antlr.Token) { s.id = v }

func (s *Inst_forContext) GetE1() IExpresionContext { return s.e1 }

func (s *Inst_forContext) GetE2() IExpresionContext { return s.e2 }

func (s *Inst_forContext) SetE1(v IExpresionContext) { s.e1 = v }

func (s *Inst_forContext) SetE2(v IExpresionContext) { s.e2 = v }

func (s *Inst_forContext) GetLinicio() string { return s.linicio }

func (s *Inst_forContext) GetLsalida() string { return s.lsalida }

func (s *Inst_forContext) SetLinicio(v string) { s.linicio = v }

func (s *Inst_forContext) SetLsalida(v string) { s.lsalida = v }

func (s *Inst_forContext) Bloque() IBloqueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBloqueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *Inst_forContext) ID() antlr.TerminalNode {
	return s.GetToken(GramaticaParserID, 0)
}

func (s *Inst_forContext) AllExpresion() []IExpresionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpresionContext)(nil)).Elem())
	var tst = make([]IExpresionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpresionContext)
		}
	}

	return tst
}

func (s *Inst_forContext) Expresion(i int) IExpresionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpresionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *Inst_forContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Inst_forContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Inst_for() (localctx IInst_forContext) {
	localctx = NewInst_forContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GramaticaParserRULE_inst_for)

	localctx.(*Inst_forContext).linicio = gen.NewEti()
	localctx.(*Inst_forContext).lsalida = gen.NewEti()

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(241)
		p.Match(GramaticaParserT__32)
	}
	{
		p.SetState(242)

		var _m = p.Match(GramaticaParserID)

		localctx.(*Inst_forContext).id = _m
	}
	{
		p.SetState(243)
		p.Match(GramaticaParserT__19)
	}
	{
		p.SetState(244)

		var _x = p.expresion(0)

		localctx.(*Inst_forContext).e1 = _x
	}
	{
		p.SetState(245)
		p.Match(GramaticaParserT__33)
	}
	{
		p.SetState(246)

		var _x = p.expresion(0)

		localctx.(*Inst_forContext).e2 = _x
	}

	gen.GenAsignacion((func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}()), localctx.(*Inst_forContext).GetE1().GetDir())
	gen.GenDestino(localctx.(*Inst_forContext).linicio)
	gen.GenIfCad((func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}())+" > "+localctx.(*Inst_forContext).GetE2().GetDir(), localctx.(*Inst_forContext).lsalida)

	{
		p.SetState(248)
		p.Match(GramaticaParserT__30)
	}
	{
		p.SetState(249)
		p.Bloque()
	}

	gen.GenAsignacion((func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}()), (func() string {
		if localctx.(*Inst_forContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Inst_forContext).GetId().GetText()
		}
	}())+"+1")
	gen.GenGoto(localctx.(*Inst_forContext).linicio)
	gen.GenDestino(localctx.(*Inst_forContext).lsalida)

	return localctx
}

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_bloque
	return p
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GramaticaParserRULE_bloque)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(GramaticaParserT__24)
	}
	{
		p.SetState(253)
		p.Instrucciones()
	}
	{
		p.SetState(254)
		p.Match(GramaticaParserT__28)
	}

	return localctx
}

// IBloqueSinLLavesContext is an interface to support dynamic dispatch.
type IBloqueSinLLavesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBloqueSinLLavesContext differentiates from other interfaces.
	IsBloqueSinLLavesContext()
}

type BloqueSinLLavesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueSinLLavesContext() *BloqueSinLLavesContext {
	var p = new(BloqueSinLLavesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GramaticaParserRULE_bloqueSinLLaves
	return p
}

func (*BloqueSinLLavesContext) IsBloqueSinLLavesContext() {}

func NewBloqueSinLLavesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueSinLLavesContext {
	var p = new(BloqueSinLLavesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GramaticaParserRULE_bloqueSinLLaves

	return p
}

func (s *BloqueSinLLavesContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueSinLLavesContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *BloqueSinLLavesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueSinLLavesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *GramaticaParser) BloqueSinLLaves() (localctx IBloqueSinLLavesContext) {
	localctx = NewBloqueSinLLavesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GramaticaParserRULE_bloqueSinLLaves)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(256)
		p.Instrucciones()
	}

	return localctx
}

func (p *GramaticaParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GramaticaParser) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
