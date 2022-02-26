// Code generated from Gramatica.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type GramaticaListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterMarcador is called when entering the marcador production.
	EnterMarcador(c *MarcadorContext)

	// EnterDec is called when entering the dec production.
	EnterDec(c *DecContext)

	// EnterTipo is called when entering the tipo production.
	EnterTipo(c *TipoContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitMarcador is called when exiting the marcador production.
	ExitMarcador(c *MarcadorContext)

	// ExitDec is called when exiting the dec production.
	ExitDec(c *DecContext)

	// ExitTipo is called when exiting the tipo production.
	ExitTipo(c *TipoContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)
}
