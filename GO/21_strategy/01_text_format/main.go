package main

import (
	"fmt"
	"strings"
)

type OutputFormat int

const (
	MARKDOWN OutputFormat = iota
	HTML
)

type ListStrategy interface {
	Start(builder *strings.Builder)
	End(builder *strings.Builder)
	AddListItem(builder *strings.Builder, item string)
}

type MarkdownListStrategy struct{}

// * one
// * two
// Markdown format doesn't need a start or end statement like HTML

func (mls *MarkdownListStrategy) Start(builder *strings.Builder) {
	// Not required in case of Markdown
}
func (mls *MarkdownListStrategy) End(builder *strings.Builder) {
	// Not required in case of Markdown
}
func (mls *MarkdownListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(fmt.Sprintf(" * %s\n", item))
}

type HtmlListStrategy struct{}

// <ul> // starting of HTML list
// <li>Item</li>
// </ul> // starting of HTML list

func (hls *HtmlListStrategy) Start(builder *strings.Builder) {
	builder.WriteString("<ul>\n")
}
func (hls *HtmlListStrategy) End(builder *strings.Builder) {
	builder.WriteString("</ul>\n")
}
func (hls *HtmlListStrategy) AddListItem(builder *strings.Builder, item string) {
	builder.WriteString(fmt.Sprintf("	<li>%s</li>\n", item))
}

type TextProcessor struct {
	build        strings.Builder
	listStrategy ListStrategy
}

func NewTextProcessor(listStrategy ListStrategy) *TextProcessor {
	return &TextProcessor{
		build:        strings.Builder{},
		listStrategy: listStrategy,
	}
}

func (tp *TextProcessor) SetOutputFormat(format OutputFormat) {
	switch format {
	case MARKDOWN:
		tp.listStrategy = &MarkdownListStrategy{}
	case HTML:
		tp.listStrategy = &HtmlListStrategy{}
	}
}

func (tp *TextProcessor) AppendList(items []string) {
	s := tp.listStrategy
	s.Start(&tp.build)
	for _, item := range items {
		s.AddListItem(&tp.build, item)
	}
	s.End(&tp.build)
}

func (tp *TextProcessor) Reset() {
	tp.build.Reset()
}

func (tp *TextProcessor) String() string {
	return tp.build.String()
}

func main() {
	tp := NewTextProcessor(&MarkdownListStrategy{})
	tp.AppendList(
		[]string{
			"Coffee",
			"Rock Music",
			"Code",
		},
	)
	fmt.Println(tp)

	tp.Reset()

	tp.SetOutputFormat(HTML)
	tp.AppendList(
		[]string{
			"Coffee",
			"Rock Music",
			"Code",
		},
	)
	fmt.Println(tp)
}
