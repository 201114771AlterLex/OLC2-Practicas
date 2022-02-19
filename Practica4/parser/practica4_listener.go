// Code generated from Practica4.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Practica4

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Practica4Listener is a complete listener for a parse tree produced by Practica4Parser.
type Practica4Listener interface {
	antlr.ParseTreeListener

	// EnterInicio is called when entering the Inicio production.
	EnterInicio(c *InicioContext)

	// EnterC1 is called when entering the C1 production.
	EnterC1(c *C1Context)

	// EnterC2 is called when entering the C2 production.
	EnterC2(c *C2Context)

	// EnterC3 is called when entering the C3 production.
	EnterC3(c *C3Context)

	// EnterCordenada is called when entering the Cordenada production.
	EnterCordenada(c *CordenadaContext)

	// ExitInicio is called when exiting the Inicio production.
	ExitInicio(c *InicioContext)

	// ExitC1 is called when exiting the C1 production.
	ExitC1(c *C1Context)

	// ExitC2 is called when exiting the C2 production.
	ExitC2(c *C2Context)

	// ExitC3 is called when exiting the C3 production.
	ExitC3(c *C3Context)

	// ExitCordenada is called when exiting the Cordenada production.
	ExitCordenada(c *CordenadaContext)
}
