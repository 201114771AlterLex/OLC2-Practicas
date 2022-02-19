// Code generated from Practica4.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Practica4

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 30, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 22, 10, 4, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 2, 2, 6, 2, 4, 6, 8, 2, 2, 2, 26, 2, 10, 3,
	2, 2, 2, 4, 14, 3, 2, 2, 2, 6, 21, 3, 2, 2, 2, 8, 23, 3, 2, 2, 2, 10, 11,
	5, 8, 5, 2, 11, 12, 5, 4, 3, 2, 12, 13, 7, 2, 2, 3, 13, 3, 3, 2, 2, 2,
	14, 15, 5, 8, 5, 2, 15, 16, 5, 6, 4, 2, 16, 5, 3, 2, 2, 2, 17, 18, 5, 8,
	5, 2, 18, 19, 5, 6, 4, 2, 19, 22, 3, 2, 2, 2, 20, 22, 3, 2, 2, 2, 21, 17,
	3, 2, 2, 2, 21, 20, 3, 2, 2, 2, 22, 7, 3, 2, 2, 2, 23, 24, 7, 3, 2, 2,
	24, 25, 7, 7, 2, 2, 25, 26, 7, 6, 2, 2, 26, 27, 7, 7, 2, 2, 27, 28, 7,
	4, 2, 2, 28, 9, 3, 2, 2, 2, 3, 21,
}
var literalNames = []string{
	"", "'('", "')'", "'-'", "','",
}
var symbolicNames = []string{
	"", "PRight", "PLeft", "Minus", "Colon", "Numbers", "Whitespace",
}

var ruleNames = []string{
	"start", "coordinates", "coor", "coordinate",
}

type Practica4Parser struct {
	*antlr.BaseParser
}

// NewPractica4Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Practica4Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPractica4Parser(input antlr.TokenStream) *Practica4Parser {
	this := new(Practica4Parser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Practica4.g4"

	return this
}

// Practica4Parser tokens.
const (
	Practica4ParserEOF        = antlr.TokenEOF
	Practica4ParserPRight     = 1
	Practica4ParserPLeft      = 2
	Practica4ParserMinus      = 3
	Practica4ParserColon      = 4
	Practica4ParserNumbers    = 5
	Practica4ParserWhitespace = 6
)

// Practica4Parser rules.
const (
	Practica4ParserRULE_start       = 0
	Practica4ParserRULE_coordinates = 1
	Practica4ParserRULE_coor        = 2
	Practica4ParserRULE_coordinate  = 3
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
	p.RuleIndex = Practica4ParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica4ParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) CopyFrom(ctx *StartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InicioContext struct {
	*StartContext
}

func NewInicioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InicioContext {
	var p = new(InicioContext)

	p.StartContext = NewEmptyStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StartContext))

	return p
}

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) Coordinate() ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *InicioContext) Coordinates() ICoordinatesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinatesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinatesContext)
}

func (s *InicioContext) EOF() antlr.TerminalNode {
	return s.GetToken(Practica4ParserEOF, 0)
}

func (s *InicioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.EnterInicio(s)
	}
}

func (s *InicioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.ExitInicio(s)
	}
}

func (p *Practica4Parser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Practica4ParserRULE_start)

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

	localctx = NewInicioContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.Coordinate()
	}
	{
		p.SetState(9)
		p.Coordinates()
	}
	{
		p.SetState(10)
		p.Match(Practica4ParserEOF)
	}

	return localctx
}

// ICoordinatesContext is an interface to support dynamic dispatch.
type ICoordinatesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCoordinatesContext differentiates from other interfaces.
	IsCoordinatesContext()
}

type CoordinatesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoordinatesContext() *CoordinatesContext {
	var p = new(CoordinatesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica4ParserRULE_coordinates
	return p
}

func (*CoordinatesContext) IsCoordinatesContext() {}

func NewCoordinatesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinatesContext {
	var p = new(CoordinatesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica4ParserRULE_coordinates

	return p
}

func (s *CoordinatesContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinatesContext) CopyFrom(ctx *CoordinatesContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CoordinatesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinatesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type C1Context struct {
	*CoordinatesContext
}

func NewC1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *C1Context {
	var p = new(C1Context)

	p.CoordinatesContext = NewEmptyCoordinatesContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CoordinatesContext))

	return p
}

func (s *C1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *C1Context) Coordinate() ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *C1Context) Coor() ICoorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoorContext)
}

func (s *C1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.EnterC1(s)
	}
}

func (s *C1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.ExitC1(s)
	}
}

func (p *Practica4Parser) Coordinates() (localctx ICoordinatesContext) {
	localctx = NewCoordinatesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Practica4ParserRULE_coordinates)

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

	localctx = NewC1Context(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Coordinate()
	}
	{
		p.SetState(13)
		p.Coor()
	}

	return localctx
}

// ICoorContext is an interface to support dynamic dispatch.
type ICoorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCoorContext differentiates from other interfaces.
	IsCoorContext()
}

type CoorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoorContext() *CoorContext {
	var p = new(CoorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica4ParserRULE_coor
	return p
}

func (*CoorContext) IsCoorContext() {}

func NewCoorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoorContext {
	var p = new(CoorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica4ParserRULE_coor

	return p
}

func (s *CoorContext) GetParser() antlr.Parser { return s.parser }

func (s *CoorContext) CopyFrom(ctx *CoorContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CoorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type C3Context struct {
	*CoorContext
}

func NewC3Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *C3Context {
	var p = new(C3Context)

	p.CoorContext = NewEmptyCoorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CoorContext))

	return p
}

func (s *C3Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *C3Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.EnterC3(s)
	}
}

func (s *C3Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.ExitC3(s)
	}
}

type C2Context struct {
	*CoorContext
}

func NewC2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *C2Context {
	var p = new(C2Context)

	p.CoorContext = NewEmptyCoorContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CoorContext))

	return p
}

func (s *C2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *C2Context) Coordinate() ICoordinateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoordinateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoordinateContext)
}

func (s *C2Context) Coor() ICoorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICoorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICoorContext)
}

func (s *C2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.EnterC2(s)
	}
}

func (s *C2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.ExitC2(s)
	}
}

func (p *Practica4Parser) Coor() (localctx ICoorContext) {
	localctx = NewCoorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Practica4ParserRULE_coor)

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

	p.SetState(19)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Practica4ParserPRight:
		localctx = NewC2Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(15)
			p.Coordinate()
		}
		{
			p.SetState(16)
			p.Coor()
		}

	case Practica4ParserEOF:
		localctx = NewC3Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ICoordinateContext is an interface to support dynamic dispatch.
type ICoordinateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCoordinateContext differentiates from other interfaces.
	IsCoordinateContext()
}

type CoordinateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCoordinateContext() *CoordinateContext {
	var p = new(CoordinateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica4ParserRULE_coordinate
	return p
}

func (*CoordinateContext) IsCoordinateContext() {}

func NewCoordinateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CoordinateContext {
	var p = new(CoordinateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica4ParserRULE_coordinate

	return p
}

func (s *CoordinateContext) GetParser() antlr.Parser { return s.parser }

func (s *CoordinateContext) CopyFrom(ctx *CoordinateContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CoordinateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoordinateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CordenadaContext struct {
	*CoordinateContext
}

func NewCordenadaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CordenadaContext {
	var p = new(CordenadaContext)

	p.CoordinateContext = NewEmptyCoordinateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CoordinateContext))

	return p
}

func (s *CordenadaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CordenadaContext) PRight() antlr.TerminalNode {
	return s.GetToken(Practica4ParserPRight, 0)
}

func (s *CordenadaContext) AllNumbers() []antlr.TerminalNode {
	return s.GetTokens(Practica4ParserNumbers)
}

func (s *CordenadaContext) Numbers(i int) antlr.TerminalNode {
	return s.GetToken(Practica4ParserNumbers, i)
}

func (s *CordenadaContext) Colon() antlr.TerminalNode {
	return s.GetToken(Practica4ParserColon, 0)
}

func (s *CordenadaContext) PLeft() antlr.TerminalNode {
	return s.GetToken(Practica4ParserPLeft, 0)
}

func (s *CordenadaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.EnterCordenada(s)
	}
}

func (s *CordenadaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica4Listener); ok {
		listenerT.ExitCordenada(s)
	}
}

func (p *Practica4Parser) Coordinate() (localctx ICoordinateContext) {
	localctx = NewCoordinateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Practica4ParserRULE_coordinate)

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

	localctx = NewCordenadaContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(21)
		p.Match(Practica4ParserPRight)
	}
	{
		p.SetState(22)
		p.Match(Practica4ParserNumbers)
	}
	{
		p.SetState(23)
		p.Match(Practica4ParserColon)
	}
	{
		p.SetState(24)
		p.Match(Practica4ParserNumbers)
	}
	{
		p.SetState(25)
		p.Match(Practica4ParserPLeft)
	}

	return localctx
}
