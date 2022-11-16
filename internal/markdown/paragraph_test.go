package markdown

import (
	"github.com/dstotijn/go-notion"
	"os"
	"testing"
)

func TestPage_AddParagraphToPage(t *testing.T) {
	os.Setenv("BLOCKS_DIRECTORY", "../../blocks")

	sampleURL := "https://ryanwise.me"
	inputs := []struct {
		notionParagraph notion.ParagraphBlock
		expected        string
		testString      string
	}{
		{
			notionParagraph: notion.ParagraphBlock{
				RichText: []notion.RichText{
					{PlainText: "Hello, ", Annotations: &notion.Annotations{}},
					{PlainText: "World!", Annotations: &notion.Annotations{}},
				},
			},
			expected:   "\nHello, World!\n",
			testString: "Test basic paragraph structure no annotations",
		},
		{
			notionParagraph: notion.ParagraphBlock{
				RichText: []notion.RichText{
					{PlainText: "Hello, ", Annotations: &notion.Annotations{Bold: true}},
					{PlainText: "World!", Annotations: &notion.Annotations{Strikethrough: true}},
				},
			},
			expected:   "\n**Hello, **~~World!~~\n",
			testString: "Test paragraph with annotations",
		},
		{
			notionParagraph: notion.ParagraphBlock{
				RichText: []notion.RichText{
					{PlainText: "Hello, ", Annotations: &notion.Annotations{Bold: true}},
					{PlainText: "World!", Annotations: &notion.Annotations{}, HRef: &sampleURL},
				},
			},
			expected:   "\n**Hello, **[World!](https://ryanwise.me)\n",
			testString: "Test paragraph with annotations and link",
		},
	}

	for _, item := range inputs {
		expected := item.expected
		testString := item.testString

		page := NewPage("Test Page 1")
		err := page.AddParagraphToPage(&item.notionParagraph)
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
