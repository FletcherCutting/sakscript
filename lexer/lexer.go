package lexer

import (
	"../token"
)

func Start(fileContents []byte) []token.Token {
	tokens := make([]token.Token, 0)
    
    i := 0
    
    for i < len(fileContents) {
        if isWhitespace(fileContents[i]) {
            i += 1
            continue
        } else if fileContents[i] == '"' {
            t := readString(fileContents, i + 1)
            tokens = append(tokens, t)
            i += len(t.Value) + 2
        } else if isLetter(fileContents[i]) {
            t := readIdent(fileContents, i)
            tokens = append(tokens, t)
            i += len(t.Value)
        } else if isDigit(fileContents[i]) {
            t := readDigit(fileContents, i)
            tokens = append(tokens, t)
            i += len(t.Value)
        } else {
            t, j := readSymbol(fileContents, i)
            tokens = append(tokens, t)
            i += j
        }
    }
    
    var eof token.Token
    eof.Type = token.EOF
    tokens = append(tokens, eof)
    
    return tokens
}

func readString(fileContents []byte, i int) token.Token {
	var t token.Token
	value := ""

	for i < len(fileContents) {
		if fileContents[i] == '"' {
			break
		}

		value += string(fileContents[i])
		i += 1
	}

	t.Type = token.StringLiteral
	t.Value = value

	return t
}

func readIdent(fileContents []byte, i int) token.Token {
	var t token.Token
	value := ""

	for i < len(fileContents) {
		if !(isLetter(fileContents[i]) && !(isDigit(fileContents[i]))) {
			break
		}

		value += string(fileContents[i])
		i += 1
	}

	t.Type = token.Ident
	t.Value = value

	return t
}

func readDigit(fileContents []byte, i int) token.Token {
	var t token.Token
	value := ""

	for i < len(fileContents) {
		if !(isDigit(fileContents[i])) {
			break
		}

		value += string(fileContents[i])
		i += 1
	}

	t.Type = token.IntLiteral
	t.Value = value

	return t
}

func readSymbol(fileContents []byte, i int) (token.Token, int) {
    var t token.Token
    j := 1
    
    switch fileContents[i] {
    case '+':
        t.Type = token.Add
    case '-':
        t.Type = token.Sub
    case '*':
        t.Type = token.Mul
    case '/':
        t.Type = token.Div
    case '%':
        t.Type = token.Mod
    case '^':
        t.Type = token.Hat
    case '#':
        t.Type = token.Hash
    case '.':
        t.Type = token.Period
    case '(':
        t.Type = token.LeftParan
    case ')':
        t.Type = token.RightParan
    case '{':
        t.Type = token.LeftBrace
    case '}':
        t.Type = token.RightBrace
    case '[':
        t.Type = token.LeftBracket
    case ']':
        t.Type = token.RightBracket
    case '=':
        if i < len(fileContents) - 1 && fileContents[i + 1] == '=' {
            t.Type = token.DoubleEqual
            j = 2
        } else {
            t.Type = token.Equal
        }
    case '!':
        if i < len(fileContents) - 1 && fileContents[i + 1] == '=' {
            t.Type = token.NotEqual
            j = 2
        } else {
            t.Type = token.Bang
        }
    case '<':
        if i < len(fileContents) - 1 && fileContents[i + 1] == '=' {
            t.Type = token.LesserThanOrEqual
            j = 2
        } else {
            t.Type = token.LesserThan
        }
    case '>':
        if i < len(fileContents) - 1 && fileContents[i + 1] == '-' {
            t.Type = token.GreaterThanOrEqual
            j = 2
        } else {
            t.Type = token.GreaterThan
        }
    default:
        t.Type = token.Illegal
        t.Value = string(fileContents[i])
    }
    
    return t, j
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\n' || ch == '\t'
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '-'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}