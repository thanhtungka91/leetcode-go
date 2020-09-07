package main

import (
	"fmt"
	"strings"
)

type HtmlElement struct {
	name, text string
	element []HtmlElement
}



func main() {

	hello := "hello"

	sb := strings.Builder{}

	sb.WriteString("<p>")
	sb.WriteString(hello)

	words := []string{"heelo", "work"}

	sb.Reset()

	for _, v := range words {
		sb.WriteString("<li>")
		sb.WriteString(v)
		sb.WriteString("</li>")
	}

	sb.WriteString("<ul>")


	b := NewHtmlBuilder("ul")
	b.AddChild("li","hello")
	fmt.Println(b.String())

}


type HtmlBuilder struct {
	rootName string
	root HtmlElement
}
func NewHtmlBuilder(rootName string) *HtmlBuilder {
	return &HtmlBuilder{
		rootName: rootName,
		root: HtmlElement{
			rootName,"", []HtmlElement{}}}
}

func (b *HtmlBuilder)AddChild(childName, childText string)  *HtmlBuilder{
	e := HtmlElement{
		childName,
		childText,
		[]HtmlElement{},
	}
	b.root.element = append(b.root.element, e)
	return b
}

func (e *HtmlElement) String() string {
	return e.string(0)
}

func (e *HtmlElement) string(indent int) string {
	sb := strings.Builder{}
	indentSize := 2
	i := strings.Repeat(" ", indentSize * indent)
	sb.WriteString(fmt.Sprintf("%s<%s>\n",
		i, e.name))
	if len(e.text) > 0 {
		sb.WriteString(strings.Repeat(" ",
			indentSize * (indent + 1)))
		sb.WriteString(e.text)
		sb.WriteString("\n")
	}

	for _, el := range e.element {
		sb.WriteString(el.string(indent+1))
	}
	sb.WriteString(fmt.Sprintf("%s</%s>\n",
		i, e.name))
	return sb.String()
}


func (b *HtmlBuilder) String() string {
	return b.root.String()
}