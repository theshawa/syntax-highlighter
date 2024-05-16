package types

import "regexp"

type Pattern string

func (p *Pattern) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(string(*p))
}
