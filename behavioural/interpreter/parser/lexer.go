package parser

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
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
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '(':
		tok = Token{Type: LPAREN, Literal: string(l.ch), Pos: l.position}
	case ')':
		tok = Token{Type: RPAREN, Literal: string(l.ch), Pos: l.position}
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = Token{Type: EQUALS, Literal: string(ch) + string(l.ch), Pos: l.position - 1}
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch), Pos: l.position}
		}
	case '<':
		tok = Token{Type: LESS_THAN, Literal: string(l.ch), Pos: l.position}
	case 0:
		tok.Literal = ""
		tok.Type = EOF
		tok.Pos = l.position
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = lookupIdent(tok.Literal)
			tok.Pos = l.position - len(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = NUMBER
			tok.Literal = l.readNumber()
			tok.Pos = l.position - len(tok.Literal)
			return tok
		} else {
			tok = Token{Type: ILLEGAL, Literal: string(l.ch), Pos: l.position}
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
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

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func lookupIdent(ident string) TokenType {
	switch ident {
	case "AND":
		return AND
	case "OR":
		return OR
	case "true":
		return IDENTIFIER
	case "false":
		return IDENTIFIER
	default:
		return IDENTIFIER
	}
}
