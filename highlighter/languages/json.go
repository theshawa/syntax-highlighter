package languages

import "github.com/Theshawa/syntax-highlighter/highlighter/types"

var JSONPatterns = map[types.TokenType]types.Pattern{
	types.TokenType("key"):           types.Pattern("(?<match>\\\"\\w+\\\"):\\s*"),
	types.TokenType("value-string"):  types.Pattern(":\\s*(?<match>\"([^\"\\\\]*(?:\\\\.[^\"\\\\]*)*)\")"),
	types.TokenType("value-number"):  types.Pattern(":\\s*(?<match>-?\\d+(\\.\\d+)?([eE][+-]?\\d+)?)"),
	types.TokenType("value-boolean"): types.Pattern(":\\s*\\b(true|false)\\b"),
	types.TokenType("value-null"):    types.Pattern(":\\s*\\b(null)\\b"),
}

var JSON = types.Language{
	Patterns: JSONPatterns,
	Name:     "JSON",
}
