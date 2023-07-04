package main

import (
	"fmt"
	"strconv"

	"parser/arithmetic_parser/parser/lexer"

	"github.com/hsputra/go-cmake/parser/arithmetic_parser/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type evalListener struct {
	*parser.BaseArithmeticListener
	result float64
}

func newEvalListener() *evalListener {
	return &evalListener{}
}

func (l *evalListener) EnterNumberExpression(ctx *parser.NumberExpressionContext) {
	number, _ := strconv.ParseFloat(ctx.GetText(), 64)
	l.result = number
}

func (l *evalListener) EnterMulDivExpression(ctx *parser.MulDivExpressionContext) {
	left := ctx.GetExpression(0).(*parser.ExpressionContext)
	right := ctx.GetExpression(1).(*parser.ExpressionContext)

	left.Accept(l)
	right.Accept(l)

	if ctx.GetOp().GetTokenType() == lexer.ArithmeticLexerMUL {
		l.result = left.result * right.result
	} else {
		l.result = left.result / right.result
	}
}

func (l *evalListener) EnterAddSubExpression(ctx *parser.AddSubExpressionContext) {
	left := ctx.GetExpression(0).(*parser.ExpressionContext)
	right := ctx.GetExpression(1).(*parser.ExpressionContext)

	left.Accept(l)
	right.Accept(l)

	if ctx.GetOp().GetTokenType() == lexer.ArithmeticLexerADD {
		l.result = left.result + right.result
	} else {
		l.result = left.result - right.result
	}
}

func main() {
	input := "2 * (3 + 4)"
	inputStream := antlr.NewInputStream(input)
	lexer := lexer.NewArithmeticLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewArithmeticParser(tokenStream)

	tree := parser.Expression()

	listener := newEvalListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	fmt.Println(listener.result)
}
