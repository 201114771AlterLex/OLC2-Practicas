// Code generated from Practica3.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Practica3

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 56, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 5, 4, 30, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 43, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 54, 10, 7, 3, 7, 2, 2, 8, 2, 4, 6, 8,
	10, 12, 2, 2, 2, 53, 2, 14, 3, 2, 2, 2, 4, 18, 3, 2, 2, 2, 6, 29, 3, 2,
	2, 2, 8, 31, 3, 2, 2, 2, 10, 42, 3, 2, 2, 2, 12, 53, 3, 2, 2, 2, 14, 15,
	5, 4, 3, 2, 15, 16, 7, 2, 2, 3, 16, 17, 8, 2, 1, 2, 17, 3, 3, 2, 2, 2,
	18, 19, 5, 8, 5, 2, 19, 20, 5, 6, 4, 2, 20, 21, 8, 3, 1, 2, 21, 5, 3, 2,
	2, 2, 22, 23, 7, 5, 2, 2, 23, 24, 5, 8, 5, 2, 24, 25, 8, 4, 1, 2, 25, 26,
	5, 6, 4, 2, 26, 27, 8, 4, 1, 2, 27, 30, 3, 2, 2, 2, 28, 30, 8, 4, 1, 2,
	29, 22, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30, 7, 3, 2, 2, 2, 31, 32, 5, 12,
	7, 2, 32, 33, 5, 10, 6, 2, 33, 34, 8, 5, 1, 2, 34, 9, 3, 2, 2, 2, 35, 36,
	7, 6, 2, 2, 36, 37, 5, 12, 7, 2, 37, 38, 8, 6, 1, 2, 38, 39, 5, 10, 6,
	2, 39, 40, 8, 6, 1, 2, 40, 43, 3, 2, 2, 2, 41, 43, 8, 6, 1, 2, 42, 35,
	3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 11, 3, 2, 2, 2, 44, 45, 7, 4, 2, 2,
	45, 46, 5, 4, 3, 2, 46, 47, 7, 3, 2, 2, 47, 48, 8, 7, 1, 2, 48, 54, 3,
	2, 2, 2, 49, 50, 7, 8, 2, 2, 50, 54, 8, 7, 1, 2, 51, 52, 7, 9, 2, 2, 52,
	54, 8, 7, 1, 2, 53, 44, 3, 2, 2, 2, 53, 49, 3, 2, 2, 2, 53, 51, 3, 2, 2,
	2, 54, 13, 3, 2, 2, 2, 5, 29, 42, 53,
}
var literalNames = []string{
	"", "')'", "'('", "'+'", "'*'", "'.'",
}
var symbolicNames = []string{
	"", "PRight", "PLeft", "Plus", "Times", "Point", "Numbers", "Decimal",
	"Whitespace",
}

var ruleNames = []string{
	"start", "e", "ep", "t", "tp", "f",
}

type Practica3Parser struct {
	*antlr.BaseParser
}

// NewPractica3Parser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Practica3Parser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewPractica3Parser(input antlr.TokenStream) *Practica3Parser {
	this := new(Practica3Parser)
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
	this.GrammarFileName = "Practica3.g4"

	return this
}

type Data struct {
	DataType  string
	DataValue string
}

// Practica3Parser tokens.
const (
	Practica3ParserEOF        = antlr.TokenEOF
	Practica3ParserPRight     = 1
	Practica3ParserPLeft      = 2
	Practica3ParserPlus       = 3
	Practica3ParserTimes      = 4
	Practica3ParserPoint      = 5
	Practica3ParserNumbers    = 6
	Practica3ParserDecimal    = 7
	Practica3ParserWhitespace = 8
)

// Practica3Parser rules.
const (
	Practica3ParserRULE_start = 0
	Practica3ParserRULE_e     = 1
	Practica3ParserRULE_ep    = 2
	Practica3ParserRULE_t     = 3
	Practica3ParserRULE_tp    = 4
	Practica3ParserRULE_f     = 5
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_e returns the _e rule contexts.
	Get_e() IEContext

	// Set_e sets the _e rule contexts.
	Set_e(IEContext)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	_e     IEContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3ParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3ParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_e() IEContext { return s._e }

func (s *StartContext) Set_e(v IEContext) { s._e = v }

func (s *StartContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *StartContext) EOF() antlr.TerminalNode {
	return s.GetToken(Practica3ParserEOF, 0)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Practica3Parser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Practica3ParserRULE_start)

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
		p.SetState(12)

		var _x = p.E()

		localctx.(*StartContext)._e = _x
	}
	{
		p.SetState(13)
		p.Match(Practica3ParserEOF)
	}

	//fmt.Println(localctx.(*StartContext).Get_e().GetD1().DataValue)

	return localctx
}

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT1 returns the t1 rule contexts.
	GetT1() ITContext

	// GetEp1 returns the ep1 rule contexts.
	GetEp1() IEpContext

	// SetT1 sets the t1 rule contexts.
	SetT1(ITContext)

	// SetEp1 sets the ep1 rule contexts.
	SetEp1(IEpContext)

	// GetD1 returns the D1 attribute.
	GetD1() *Data

	// SetD1 sets the D1 attribute.
	SetD1(*Data)

	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	D1     *Data
	t1     ITContext
	ep1    IEpContext
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3ParserRULE_e
	return p
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3ParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) GetT1() ITContext { return s.t1 }

func (s *EContext) GetEp1() IEpContext { return s.ep1 }

func (s *EContext) SetT1(v ITContext) { s.t1 = v }

func (s *EContext) SetEp1(v IEpContext) { s.ep1 = v }

func (s *EContext) GetD1() *Data { return s.D1 }

func (s *EContext) SetD1(v *Data) { s.D1 = v }

func (s *EContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *EContext) Ep() IEpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEpContext)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.EnterE(s)
	}
}

func (s *EContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.ExitE(s)
	}
}

func (p *Practica3Parser) E() (localctx IEContext) {
	localctx = NewEContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Practica3ParserRULE_e)

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
		p.SetState(16)

		var _x = p.T()

		localctx.(*EContext).t1 = _x
	}
	{
		p.SetState(17)

		var _x = p.Ep(localctx.(*EContext).GetT1().GetD1())

		localctx.(*EContext).ep1 = _x
	}

	localctx.(*EContext).D1 = localctx.(*EContext).GetEp1().GetD1()

	return localctx
}

// IEpContext is an interface to support dynamic dispatch.
type IEpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT1 returns the t1 rule contexts.
	GetT1() ITContext

	// GetEp1 returns the ep1 rule contexts.
	GetEp1() IEpContext

	// SetT1 sets the t1 rule contexts.
	SetT1(ITContext)

	// SetEp1 sets the ep1 rule contexts.
	SetEp1(IEpContext)

	// GetD returns the D attribute.
	GetD() *Data

	// GetD1 returns the D1 attribute.
	GetD1() *Data

	// GetDTemp returns the DTemp attribute.
	GetDTemp() *Data

	// SetD sets the D attribute.
	SetD(*Data)

	// SetD1 sets the D1 attribute.
	SetD1(*Data)

	// SetDTemp sets the DTemp attribute.
	SetDTemp(*Data)

	// IsEpContext differentiates from other interfaces.
	IsEpContext()
}

type EpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	D      *Data
	D1     *Data
	DTemp  *Data
	t1     ITContext
	ep1    IEpContext
}

func NewEmptyEpContext() *EpContext {
	var p = new(EpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3ParserRULE_ep
	return p
}

func (*EpContext) IsEpContext() {}

func NewEpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, D *Data) *EpContext {
	var p = new(EpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3ParserRULE_ep

	p.D = D

	return p
}

func (s *EpContext) GetParser() antlr.Parser { return s.parser }

func (s *EpContext) GetT1() ITContext { return s.t1 }

func (s *EpContext) GetEp1() IEpContext { return s.ep1 }

func (s *EpContext) SetT1(v ITContext) { s.t1 = v }

func (s *EpContext) SetEp1(v IEpContext) { s.ep1 = v }

func (s *EpContext) GetD() *Data { return s.D }

func (s *EpContext) GetD1() *Data { return s.D1 }

func (s *EpContext) GetDTemp() *Data { return s.DTemp }

func (s *EpContext) SetD(v *Data) { s.D = v }

func (s *EpContext) SetD1(v *Data) { s.D1 = v }

func (s *EpContext) SetDTemp(v *Data) { s.DTemp = v }

func (s *EpContext) Plus() antlr.TerminalNode {
	return s.GetToken(Practica3ParserPlus, 0)
}

func (s *EpContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *EpContext) Ep() IEpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEpContext)
}

func (s *EpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.EnterEp(s)
	}
}

func (s *EpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.ExitEp(s)
	}
}

func (p *Practica3Parser) Ep(D *Data) (localctx IEpContext) {
	localctx = NewEpContext(p, p.GetParserRuleContext(), p.GetState(), D)
	p.EnterRule(localctx, 4, Practica3ParserRULE_ep)

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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Practica3ParserPlus:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Match(Practica3ParserPlus)
		}
		{
			p.SetState(21)

			var _x = p.T()

			localctx.(*EpContext).t1 = _x
		}

		localctx.(*EpContext).DTemp = &Data{}
		if localctx.(*EpContext).D.DataType == "D" || localctx.(*EpContext).GetT1().GetD1().DataType == "D" {
			localctx.(*EpContext).DTemp.DataType = "D"
			NT1, _ := strconv.ParseFloat(localctx.(*EpContext).D.DataValue, 64)
			NT2, _ := strconv.ParseFloat(localctx.(*EpContext).GetT1().GetD1().DataValue, 64)
			Result := NT1 + NT2
			localctx.(*EpContext).DTemp.DataValue = fmt.Sprintf("%g", Result)
		} else {
			localctx.(*EpContext).DTemp.DataType = "E"
			NT1, _ := strconv.Atoi(localctx.(*EpContext).D.DataValue)
			NT2, _ := strconv.Atoi(localctx.(*EpContext).GetT1().GetD1().DataValue)
			Result := NT1 + NT2
			localctx.(*EpContext).DTemp.DataValue = strconv.Itoa(Result)
		}

		{
			p.SetState(23)

			var _x = p.Ep(localctx.(*EpContext).DTemp)

			localctx.(*EpContext).ep1 = _x
		}

		localctx.(*EpContext).D1 = localctx.(*EpContext).GetEp1().GetD1()

	case Practica3ParserEOF, Practica3ParserPRight:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*EpContext).D1 = localctx.(*EpContext).D

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITContext is an interface to support dynamic dispatch.
type ITContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF1 returns the f1 rule contexts.
	GetF1() IFContext

	// GetTp1 returns the tp1 rule contexts.
	GetTp1() ITpContext

	// SetF1 sets the f1 rule contexts.
	SetF1(IFContext)

	// SetTp1 sets the tp1 rule contexts.
	SetTp1(ITpContext)

	// GetD1 returns the D1 attribute.
	GetD1() *Data

	// SetD1 sets the D1 attribute.
	SetD1(*Data)

	// IsTContext differentiates from other interfaces.
	IsTContext()
}

type TContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	D1     *Data
	f1     IFContext
	tp1    ITpContext
}

func NewEmptyTContext() *TContext {
	var p = new(TContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3ParserRULE_t
	return p
}

func (*TContext) IsTContext() {}

func NewTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TContext {
	var p = new(TContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3ParserRULE_t

	return p
}

func (s *TContext) GetParser() antlr.Parser { return s.parser }

func (s *TContext) GetF1() IFContext { return s.f1 }

func (s *TContext) GetTp1() ITpContext { return s.tp1 }

func (s *TContext) SetF1(v IFContext) { s.f1 = v }

func (s *TContext) SetTp1(v ITpContext) { s.tp1 = v }

func (s *TContext) GetD1() *Data { return s.D1 }

func (s *TContext) SetD1(v *Data) { s.D1 = v }

func (s *TContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *TContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
}

func (s *TContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.EnterT(s)
	}
}

func (s *TContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.ExitT(s)
	}
}

func (p *Practica3Parser) T() (localctx ITContext) {
	localctx = NewTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Practica3ParserRULE_t)

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
		p.SetState(29)

		var _x = p.F()

		localctx.(*TContext).f1 = _x
	}
	{
		p.SetState(30)

		var _x = p.Tp(localctx.(*TContext).GetF1().GetD1())

		localctx.(*TContext).tp1 = _x
	}

	localctx.(*TContext).D1 = localctx.(*TContext).GetTp1().GetD1()

	return localctx
}

// ITpContext is an interface to support dynamic dispatch.
type ITpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetF1 returns the f1 rule contexts.
	GetF1() IFContext

	// GetTp1 returns the tp1 rule contexts.
	GetTp1() ITpContext

	// SetF1 sets the f1 rule contexts.
	SetF1(IFContext)

	// SetTp1 sets the tp1 rule contexts.
	SetTp1(ITpContext)

	// GetD returns the D attribute.
	GetD() *Data

	// GetD1 returns the D1 attribute.
	GetD1() *Data

	// GetDTemp returns the DTemp attribute.
	GetDTemp() *Data

	// SetD sets the D attribute.
	SetD(*Data)

	// SetD1 sets the D1 attribute.
	SetD1(*Data)

	// SetDTemp sets the DTemp attribute.
	SetDTemp(*Data)

	// IsTpContext differentiates from other interfaces.
	IsTpContext()
}

type TpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	D      *Data
	D1     *Data
	DTemp  *Data
	f1     IFContext
	tp1    ITpContext
}

func NewEmptyTpContext() *TpContext {
	var p = new(TpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3ParserRULE_tp
	return p
}

func (*TpContext) IsTpContext() {}

func NewTpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, D *Data) *TpContext {
	var p = new(TpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3ParserRULE_tp

	p.D = D

	return p
}

func (s *TpContext) GetParser() antlr.Parser { return s.parser }

func (s *TpContext) GetF1() IFContext { return s.f1 }

func (s *TpContext) GetTp1() ITpContext { return s.tp1 }

func (s *TpContext) SetF1(v IFContext) { s.f1 = v }

func (s *TpContext) SetTp1(v ITpContext) { s.tp1 = v }

func (s *TpContext) GetD() *Data { return s.D }

func (s *TpContext) GetD1() *Data { return s.D1 }

func (s *TpContext) GetDTemp() *Data { return s.DTemp }

func (s *TpContext) SetD(v *Data) { s.D = v }

func (s *TpContext) SetD1(v *Data) { s.D1 = v }

func (s *TpContext) SetDTemp(v *Data) { s.DTemp = v }

func (s *TpContext) Times() antlr.TerminalNode {
	return s.GetToken(Practica3ParserTimes, 0)
}

func (s *TpContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *TpContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
}

func (s *TpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.EnterTp(s)
	}
}

func (s *TpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.ExitTp(s)
	}
}

func (p *Practica3Parser) Tp(D *Data) (localctx ITpContext) {
	localctx = NewTpContext(p, p.GetParserRuleContext(), p.GetState(), D)
	p.EnterRule(localctx, 8, Practica3ParserRULE_tp)

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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Practica3ParserTimes:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(Practica3ParserTimes)
		}
		{
			p.SetState(34)

			var _x = p.F()

			localctx.(*TpContext).f1 = _x
		}

		localctx.(*TpContext).DTemp = &Data{}
		if localctx.(*TpContext).D.DataType == "D" || localctx.(*TpContext).GetF1().GetD1().DataType == "D" {
			localctx.(*TpContext).DTemp.DataType = "D"
			NT1, _ := strconv.ParseFloat(localctx.(*TpContext).D.DataValue, 64)
			NT2, _ := strconv.ParseFloat(localctx.(*TpContext).GetF1().GetD1().DataValue, 64)
			Result := NT1 * NT2
			localctx.(*TpContext).DTemp.DataValue = fmt.Sprintf("%g", Result)
		} else {
			localctx.(*TpContext).DTemp.DataType = "E"
			NT1, _ := strconv.Atoi(localctx.(*TpContext).D.DataValue)
			NT2, _ := strconv.Atoi(localctx.(*TpContext).GetF1().GetD1().DataValue)
			Result := NT1 * NT2
			localctx.(*TpContext).DTemp.DataValue = strconv.Itoa(Result)
		}

		{
			p.SetState(36)

			var _x = p.Tp(localctx.(*TpContext).DTemp)

			localctx.(*TpContext).tp1 = _x
		}

		localctx.(*TpContext).D1 = localctx.(*TpContext).GetTp1().GetD1()

	case Practica3ParserEOF, Practica3ParserPRight, Practica3ParserPlus:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*TpContext).D1 = localctx.(*TpContext).D

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFContext is an interface to support dynamic dispatch.
type IFContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetNe returns the ne token.
	GetNe() antlr.Token

	// GetNd returns the nd token.
	GetNd() antlr.Token

	// SetNe sets the ne token.
	SetNe(antlr.Token)

	// SetNd sets the nd token.
	SetNd(antlr.Token)

	// Get_e returns the _e rule contexts.
	Get_e() IEContext

	// Set_e sets the _e rule contexts.
	Set_e(IEContext)

	// GetD1 returns the D1 attribute.
	GetD1() *Data

	// SetD1 sets the D1 attribute.
	SetD1(*Data)

	// IsFContext differentiates from other interfaces.
	IsFContext()
}

type FContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	D1     *Data
	_e     IEContext
	ne     antlr.Token
	nd     antlr.Token
}

func NewEmptyFContext() *FContext {
	var p = new(FContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3ParserRULE_f
	return p
}

func (*FContext) IsFContext() {}

func NewFContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FContext {
	var p = new(FContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3ParserRULE_f

	return p
}

func (s *FContext) GetParser() antlr.Parser { return s.parser }

func (s *FContext) GetNe() antlr.Token { return s.ne }

func (s *FContext) GetNd() antlr.Token { return s.nd }

func (s *FContext) SetNe(v antlr.Token) { s.ne = v }

func (s *FContext) SetNd(v antlr.Token) { s.nd = v }

func (s *FContext) Get_e() IEContext { return s._e }

func (s *FContext) Set_e(v IEContext) { s._e = v }

func (s *FContext) GetD1() *Data { return s.D1 }

func (s *FContext) SetD1(v *Data) { s.D1 = v }

func (s *FContext) PLeft() antlr.TerminalNode {
	return s.GetToken(Practica3ParserPLeft, 0)
}

func (s *FContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *FContext) PRight() antlr.TerminalNode {
	return s.GetToken(Practica3ParserPRight, 0)
}

func (s *FContext) Numbers() antlr.TerminalNode {
	return s.GetToken(Practica3ParserNumbers, 0)
}

func (s *FContext) Decimal() antlr.TerminalNode {
	return s.GetToken(Practica3ParserDecimal, 0)
}

func (s *FContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.EnterF(s)
	}
}

func (s *FContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(Practica3Listener); ok {
		listenerT.ExitF(s)
	}
}

func (p *Practica3Parser) F() (localctx IFContext) {
	localctx = NewFContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Practica3ParserRULE_f)

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

	p.SetState(51)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Practica3ParserPLeft:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Match(Practica3ParserPLeft)
		}
		{
			p.SetState(43)

			var _x = p.E()

			localctx.(*FContext)._e = _x
		}
		{
			p.SetState(44)
			p.Match(Practica3ParserPRight)
		}

		localctx.(*FContext).D1 = localctx.(*FContext).Get_e().GetD1()

	case Practica3ParserNumbers:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)

			var _m = p.Match(Practica3ParserNumbers)

			localctx.(*FContext).ne = _m
		}

		localctx.(*FContext).D1 = &Data{}
		localctx.(*FContext).D1.DataType = "E"
		localctx.(*FContext).D1.DataValue = (func() string {
			if localctx.(*FContext).GetNe() == nil {
				return ""
			} else {
				return localctx.(*FContext).GetNe().GetText()
			}
		}())

	case Practica3ParserDecimal:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)

			var _m = p.Match(Practica3ParserDecimal)

			localctx.(*FContext).nd = _m
		}

		localctx.(*FContext).D1 = &Data{}
		localctx.(*FContext).D1.DataType = "D"
		localctx.(*FContext).D1.DataValue = (func() string {
			if localctx.(*FContext).GetNd() == nil {
				return ""
			} else {
				return localctx.(*FContext).GetNd().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
