package languages

import "github.com/Theshawa/syntax-highlighter/highlighter/types"

var HTMLPatterns = map[types.TokenType]types.Pattern{
	types.TokenType("doctype-declaration"): types.Pattern(`(?i)<\!(?<match>DOCTYPE)\s`),
	types.TokenType("doctype-value"):       types.Pattern(`(?i)<\!DOCTYPE\s(?<match>\w+)`),
	types.TokenType("entity"):              types.Pattern(`(?i)(?<match>&\w+;)`),
}

func NewHTML() types.Language {
	patterns := map[types.TokenType]types.Pattern{}
	for tokenType, pattern := range XMLPatterns {
		patterns[tokenType] = pattern
	}
	for tokenType, pattern := range HTMLPatterns {
		patterns[tokenType] = pattern
	}
	return types.Language{
		Name:     "HTML",
		Patterns: patterns,
	}
}

var HTML = NewHTML()
