package main

import (
	"fmt"
	"strings"
)

const (
	indentSize = 2
)

type HtmlElement struct {
	name, text string
	elements   []HtmlElement
}

// This is helper method to call the recursive method
func (e *HtmlElement) String() string {
	return e.string(0)
}

// recursive function
func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}

	// entering required indentation
	i := strings.Repeat(" ", indentSize*indent)
	// writing the opening intdented html element
	sb.WriteString(fmt.Sprintf("%s<%s>\n", i, e.name))
	//  writing the internal text of the HTMl element
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ", indentSize*(indent+1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	// writing the inner html elements, if any / recursive call
	for _, innerEle := range e.elements {
		sb.WriteString(innerEle.string(indent + 1))
	}

	// writing the closing intdented html element
	sb.WriteString(fmt.Sprintf("%s</%s>\n", i, e.name))

	return sb.String()
}

// Builder

type HtmlBuilder struct {
	rootName string
	root     HtmlElement
}

// constructor function
func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root: HtmlElement{
			name:     rootName,
			text:     "",
			elements: []HtmlElement{},
		},
	}
}

func (hb *HtmlBuilder) String() string {
	return hb.root.String()
}

func (hb *HtmlBuilder) AddChild(childName, childText string) {
	e := HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	}

	hb.root.elements = append(hb.root.elements, e)
}

// fluent version
func (hb *HtmlBuilder) AddChildFluent(childName, childText string) *HtmlBuilder {
	e := HtmlElement{
		name:     childName,
		text:     childText,
		elements: []HtmlElement{},
	}

	hb.root.elements = append(hb.root.elements, e)
	return hb
}

func main() {
	// bad way
	// msg := "hello"

	// sb := strings.Builder{}

	// sb.WriteString("<p>")
	// sb.WriteString(msg)
	// sb.WriteString("</p>")
	// fmt.Println(sb.String())

	// words := []string{"hello", "world"}
	// sb.Reset()
	// // <ul><li>...</li></ul>
	// sb.WriteString("<ul>")
	// for _, w := range words {
	// 	sb.WriteString("<li>")
	// 	sb.WriteString(w)
	// 	sb.WriteString("</li>")
	// }
	// sb.WriteString("</ul>")
	// fmt.Println(sb.String())

	// better way
	b := NewHtmlBuilder("ul")
	b.AddChild("li", "hello")
	b.AddChildFluent("li", "world").
		AddChild("li", "Nice to meet ya")
	fmt.Println(b.String())
}
