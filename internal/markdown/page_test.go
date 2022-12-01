package markdown

import (
	"testing"
)

func Test_EmbedFS(t *testing.T) {
	page := GenPage("Test EmbedFS")

	page.FetchTemplate("blocks/CodeTemplate.md")
}
