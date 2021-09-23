package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Element interface {
	Value() int
}
type Integer struct {
	value int
}

func (i *Integer) Value() int {
	return i.value
}

func NewInteger(value int) *Integer {
	return &Integer{
		value: value,
	}
}

type Operation int

const (
	ADDITION Operation = iota
	SUBTRACTION
)

type BinaryOperation struct {
	Type        Operation
	Left, Right Element
}

func (bo *BinaryOperation) Value() int {
	switch bo.Type {
	case ADDITION:
		return bo.Left.Value() + bo.Right.Value()
	case SUBTRACTION:
		return bo.Left.Value() - bo.Right.Value()
	default:
		panic("Unsupported Operations")
	}
}

type TokenType int

const (
	INT TokenType = iota
	PLUS
	MINUS
	LPAREN
	RPAREN
)

type Token struct {
	Type TokenType
	Text string
}

func (t *Token) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string) []Token {
	var result []Token

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '+':
			result = append(result, Token{PLUS, "+"})
		case '-':
			result = append(result, Token{MINUS, "-"})
		case '(':
			result = append(result, Token{LPAREN, "("})
		case ')':
			result = append(result, Token{RPAREN, ")"})
		default:
			sb := strings.Builder{}
			for j := i; j < len(input); j++ {
				if unicode.IsDigit(rune(input[j])) {
					sb.WriteRune(rune(input[j]))
					i++
				} else {
					result = append(result, Token{
						Type: INT,
						Text: sb.String(),
					})
					i--
					break
				}
			}
		}
	}
	return result
}

func Parse(tokens []Token) Element {
	result := BinaryOperation{}
	haveLHS := false
	for i := 0; i < len(tokens); i++ {
		token := &tokens[i]
		switch token.Type {
		case INT:
			n, _ := strconv.Atoi(token.Text)
			integer := &Integer{
				value: n,
			}
			if !haveLHS {
				result.Left = integer
				haveLHS = true
			} else {
				result.Right = integer
			}
		case PLUS:
			result.Type = ADDITION
		case MINUS:
			result.Type = SUBTRACTION
		case LPAREN:
			j := i
			for ; j < len(tokens); j++ {
				if tokens[j].Type == RPAREN {
					break
				}
			}
			var subExpression []Token
			for k := i + 1; k < j; k++ {
				subExpression = append(subExpression, tokens[k])
			}
			element := Parse(subExpression)
			if !haveLHS {
				result.Left = element
				haveLHS = true
			} else {
				result.Right = element
			}
			i = j
		}
	}
	return &result
}

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	for _, t := range tokens {
		fmt.Printf("%s ", t.String())
	}
	fmt.Println()

	parsed := Parse(tokens)
	fmt.Printf("%s = %d\n", input, parsed.Value())
}
