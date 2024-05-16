package types

type Language struct {
	Name     string
	Patterns map[TokenType]Pattern
}
