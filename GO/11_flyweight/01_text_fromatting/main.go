package main

import (
	"fmt"
	"strings"
	"unicode"
)

//-----------------------------------------------------------------------------------

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
//---------------------------------------------------------------------------------

// Flyweight
type TextRange struct {
	Start, End               int
	Capitalise, Bold, Italic bool
}

// check whether range is covering a particular position
func (t *TextRange) Covers(position int) bool {
	return position >= t.Start && position <= t.End
}

type BetterFormattedText struct {
	plainText  string
	formatting []*TextRange // pointer as we want to share these text Ranges
}

// factory
func NewBetterFormattedText(plainText string) *BetterFormattedText {
	return &BetterFormattedText{
		plainText:  plainText,
		formatting: make([]*TextRange, 0),
	}
}

func (bft *BetterFormattedText) Range(start, end int) *TextRange {
	tr := &TextRange{
		Start:      start,
		End:        end,
		Capitalise: false,
		Bold:       false,
		Italic:     false,
	}
	bft.formatting = append(bft.formatting, tr)
	return tr
}

func (bft *BetterFormattedText) String() string {
	sb := strings.Builder{}

	for i := 0; i < len(bft.plainText); i++ {
		c := bft.plainText[i]
		for _, tr := range bft.formatting {
			if tr.Covers(i) {
				if tr.Capitalise {
					c = uint8(unicode.ToUpper(rune(c)))
				}
				// logic for bold check
				// logic for italics check
			}
		}
		sb.WriteRune(rune(c))
	}

	return sb.String()
}

func main() {
	text := "This is a brave new world"

	// ft := NewFormattedText(text)
	// ft.Capitalise(10, 15)
	// fmt.Println(ft.String())

	bft := NewBetterFormattedText(text)
	bft.Range(10, 15).Capitalise = true
	fmt.Println(bft.String())
}
