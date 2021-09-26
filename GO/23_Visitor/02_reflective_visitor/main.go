package main

import (
	"fmt"
	"strings"
)

type Expression interface {
}

type DoubleExpression struct {
	value float64
}

type AdditionExpression struct {
	left, right Expression
}

func Print(ex Expression, sb *strings.Builder) {
	if dex, ok := ex.(*DoubleExpression); ok {
		sb.WriteString(fmt.Sprintf("%g", dex.value))
	} else if aex, ok := ex.(*AdditionExpression); ok {
		sb.WriteRune('(')
		Print(aex.left, sb)
		sb.WriteRune('+')
		Print(aex.right, sb)
		sb.WriteRune(')')
	}

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
	sb := &strings.Builder{}
	Print(ex, sb)
	fmt.Println(sb.String())
}
