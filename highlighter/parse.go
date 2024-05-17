package highlighter

import (
	"github.com/Theshawa/syntax-highlighter/highlighter/types"
	"sort"
)

func (h *Highlighter) Parse(content string, language types.Language) []types.Token {
	tokens := make([]types.Token, 0)
	for _, pattern := range language.Patterns {
		re := pattern.GetRegex()
		for _, match := range re.FindAllStringSubmatchIndex(content, -1) {
			if len(match) < 4 {
				continue
			}
			if match[2] != match[3] {
				tokens = append(tokens, types.Token{
					Class:   pattern.Class,
					Start:   match[2],
					End:     match[3],
					Content: content[match[2]:match[3]],
				})
			}
		}
	}

	sort.Slice(tokens, func(i, j int) bool {
		return tokens[i].Start < tokens[j].Start
	})

	return tokens
}
