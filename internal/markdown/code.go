package markdown

import (
	"bytes"
	"fmt"
	"github.com/dstotijn/go-notion"
	"text/template"
)

type Code struct {
	Code     string
	Language string
}

func newCode(block notion.CodeBlock) Code {
	s := ""

	for _, text := range block.RichText {
		s += text.PlainText
	}

	return Code{
		Code:     s,
		Language: *block.Language,
	}
}

func (p Page) AddCodeToPage(block *notion.CodeBlock) error {
	md := newCode(*block)

	template, err := template.ParseFiles("blocks/CodeTemplate.md")
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	err = template.Execute(&mdBuffer, md)
	if err != nil {
		return err
	}

	p.sbContent.WriteString(mdBuffer.String())
	fmt.Println(mdBuffer.String())

	return nil
}
