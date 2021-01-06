package lexer

import "github.com/konojunya/monkey-language/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case 61: // =
		if l.peekChar() == 61 {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			t = token.Token{Type: token.EQ, Literal: literal}
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
	case 59: // ;
		t = newToken(token.SEMICOLON, l.ch)
	case 40: // (
		t = newToken(token.LPAREN, l.ch)
	case 41: // )
		t = newToken(token.RPAREN, l.ch)
	case 44: // ,
		t = newToken(token.COMMA, l.ch)
	case 43: // +
		t = newToken(token.PLUS, l.ch)
	case 45: // -
		t = newToken(token.MINUS, l.ch)
	case 33: // !
		if l.peekChar() == 61 {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			t = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case 42: // *
		t = newToken(token.ASTERISK, l.ch)
	case 47: // /
		t = newToken(token.SLASH, l.ch)
	case 60: // <
		t = newToken(token.LT, l.ch)
	case 62: // >
		t = newToken(token.GT, l.ch)
	case 123: // {
		t = newToken(token.LBRACE, l.ch)
	case 125: // }
		t = newToken(token.RBRACE, l.ch)
	case 0: // EOF
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)

			return t
		} else if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()

			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return t
}

func (l *Lexer) skipWhitespace() {
	// whitespace or \t or \n or \r
	for l.ch == 32 || l.ch == 9 || l.ch == 10 || l.ch == 13 {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	// a <= ch <= z or A <= ch <= Z or _
	return 97 <= ch && ch <= 122 || 65 <= ch && ch <= 90 || ch == 95
}

func isDigit(ch byte) bool {
	// 0 <= ch <= 9
	return 48 <= ch && ch <= 57
}
