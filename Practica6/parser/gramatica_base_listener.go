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

// EnterExpresion is called when production expresion is entered.
func (s *BaseGramaticaListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BaseGramaticaListener) ExitExpresion(ctx *ExpresionContext) {}

// EnterOprel is called when production oprel is entered.
func (s *BaseGramaticaListener) EnterOprel(ctx *OprelContext) {}

// ExitOprel is called when production oprel is exited.
func (s *BaseGramaticaListener) ExitOprel(ctx *OprelContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseGramaticaListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseGramaticaListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseGramaticaListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseGramaticaListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterInst_asignacion is called when production inst_asignacion is entered.
func (s *BaseGramaticaListener) EnterInst_asignacion(ctx *Inst_asignacionContext) {}

// ExitInst_asignacion is called when production inst_asignacion is exited.
func (s *BaseGramaticaListener) ExitInst_asignacion(ctx *Inst_asignacionContext) {}

// EnterInst_if is called when production inst_if is entered.
func (s *BaseGramaticaListener) EnterInst_if(ctx *Inst_ifContext) {}

// ExitInst_if is called when production inst_if is exited.
func (s *BaseGramaticaListener) ExitInst_if(ctx *Inst_ifContext) {}

// EnterInst_switch_propuesta1 is called when production inst_switch_propuesta1 is entered.
func (s *BaseGramaticaListener) EnterInst_switch_propuesta1(ctx *Inst_switch_propuesta1Context) {}

// ExitInst_switch_propuesta1 is called when production inst_switch_propuesta1 is exited.
func (s *BaseGramaticaListener) ExitInst_switch_propuesta1(ctx *Inst_switch_propuesta1Context) {}

// EnterInst_switch_propuesta2 is called when production inst_switch_propuesta2 is entered.
func (s *BaseGramaticaListener) EnterInst_switch_propuesta2(ctx *Inst_switch_propuesta2Context) {}

// ExitInst_switch_propuesta2 is called when production inst_switch_propuesta2 is exited.
func (s *BaseGramaticaListener) ExitInst_switch_propuesta2(ctx *Inst_switch_propuesta2Context) {}

// EnterInst_while is called when production inst_while is entered.
func (s *BaseGramaticaListener) EnterInst_while(ctx *Inst_whileContext) {}

// ExitInst_while is called when production inst_while is exited.
func (s *BaseGramaticaListener) ExitInst_while(ctx *Inst_whileContext) {}

// EnterInst_doWhile is called when production inst_doWhile is entered.
func (s *BaseGramaticaListener) EnterInst_doWhile(ctx *Inst_doWhileContext) {}

// ExitInst_doWhile is called when production inst_doWhile is exited.
func (s *BaseGramaticaListener) ExitInst_doWhile(ctx *Inst_doWhileContext) {}

// EnterInst_loop is called when production inst_loop is entered.
func (s *BaseGramaticaListener) EnterInst_loop(ctx *Inst_loopContext) {}

// ExitInst_loop is called when production inst_loop is exited.
func (s *BaseGramaticaListener) ExitInst_loop(ctx *Inst_loopContext) {}

// EnterInst_for is called when production inst_for is entered.
func (s *BaseGramaticaListener) EnterInst_for(ctx *Inst_forContext) {}

// ExitInst_for is called when production inst_for is exited.
func (s *BaseGramaticaListener) ExitInst_for(ctx *Inst_forContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BaseGramaticaListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BaseGramaticaListener) ExitBloque(ctx *BloqueContext) {}

// EnterBloqueSinLLaves is called when production bloqueSinLLaves is entered.
func (s *BaseGramaticaListener) EnterBloqueSinLLaves(ctx *BloqueSinLLavesContext) {}

// ExitBloqueSinLLaves is called when production bloqueSinLLaves is exited.
func (s *BaseGramaticaListener) ExitBloqueSinLLaves(ctx *BloqueSinLLavesContext) {}
