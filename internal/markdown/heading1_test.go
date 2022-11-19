package markdown

import (
	"github.com/dstotijn/go-notion"
	"os"
	"testing"
)

func TestPage_AddHeading1ToPage(t *testing.T) {
	os.Setenv("BLOCKS_DIRECTORY", "../../blocks")

	sampleURL := "https://ryanwise.me"
	inputs := []struct {
		notionHeading notion.Heading1Block
		expected      string
		testString    string
	}{
		{
			notionHeading: notion.Heading1Block{
				RichText: []notion.RichText{
					{PlainText: "Hello, ", Annotations: &notion.Annotations{}},
					{PlainText: "World!", Annotations: &notion.Annotations{}},
				},
			},
			expected:   "\n# Hello, World!\n",
			testString: "Test basic heading structure no annotations",
		},
		{
			notionHeading: notion.Heading1Block{
				RichText: []notion.RichText{
					{PlainText: "Hello, ", Annotations: &notion.Annotations{Bold: true}},
					{PlainText: "World!", Annotations: &notion.Annotations{Strikethrough: true}},
				},
			},
			expected:   "\n# **Hello, **~~World!~~\n",
			testString: "Test heading with annotations",
		},
		{
			notionHeading: notion.Heading1Block{
				RichText: []notion.RichText{
					{PlainText: "Hello, ", Annotations: &notion.Annotations{Bold: true}},
					{PlainText: "World!", Annotations: &notion.Annotations{}, HRef: &sampleURL},
				},
			},
			expected:   "\n# **Hello, **[World!](https://ryanwise.me)\n",
			testString: "Test heading with annotations and link",
		},
	}

	for _, item := range inputs {
		expected := item.expected
		testString := item.testString

		page := NewPage(GenPageContext(), "Test Page 1", "")
		err := page.AddHeading1ToPage(&item.notionHeading)
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
