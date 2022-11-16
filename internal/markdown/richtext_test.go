package markdown

import (
	"github.com/dstotijn/go-notion"
	"os"
	"testing"
)

func TestRichTextToString(t *testing.T) {
	os.Setenv("BLOCKS_DIRECTORY", "../../blocks")

	sampleURL := "https://google.com"
	inputs := []struct {
		input      notion.RichText
		expected   string
		testString string
	}{
		{
			input: notion.RichText{
				PlainText:   "Hahaha!",
				Annotations: &notion.Annotations{Bold: true, Italic: true},
			},
			expected:   "***Hahaha!***\n",
			testString: "Testing italics and bold",
		},
		{
			input: notion.RichText{
				PlainText:   "Hahaha!",
				Annotations: &notion.Annotations{Strikethrough: true},
			},
			expected:   "~~Hahaha!~~\n",
			testString: "Testing strikethrough",
		},
		{
			input: notion.RichText{
				PlainText:   "Hahaha!",
				Annotations: &notion.Annotations{},
				HRef:        &sampleURL,
			},
			expected:   "[Hahaha!](https://google.com)\n",
			testString: "Testing URL",
		},
		{
			input: notion.RichText{
				PlainText:   "Hahaha!",
				Annotations: &notion.Annotations{Bold: true, Strikethrough: true, Italic: true},
				HRef:        &sampleURL,
			},
			expected:   "[***~~Hahaha!~~***](https://google.com)\n",
			testString: "Testing all",
		},
	}

	for _, item := range inputs {
		result, err := RichTextToString(item.input)
		if err != nil {
			t.Fatalf("ERR: %s \n error: %s\n", item.testString, err)
		}

		if result != item.expected {
			t.Errorf("FAILED: %s \nExpected: %s\nActual: %s\n", item.testString, item.expected, result)
		} else {
			t.Logf("PASSED: %s\n", item.testString)
		}
	}
}

func TestRichTextArrToString(t *testing.T) {
	os.Setenv("BLOCKS_DIRECTORY", "../../blocks")

	inputs := []struct {
		input      []notion.RichText
		expected   string
		testString string
	}{
		{
			input: []notion.RichText{
				{Annotations: &notion.Annotations{Bold: true}, PlainText: "Hello"},
				{Annotations: &notion.Annotations{Strikethrough: true}, PlainText: " World"},
				{Annotations: &notion.Annotations{Italic: true}, PlainText: "!"},
			},
			expected:   "**Hello**~~ World~~*!*\n",
			testString: "Testing multiple rich text on one line",
		},
	}

	for _, item := range inputs {
		result, err := RichTextArrToString(item.input)
		if err != nil {
			t.Fatalf("ERR: %s \n error: %s\n", item.testString, err)
		}

		if result != item.expected {
			t.Errorf("FAILED: %s \nExpected: %s\nActual: %s\n", item.testString, item.expected, result)
		} else {
			t.Logf("PASSED: %s\n", item.testString)
		}
	}
}
