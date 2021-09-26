package main

import (
	"fmt"
	"strings"
)

type ExpressionVisitor interface {
	VisitDoubleExpression(dex *DoubleExpression)
	VisitAdditionExpression(aex *AdditionExpression)
}

type Expression interface {
	Accept(ev ExpressionVisitor)
}

type DoubleExpression struct {
	value float64
}

// double jump
func (dex *DoubleExpression) Accept(ev ExpressionVisitor) {
	ev.VisitDoubleExpression(dex)
}

type AdditionExpression struct {
	left, right Expression
}

// double jump
func (aex *AdditionExpression) Accept(ev ExpressionVisitor) {
	ev.VisitAdditionExpression(aex)
}

type ExpressionPrinter struct {
	sb strings.Builder
}

func (ep *ExpressionPrinter) VisitDoubleExpression(dex *DoubleExpression) {
	ep.sb.WriteString(fmt.Sprintf("%g", dex.value))
}
func (ep *ExpressionPrinter) VisitAdditionExpression(aex *AdditionExpression) {
	ep.sb.WriteRune('(')
	aex.left.Accept(ep)
	ep.sb.WriteRune('+')
	aex.right.Accept(ep)
	ep.sb.WriteRune(')')
}

func (ep *ExpressionPrinter) String() string {
	return ep.sb.String()
}

func NewExpressionPrinter() *ExpressionPrinter {
	return &ExpressionPrinter{
		sb: strings.Builder{},
	}
}

type ExpressionEvaluator struct {
	result float64
}

func (ee *ExpressionEvaluator) VisitDoubleExpression(dex *DoubleExpression) {
	ee.result = dex.value
}
func (ee *ExpressionEvaluator) VisitAdditionExpression(aex *AdditionExpression) {
	aex.left.Accept(ee)
	res := ee.result

	aex.right.Accept(ee)
	res += ee.result
	ee.result = res
}

func main() {
	// 1 + (2+3)
	ex := &AdditionExpression{
		left: &DoubleExpression{1},
		right: &AdditionExpression{
			left:  &DoubleExpression{2},
			right: &DoubleExpression{3},
		},
	}
	ep := NewExpressionPrinter()
	ex.Accept(ep)

	ee := &ExpressionEvaluator{}
	ex.Accept(ee)
	fmt.Printf("%s = %g\n", ep, ee.result)
}
