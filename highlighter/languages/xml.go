package languages

import (
	types2 "github.com/Theshawa/syntax-highlighter/highlighter/types"
)

var XMLPatterns = map[types2.TokenType]types2.Pattern{
	types2.TokenType("tag"):        types2.Pattern("<\\/?(?<match>[\\w\\-]+)"),
	types2.TokenType("attr-name"):  types2.Pattern(`(?<match>[\w\-]+)=["'][^"']*["']`),
	types2.TokenType("attr-value"): types2.Pattern(`[\w]+=(?<match>["'][^"']*["'])`),
	types2.TokenType("comment"):    types2.Pattern(`(?m)(?<match>\<!--(.|\s)*-->)`),
}

var XML = types2.Language{
	Patterns: XMLPatterns,
	Name:     "XML",
}
