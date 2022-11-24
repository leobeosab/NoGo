package markdown

import (
	"testing"
)

func Test_EmbedFS(t *testing.T) {
	page := NewPage(GenPageContext(), "", "")

	page.FetchTemplate("blocks/CodeTemplate.md")
}
