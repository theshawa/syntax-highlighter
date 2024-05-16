package highlighter

import (
	"errors"
	types2 "github.com/Theshawa/syntax-highlighter/highlighter/types"
)

type Highlighter struct {
	Languages map[string]types2.Language
}

func (h *Highlighter) Highlight(content string, languageName string) ([]types2.Token, error) {
	if language, found := h.Languages[languageName]; found {
		tokens := h.Parse(content, language)
		return tokens, nil
	} else {
		return nil, errors.New("language not found")
	}
}
