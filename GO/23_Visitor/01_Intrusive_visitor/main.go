package main

import (
	"fmt"
	"strings"
)

type Expression interface {
	Print(sb *strings.Builder) // sb i.e., the *strings.Builder is the visitor here
}

type DoubleExpression struct {
	value float64
}

func (dex *DoubleExpression) Print(sb *strings.Builder) {
	sb.WriteString(fmt.Sprintf("%g", dex.value))
}

type AdditionExpression struct {
	left, right Expression
}

func (aex *AdditionExpression) Print(sb *strings.Builder) {
	sb.WriteRune('(')
	aex.left.Print(sb)
	sb.WriteRune('+')
	aex.right.Print(sb)
	sb.WriteRune(')')
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
	sb := strings.Builder{}
	ex.Print(&sb)
	fmt.Println(sb.String())
}
