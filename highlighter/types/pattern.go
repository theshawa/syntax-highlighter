package types

import "regexp"

type Pattern struct {
	regex string
	Class Class
}

func (p *Pattern) GetRegex() *regexp.Regexp {
	return regexp.MustCompile(p.regex)
}

func NewPattern(class Class, regex string) *Pattern {
	return &Pattern{
		regex: regex,
		Class: class,
	}
}
