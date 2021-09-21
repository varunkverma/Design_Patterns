package main

// import (
// 	"fmt"
// 	"strings"
// 	"unicode"
// )

// type FormattedText struct {
// 	plainText  string
// 	capitalise []bool // same size as plainText, for each element we store bool value to incate whether the corresponding plainText element is capitalised or not
// }

// // factory
// func NewFormattedText(plainText string) *FormattedText {
// 	return &FormattedText{
// 		plainText:  plainText,
// 		capitalise: make([]bool, len(plainText)),
// 	}
// }

// func (f *FormattedText) String() string {
// 	sb := strings.Builder{}

// 	for i := 0; i < len(f.plainText); i++ {
// 		c := f.plainText[i]
// 		if f.capitalise[i] {
// 			sb.WriteRune(unicode.ToUpper(rune(c)))
// 		} else {
// 			sb.WriteRune(rune(c))
// 		}
// 	}
// 	return sb.String()
// }

// func (f *FormattedText) Capitalise(start, end int) {
// 	for i := start; i <= end; i++ {
// 		f.capitalise[i] = true
// 	}
// }

// func main() {
// 	text := "This is a brave new world"

// 	ft := NewFormattedText(text)
// 	ft.Capitalise(10, 15)
// 	fmt.Println(ft.String())
// }
