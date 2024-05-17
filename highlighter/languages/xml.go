package languages

import (
	"github.com/Theshawa/syntax-highlighter/highlighter/types"
)

var XMLPatterns = map[types.TokenType]types.Pattern{
	types.TokenType("tag"):             types.Pattern("(?i)<\\/?(?<match>[\\w\\-]+)"),
	types.TokenType("attr-name"):       types.Pattern(`(?i)(?<match>[\w\-]+)=["'][^"']*["']`),
	types.TokenType("attr-value"):      types.Pattern(`(?i)[\w]+=(?<match>["'][^"']*["'])`),
	types.TokenType("comment"):         types.Pattern(`(?m)(?<match>\<!--(.|\s)*-->)`),
	types.TokenType("xml-declaration"): types.Pattern(`(?i)<\?(?<match>xml)\s`),
}

var XML = types.Language{
	Patterns: XMLPatterns,
	Name:     "XML",
}
