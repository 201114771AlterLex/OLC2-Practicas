// Code generated from /home/alterlex/go_projects/src/github.com/OLC2-Practicas/Evaluacion1/Evaluacion1.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Evaluacion1

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 35, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 25, 10,
	4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 5, 6, 33, 10, 6, 3, 6, 2, 2, 7,
	2, 4, 6, 8, 10, 2, 2, 2, 31, 2, 12, 3, 2, 2, 2, 4, 15, 3, 2, 2, 2, 6, 24,
	3, 2, 2, 2, 8, 26, 3, 2, 2, 2, 10, 32, 3, 2, 2, 2, 12, 13, 5, 4, 3, 2,
	13, 14, 7, 2, 2, 3, 14, 3, 3, 2, 2, 2, 15, 16, 5, 6, 4, 2, 16, 17, 7, 4,
	2, 2, 17, 18, 5, 6, 4, 2, 18, 5, 3, 2, 2, 2, 19, 20, 5, 8, 5, 2, 20, 21,
	7, 3, 2, 2, 21, 22, 5, 8, 5, 2, 22, 25, 3, 2, 2, 2, 23, 25, 5, 8, 5, 2,
	24, 19, 3, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 7, 3, 2, 2, 2, 26, 27, 7, 5,
	2, 2, 27, 28, 5, 10, 6, 2, 28, 9, 3, 2, 2, 2, 29, 30, 7, 5, 2, 2, 30, 33,
	5, 10, 6, 2, 31, 33, 3, 2, 2, 2, 32, 29, 3, 2, 2, 2, 32, 31, 3, 2, 2, 2,
	33, 11, 3, 2, 2, 2, 4, 24, 32,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'.'", "','",
}
var symbolicNames = []string{
	"", "Point", "Colon", "Hex", "Number", "Letter", "Whitespace",
}

var ruleNames = []string{
	"start", "list", "h", "t", "tp",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Evaluacion1Parser struct {
	*antlr.BaseParser
}

func NewEvaluacion1Parser(input antlr.TokenStream) *Evaluacion1Parser {
	this := new(Evaluacion1Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Evaluacion1.g4"

	return this
}

type DataHex struct {
	Valor    int
	Posicion int
}

type Decimal struct {
	Decimal1 float64
	Decimal2 float64
}

// Evaluacion1Parser tokens.
const (
	Evaluacion1ParserEOF        = antlr.TokenEOF
	Evaluacion1ParserPoint      = 1
	Evaluacion1ParserColon      = 2
	Evaluacion1ParserHex        = 3
	Evaluacion1ParserNumber     = 4
	Evaluacion1ParserLetter     = 5
	Evaluacion1ParserWhitespace = 6
)

// Evaluacion1Parser rules.
const (
	Evaluacion1ParserRULE_start = 0
	Evaluacion1ParserRULE_list  = 1
	Evaluacion1ParserRULE_h     = 2
	Evaluacion1ParserRULE_t     = 3
	Evaluacion1ParserRULE_tp    = 4
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
	p.RuleIndex = Evaluacion1ParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Evaluacion1ParserRULE_start

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

func (s *InicioContext) List() IListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *InicioContext) EOF() antlr.TerminalNode {
	return s.GetToken(Evaluacion1ParserEOF, 0)
}

func (p *Evaluacion1Parser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Evaluacion1ParserRULE_start)

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
		p.SetState(10)
		p.List()
	}
	{
		p.SetState(11)
		p.Match(Evaluacion1ParserEOF)
	}

	return localctx
}

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Evaluacion1ParserRULE_list
	return p
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Evaluacion1ParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) CopyFrom(ctx *ListContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DosNumeroContext struct {
	*ListContext
}

func NewDosNumeroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DosNumeroContext {
	var p = new(DosNumeroContext)

	p.ListContext = NewEmptyListContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ListContext))

	return p
}

func (s *DosNumeroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DosNumeroContext) AllH() []IHContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHContext)(nil)).Elem())
	var tst = make([]IHContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHContext)
		}
	}

	return tst
}

func (s *DosNumeroContext) H(i int) IHContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHContext)
}

func (s *DosNumeroContext) Colon() antlr.TerminalNode {
	return s.GetToken(Evaluacion1ParserColon, 0)
}

func (p *Evaluacion1Parser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Evaluacion1ParserRULE_list)

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

	localctx = NewDosNumeroContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		p.H()
	}
	{
		p.SetState(14)
		p.Match(Evaluacion1ParserColon)
	}
	{
		p.SetState(15)
		p.H()
	}

	return localctx
}

// IHContext is an interface to support dynamic dispatch.
type IHContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHDecimal returns the HDecimal attribute.
	GetHDecimal() float64

	// SetHDecimal sets the HDecimal attribute.
	SetHDecimal(float64)

	// IsHContext differentiates from other interfaces.
	IsHContext()
}

type HContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	HDecimal float64
}

func NewEmptyHContext() *HContext {
	var p = new(HContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Evaluacion1ParserRULE_h
	return p
}

func (*HContext) IsHContext() {}

func NewHContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HContext {
	var p = new(HContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Evaluacion1ParserRULE_h

	return p
}

func (s *HContext) GetParser() antlr.Parser { return s.parser }

func (s *HContext) GetHDecimal() float64 { return s.HDecimal }

func (s *HContext) SetHDecimal(v float64) { s.HDecimal = v }

func (s *HContext) CopyFrom(ctx *HContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.HDecimal = ctx.HDecimal
}

func (s *HContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type HTPTContext struct {
	*HContext
}

func NewHTPTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HTPTContext {
	var p = new(HTPTContext)

	p.HContext = NewEmptyHContext()
	p.parser = parser
	p.CopyFrom(ctx.(*HContext))

	return p
}

func (s *HTPTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HTPTContext) AllT() []ITContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITContext)(nil)).Elem())
	var tst = make([]ITContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITContext)
		}
	}

	return tst
}

func (s *HTPTContext) T(i int) ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *HTPTContext) Point() antlr.TerminalNode {
	return s.GetToken(Evaluacion1ParserPoint, 0)
}

type HTContext struct {
	*HContext
}

func NewHTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HTContext {
	var p = new(HTContext)

	p.HContext = NewEmptyHContext()
	p.parser = parser
	p.CopyFrom(ctx.(*HContext))

	return p
}

func (s *HTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HTContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (p *Evaluacion1Parser) H() (localctx IHContext) {
	localctx = NewHContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Evaluacion1ParserRULE_h)

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

	p.SetState(22)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewHTPTContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(17)
			p.T()
		}
		{
			p.SetState(18)
			p.Match(Evaluacion1ParserPoint)
		}
		{
			p.SetState(19)
			p.T()
		}

	case 2:
		localctx = NewHTContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
			p.T()
		}

	}

	return localctx
}

// ITContext is an interface to support dynamic dispatch.
type ITContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDhex returns the Dhex attribute.
	GetDhex() *DataHex

	// SetDhex sets the Dhex attribute.
	SetDhex(*DataHex)

	// IsTContext differentiates from other interfaces.
	IsTContext()
}

type TContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Dhex   *DataHex
}

func NewEmptyTContext() *TContext {
	var p = new(TContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Evaluacion1ParserRULE_t
	return p
}

func (*TContext) IsTContext() {}

func NewTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TContext {
	var p = new(TContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Evaluacion1ParserRULE_t

	return p
}

func (s *TContext) GetParser() antlr.Parser { return s.parser }

func (s *TContext) GetDhex() *DataHex { return s.Dhex }

func (s *TContext) SetDhex(v *DataHex) { s.Dhex = v }

func (s *TContext) CopyFrom(ctx *TContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.Dhex = ctx.Dhex
}

func (s *TContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type THexContext struct {
	*TContext
}

func NewTHexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *THexContext {
	var p = new(THexContext)

	p.TContext = NewEmptyTContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TContext))

	return p
}

func (s *THexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *THexContext) Hex() antlr.TerminalNode {
	return s.GetToken(Evaluacion1ParserHex, 0)
}

func (s *THexContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
}

func (p *Evaluacion1Parser) T() (localctx ITContext) {
	localctx = NewTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Evaluacion1ParserRULE_t)

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

	localctx = NewTHexContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(24)
		p.Match(Evaluacion1ParserHex)
	}
	{
		p.SetState(25)
		p.Tp()
	}

	return localctx
}

// ITpContext is an interface to support dynamic dispatch.
type ITpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDhex returns the Dhex attribute.
	GetDhex() *DataHex

	// SetDhex sets the Dhex attribute.
	SetDhex(*DataHex)

	// IsTpContext differentiates from other interfaces.
	IsTpContext()
}

type TpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	Dhex   *DataHex
}

func NewEmptyTpContext() *TpContext {
	var p = new(TpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Evaluacion1ParserRULE_tp
	return p
}

func (*TpContext) IsTpContext() {}

func NewTpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TpContext {
	var p = new(TpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Evaluacion1ParserRULE_tp

	return p
}

func (s *TpContext) GetParser() antlr.Parser { return s.parser }

func (s *TpContext) GetDhex() *DataHex { return s.Dhex }

func (s *TpContext) SetDhex(v *DataHex) { s.Dhex = v }

func (s *TpContext) CopyFrom(ctx *TpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.Dhex = ctx.Dhex
}

func (s *TpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EpTpContext struct {
	*TpContext
}

func NewEpTpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpTpContext {
	var p = new(EpTpContext)

	p.TpContext = NewEmptyTpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TpContext))

	return p
}

func (s *EpTpContext) GetRuleContext() antlr.RuleContext {
	return s
}

type TpHexContext struct {
	*TpContext
}

func NewTpHexContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TpHexContext {
	var p = new(TpHexContext)

	p.TpContext = NewEmptyTpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TpContext))

	return p
}

func (s *TpHexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TpHexContext) Hex() antlr.TerminalNode {
	return s.GetToken(Evaluacion1ParserHex, 0)
}

func (s *TpHexContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
}

func (p *Evaluacion1Parser) Tp() (localctx ITpContext) {
	localctx = NewTpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Evaluacion1ParserRULE_tp)

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

	p.SetState(30)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Evaluacion1ParserHex:
		localctx = NewTpHexContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Match(Evaluacion1ParserHex)
		}
		{
			p.SetState(28)
			p.Tp()
		}

	case Evaluacion1ParserEOF, Evaluacion1ParserPoint, Evaluacion1ParserColon:
		localctx = NewEpTpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
