package markdown

import (
	"github.com/dstotijn/go-notion"
	"os"
	"testing"
)

func TestPage_AddCodeToPage(t *testing.T) {
	os.Setenv("BLOCKS_DIRECTORY", "../../blocks")

	sampleLanguage := "Go"
	inputs := []struct {
		notionCode notion.CodeBlock
		expected   string
		testString string
	}{
		{
			notionCode: notion.CodeBlock{
				RichText: []notion.RichText{
					{PlainText: "cat ~/bash.rc"},
				},
			},
			expected:   "\n```\ncat ~/bash.rc\n```",
			testString: "Test code block with no language",
		},
		{
			notionCode: notion.CodeBlock{
				RichText: []notion.RichText{
					{PlainText: "fmt.Println(\"Hello, World!\")"},
				},
				Language: &sampleLanguage,
			},
			expected:   "\n```Go\nfmt.Println(\"Hello, World!\")\n```",
			testString: "Test code block with language",
		},
	}

	for _, item := range inputs {
		expected := item.expected
		testString := item.testString

		page := NewPage("Test Page 1")
		err := page.AddCodeToPage(&item.notionCode)
		if err != nil {
			t.Fatalf("ERR: %s \n error: %s\n", testString, err)
		}

		result := page.Build()

		if result != item.expected {
			t.Errorf("FAILED: %s \nExpected: %s\nActual: %s\n", testString, expected, result)
		} else {
			t.Logf("PASSED: %s\n", item.testString)
		}
	}
}
