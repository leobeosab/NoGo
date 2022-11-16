package markdown

import (
	"bytes"
	"github.com/dstotijn/go-notion"
	"os"
	"text/template"
)

type heading1markdown struct {
	Text string
}

func newHeading1Markdown(block notion.Heading1Block) heading1markdown {
	s, err := RichTextArrToString(block.RichText)
	if err != nil {
		panic(err)
	}

	return heading1markdown{
		Text: s,
	}
}

func (p *Page) AddHeading1ToPage(block *notion.Heading1Block) error {
	md := newHeading1Markdown(*block)

	template, err := template.ParseFiles(os.Getenv("BLOCKS_DIRECTORY") + "/Heading1Template.md")
	if err != nil {
		return err
	}

	var mdBuffer bytes.Buffer
	err = template.Execute(&mdBuffer, md)
	if err != nil {
		return err
	}

	p.AddBlock(mdBuffer.String())

	return nil
}
