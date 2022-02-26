// Code generated from Gramatica.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Gramatica

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseGramaticaListener is a complete listener for a parse tree produced by GramaticaParser.
type BaseGramaticaListener struct{}

var _ GramaticaListener = &BaseGramaticaListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGramaticaListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGramaticaListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGramaticaListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGramaticaListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseGramaticaListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseGramaticaListener) ExitStart(ctx *StartContext) {}

// EnterMarcador is called when production marcador is entered.
func (s *BaseGramaticaListener) EnterMarcador(ctx *MarcadorContext) {}

// ExitMarcador is called when production marcador is exited.
func (s *BaseGramaticaListener) ExitMarcador(ctx *MarcadorContext) {}

// EnterDec is called when production dec is entered.
func (s *BaseGramaticaListener) EnterDec(ctx *DecContext) {}

// ExitDec is called when production dec is exited.
func (s *BaseGramaticaListener) ExitDec(ctx *DecContext) {}

// EnterTipo is called when production tipo is entered.
func (s *BaseGramaticaListener) EnterTipo(ctx *TipoContext) {}

// ExitTipo is called when production tipo is exited.
func (s *BaseGramaticaListener) ExitTipo(ctx *TipoContext) {}

// EnterExp is called when production exp is entered.
func (s *BaseGramaticaListener) EnterExp(ctx *ExpContext) {}

// ExitExp is called when production exp is exited.
func (s *BaseGramaticaListener) ExitExp(ctx *ExpContext) {}
