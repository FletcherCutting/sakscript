package token

type Token struct {
	Type int
	Value string
}

const (
	Illegal = 0
	EOF = 1

	StringLiteral = 2
	IntLiteral = 3
	Ident = 4

	Add = 5
	Sub = 6
	Mul = 7
	Div = 8
	Mod = 9
	Hat = 10
	Hash = 11
	Period = 12

	LeftParan = 13
	RightParan = 14
	LeftBrace = 15
	RightBrace = 16
	LeftBracket = 17
	RightBracket = 18

	Equal = 19
	DoubleEqual = 20
	Bang = 21
	NotEqual = 22
	LesserThan = 23
	LesserThanOrEqual = 24
	GreaterThan = 26
	GreaterThanOrEqual = 27
)