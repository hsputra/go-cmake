// Code generated from antlr4/Arithmetic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Arithmetic

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 9, 24, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 11, 10, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22, 11, 2, 3,
	2, 2, 3, 2, 3, 2, 2, 4, 3, 2, 5, 6, 3, 2, 7, 8, 2, 25, 2, 10, 3, 2, 2,
	2, 4, 5, 8, 2, 1, 2, 5, 6, 7, 3, 2, 2, 6, 7, 5, 2, 2, 2, 7, 8, 7, 4, 2,
	2, 8, 11, 3, 2, 2, 2, 9, 11, 7, 9, 2, 2, 10, 4, 3, 2, 2, 2, 10, 9, 3, 2,
	2, 2, 11, 20, 3, 2, 2, 2, 12, 13, 12, 5, 2, 2, 13, 14, 9, 2, 2, 2, 14,
	19, 5, 2, 2, 6, 15, 16, 12, 4, 2, 2, 16, 17, 9, 3, 2, 2, 17, 19, 5, 2,
	2, 5, 18, 12, 3, 2, 2, 2, 18, 15, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18,
	3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 3, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2,
	5, 10, 18, 20,
}
var literalNames = []string{
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "NUMBER",
}

var ruleNames = []string{
	"expression",
}

type ArithmeticParser struct {
	*antlr.BaseParser
}

// NewArithmeticParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *ArithmeticParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewArithmeticParser(input antlr.TokenStream) *ArithmeticParser {
	this := new(ArithmeticParser)
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
	this.GrammarFileName = "Arithmetic.g4"

	return this
}

// ArithmeticParser tokens.
const (
	ArithmeticParserEOF    = antlr.TokenEOF
	ArithmeticParserT__0   = 1
	ArithmeticParserT__1   = 2
	ArithmeticParserT__2   = 3
	ArithmeticParserT__3   = 4
	ArithmeticParserT__4   = 5
	ArithmeticParserT__5   = 6
	ArithmeticParserNUMBER = 7
)

// ArithmeticParserRULE_expression is the ArithmeticParser rule.
const ArithmeticParserRULE_expression = 0

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ArithmeticParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ArithmeticParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulDivExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMulDivExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivExpressionContext {
	var p = new(MulDivExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivExpressionContext) GetOp() antlr.Token { return s.op }

func (s *MulDivExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *MulDivExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArithmeticListener); ok {
		listenerT.EnterMulDivExpression(s)
	}
}

func (s *MulDivExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArithmeticListener); ok {
		listenerT.ExitMulDivExpression(s)
	}
}

type AddSubExpressionContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddSubExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubExpressionContext {
	var p = new(AddSubExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubExpressionContext) GetOp() antlr.Token { return s.op }

func (s *AddSubExpressionContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AddSubExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArithmeticListener); ok {
		listenerT.EnterAddSubExpression(s)
	}
}

func (s *AddSubExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArithmeticListener); ok {
		listenerT.ExitAddSubExpression(s)
	}
}

type NumberExpressionContext struct {
	*ExpressionContext
}

func NewNumberExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberExpressionContext {
	var p = new(NumberExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ArithmeticParserNUMBER, 0)
}

func (s *NumberExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArithmeticListener); ok {
		listenerT.EnterNumberExpression(s)
	}
}

func (s *NumberExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArithmeticListener); ok {
		listenerT.ExitNumberExpression(s)
	}
}

type ParensExpressionContext struct {
	*ExpressionContext
}

func NewParensExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensExpressionContext {
	var p = new(ParensExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParensExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParensExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArithmeticListener); ok {
		listenerT.EnterParensExpression(s)
	}
}

func (s *ParensExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ArithmeticListener); ok {
		listenerT.ExitParensExpression(s)
	}
}

func (p *ArithmeticParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *ArithmeticParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, ArithmeticParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(8)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ArithmeticParserT__0:
		localctx = NewParensExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(3)
			p.Match(ArithmeticParserT__0)
		}
		{
			p.SetState(4)
			p.expression(0)
		}
		{
			p.SetState(5)
			p.Match(ArithmeticParserT__1)
		}

	case ArithmeticParserNUMBER:
		localctx = NewNumberExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(7)
			p.Match(ArithmeticParserNUMBER)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(16)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ArithmeticParserRULE_expression)
				p.SetState(10)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(11)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ArithmeticParserT__2 || _la == ArithmeticParserT__3) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(12)
					p.expression(4)
				}

			case 2:
				localctx = NewAddSubExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, ArithmeticParserRULE_expression)
				p.SetState(13)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(14)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubExpressionContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ArithmeticParserT__4 || _la == ArithmeticParserT__5) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubExpressionContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(15)
					p.expression(3)
				}

			}

		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *ArithmeticParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ArithmeticParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
