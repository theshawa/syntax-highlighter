package types

type Token struct {
	Type    TokenType
	Start   int
	End     int
	Content string
}
