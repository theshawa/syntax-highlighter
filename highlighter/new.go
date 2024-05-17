package highlighter

import (
	"github.com/Theshawa/syntax-highlighter/highlighter/languages"
	"github.com/Theshawa/syntax-highlighter/highlighter/types"
)

func New() *Highlighter {
	return &Highlighter{
		Languages: map[string]types.Language{
			"xml":  languages.XML,
			"html": languages.HTML,
			"json": languages.JSON,
			"yaml": languages.YAML,
		},
	}
}
