package markdown

import (
	"github.com/dstotijn/go-notion"
	"os"
	"testing"
)

func TestPage_AddImageToPage(t *testing.T) {
	os.Setenv("BLOCKS_DIRECTORY", "../../blocks")
	os.Setenv("ASSET_URL", "/assets/img/$PAGE_URI$/")

	// Shameless plug
	sampleImageURL := "https://ryanwise.me/assets/img/main/logo.jpg"
	inputs := []struct {
		notionImage notion.ImageBlock
		expected    string
		testString  string
	}{
		{
			notionImage: notion.ImageBlock{
				File: &notion.FileFile{
					URL: sampleImageURL,
				},
				Caption: []notion.RichText{
					{PlainText: "Hello there"},
				},
			},
			expected:   "<figure><img src=\"/assets/img/test-page-1/logo.jpg\" /><figcaption>Hello there</figcaption></figure>",
			testString: "Test non external image",
		},
	}

	for _, item := range inputs {
		expected := item.expected
		testString := item.testString

		page := NewPage(GenPageContext(), "Test Page 1", "")
		err := page.AddImageToPage(&item.notionImage)
		if err != nil {
			t.Fatalf("ERR: %s \n error: %s\n", testString, err)
		}

		result := StripHTMLWhitespace(page.Build())

		if result != item.expected {
			t.Errorf("FAILED: %s \nExpected: %s\nActual: %s\n", testString, expected, result)
		} else {
			t.Logf("PASSED: %s\n", item.testString)
		}
	}
}
