package markdown

import "regexp"

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
