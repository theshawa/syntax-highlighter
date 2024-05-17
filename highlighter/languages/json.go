package languages

import "github.com/Theshawa/syntax-highlighter/highlighter/types"

var JSONPatterns = []*types.Pattern{
	types.NewPattern("key", "(?<match>\\\"\\w+\\\"):\\s*"),
	types.NewPattern("value-string", ":\\s*(?<match>\"([^\"\\\\]*(?:\\\\.[^\"\\\\]*)*)\")"),
	types.NewPattern("value-number", ":\\s*(?<match>-?\\d+(\\.\\d+)?([eE][+-]?\\d+)?)"),
	types.NewPattern("value-inbuilt", ":\\s*\\b(true|false)\\b"),
	types.NewPattern("value-inbuilt", ":\\s*\\b(null)\\b"),
}

var JSON = types.Language{
	Patterns: JSONPatterns,
	Name:     "JSON",
}
