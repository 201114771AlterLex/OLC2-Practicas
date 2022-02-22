// Code generated from /home/alterlex/go_projects/src/github.com/OLC2-Practicas/Practica3/Practica3copy.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Practica3copy

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
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

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
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Practica3copyParser struct {
	*antlr.BaseParser
}

func NewPractica3copyParser(input antlr.TokenStream) *Practica3copyParser {
	this := new(Practica3copyParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Practica3copy.g4"

	return this
}

type Data struct {
	DataType  string
	DataValue string
}

// Practica3copyParser tokens.
const (
	Practica3copyParserEOF        = antlr.TokenEOF
	Practica3copyParserPRight     = 1
	Practica3copyParserPLeft      = 2
	Practica3copyParserPlus       = 3
	Practica3copyParserTimes      = 4
	Practica3copyParserPoint      = 5
	Practica3copyParserNumbers    = 6
	Practica3copyParserDecimal    = 7
	Practica3copyParserWhitespace = 8
)

// Practica3copyParser rules.
const (
	Practica3copyParserRULE_start = 0
	Practica3copyParserRULE_e     = 1
	Practica3copyParserRULE_ep    = 2
	Practica3copyParserRULE_t     = 3
	Practica3copyParserRULE_tp    = 4
	Practica3copyParserRULE_f     = 5
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
	p.RuleIndex = Practica3copyParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3copyParserRULE_start

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
	_e IEContext
}

func NewInicioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InicioContext {
	var p = new(InicioContext)

	p.StartContext = NewEmptyStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StartContext))

	return p
}

func (s *InicioContext) Get_e() IEContext { return s._e }

func (s *InicioContext) Set_e(v IEContext) { s._e = v }

func (s *InicioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InicioContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *InicioContext) EOF() antlr.TerminalNode {
	return s.GetToken(Practica3copyParserEOF, 0)
}

func (p *Practica3copyParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Practica3copyParserRULE_start)

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
		p.SetState(12)

		var _x = p.E()

		localctx.(*InicioContext)._e = _x
	}
	{
		p.SetState(13)
		p.Match(Practica3copyParserEOF)
	}

	fmt.Println(localctx.(*InicioContext).Get_e().GetD1().DataValue)

	return localctx
}

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3copyParserRULE_e
	return p
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3copyParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) GetD1() *Data { return s.D1 }

func (s *EContext) SetD1(v *Data) { s.D1 = v }

func (s *EContext) CopyFrom(ctx *EContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.D1 = ctx.D1
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ETContext struct {
	*EContext
	t1  ITContext
	ep1 IEpContext
}

func NewETContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ETContext {
	var p = new(ETContext)

	p.EContext = NewEmptyEContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EContext))

	return p
}

func (s *ETContext) GetT1() ITContext { return s.t1 }

func (s *ETContext) GetEp1() IEpContext { return s.ep1 }

func (s *ETContext) SetT1(v ITContext) { s.t1 = v }

func (s *ETContext) SetEp1(v IEpContext) { s.ep1 = v }

func (s *ETContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ETContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *ETContext) Ep() IEpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEpContext)
}

func (p *Practica3copyParser) E() (localctx IEContext) {
	localctx = NewEContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Practica3copyParserRULE_e)

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

	localctx = NewETContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)

		var _x = p.T()

		localctx.(*ETContext).t1 = _x
	}
	{
		p.SetState(17)

		var _x = p.Ep(localctx.(*ETContext).GetT1().GetD1())

		localctx.(*ETContext).ep1 = _x
	}

	localctx.(*EContext).D1 = localctx.(*ETContext).GetEp1().GetD1()

	return localctx
}

// IEpContext is an interface to support dynamic dispatch.
type IEpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
}

func NewEmptyEpContext() *EpContext {
	var p = new(EpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3copyParserRULE_ep
	return p
}

func (*EpContext) IsEpContext() {}

func NewEpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, D *Data) *EpContext {
	var p = new(EpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3copyParserRULE_ep

	p.D = D

	return p
}

func (s *EpContext) GetParser() antlr.Parser { return s.parser }

func (s *EpContext) GetD() *Data { return s.D }

func (s *EpContext) GetD1() *Data { return s.D1 }

func (s *EpContext) GetDTemp() *Data { return s.DTemp }

func (s *EpContext) SetD(v *Data) { s.D = v }

func (s *EpContext) SetD1(v *Data) { s.D1 = v }

func (s *EpContext) SetDTemp(v *Data) { s.DTemp = v }

func (s *EpContext) CopyFrom(ctx *EpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.D = ctx.D
	s.D1 = ctx.D1
	s.DTemp = ctx.DTemp
}

func (s *EpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EpEpContext struct {
	*EpContext
}

func NewEpEpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpEpContext {
	var p = new(EpEpContext)

	p.EpContext = NewEmptyEpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EpContext))

	return p
}

func (s *EpEpContext) GetRuleContext() antlr.RuleContext {
	return s
}

type EpPlusContext struct {
	*EpContext
	t1  ITContext
	ep1 IEpContext
}

func NewEpPlusContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpPlusContext {
	var p = new(EpPlusContext)

	p.EpContext = NewEmptyEpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EpContext))

	return p
}

func (s *EpPlusContext) GetT1() ITContext { return s.t1 }

func (s *EpPlusContext) GetEp1() IEpContext { return s.ep1 }

func (s *EpPlusContext) SetT1(v ITContext) { s.t1 = v }

func (s *EpPlusContext) SetEp1(v IEpContext) { s.ep1 = v }

func (s *EpPlusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpPlusContext) Plus() antlr.TerminalNode {
	return s.GetToken(Practica3copyParserPlus, 0)
}

func (s *EpPlusContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *EpPlusContext) Ep() IEpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEpContext)
}

func (p *Practica3copyParser) Ep(D *Data) (localctx IEpContext) {
	localctx = NewEpContext(p, p.GetParserRuleContext(), p.GetState(), D)
	p.EnterRule(localctx, 4, Practica3copyParserRULE_ep)

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
	case Practica3copyParserPlus:
		localctx = NewEpPlusContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Match(Practica3copyParserPlus)
		}
		{
			p.SetState(21)

			var _x = p.T()

			localctx.(*EpPlusContext).t1 = _x
		}

		localctx.(*EpPlusContext).DTemp = &Data{}
		if localctx.(*EpPlusContext).D.DataType == "D" || localctx.(*EpPlusContext).GetT1().GetD1().DataType == "D" {
			localctx.(*EpPlusContext).DTemp.DataType = "D"
			NT1, _ := strconv.ParseFloat(localctx.(*EpPlusContext).D.DataValue, 64)
			NT2, _ := strconv.ParseFloat(localctx.(*EpPlusContext).GetT1().GetD1().DataValue, 64)
			Result := NT1 + NT2
			localctx.(*EpPlusContext).DTemp.DataValue = fmt.Sprintf("%g", Result)
		} else {
			localctx.(*EpPlusContext).DTemp.DataType = "E"
			NT1, _ := strconv.Atoi(localctx.(*EpPlusContext).D.DataValue)
			NT2, _ := strconv.Atoi(localctx.(*EpPlusContext).GetT1().GetD1().DataValue)
			Result := NT1 + NT2
			localctx.(*EpPlusContext).DTemp.DataValue = strconv.Itoa(Result)
		}

		{
			p.SetState(23)

			var _x = p.Ep(localctx.(*EpPlusContext).DTemp)

			localctx.(*EpPlusContext).ep1 = _x
		}

		localctx.(*EpContext).D1 = localctx.(*EpPlusContext).GetEp1().GetD1()

	case Practica3copyParserEOF, Practica3copyParserPRight:
		localctx = NewEpEpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		localctx.(*EpContext).D1 = localctx.(*EpEpContext).D

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
}

func NewEmptyTContext() *TContext {
	var p = new(TContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3copyParserRULE_t
	return p
}

func (*TContext) IsTContext() {}

func NewTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TContext {
	var p = new(TContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3copyParserRULE_t

	return p
}

func (s *TContext) GetParser() antlr.Parser { return s.parser }

func (s *TContext) GetD1() *Data { return s.D1 }

func (s *TContext) SetD1(v *Data) { s.D1 = v }

func (s *TContext) CopyFrom(ctx *TContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.D1 = ctx.D1
}

func (s *TContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TFTPContext struct {
	*TContext
	f1  IFContext
	tp1 ITpContext
}

func NewTFTPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TFTPContext {
	var p = new(TFTPContext)

	p.TContext = NewEmptyTContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TContext))

	return p
}

func (s *TFTPContext) GetF1() IFContext { return s.f1 }

func (s *TFTPContext) GetTp1() ITpContext { return s.tp1 }

func (s *TFTPContext) SetF1(v IFContext) { s.f1 = v }

func (s *TFTPContext) SetTp1(v ITpContext) { s.tp1 = v }

func (s *TFTPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TFTPContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *TFTPContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
}

func (p *Practica3copyParser) T() (localctx ITContext) {
	localctx = NewTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Practica3copyParserRULE_t)

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

	localctx = NewTFTPContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(29)

		var _x = p.F()

		localctx.(*TFTPContext).f1 = _x
	}
	{
		p.SetState(30)

		var _x = p.Tp(localctx.(*TFTPContext).GetF1().GetD1())

		localctx.(*TFTPContext).tp1 = _x
	}

	localctx.(*TContext).D1 = localctx.(*TFTPContext).GetTp1().GetD1()

	return localctx
}

// ITpContext is an interface to support dynamic dispatch.
type ITpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

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
}

func NewEmptyTpContext() *TpContext {
	var p = new(TpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3copyParserRULE_tp
	return p
}

func (*TpContext) IsTpContext() {}

func NewTpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int, D *Data) *TpContext {
	var p = new(TpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3copyParserRULE_tp

	p.D = D

	return p
}

func (s *TpContext) GetParser() antlr.Parser { return s.parser }

func (s *TpContext) GetD() *Data { return s.D }

func (s *TpContext) GetD1() *Data { return s.D1 }

func (s *TpContext) GetDTemp() *Data { return s.DTemp }

func (s *TpContext) SetD(v *Data) { s.D = v }

func (s *TpContext) SetD1(v *Data) { s.D1 = v }

func (s *TpContext) SetDTemp(v *Data) { s.DTemp = v }

func (s *TpContext) CopyFrom(ctx *TpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.D = ctx.D
	s.D1 = ctx.D1
	s.DTemp = ctx.DTemp
}

func (s *TpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TpTimeContext struct {
	*TpContext
	f1  IFContext
	tp1 ITpContext
}

func NewTpTimeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TpTimeContext {
	var p = new(TpTimeContext)

	p.TpContext = NewEmptyTpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TpContext))

	return p
}

func (s *TpTimeContext) GetF1() IFContext { return s.f1 }

func (s *TpTimeContext) GetTp1() ITpContext { return s.tp1 }

func (s *TpTimeContext) SetF1(v IFContext) { s.f1 = v }

func (s *TpTimeContext) SetTp1(v ITpContext) { s.tp1 = v }

func (s *TpTimeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TpTimeContext) Times() antlr.TerminalNode {
	return s.GetToken(Practica3copyParserTimes, 0)
}

func (s *TpTimeContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *TpTimeContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
}

type TpEpContext struct {
	*TpContext
}

func NewTpEpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TpEpContext {
	var p = new(TpEpContext)

	p.TpContext = NewEmptyTpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TpContext))

	return p
}

func (s *TpEpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (p *Practica3copyParser) Tp(D *Data) (localctx ITpContext) {
	localctx = NewTpContext(p, p.GetParserRuleContext(), p.GetState(), D)
	p.EnterRule(localctx, 8, Practica3copyParserRULE_tp)

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
	case Practica3copyParserTimes:
		localctx = NewTpTimeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(Practica3copyParserTimes)
		}
		{
			p.SetState(34)

			var _x = p.F()

			localctx.(*TpTimeContext).f1 = _x
		}

		localctx.(*TpTimeContext).DTemp = &Data{}
		if localctx.(*TpTimeContext).D.DataType == "D" || localctx.(*TpTimeContext).GetF1().GetD1().DataType == "D" {
			localctx.(*TpTimeContext).DTemp.DataType = "D"
			NT1, _ := strconv.ParseFloat(localctx.(*TpTimeContext).D.DataValue, 64)
			NT2, _ := strconv.ParseFloat(localctx.(*TpTimeContext).GetF1().GetD1().DataValue, 64)
			Result := NT1 * NT2
			localctx.(*TpTimeContext).DTemp.DataValue = fmt.Sprintf("%g", Result)
		} else {
			localctx.(*TpTimeContext).DTemp.DataType = "E"
			NT1, _ := strconv.Atoi(localctx.(*TpTimeContext).D.DataValue)
			NT2, _ := strconv.Atoi(localctx.(*TpTimeContext).GetF1().GetD1().DataValue)
			Result := NT1 * NT2
			localctx.(*TpTimeContext).DTemp.DataValue = strconv.Itoa(Result)
		}

		{
			p.SetState(36)

			var _x = p.Tp(localctx.(*TpTimeContext).DTemp)

			localctx.(*TpTimeContext).tp1 = _x
		}

		localctx.(*TpContext).D1 = localctx.(*TpTimeContext).GetTp1().GetD1()

	case Practica3copyParserEOF, Practica3copyParserPRight, Practica3copyParserPlus:
		localctx = NewTpEpContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		localctx.(*TpContext).D1 = localctx.(*TpEpContext).D

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
}

func NewEmptyFContext() *FContext {
	var p = new(FContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Practica3copyParserRULE_f
	return p
}

func (*FContext) IsFContext() {}

func NewFContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FContext {
	var p = new(FContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Practica3copyParserRULE_f

	return p
}

func (s *FContext) GetParser() antlr.Parser { return s.parser }

func (s *FContext) GetD1() *Data { return s.D1 }

func (s *FContext) SetD1(v *Data) { s.D1 = v }

func (s *FContext) CopyFrom(ctx *FContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
	s.D1 = ctx.D1
}

func (s *FContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FPEPContext struct {
	*FContext
	_e IEContext
}

func NewFPEPContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FPEPContext {
	var p = new(FPEPContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *FPEPContext) Get_e() IEContext { return s._e }

func (s *FPEPContext) Set_e(v IEContext) { s._e = v }

func (s *FPEPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FPEPContext) PLeft() antlr.TerminalNode {
	return s.GetToken(Practica3copyParserPLeft, 0)
}

func (s *FPEPContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *FPEPContext) PRight() antlr.TerminalNode {
	return s.GetToken(Practica3copyParserPRight, 0)
}

type FEntContext struct {
	*FContext
	ne antlr.Token
}

func NewFEntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FEntContext {
	var p = new(FEntContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *FEntContext) GetNe() antlr.Token { return s.ne }

func (s *FEntContext) SetNe(v antlr.Token) { s.ne = v }

func (s *FEntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FEntContext) Numbers() antlr.TerminalNode {
	return s.GetToken(Practica3copyParserNumbers, 0)
}

type FDecContext struct {
	*FContext
	nd antlr.Token
}

func NewFDecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FDecContext {
	var p = new(FDecContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *FDecContext) GetNd() antlr.Token { return s.nd }

func (s *FDecContext) SetNd(v antlr.Token) { s.nd = v }

func (s *FDecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FDecContext) Decimal() antlr.TerminalNode {
	return s.GetToken(Practica3copyParserDecimal, 0)
}

func (p *Practica3copyParser) F() (localctx IFContext) {
	localctx = NewFContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Practica3copyParserRULE_f)

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
	case Practica3copyParserPLeft:
		localctx = NewFPEPContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(42)
			p.Match(Practica3copyParserPLeft)
		}
		{
			p.SetState(43)

			var _x = p.E()

			localctx.(*FPEPContext)._e = _x
		}
		{
			p.SetState(44)
			p.Match(Practica3copyParserPRight)
		}

		localctx.(*FContext).D1 = localctx.(*FPEPContext).Get_e().GetD1()

	case Practica3copyParserNumbers:
		localctx = NewFEntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(47)

			var _m = p.Match(Practica3copyParserNumbers)

			localctx.(*FEntContext).ne = _m
		}

		localctx.(*FContext).D1 = &Data{}
		localctx.(*FContext).D1.DataType = "E"
		localctx.(*FContext).D1.DataValue = (func() string {
			if localctx.(*FEntContext).GetNe() == nil {
				return ""
			} else {
				return localctx.(*FEntContext).GetNe().GetText()
			}
		}())

	case Practica3copyParserDecimal:
		localctx = NewFDecContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(49)

			var _m = p.Match(Practica3copyParserDecimal)

			localctx.(*FDecContext).nd = _m
		}

		localctx.(*FContext).D1 = &Data{}
		localctx.(*FContext).D1.DataType = "D"
		localctx.(*FContext).D1.DataValue = (func() string {
			if localctx.(*FDecContext).GetNd() == nil {
				return ""
			} else {
				return localctx.(*FDecContext).GetNd().GetText()
			}
		}())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
