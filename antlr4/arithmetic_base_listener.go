// Code generated from antlr4/Arithmetic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Arithmetic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseArithmeticListener is a complete listener for a parse tree produced by ArithmeticParser.
type BaseArithmeticListener struct{}

var _ ArithmeticListener = &BaseArithmeticListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseArithmeticListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseArithmeticListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseArithmeticListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseArithmeticListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterMulDivExpression is called when production mulDivExpression is entered.
func (s *BaseArithmeticListener) EnterMulDivExpression(ctx *MulDivExpressionContext) {}

// ExitMulDivExpression is called when production mulDivExpression is exited.
func (s *BaseArithmeticListener) ExitMulDivExpression(ctx *MulDivExpressionContext) {}

// EnterAddSubExpression is called when production addSubExpression is entered.
func (s *BaseArithmeticListener) EnterAddSubExpression(ctx *AddSubExpressionContext) {}

// ExitAddSubExpression is called when production addSubExpression is exited.
func (s *BaseArithmeticListener) ExitAddSubExpression(ctx *AddSubExpressionContext) {}

// EnterNumberExpression is called when production numberExpression is entered.
func (s *BaseArithmeticListener) EnterNumberExpression(ctx *NumberExpressionContext) {}

// ExitNumberExpression is called when production numberExpression is exited.
func (s *BaseArithmeticListener) ExitNumberExpression(ctx *NumberExpressionContext) {}

// EnterParensExpression is called when production parensExpression is entered.
func (s *BaseArithmeticListener) EnterParensExpression(ctx *ParensExpressionContext) {}

// ExitParensExpression is called when production parensExpression is exited.
func (s *BaseArithmeticListener) ExitParensExpression(ctx *ParensExpressionContext) {}
