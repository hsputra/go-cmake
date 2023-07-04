// Code generated from antlr4/Arithmetic.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Arithmetic

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ArithmeticListener is a complete listener for a parse tree produced by ArithmeticParser.
type ArithmeticListener interface {
	antlr.ParseTreeListener

	// EnterMulDivExpression is called when entering the mulDivExpression production.
	EnterMulDivExpression(c *MulDivExpressionContext)

	// EnterAddSubExpression is called when entering the addSubExpression production.
	EnterAddSubExpression(c *AddSubExpressionContext)

	// EnterNumberExpression is called when entering the numberExpression production.
	EnterNumberExpression(c *NumberExpressionContext)

	// EnterParensExpression is called when entering the parensExpression production.
	EnterParensExpression(c *ParensExpressionContext)

	// ExitMulDivExpression is called when exiting the mulDivExpression production.
	ExitMulDivExpression(c *MulDivExpressionContext)

	// ExitAddSubExpression is called when exiting the addSubExpression production.
	ExitAddSubExpression(c *AddSubExpressionContext)

	// ExitNumberExpression is called when exiting the numberExpression production.
	ExitNumberExpression(c *NumberExpressionContext)

	// ExitParensExpression is called when exiting the parensExpression production.
	ExitParensExpression(c *ParensExpressionContext)
}
