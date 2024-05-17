package languages

import (
	"github.com/Theshawa/syntax-highlighter/highlighter/types"
)

var XMLPatterns = []*types.Pattern{
	types.NewPattern("tag", "(?i)<\\/?(?<match>[\\w\\-]+)"),
	types.NewPattern("attr-name", "(?i)(?<match>[\\w\\-]+)=[\"'][^\"']*[\"']"),
	types.NewPattern("attr-value", "(?i)[\\w]+=(?<match>[\"'][^\"']*[\"'])"),
	types.NewPattern("comment", "(?m)(?<match>\\<!--(.|\\s)*-->)"),
	types.NewPattern("xml-declaration", "(?i)<\\?(?<match>xml)\\s"),
}

var XML = types.Language{
	Patterns: XMLPatterns,
	Name:     "XML",
}
