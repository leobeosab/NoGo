//TODO: Implement color
package markdown

import (
	"bytes"
	"github.com/dstotijn/go-notion"
)

type RichText struct {
	IsLink      bool
	Link        string
	Annotations RichTextAnnotations
	Text        string
}

type RichTextAnnotations struct {
	Bold          bool
	Italic        bool
	Code          bool
	Strikethrough bool
}

func (p Page) RichTextToString(rt notion.RichText) (string, error) {

	link := ""
	if rt.HRef != nil {
		link = *rt.HRef
	}

	var annotations RichTextAnnotations
	if rt.Annotations == nil {
		annotations = RichTextAnnotations{}
	} else {
		annotations = RichTextAnnotations{
			Bold:          rt.Annotations.Bold,
			Italic:        rt.Annotations.Italic,
			Strikethrough: rt.Annotations.Strikethrough,
			Code:          rt.Annotations.Code,
		}
	}

	mdRt := RichText{
		IsLink:      rt.HRef != nil,
		Link:        link,
		Text:        rt.PlainText,
		Annotations: annotations,
	}

	t, err := p.FetchTemplate("RichTextTemplate.md")
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = t.Execute(&result, mdRt)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}

func (p Page) RichTextArrToString(rt []notion.RichText) (string, error) {

	outputStr := ""
	for _, t := range rt {
		result, err := p.RichTextToString(t)
		if err != nil {
			return "", err
		}

		outputStr += result
	}

	return outputStr + "\n", nil
}
