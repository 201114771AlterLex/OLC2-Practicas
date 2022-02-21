// Code generated from Evaluacion1.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Evaluacion1

import "github.com/antlr/antlr4/runtime/Go/antlr"

// Evaluacion1Listener is a complete listener for a parse tree produced by Evaluacion1Parser.
type Evaluacion1Listener interface {
	antlr.ParseTreeListener

	// EnterInicio is called when entering the Inicio production.
	EnterInicio(c *InicioContext)

	// EnterDosNumero is called when entering the DosNumero production.
	EnterDosNumero(c *DosNumeroContext)

	// EnterHTPT is called when entering the HTPT production.
	EnterHTPT(c *HTPTContext)

	// EnterHT is called when entering the HT production.
	EnterHT(c *HTContext)

	// EnterTHex is called when entering the THex production.
	EnterTHex(c *THexContext)

	// EnterTpHex is called when entering the TpHex production.
	EnterTpHex(c *TpHexContext)

	// EnterEpTp is called when entering the EpTp production.
	EnterEpTp(c *EpTpContext)

	// ExitInicio is called when exiting the Inicio production.
	ExitInicio(c *InicioContext)

	// ExitDosNumero is called when exiting the DosNumero production.
	ExitDosNumero(c *DosNumeroContext)

	// ExitHTPT is called when exiting the HTPT production.
	ExitHTPT(c *HTPTContext)

	// ExitHT is called when exiting the HT production.
	ExitHT(c *HTContext)

	// ExitTHex is called when exiting the THex production.
	ExitTHex(c *THexContext)

	// ExitTpHex is called when exiting the TpHex production.
	ExitTpHex(c *TpHexContext)

	// ExitEpTp is called when exiting the EpTp production.
	ExitEpTp(c *EpTpContext)
}
