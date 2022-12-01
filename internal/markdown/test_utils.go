package markdown

import (
	"context"
	"regexp"
)

/* StripHTMLWhitespace
Turns:
<body>
	<div>
		<p>
			Hello
		</p>
	</div>
</body>
Into:
<body><div><p>Hello</p></div></body>
*/
func StripHTMLWhitespace(html string) string {
	r, err := regexp.Compile("(?:^ +\\s)|(?: *\\n+(?: |\\t)*)")
	if err != nil {
		panic(err)
	}

	return r.ReplaceAllString(html, "")
}

func GenPageContext() *PageContext {
	return &PageContext{
		C: context.Background(),
		Config: &PageConfig{
			BlocksDirectory: "",
			AssetURL:        "assets/img/posts/$PAGE_URI$/",
			AssetDirectory:  "static/assets/img/posts/$PAGE_URI$/",
		},
	}
}

func GenPage(testName string) *PageBuilder {
	return NewPageBuilder(GenPageContext(), "test-page-1", "", "PageBuilder for: "+testName, "11-11-2011")
}
