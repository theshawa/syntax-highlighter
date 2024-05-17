package languages

import "github.com/Theshawa/syntax-highlighter/highlighter/types"

var HTMLPatterns = append([]*types.Pattern{
	types.NewPattern("doctype-declaration", "(?i)<\\!(?<match>DOCTYPE)\\s"),
	types.NewPattern("doctype-declaration-value", "(?i)<\\!DOCTYPE\\s(?<match>\\w+)"),
	types.NewPattern("entity", "(?i)(?<match>&\\w+;)"),
}, XMLPatterns...)

var HTML = types.Language{
	Name:     "HTML",
	Patterns: HTMLPatterns,
}
