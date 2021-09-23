package main

import (
	"fmt"
	"strings"
	"unicode"
)

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

func main() {
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	for _, t := range tokens {
		fmt.Printf("%s ", t.String())
	}
	fmt.Println()
}
