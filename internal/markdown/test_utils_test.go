package markdown

import (
	"fmt"
	"testing"
)

func TestStripHTMLWhitespace(t *testing.T) {
	inputs := []struct {
		s        string
		expected string
		testCase string
	}{
		{
			s:        "<div>\n\t<p>\n\t\tHello\n\t</p>\n</div>",
			expected: "<div><p>Hello</p></div>",
			testCase: "Test basic whitespace removal",
		},
	}

	for _, item := range inputs {
		expected := item.expected
		result := StripHTMLWhitespace(item.s)
		testString := item.testCase

		fmt.Println(testString)

		if result != item.expected {
			t.Errorf("FAILED: %s \nExpected: %s\nActual: %s\n", testString, expected, result)
		}
	}
}
