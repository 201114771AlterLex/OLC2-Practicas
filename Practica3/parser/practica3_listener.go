// Code generated from Practica3.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Practica3

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Practica3Listener is a complete listener for a parse tree produced by Practica3Parser.
type Practica3Listener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterE is called when entering the e production.
	EnterE(c *EContext)

	// EnterEp is called when entering the ep production.
	EnterEp(c *EpContext)

	// EnterT is called when entering the t production.
	EnterT(c *TContext)

	// EnterTp is called when entering the tp production.
	EnterTp(c *TpContext)

	// EnterF is called when entering the f production.
	EnterF(c *FContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitE is called when exiting the e production.
	ExitE(c *EContext)

	// ExitEp is called when exiting the ep production.
	ExitEp(c *EpContext)

	// ExitT is called when exiting the t production.
	ExitT(c *TContext)

	// ExitTp is called when exiting the tp production.
	ExitTp(c *TpContext)

	// ExitF is called when exiting the f production.
	ExitF(c *FContext)
}
